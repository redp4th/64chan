<template>
    <div class="message" :class="who">
        <div class="message-bubble" :class="who">
            {{ message.payload }}
        </div>
        <div v-if="who == 'other'" style="font-size: 0.9em;">{{ message.sender }}</div>
    </div>
</template>

<script>
export default {
    name: 'Message',
    props: {
        message: Object,
        who: String,
    },
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


</style>
