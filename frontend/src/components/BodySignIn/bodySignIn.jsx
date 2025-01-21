import './bodySignIn.css'
import { Link } from "react-router-dom";

function BodySignIn() {
    return (
        <section className="body-sign-in-container">
            <div className="form-container">
                <h3>
                    Sign in to Cinema Chain Name
                </h3>
                <form>
                    <label>
                        Email adress
                    </label>
                    <input>

                    </input>
                    <label>
                        Password
                    </label>
                    <input>

                    </input>
                    <button className="sign-in-form-button">
                        Sign in
                    </button>
                    <span>Don't have an account? <Link to="/register">Register</Link></span>
                </form>
            </div>
        </section>
    )
}

export default BodySignIn;