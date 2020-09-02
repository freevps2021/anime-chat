<template>
  <div id="app">
    <el-container>
      <el-aside width="200px">
        <el-row class="me user">
          <el-col :span="6">
            <el-avatar shape="circle" :size="meAvatarsize" :src="squareUrl"></el-avatar>
          </el-col>
          <el-col :span="18">
            <div>
              <div>
                <span class="user-name">伊藤诚</span>
              </div>
              <div class="user-signature color-dark-light">诚信交友😍</div>
            </div>
          </el-col>
        </el-row>
        <div class="line"></div>
        <el-row class="user">
          <el-col :span="6">
            <el-avatar shape="circle" :size="userAvatarSize" :src="squareUrl"></el-avatar>
          </el-col>
          <el-col :span="18">
            <div>
              <div>
                <span class="user-name">伊藤诚</span>
                <span class="user-status">●</span>
              </div>
              <div class="user-signature color-dark-light">诚信交友😍</div>
            </div>
          </el-col>
        </el-row>
      </el-aside>
      <el-main>
        <div class="msgBox">
          <el-row type="flex" justify="left" class="msg" v-for="m in msgList" :key="m.lable" :value="m.value">
            <el-col :span="1">
              <el-avatar shape="circle" :size="msgBoxAvatarsize" :src="m.avatar"></el-avatar>
            </el-col>
            <el-col>
              <div class="name">{{ m.name }}</div>
              <div class="content">{{ m.content }}</div>
            </el-col>
          </el-row>
        </div>
        <el-input
          class="msgInput"
          type="textarea"
          v-model="msg"
          resize="none"
          style="border-radius:none;"
          v-on:keyup.enter.native="sendMsg()"
        ></el-input>
      </el-main>
    </el-container>
  </div>
</template>

<script>
6;
export default {
  name: "app",
  data() {
    return {
      msg: "",
      squareUrl: require("./assets/VVosbJSBNBn6eesWgk0YuT3nPP4hUndKYat7aJB6CmqO5dHFS1blruaKeT0tQp4n.png"),
      userAvatarSize: "medium",
      meAvatarsize: "large",
      msgBoxAvatarsize:"small",
      msgList: [],
      websock: null,
    };
  },
  created() {
    this.initWebSocket();
  },
  destroyed() {
    this.websock.close(); //离开路由之后断开websocket连接
  },
  methods: {
    sendMsg: function () {
      this.websocketsend(this.msg);
      this.msg = "";
    },
    initWebSocket() {
      //初始化weosocket
      const wsUri = "ws://192.168.226.129:8081";
      this.websock = new WebSocket(wsUri);
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onopen = this.websocketonopen;
      this.websock.onerror = this.websocketonerror;
      this.websock.onclose = this.websocketclose;
    },
    websocketonopen() {
      //连接建立之后执行send方法发送数据
      // let actions = {"test":"12345"};
      // this.websocketsend(JSON.stringify(actions));
    },
    websocketonerror() {
      //连接建立失败重连
      // this.initWebSocket();
    },
    websocketonmessage(e) {
      //数据接收
      const redata = JSON.parse(e.data);
      console.log(redata);
      this.msgList.push(redata);
    },
    websocketsend(data) {
      //数据发送
      this.websock.send(data);
    },
    websocketclose(e) {
      //关闭
      console.log("断开连接", e);
    },
  },
};
</script>

<style>
#app {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
    "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 100px;
  font-size: 14px;
}
body {
  background-image: url("./assets/bg.jpg");
  background-size: 100%;
}

.el-aside {
  border-top-left-radius: 3px;
  border-bottom-left-radius: 3px;
  color: #333;
}

.me,
.el-aside {
  background-color: #2e3238 !important;
}

.el-main {
  padding: 0 !important;
  border-top-right-radius: 3px;
  border-bottom-right-radius: 3px;
  background-color: #eeeeee;
  color: #333;
}

.el-container {
  margin-bottom: 40px;
  width: 1000px;
  margin: 0 auto;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  /* line-height: 260px; */
}

.el-container:nth-child(7) .el-aside {
  /* line-height: 320px; */
}

.user {
  padding: 3px;
  border-radius: 1px;
  margin: 5px;
  background-color: #43474c;
}

.user-name {
  font-weight: bold;
}

.user-signature {
  line-height: 20px;
  font-size: 12px;
  color: #909399;
}

.user-name,
.user-signature {
  color: #eee;
}

.user-status {
  font-size: 13px;
  float: right;
  color: lightgreen;
}

.msgBox {
  height: 400px;
  padding: 15px;
  margin-bottom: 5px;
}

.line {
  border-bottom: #666 solid 1px;
  width: 100%;
}

textarea {
  border-radius: 0 !important;
  border-left: none;
  border-right: none;
  border-bottom: none;
}

.msgBox .content {
  background-color: #fff;
  padding: 5px 8px;
  display: inline-block;
  border-radius: 4px;
  position: relative;
  font-size: 12px;
  margin-top:3px;
}
.msgBox .content::after {
  content: "";
  border: 8px solid #ffffff00;
  border-right: 8px solid #fff;
  position: absolute;
  top: 6px;
  left: -16px;
}
.msgBox .name{
  color:#888;
  font-size: 12px;
}
</style>
