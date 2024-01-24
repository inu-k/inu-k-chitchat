export async function fetchData(url, options = {}) {
    try {
        const response = await fetch(url, options);
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
    const minutes = ('0' + createdAt.getMinutes()).slice(-2);
    const seconds = ('0' + createdAt.getSeconds()).slice(-2);

    const formattedData = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedData;
}

export async function fetchUserInfo() {
    try {
        const data = await fetchData('http://localhost:8999/users/me', {
            mode: 'cors',
            credentials: 'include',
        });
        console.log('response data in utils: ', data);

        return data;
    } catch (error) {
        console.error('Error fetching data in fetchUserInfo: ', error);
        throw error;
    }
}
