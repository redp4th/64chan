<template>
    <div id="chatroom">
        <div v-if="focusing == null">
                <el-button type="text" @click="createChannel">创建Channel</el-button>
                <div v-if="channels.length > 0">
                    <el-divider>Channel列表</el-divider>
                    <el-menu>
                        <channel-button
                            v-for="channel in channels"
                            socketURL="https://10.128.196.86:3000/"
                            :key="channel"
                            :room="channel"
                            @change-focus="switchFocus"
                        ></channel-button>
                    </el-menu>
                </div>
                <div v-else style="">
                    Channel为空    
                </div>
        </div>
        <div class="chat-panel" v-else>
            <el-header id="dialog-header">
                <el-page-header @back="goBack" :content="focusing">
                </el-page-header>
            </el-header>
            <vue-webrtc ref="webrtc"
                width="100%"
                :roomId="focusing"
                @error="onError" />
            <el-main id="dialog">
                <message
                    v-for="message in filterMessage(focusing)"
                    :key="JSON.stringify(message)"
                    :message="message"
                    :from-me="cookie == message.sender"
                    >
                </message>
            </el-main>
            <el-footer id="dialog-footer">
                <el-row class="footer-line">
                    <el-col :span="2">
                        <el-button icon="el-icon-video-camera" size="mini" circle @click="onJoin" v-if="!recording"></el-button>
                        <el-button icon="el-icon-video-camera" size="mini" circle type="danger" @click="onLeave" v-else></el-button>
                    </el-col>
                    <el-col :span="2">
                        <el-button icon="el-icon-monitor" size="mini" circle :type="screenSharing ? 'danger' : ''" :disabled="!recording" @click="shareScreen"></el-button>
                    </el-col>
                    <el-col :span="2">
                        <el-upload
                            accept="image/png, image/jpg, image/jpeg"
                            :show-file-list="false"
                            :data="{ sender: this.cookie, channel: this.focusing }"
                            action="https://10.128.196.86:8000/api/v1/pic"
                        >
                        <el-button icon="el-icon-picture-outline" size="mini" circle></el-button>
                        </el-upload>
                    </el-col>
                    <el-col :span="2">
                        <el-upload
                            :show-file-list="false"
                            :data="{ sender: this.cookie, channel: this.focusing }"
                            action="https://10.128.196.86:8000/api/v1/upload">
                            <el-button icon="el-icon-link" size="mini" circle></el-button>
                        </el-upload>
                    </el-col>
                </el-row>
                <el-input
                    type="textarea"
                    :row="6"
                    placeholder="请输入内容"
                    v-model="textarea"
                    resize="none"
                    size="medium"
                    class="footer-line"
                >
                </el-input>
                <el-row class="footer-line">
                    <el-col :offset="23">
                        <el-button size="mini" plain @click="sendText">发送</el-button>
                    </el-col>
                </el-row>
            </el-footer>
        </div>
    </div>
</template>

<script>
import ChannelButton from './ChannelButton.vue'
import Message from './Message.vue'

import { code, api } from '../service'

import { mapGetters, mapState } from 'vuex'

export default {
    name: 'Chatroom',
    data() {
        return {
            focusing: null,
            textarea: '',
            recording: false,
            screenSharing: false,
        }
    },
    computed: {
        ...mapState('jar', ['cookie']),
        ...mapGetters('message', ['filterMessage', 'channels']),
    },
    components: {
        ChannelButton,
        Message,
    },
    methods:{ 
        goBack() {
            this.focusing = null
        },
        switchFocus(newFocus) {
            this.focusing = newFocus
        },
        createChannel() {
            this.$prompt('请输入Channel名', '提示', {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                inputPattern: /.+/,
                inputErrorMessage: 'Channel名不能为空'
            })
            .then( ({ value }) => {
                api.createChannel({ channel: value })
            })
            .catch( () => {
                this.$message({
                    type: 'info',
                    message: '取消创建'
                })
            })
        },
        sendText() {
            if (this.textarea.trim() != '') {
                let message = {
                    sender: this.cookie,
                    channel: this.focusing,
                    kind: code.text,
                    payload: this.textarea,
                    timestamp: new Date().toLocaleString()
                }
                this.$bus.$emit('send-message', JSON.stringify(message))
            } else {
                this.$message.error('发送内容为空')
            }
            this.textarea = ''
        },
        onJoin() {
            this.recording = true
            this.$refs.webrtc.join()
        },
        onLeave() {
            this.recording = false
            this.$refs.webrtc.leave()
        },
        onError(error, stream) {
            console.log('On Error Event', error, stream)
        },
        shareScreen() {
            if (!this.screenSharing) {
                this.$refs.webrtc.shareScreen()
            }
            this.screenSharing = !this.screenSharing
        }
    },
}

</script>

<style scoped>
#chatroom {
    height: 100%;
    width: 100%;
    overflow-x: hidden;
}


.chat-panel {
    display: flex;
    flex-direction: column;
    height: 80%;
}

#dialog-header {
    height: 20%;
}

#dialog {
    height: 30%;
}

#dialog-footer {
    height: 40%;
}

.footer-line {
    margin: 0.5em;
}

#video {
    z-index: 1;
}

#messages {
    height: auto;
    overflow: scroll;
}
</style>