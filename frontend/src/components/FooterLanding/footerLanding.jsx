import './footerLanding.css'
import { CinemaChainName } from '../../utils';

function FooterLanding() {
    return (
        <footer className='footer-landing'>
            <p>
                © {CinemaChainName}. All rights reserved.
            </p>
        </footer>
    )
}

export default FooterLanding;