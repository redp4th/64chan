<template>
    <div id="welcome">
        <div id="slogan">
            64Chan<br>
            匿名版
        </div>
        <div id="ad">
            无需登录,凭借Cookie即可随时接入聊天室<br>
            匿名聊天,匿名视频通话,匿名桌面共享
        </div>
        <button @click="handleCookie" :disabled="disabled">Enter</button>
    </div>
</template>

<script>
import { router } from '../router'
import { api } from '../service'


export default {
    name: 'Welcome',
    data() {
        return {
            disabled: false,
        }
    },
    methods: {
        handleCookie() {
            if ((localStorage.getItem('token')) != null) {
                router.push('/chat')
                return
            }
            api.requestCookie()
                .then(
                    data => {
                        localStorage.setItem('token', data.token)
                        router.push('/chat')
                    },
                    err => {
                        this.$notify.error({
                            title: '连接失败',
                            message: err
                        })
                        this.disabled = true
                    }
                )
        }
    }
}

</script>

<style scoped>
#welcome {
    padding-left: 100px;
}

#slogan {
    padding-top: 150px;
    font-size: 3em;
    color: black;
}

#ad {
    font-size: 2em;
}

button {
    font-size: 1.5em;
    font-weight: bold;
    height: 2em;
    width: 4em;
    margin-top: 2em;
    cursor: pointer;
    border-radius: 0.3em;
    border: 2px solid black;
    background-color: white;
}

button:hover {
    background-color: #eee;
}
</style>