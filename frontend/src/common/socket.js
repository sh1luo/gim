var websock = null
var againSend = null
//var globalCallback = null
var msgList = []
var username = ''

function initWebSocket(nickname) {
  //const info = JSON.parse(localStorage.getItem('loginInfo'))
  //if (!localStorage.getItem('loginInfo')) return false
  // var ws= `ws://10.2.0.60:7002/socket?name=${info.name}&userId=${info.id}&userType=kefu&deviceType=web`

  //测试服务器
  var ws = `ws://127.0.0.1:8000/?username=${nickname}`
  username = nickname

  //线上服务器
  // var ws= `ws://103.9.230.66:10010/socket?name=${info.name}&userId=${info.id}&userType=kefu&deviceType=web`

  websock = new WebSocket(ws)
  websock.onmessage = function (e) {
    msgList.push(JSON.parse(e.data))
  }
  websock.onopen = function () {
    websocketOpen()
  }
  websock.onerror = function () {
    alert('连接发生错误，即将重新连接。如果还是失败，请手动连接')
    location.reload()
  }
   websock.onclose = function(e){
     websock.close(); //关闭TCP连接
   };
}

function websocketOpen() {
  //alert('连接成功  ', websock.readyState)

  // againSend = setInterval(function () {
  //   var timer = new Date().getTime()
  //   //localStorage.setItem('timer', timer)
  //   websock.send(JSON.stringify({"chatType": 999, "msgId": timer}))
  // }, 60000)
}

function globalSocketOpen(nickname) {
  return initWebSocket(nickname)
}

function globalSocketClose() {
  return websock.close()
}

function globalSetIntervalClose() {
  return clearInterval(againSend)
}

function globalSocketSend(obj) {
  if (websock.readyState === websock.OPEN) {
    websock.send(JSON.stringify(obj))
  } else if (websock.readyState === websock.CONNECTING) {
    setTimeout(function () {
      websock.send(JSON.stringify(obj))
    }, 1000)
  } else {
    setTimeout(function () {
      websock.send(JSON.stringify(obj))
    }, 1000)
  }
}
/*
function globalSocketReceive(callback) {
  globalCallback = callback
  setTimeout(function () {
    globalSocketReceive(callback)
  }, 1000)
}*/

export {
  globalSocketOpen,
  globalSocketClose,
  globalSetIntervalClose,
  globalSocketSend,
  msgList
  //globalSocketReceive
}