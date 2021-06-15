<template>
    <div id="about">
        <div>Cookie: {{ cookie }}</div>
        <el-button @click="replaceCookie">更换Cookie</el-button>
    </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'
import axios from 'axios'

export default {
    name: 'About',
    computed: {
        ...mapState('jar', ['cookie'])
    },
    methods: {
        ...mapMutations('jar', ['setCookie']),
        replaceCookie() {
            axios.get('http://localhost:8000/api/v1/getcookie')
                .then(response => {
                    if (response.status != 200)
                        return Promise.reject(response.data.message)
                    return response.data
                })
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
