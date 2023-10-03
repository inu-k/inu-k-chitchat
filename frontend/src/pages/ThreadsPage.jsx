import { useEffect, useState } from "react"
import ThreadPanel from "../components/ThreadPanel.jsx"
import { fetchData } from "../functions/utils.jsx";

// threads page
export default function ThreadsPage() {
    const [threads, setThreads] = useState([]);

    useEffect(() => {
        fetchData('http://localhost:8999/threads')
            .then(data => setThreads(data))
            .catch(error => console.error('Error fetching data: ', error));
    }, []);

    console.log(threads);
    return (
        <div className="container">
            <h1>Threads</h1>
            {threads.map((thread) => {
                return (
                    <ThreadPanel key={thread.id} thread={thread} />
                )
            })}
        </div>
    )
}