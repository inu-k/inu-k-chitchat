import '../App.css'
import { Link } from 'react-router-dom'

export default function ThreadPanel({ thread }) {
    return (
        <div className='thread-panel'>
            <div className='thread-panel-heading'>
                {thread.topic}
            </div>
            <div className='thread-panel-body'>
                {thread.userId} {thread.createdAt} {thread.uuid}
                <Link to={`/posts?threads_uuid=${thread.uuid}`}>
                    Read more
                </Link>
            </div>
        </div>
    )
}