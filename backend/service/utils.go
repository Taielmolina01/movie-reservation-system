package service

import (
	"reflect"
	"strings"
	"movie-reservation-system/models"
	"movie-reservation-system/errors"
)

func trimStructFields(s interface{}) {
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := val.Type().Field(i).Name

			if field.Kind() == reflect.String && fieldName != "Password" {
				field.SetString(strings.TrimSpace(field.String()))
			}
		}
	}
}

func ValidateUserFields(req *models.UserRequest) error {
	trimStructFields(req)

	if req.Email == "" {
		return errors.ErrorUserMustHaveEmail{}
	}

	if req.Name == "" {
		return errors.ErrorUserMustHaveName{}
	}

	if len(req.Password) < 8 {
		return errors.ErrorPasswordMustHaveLenght8{}
	}

	if !Contains(models.GetRoles(), string(req.Role)) {
		return errors.ErrorUserRoleInvalid{Role: string(req.Role)}
	}

	return nil
}

func ValidateUserUpdateFields(req *models.UserUpdateRequest) error {
	trimStructFields(req)

	if req.Name != nil && *req.Name == "" {
		return errors.ErrorUserMustHaveName{}
	}

	if req.Password != nil && len(*req.Password) < 8 {
		return errors.ErrorPasswordMustHaveLenght8{}
	}

	if req.Role != nil && !Contains(models.GetRoles(), string(*req.Role)) {
		return errors.ErrorUserRoleInvalid{Role: string(*req.Role)}
	}

	return nil
}

func Contains(slice []string, value string) bool {
    for _, v := range slice {
        if v == value {
            return true
        }
    }
    return false
}