import NavbarNotLanding from '../../components/NavbarNotLanding/navbarNotLanding'
import BodyRegister from '../../components/BodyRegister/bodyRegister'
import FooterLanding from '../../components/FooterLanding/footerLanding'
import '../landingPage/landingPage.css'

function RegisterPage() {
    return (
        <div className="container-landing">
            <NavbarNotLanding/>
            <BodyRegister />
            <FooterLanding/>
        </div>
    )
}

 export default RegisterPage;