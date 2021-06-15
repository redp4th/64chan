<template>
    <el-col :span="24">
        <el-row justify="center">
            <h3 style="margin-left: 2em;">菜单</h3>
        </el-row>
        <el-menu
        :router="true">
        <side-options
            v-for="item in options"
            :key="item.name"
            :config="item"
        ></side-options>

        </el-menu>
    </el-col>
</template>

<script>
import SideOptions from './SideOptions.vue'

export default {
    name: 'Sidebar',
    components: {
        SideOptions,
    },
    data() {
        return {
            options: [
                {
                    text: '广场',
                    index: '/chat',
                    class: 'el-icon-chat-dot-round',
                },
                {
                    text: '关于',
                    index: '/about',
                    class: 'el-icon-document',
                },
                {
                    text: '退出',
                    index: '/welcome',
                    class: 'el-icon-switch-button',
                }
            ]
        }
    }
}
</script>

<style scoped>

</style>
