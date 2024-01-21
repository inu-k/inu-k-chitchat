import { Link } from 'react-router-dom';
import Cookie from 'js-cookie';
import { useEffect, useState } from 'react';
import { fetchData } from '../functions/utils';

// header
export default function NavBar() {
    const [userInfo, setUserInfo] = useState([]);
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
        fetchData('http://localhost:8999/users/me', {
            mode: 'cors',
            credentials: 'include',
        })
            .then(data => setUserInfo(data))
            .catch(error => console.error('Error fetching data: ', error));

        if (userInfo.length !== null) {
            setIsLoggedIn(true);
        }
    }, []);



    console.log('Cookie: ', Cookie.get('cookie'));
    console.log('isLoggedIn: ', isLoggedIn);

    return (
        <div className='NavBar'>
            <div className='NavBar-bland'>
                <Link style={{ textDecoration: 'none', color: '#666666' }} to='/'>
                    <p>inu-k-ChitChat</p>
                </Link>
            </div>
            {isLoggedIn ? (
                <div className='login-button'>
                    Hello, {userInfo.name}!
                    <Link style={{ textDecoration: 'none', color: '#666666' }} to='/logout'>
                        <p>Logout</p>
                    </Link>
                </div>
            ) : (
                <div className='login-button'>
                    <Link style={{ textDecoration: 'none', color: '#666666' }} to='/login'>
                        <p>Login</p>
                    </Link>
                </div>
            )}
        </div >
    )
}