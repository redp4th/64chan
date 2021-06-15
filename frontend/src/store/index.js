import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

import message from './message'
import jar from './jar'

export const store = new Vuex.Store({
    modules: {
        jar,
        message,
    }
})