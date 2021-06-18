import axios from 'axios'
import https from 'https'

const instance = axios.create({
    httpsAgent: new https.Agent({
        rejectUnauthorized: false
    })
})

const v1 = '10.128.196.86:8000/api/v1'

export const api = {
    requestCookie,
    createChannel,
    newSecureSocket,
}

async function requestCookie() {
    return instance.get(`https://${v1}/getcookie`)
        .then(data => {
            if (data.status != 200) {
                return Promise.reject(data.data.message)
            }
            console.log(data.data)
            return data.data
        })
}

function createChannel({ channel }) {
    instance.post(`https://${v1}/channel`, {
        channel
    })
}

function newSecureSocket() {
    return new WebSocket(`wss://${v1}/ws`)
}