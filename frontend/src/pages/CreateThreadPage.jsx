import { useState } from 'react';

function SetTopicName() {
}

export function CreateThreadPage() {
    const [topicName, setTopicName] = useState('');

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
                body: JSON.stringify(data)
            })
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
        } catch (error) {
            console.error('Error fetching data: ', error);
            throw error;
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