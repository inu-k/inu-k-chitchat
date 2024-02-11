import '../App.css';
import { Link } from 'react-router-dom';
import { formatDate } from '../functions/utils.jsx';

// show threads
export default function ThreadPanel({ thread }) {
    return (
        <div className='thread-panel'>
            <div className='thread-panel-heading'>
                {thread.topic}
            </div>
            <div className='thread-panel-body'>
                {`started by: ${thread.userId} - created at: ${formatDate(thread.createdAt)} - number of posts: ${thread.postsNum}.  `}
                <Link to={`/posts?threads_uuid=${thread.uuid}`}>
                    Read more
                </Link>
            </div>
        </div>
    )
}