import { Link } from 'react-router-dom';

// header
export default function NavBar() {
    return (
        <div className='NavBar'>
            <div className='NavBar-bland'>
                <Link style={{ textDecoration: 'none', color: '#666666' }} to='/'>
                    <p >inu-k-ChitChat</p>
                </Link>
            </div>
        </div >
    )
}