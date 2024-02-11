import { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';

export function CreateThreadPage() {
    const [topicName, setTopicName] = useState('')
    const navigate = useNavigate();

    const handleChange = (e) => {
        setTopicName(e.target.value);
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = { 'topic': topicName };

        try {
            const response = await fetch('http://localhost:8999/threads', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data),
                credentials: 'include',
            })
            const retdata = await response.json();
            console.log('retdata: ', retdata);
            navigate(`/posts?threads_uuid=${retdata.uuid}`);
        } catch (error) {
            console.error('Error creating a thread: ', error);
            // throw error;
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <div>Topic name: </div>
                <input type='text' value={topicName} onChange={handleChange} />
                <button type='submit'>Create thread</button>
            </div>
        </form>
    );
}