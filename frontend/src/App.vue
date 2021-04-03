<template>
  <div id="app">
    <login-page class="login-page" @onLogin="onLoginClick"></login-page>
    <chat-page class="chat-page" ref="send"></chat-page>
  </div>
</template>

<script>
import loginPage from "./components/login";
import chatPage from "./components/chat";
import $ from './common/jquery-3.4.1.js'

export default {
  name: "app",
  components: {
    loginPage,
    chatPage
  },
  methods:{
    onLoginClick() {
      $('.login-page').toggleClass("hiden");
      $('.chat-page').toggleClass("showed");
      this.$refs.send.setSendBoxPosition();
    }
  }
};
</script>

<style>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

#app {
  font-weight: 300;
  -webkit-font-smoothing: antialiased;
}

#app,
input {
  font-family: "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue",
    Helvetica, Arial, "Lucida Grande", sans-serif;
}

#app {
  height: 100%;
}

.login-page {
  transition: .6s;
  left: 0;
}

.chat-page {
  left: -100%;
}

.showed {
  left: 0;
}

.hiden {
  left: -100%
}

</style>
