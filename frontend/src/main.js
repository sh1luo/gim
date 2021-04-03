import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false
import * as socketApi from './common/socket.js'
Vue.prototype.socketApi = socketApi

new Vue({
  render: h => h(App),
}).$mount('#app')
