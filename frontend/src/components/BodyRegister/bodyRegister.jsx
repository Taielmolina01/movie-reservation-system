import './bodyRegister.css'
import { Link } from "react-router-dom";

function BodyRegister() {
    return (
        <section className="body-register-container">
            <div className="form-container">
                <h3>
                    Register in Cinema Chain Name
                </h3>
                <form>
                    <label>
                        User name
                    </label>
                    <input>
                    
                    </input>
                    <label>
                        Email adress
                    </label>
                    <input>

                    </input>
                    <label>
                        Password (min 8 characters)
                    </label>
                    <input>

                    </input>
                    <label>
                        Confirm your password
                    </label>
                    <input>
                    
                    </input>
                    <button className="sign-in-form-button">
                        Register
                    </button>
                </form>
                <span>
                    Already have an account? <Link to="/sign-in">Sign in</Link>
                </span>
            </div>
        </section>
    )
}

export default BodyRegister;