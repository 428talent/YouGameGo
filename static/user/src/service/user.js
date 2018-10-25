import fetch from 'dva/fetch';

export async function FetchUser(userId) {
    console.log(`fetch user id = ${userId}`);
    const response = await fetch(`http://localhost:8080/v1/user/${userId}`);
    const data = await response.json();
    return data
}