import { Link, useNavigate } from 'react-router-dom';
import Cookie from 'js-cookie';
import { useEffect, useState } from 'react';
import { fetchUserInfo } from '../functions/utils';

// header
export default function NavBar({ isLoggedIn, setIsLoggedIn }) {
    const [userInfo, setUserInfo] = useState({});
    const navigate = useNavigate();

    async function getUserInfoAndUpdateStatus() {
        try {
            const data = await fetchUserInfo()

            console.log('data: ', data);
            setUserInfo(data);

            if (Object.keys(data).length !== 0) {
                setIsLoggedIn(true);
            } else {
                setIsLoggedIn(false);
            }
        }
        catch (error) {
            console.error('Error fetching data: ', error);
        }
    }

    useEffect(() => {
        getUserInfoAndUpdateStatus();
    }, [])

    useEffect(() => {
        console.log('useEffect users/me');
        if (isLoggedIn) {
            getUserInfoAndUpdateStatus();
        }
    }, [isLoggedIn]);

    const HandleLogout = async (e) => {
        e.preventDefault();
        setIsLoggedIn(false);

        try {
            const response = await fetch('http://localhost:8999/sessions/me', {
                method: 'DELETE',
                mode: 'cors',
                credentials: 'include',
            });

            if (response.ok) {
                navigate('/');
            } else {
                console.error('Network response was not ok', response);
            }
        } catch (error) {
            console.error('Error fetching data: ', error);
        }
    }

    console.log('Cookie: ', Cookie.get('_cookie'));
    console.log('isLoggedIn: ', isLoggedIn);

    return (
        <div className='NavBar'>
            <div className='NavBar-bland'>
                <Link style={{ textDecoration: 'none', color: '#666666' }} to='/'>
                    <p>inu-k-ChitChat</p>
                </Link>
            </div>
            {isLoggedIn ? (
                <div className='user-info-and-logout-button'>
                    <span className='NavBar-user-info'>Hello, {userInfo.name}!</span>
                    <Link className='logout-button' onClick={HandleLogout} style={{ textDecoration: 'none', color: '#666666' }} to='/logout'>
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