import '../App.js'
import { Link, useNavigate } from 'react-router-dom';

export default function LoginForm({ setIsLoggedIn }) {
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);

        try {
            const response = await fetch('http://localhost:8999/sessions', {
                method: 'POST',
                body: formData,
                credentials: 'include',
            });

            console.log('cookie: ', document.cookie)

            if (response.ok) {
                setIsLoggedIn(true);
                navigate('/');
            } else {
                throw new Error('Network response was not ok');
            }
        } catch (error) {
            console.error('Error fetching data: ', error);
            throw error;
        }
    }
    return (
        <div className='login-form'>
            <form onSubmit={handleSubmit}>
                <input type='email' name='email' placeholder='Email' required />
                <br />
                <input type='password' name='password' placeholder='Password' required />
                <br />
                <button type='submit'>Login</button>
            </form>

            <div className='signup-button'>
                Don't have an account? <Link to='/signup'>Signup now!</Link>
            </div>
        </div>
    )
}