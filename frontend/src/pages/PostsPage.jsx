import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import PostsPanel from '../components/PostsPanel.jsx';

// show posts page
// posts to a thread with thread information
export default function PostsPage() {
    const location = useLocation();
    const queryParams = new URLSearchParams(location.search);
    const threadsUuid = queryParams.get('threads_uuid');
    console.log('threadsUuid: ', threadsUuid);

    const [posts, setPosts] = useState([]);

    useEffect(() => {
        async function fetchData() {
            try {
                const response = await fetch(`http://localhost:8999/posts?thread_uuid=${threadsUuid}`);

                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const data = await response.json();
                setPosts(data);
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        }

        fetchData();
    }, []);

    return (
        <div className="container">
            <h1>Posts</h1>
            <PostsPanel posts={posts} />
        </div>
    )
}