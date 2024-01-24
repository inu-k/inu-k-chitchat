import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import PostsPanel from '../components/PostsPanel.jsx';
import { fetchData } from '../functions/utils.jsx';

// show posts page
// posts to a thread with thread information
export default function PostsPage({ isLoggedIn }) {
    const location = useLocation();
    const queryParams = new URLSearchParams(location.search);
    const threadsUuid = queryParams.get('threads_uuid');
    console.log('threadsUuid: ', threadsUuid);

    const [posts, setPosts] = useState([]);
    const [thread, setThread] = useState({});

    useEffect(() => {
        fetchData(`http://localhost:8999/posts?thread_uuid=${threadsUuid}`)
            .then(data => setPosts(data))
            .catch(error => console.error('Error fetching data: ', error));
    }, []);

    useEffect(() => {
        fetchData(`http://localhost:8999/threads/${threadsUuid}`)
            .then(data => setThread(data))
            .catch(error => console.error('Error fetching data: ', error));
    }, []);

    return (
        <div className="container">
            <h1>Posts</h1>
            <PostsPanel posts={posts} thread={thread} isLoggedIn={isLoggedIn} />
        </div>
    )
}