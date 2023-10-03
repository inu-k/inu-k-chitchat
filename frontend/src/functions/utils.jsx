export async function fetchData(url) {
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data: ', error);
        throw error;
    }
}

export function formatDate(date) {
    const createdAt = new Date(date);

    const year = createdAt.getFullYear();
    const month = createdAt.getMonth() + 1;
    const day = createdAt.getDate();
    const hours = createdAt.getHours();
    const minutes = createdAt.getMinutes();
    const seconds = createdAt.getSeconds();

    const formattedData = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedData;
}
