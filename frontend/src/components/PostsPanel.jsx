// show posts to a thread
export default function PostsPanel({ posts }) {

    return (
        <div className="posts-panel">
            <div className="posts-panel-heading">
                aaa
            </div>
            {
                posts.map((post) => {
                    return (
                        <div className='posts-panel-body' style={{ display: "flex", justifyContent: "space-between" }}>
                            <div>
                                {post.body}
                            </div>
                            <div>
                                {`作成者ID: ${post.userId} - 作成日時: ${post.createdAt}`}
                            </div>
                        </div>
                    )
                })
            }
        </div>
    )
}