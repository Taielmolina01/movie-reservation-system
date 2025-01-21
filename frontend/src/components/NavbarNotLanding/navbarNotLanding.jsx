import '../NavbarLanding/navbarLanding.css'
import { Link } from "react-router-dom";

function NavbarNotLanding() {
    return (
        <nav>
            <div className='nav-real-content'>
                <Link to="/">
                    <h1>CinemaChainName</h1>
                </Link>
            </div>
        </nav>
    )
}

export default NavbarNotLanding;