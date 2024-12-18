import NavbarLanding from '../components/NavbarLanding/navbarLanding'
import BodyLanding from '../components/BodyLanding/bodyLanding'
import FooterLanding from '../components/FooterLanding/footerLanding'
import './landingPage.css'

function LandingPage() {
    return (
        <div className="container-landing">
            <NavbarLanding />
            <BodyLanding />
            <FooterLanding />
        </div>
    )
}

export default LandingPage;