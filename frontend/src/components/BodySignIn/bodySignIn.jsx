import './bodySignIn.css'
import { Link } from "react-router-dom";
import { useState } from 'react';
import { CinemaChainName } from '../../utils';
import { useNavigate } from 'react-router-dom';

function BodySignIn() {


    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(false)
    const [formData, setFormData] = useState({
        email: "",
        password: "",
    });
    const navigate = useNavigate();


    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

    const handleSubmit = async(e) => {
        e.PreventDefault();

        setError(null);

        await HandleSignIn(formData.email, formData.password, setLoading, setError, navigate);
    }
    

    return (
        <section className="body-sign-in-container">
            <div className="form-sign-in-container">
                <h3>
                    Sign in to {CinemaChainName}
                </h3>
                <form>
                    <label>
                        Email adress
                    </label>
                    <input
                        type="email"
                        value={formData.email}
                        onChange={handleChange}
                    >

                    </input>
                    <label>
                        Password
                    </label>
                    <input
                        type="password"
                        value={formData.password}
                        onChange={handleChange}
                    >

                    </input>
                    <button
                        type="submit"
                        className={loading ? "sign-in-form-button-disabled" : "sign-in-form-button"}
                        onChange={handleSubmit}
                    >
                        {loading ? <LoadingSpinner/> : "Sign in"}
                    </button>
                    {error && <p style={{ color: 'red', maxWidth: '255px', textAlign: 'center' }}>{error}</p>}
                </form>
                <span className="dont-have-an-account">
                    Don't have an account? <Link to="/register">Register</Link>
                </span>
            </div>
        </section>
    )
}

export default BodySignIn;