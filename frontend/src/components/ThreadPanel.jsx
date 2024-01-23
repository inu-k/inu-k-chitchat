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
                {`作成者ID: ${thread.userId} - 作成日時: ${formatDate(thread.createdAt)} - 投稿数: ${thread.postsNum}.  `}
                <Link to={`/posts?threads_uuid=${thread.uuid}`}>
                    Read more
                </Link>
            </div>
        </div>
    )
}