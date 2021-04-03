<template>
  <div class="input-box">
    <input type="text" placeholder="Type here..." @keyup.enter="send" v-model="msg" />
  </div>
</template>

<script>
export default {
  name: "Send",
  data() {
    return {
      msg: ""
    };
  },
  methods: {
    send() {
      if (this.msg !== "") {
        const obj = {
          type: 1,
          username: this.socketApi.username,
          content: this.msg
        };
        this.socketApi.globalSocketSend(obj);
        this.msg = "";
      }
    }
  }
};
</script>

<style scoped>
.input-box {
  /* position: absolute; */
  width: 100%;
  height: 60px;
  /* bottom: 0; */
  border: 10px solid #000;
}

input {
  outline: none;
  border: none;
  font-size: 100%;
  padding: 10px 0 5px 10px;
  width: 100%;
}
</style>
