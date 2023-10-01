import '../App.css'
import { Link } from 'react-router-dom'

// show threads
export default function ThreadPanel({ thread }) {
    const createdAt = new Date(thread.createdAt);

    const year = createdAt.getFullYear();
    const month = createdAt.getMonth() + 1;
    const day = createdAt.getDate();
    const hours = createdAt.getHours();
    const minutes = createdAt.getMinutes();
    const seconds = createdAt.getSeconds();

    const formattedData = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;

    return (
        <div className='thread-panel'>
            <div className='thread-panel-heading'>
                {thread.topic}
            </div>
            <div className='thread-panel-body'>
                {`作成者ID: ${thread.userId} - 作成日時: ${formattedData} - 投稿数: ${thread.postsNum}.  `}
                <Link to={`/posts?threads_uuid=${thread.uuid}`}>
                    Read more
                </Link>
            </div>
        </div>
    )
}