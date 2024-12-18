import './navbarLanding.css'

function NavbarLanding() {
    return (
        <nav>
            <div className='nav-real-content'>
                <h1>CinemaChainName</h1>
                <div className='box-right'>
                    <button className='sign-in-button'>
                        SIGN IN
                    </button>
                    <button className='register-button'>
                        REGISTER
                    </button>
                </div>
            </div>
        </nav>
    )
}

export default NavbarLanding;