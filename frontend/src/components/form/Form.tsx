import axios from 'axios';
import { useEffect, useState } from 'react';

type cloneData = {
    "message": string,
    "id": string,
    "repo_url": string,
}
export const Form = () => {
    const [logs, setLogs] = useState<string[]>([]);
    const [queueName, setQueueName] = useState<string>(''); // default queue name

    useEffect(() => {
        if (queueName != '') {
            const eventSource = new EventSource(`http://localhost:1234/events?queue=${queueName}`);
            const now = new Date();

            const hours = now.getHours();
            const minutes = now.getMinutes();
            const seconds = now.getSeconds();

            const time = `${hours}:${minutes}:${seconds}`

            eventSource.onmessage = (event: MessageEvent) => {
                setLogs(prevLogs => [...prevLogs, time + "    " + event.data]);
            };

            eventSource.onerror = (event: Event) => {
                console.error('EventSource failed:', event);
                eventSource.close();
            };

            return () => {
                eventSource.close();
            };
        }
    }, [queueName]);

    const [data, setData] = useState<cloneData>();
    const [error, setError] = useState(null);

    const extractAfterGithub = (url: string) => {
        const regex = /https:\/\/github\.com\/(.*)/;
        const match = url.match(regex);
        return match ? match[1] : null;
    };

    const startProcess = async (after: any) => {
        try {
            console.log();
            const response = await axios.get(`${import.meta.env.VITE_BASE_URL}/clone/${after}`);
            setData(response.data);
        } catch (error: any) {
            setError(error);
        }
    };

    const isGithubRepoUrl = (url: string) => {
        const regex = /^https:\/\/github\.com\/[^/]+\/[^/]+$/;
        return regex.test(url);
    };

    const handleSubmit = (e: any) => {
        e.preventDefault();
        setLogs([]);
        const url = e.target.email.value || "";
        if (isGithubRepoUrl(url) == false) {
            alert("Not a valid URL");
            return;
        }
        const params = extractAfterGithub(url);
        startProcess(params);
    }

    useEffect(() => {
        if (data) {
            setQueueName(`log:${data.id}`);
            return;
        }
        console.log(error);

    }, [data, error]);

    return (

        <div>
            <form className="max-w-sm mx-auto mt-10 border rounded p-4 shadow" onSubmit={handleSubmit}>
                <div className="mb-5">
                    <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Github URL</label>
                    <input type="text" id="email" name="email" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="https://github.com/username/repo" required />
                </div>
                <button type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Submit</button>
            </form>
            <div className='container border rounded-3xl md:w-6/12 w-full mt-10 shadow p-0' style={{ overflow: 'hidden', background: 'whitesmoke' }}>
                <div id="logs" className='p-10' style={{ height: '300px', overflowY: 'scroll' }}>
                    {logs.map((log, index) => (
                        <p key={index}>{log}</p>
                    ))}
                </div>
            </div>
        </div>
    )

}