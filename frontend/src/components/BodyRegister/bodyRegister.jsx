import { useState } from 'react';
import './bodyRegister.css'
import { Link } from "react-router-dom";
import { CinemaChainName } from '../../utils';

function BodyRegister() {

    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(false)
    const [formData, setFormData] = useState({
        userName: "",
        email: "",
        password: "",
        confirmPassword: "",
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };


    const handleSubmit = async(e) => {
        e.preventDefault();


    }

    return (
        <section className="body-register-container">
            <section className="left-section">
                <h2>
                    Create your free account
                </h2>
                <h4>
                    Once you have an account you can buy your tickets online!
                </h4>
            </section>
            <section className="right-section">
                <div className="form-register-container">
                    <h3>
                        Register in {CinemaChainName}
                    </h3>
                    <form onSubmit={handleSubmit}>
                        <label>
                            User name
                        </label>
                        <input
                            type="text"
                            value={formData.userName}
                            onChange={handleChange}
                        >
                        
                        </input>
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
                        <span className="password-helper">
                            Password must have at least 8 characters
                        </span>
                        <label>
                            Repeat your password
                        </label>
                        <input
                            type="password"
                            value={formData.confirmPassword}
                            onChange={handleChange}
                        >
                        
                        </input>
                        <button 
                            className="sign-in-form-button"
                            type="submit"
                        >
                            Register
                        </button>
                        {error && <p style={{ color: 'red', maxWidth: '255px', textAlign: 'center' }}>{error}</p>}
                    </form>
                    <span className="have-an-account">
                        Already have an account? <Link to="/sign-in">Sign in</Link>
                    </span>
                </div>
            </section>
        </section>
    )
}

export default BodyRegister;