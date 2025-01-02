package auth

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"math/rand"
	ownErrors "movie-reservation-system/errors"
	"movie-reservation-system/models"
	authRepository "movie-reservation-system/repository/auth"
	userRepository "movie-reservation-system/repository/user"
	"movie-reservation-system/service"
	"os"
	"time"
)

type AuthServiceImpl struct {
	userRepository userRepository.UserRepository
	authRepository authRepository.AuthRepository
}

var signingMethod jwt.SigningMethod

const expiresHours = 2

func init() {
	algo := os.Getenv("JWT_ALGORITHM")
	switch algo {
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
	case "RS256":
		signingMethod = jwt.SigningMethodRS256
	default:
		log.Fatalf("Unsupported JWT algorithm: %s", algo)
	}
}

func NewLoginServiceImpl(userRepository userRepository.UserRepository, authRepository authRepository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func mapTokenRequestToTokenDB(token, refreshToken, userEmail string, expiresAt time.Time) *models.TokenDB {
	return &models.TokenDB{
		ID:           uuid.New(),
		AccessToken:  token,
		RefreshToken: refreshToken,
		UserEmail:    userEmail,
		ExpiresAt:    expiresAt,
	}
}

func (aui *AuthServiceImpl) Login(req *models.UserLoginRequest) (*models.TokenDB, error) {
	user, err := aui.userRepository.GetUser(req.Email)
	if err != nil {
		return nil, ownErrors.ErrorUserNotExist{Email: req.Email}
	}

	if !service.ValidatePassword(user.Password, req.Password) {
		return nil, ownErrors.ErrorWrongOldPassword{}
	}

	expiresAt := time.Now().Add(time.Hour * expiresHours)

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, ownErrors.ErrorSigningToken{TypeError: err}
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, ownErrors.ErrorGeneratingRefreshToken{TypeError: err}
	}

	tokenDB := mapTokenRequestToTokenDB(signedToken, refreshToken, user.Email, expiresAt)

	return aui.authRepository.CreateToken(tokenDB)
}

func (aui *AuthServiceImpl) Logout(userEmail string) (*models.TokenDB, error) {
	_, err := aui.userRepository.GetUser(userEmail)

	if err != nil {
		return nil, ownErrors.ErrorUserNotExist{}
	}

	token, err := aui.authRepository.GetToken(userEmail)

	if err != nil {
		return nil, ownErrors.ErrorUserTokenNotExist{UserEmail: userEmail}
	}

	return aui.authRepository.DeleteToken(token)
}

func generateRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
