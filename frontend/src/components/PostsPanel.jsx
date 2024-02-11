import { formatDate } from '../functions/utils.jsx';
import { CreatePostPanel } from './CreatePostPanel.jsx';

// show posts to a thread
export default function PostsPanel({ posts, thread, isLoggedIn }) {

    return (
        <div>
            <div className="posts-panel">
                <div className="posts-panel-heading">
                    Topic name: {thread.topic}
                </div>
                {
                    posts.map((post) => {
                        return (
                            <div className='posts-panel-body' style={{ display: "flex", justifyContent: "space-between" }}>
                                <div>
                                    {post.body}
                                </div>
                                <div>
                                    {`posted by: ${post.userId} - created at: ${formatDate(post.createdAt)}`}
                                </div>
                            </div>
                        )
                    })
                }
            </div>
            <CreatePostPanel thread={thread} isLoggedIn={isLoggedIn} />
        </div>
    )
}