import NavbarNotLanding from '../../components/NavbarNotLanding/navbarNotLanding'
import BodySignIn from '../../components/BodySignIn/bodySignIn'
import FooterLanding from '../../components/FooterLanding/footerLanding'
import '../landingPage/landingPage.css'

function SignInPage() {
    return (
        <div className="container-landing">
            <NavbarNotLanding/>
            <BodySignIn />
            <FooterLanding/>
        </div>
    )
}

export default SignInPage;