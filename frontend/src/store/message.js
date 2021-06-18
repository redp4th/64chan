import { code } from '../service'

const state = {
    messages: { },
}

const mutations = {
    cache(state, message) {
        if (!Object.prototype.hasOwnProperty.call(state.messages, message.channel))
            state.messages = { ...state.messages, ...{ [message.channel]: [] } }
        state.messages[message.channel].push(message)
    },
}

const getters = {
    channels: state => Object.keys(state.messages),
    filterMessage: state => channel => state.messages[channel].filter(message => message.kind != code.control)
}

export default {
    namespaced: true,
    state,
    mutations,
    getters,
}