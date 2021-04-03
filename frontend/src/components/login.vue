<template>
  <div class="Wrapper">
    <div class="form">
      <h3 class="title">What's your nickname?</h3>
      <input
        class="usernameInput"
        type="text"
        maxlength="14"
        autofocus="autofocus"
        @keyup.enter="login"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: "loginPage",
  methods: {
    login() {
      if (!("WebSocket" in window)) {
        alert("您的浏览器不支持！推荐使用Chrome，Firefox或Safari");
      } else {
        this.socketApi.globalSocketOpen(
          document.getElementsByClassName("usernameInput")[0].value
        );
        this.$emit("onLogin");
      }
    }
  }
};
</script>

<style scoped="scoped">
.Wrapper {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
  background-color: #000;
}

.title,
.usernameInput {
  color: #fff;
  font-weight: 100;
  font-size: 200%;
}

.title {
  text-align: center;
  margin-bottom: 30px;
}

.usernameInput {
  outline: none;
  border: none;
  border-bottom: 2px solid #fff;
  background-color: transparent;
  width: 400px;
  text-align: center;
  padding-bottom: 15px;
  letter-spacing: 3px;
}
</style>
