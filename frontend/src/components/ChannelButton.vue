<template>
    <el-menu-item @click="changeFocus">
        <span slot="title">{{ room }}</span>
    </el-menu-item>
</template>

<script>
export default {
    name: "ChannelButton",
    props: {
        room: String,
    },
    methods: {
        changeFocus() {
            this.$emit('change-focus', this.room)
        }
    }
}
</script>

<style scoped>

</style>