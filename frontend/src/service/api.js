import axios from 'axios'

const v1 = 'http://localhost:8000/api/v1'

export const api = {
    requestCookie,
    createChannel
}

function requestCookie() {
    axios.get(`${v1}/getcookie`)
        .then(data => {
            if (data.status != 200) {
                return Promise.reject(data.data.message)
            }
            return data.data
        })
}

function createChannel({ channel }) {
    axios.post(`${v1}/channel`, {
        channel
    })
}