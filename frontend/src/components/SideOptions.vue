<template>
    <el-menu-item :index="config.index">
        <i :class="config.class"></i>
        <span slot="title">{{ config.text }}</span>
    </el-menu-item>
</template>

<script>
export default {
    name: 'SideOptions',
    props: {
        config: Object,
    }
}
</script>
