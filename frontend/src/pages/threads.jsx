import { useEffect, useState } from "react"
import ThreadPanel from "../components/ThreadPanel.jsx"

export default function ThreadsPage() {
    const [threads, setThreads] = useState([]);

    useEffect(() => {
        async function fetchData() {
            try {
                const response = await fetch('http://localhost:8999/threads');

                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const data = await response.json();
                setThreads(data);
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        }

        fetchData();
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