import { useState } from 'react';
import { Link } from 'react-router-dom';
import { fetchUserInfo } from '../functions/utils';

export function CreatePostPanel({ thread, isLoggedIn }) {
    const [bodyText, setBodyText] = useState('');


    const handleChange = (e) => {
        setBodyText(e.target.value);
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log('thread:', thread)

        const data = {
            'body': bodyText,
            'threadUuid': thread.uuid
        };

        try {
            const response = await fetch('http://localhost:8999/posts', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data),
                credentials: 'include',
            })
            const retdata = await response.json();
            window.location.reload();
        } catch (error) {
            console.error('Error creating a post: ', error);
            // throw error;
        }
    }

    return (
        <div>
            {isLoggedIn ? (
                < form onSubmit={handleSubmit} >
                    <div>
                        <textarea type='text' value={bodyText} onChange={handleChange} />
                        < button type='submit' > Send post!</button >
                    </div >
                </form >
            ) : (
                <div style={{ padding: "10px" }}>
                    <div>Please <Link to='/login'>login</Link> to send a post.</div>
                </div>
            )}
        </div>

    );
}