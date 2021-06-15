const state = {
    cookie: null
}

const mutations = {
    setCookie(state, token) {
        let cookie = parseJWT(token).cookie
        state.cookie = cookie.toString(16)
    }
}

export default {
    namespaced: true,
    state,
    mutations,
}

function parseJWT(token) {
    const url = token.split('.')[1]
    const base64 = url.replace(/-/g, '+').replace(/_/g, '/');
    const payload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(payload)
}