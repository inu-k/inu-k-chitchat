import '../App.js'
import { Link } from 'react-router-dom';

export default function LoginForm() {
    return (
        <div className='login-form'>
            <form action='http://localhost:8999/users' method='post'>
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