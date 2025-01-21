import '../NavbarLanding/navbarLanding.css'
import './navbarNotLanding.css'
import { Link } from "react-router-dom";

function NavbarNotLanding() {
    return (
        <nav>
            <div className='nav-real-content'>
                <Link to="/" style={{textDecoration: 'none', color: 'inherit'}}>
                    <h1>CinemaChainName</h1>
                </Link>
            </div>
        </nav>
    )
}

export default NavbarNotLanding;