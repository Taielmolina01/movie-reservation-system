import './navbarLanding.css'
import { Link } from "react-router-dom";

function NavbarLanding() {
    return (
        <nav>
            <div className='nav-real-content'>
                <h1>CinemaChainName</h1>
                <div className='box-right'>
                    <Link to="/sign-in">           
                        <button className='sign-in-button'>
                            SIGN IN
                        </button>
                    </Link>
                    <Link to="/register">           
                        <button className='register-button'>
                            REGISTER
                        </button>
                    </Link>
                </div>
            </div>
        </nav>
    )
}

export default NavbarLanding;