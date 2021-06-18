<template>
    <div class="message" :class="{ 'other': !fromMe }">
        <div class="message-bubble" :class="{ 'other': !fromMe }">
            <div v-if="message.kind == code.text">
                {{ message.payload }}
            </div>
            <div v-else-if="message.kind == code.picture">
                <img :src="`data:image/jpeg;base64,${message.payload}`" />
            </div>
            <div v-else-if="message.kind == code.file">
                <el-link icon="el-icon-paperclip" :href="message.payload" target="_blank">{{ message.payload.split('/').pop() }}</el-link>
            </div>
            <div v-else>
                ERROR MESSAGE
            </div>
        </div>
        <div v-if="!fromMe" style="font-size: 0.9em;">{{ message.sender }}</div>
    </div>
</template>

<script>
import { code } from '../service'

export default {
    name: 'Message',
    props: {
        message: Object,
        fromMe: Boolean,
    },
    data() {
        return {
            code,
        }
    }
}
</script>

<style scoped>
.message {
    width: auto;
    max-width: 500px;
    position: relative;
    margin: 5px;
    float: right;
    clear: both;
}

.message-bubble {
    padding: 0.3em 0.5em 0.3em 0.5em;
    border-radius: 10px;
    background-color: dodgerblue;
    color: white;
    width: auto;
    position: relative;
    font-size: 1.2em;
}

.message.other {
    float: left;
}

.message-bubble.other {
    background-color: #eee;
    color: black;
    float: left;
}

img {
    max-width: 450px;
    max-height: 450px;
}

</style>
