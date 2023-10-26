import { useState } from 'react';
import { Navigate, useNavigate, useLocation } from 'react-router-dom';


export function CreatePostPanel({ thread }) {
    const [bodyText, setBodyText] = useState('');
    const navigate = useNavigate();


    const handleChange = (e) => {
        setBodyText(e.target.value);
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log('thread:', thread)
        const data = {
            'body': bodyText,
            'userId': 1,
            'threadUuid': thread.uuid
        };

        try {
            const response = await fetch('http://localhost:8999/posts', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            })
            const retdata = await response.json();
            window.location.reload();
        } catch (error) {
            console.error('Error fetching data: ', error);
            throw error;
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <textarea type='text' value={bodyText} onChange={handleChange} />
                <button type='submit'>Send post!</button>
            </div>
        </form>
    );
}