import '../App.js'
import { Link, useNavigate } from 'react-router-dom';

export default function LoginForm() {
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);

        try {
            const response = await fetch('http://localhost:8999/users', {
                method: 'POST',
                body: formData
            });

            if (response.ok) {
                navigate('/login');
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
                <input type='name' name='user-name' placeholder='User name' required />
                <br />
                <input type='email' name='email' placeholder='Email' required />
                <br />
                <input type='password' name='password' placeholder='Password' required />
                <br />
                <button type='submit'>Signup</button>
            </form>

            <div className='signup-button'>
                Already have an account? <Link to='/login'>Login here!</Link>
            </div>
        </div>
    )
}