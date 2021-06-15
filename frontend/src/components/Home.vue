<template>
    <el-container id="container">
        <el-header>
            <h1>64Chan</h1>
        </el-header>
        <el-container id="display">
            <el-aside width="20%">
                <Sidebar></Sidebar>
            </el-aside>
            <el-main>
                <keep-alive>
                    <router-view></router-view>
                </keep-alive>
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
import Sidebar from './Sidebar.vue'

import { mapMutations } from 'vuex'

import { router } from '../router'

export default {
    name: 'Home',
    components: {
        Sidebar,
    },
    data() {
        return {
            socket: null,
        }
    },
    methods: {
        ...mapMutations('message', ['cache', 'newChannel']),
        ...mapMutations('jar', ['setCookie']),
    },
    mounted() {
        this.socket = new WebSocket('ws://10.128.196.86:8000/api/v1/ws')
        this.socket.onerror = (evt) => {
            console.log(evt)
            this.$notify.error({
                title: '建立WebSocket失败',
                message: evt.data,
            })
            router.push('/welcome')
        }

        this.socket.onclose = (evt) => {
            console.log(evt)
            this.$notify.error({
                title: 'WebSocket连接关闭',
                message: evt.data,
            })
            router.push('/welcome')
        }

        this.socket.onmessage = (evt) => {
            let message = JSON.parse(evt.data)
            this.cache(message)
        }

        this.setCookie(localStorage.getItem('token'))
    },
    created() {
        this.$bus.$on('send-message', (message) => this.socket.send(message))
    },
    beforeDestroy() {
        this.$bus.$off('send-message', (message) => this.socket.send(message))
    }
}
</script>

<style scoped>
.el-header {
    background-color: #409EFF;
    height: 20%;
}

h1 {
    color: white;
}

#container {
    height: 100%;
    border: 1px solid #eee;
}

#display {
    height: 75%;
    width: 100%;
}

.el-main {
    height: 100%;
}
</style>