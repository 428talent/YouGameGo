import Cookies from 'js-cookie'

const request = ({url, method, data, queryParams, pathParams}) => {
    console.log(pathParams)
    console.log(data)
    if (queryParams) {
        url = url + "?" + Object.getOwnPropertyNames(queryParams).map(value => {
            return `${value}=${queryParams[value]}`
        }).join("&")
    }
    if (pathParams) {
        Object.getOwnPropertyNames(pathParams).forEach(value => {
            url = url.replace(`:${value}`, pathParams[value])
        })
    }
    return fetch(url, {
        method,
        mode: "cors",
        headers: {
            "Authorization": Cookies.get("yougame_token"),
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
};

export default request