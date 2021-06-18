<template>
    <div id="about">
        <div>Cookie: {{ cookie }}</div>
        <el-button @click="replaceCookie">更换Cookie</el-button>
    </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'
import { api } from '../service'

export default {
    name: 'About',
    computed: {
        ...mapState('jar', ['cookie'])
    },
    methods: {
        ...mapMutations('jar', ['setCookie']),
        replaceCookie() {
            api.requestCookie()
                .then(
                    data => {
                        localStorage.setItem('token', data.token)
                        this.$message({
                            message: '更换成功',
                            type: 'success',
                        })
                        this.setCookie(data.token)
                    },
                    err => {
                        this.$message.error(err)
                    }
                )
        }
    }
}
</script>

<style scoped>
#about {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}
</style>
