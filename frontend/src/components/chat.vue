<template>
  <div class="chat">
    <div class="chatArea" ref="chatArea" v-for="(msg,index) in messages" :key="index">
      <template v-if="msg.type === 1">
        <div class="message">
          <span class="nickname">{{msg.nickname}}</span>
          <span class="msg">{{msg.body}}</span>
        </div>
      </template>
      <template v-else>
        <div class="notice">{{msg.body}}</div>
      </template>
    </div>
    <Send class="send-box" />
  </div>
</template>

<script>
import Send from "./msgInput";
import $ from "../common/jquery-3.4.1";

export default {
  name: "chatPage",
  components: {
    Send
  },
  data() {
    return {
      messages: this.socketApi.msgList
    };
  },
  mounted() {
    $(".chat").css("height", $(window).height() - 68);
  },
  updated() {
    this.scrollToBottom();
  },
  methods: {
    scrollToBottom() {
      var chat = document.getElementsByClassName("chat")[0];
      chat.scrollTop = chat.scrollHeight;
    },
    setSendBoxPosition() {
      $(".send-box")
        .css("position", "absolute")
        .css("bottom", 0);
    }
  }
};
</script>

<style scoped="scoped">
.nickname,
.msg {
  font-size: 150%;
}

.nickname {
  margin: 15px;
  font-weight: 700;
}

.notice {
  font-size: 100%;
  color: gray;
  text-align: center;
  margin: 5px;
}

.chat {
  overflow-y: auto;
  width: 100%;
}

.message{
  margin: 5px 5px;
  max-width: 1400px;
}
</style>
