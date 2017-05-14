<template>
  <div class="headerindex">
  <div class="zhezhao" v-show="zhezhao"></div>
    <header class="index-header">
    <span class="titlelist">ArtChain艺术品版权交易</span>
      <ul v-show="!logintue">
        <li @click="registerdialog()">注册</li>
        <li @click="logindialog()">登录</li>
        <li @click="entrylogindialog()">机构登录</li>
      </ul>
       <ul class="right" v-show="logintue">
          <li class="tuichu" @click="goout()">退出</li>
          <li v-if="userType=='org'"><router-link to="/entery">用户中心</router-link></li>
          <li v-if="userType=='user'"><router-link to="/user">用户中心</router-link></li>
          <li class="name">
            <p>你好:{{userName}}</p>
          </li>
        </ul>
    </header>
    <div class="index-content">
      <div class="index-con1"><span>在线IP许可权</span></div>
      <div class="table1">
        <table>
          <thead>
            <tr>
              <th>原作名称</th>
              <th>原作者</th>
              <th>备案信息</th>
              <th>商品子代码</th>
              <th>照片</th>
              <th>价格</th>
              <th v-show="userType=='user'">购买</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item,index) in pageIpList">
              <td>{{item.ipName}}</td>
              <td>{{item.author}}</td>
              <td><a download :href="apiurl+'download/file/'+item.proposalUrl">{{item.proposalUrl}}</a></td>
              <td>{{item.ipId}}</td>
              <td @click='showimgss(item.pictureUrl)'><img class="showimg" :src="apiurl+'download/file/'+item.pictureUrl"></td>
              <td>{{item.Price}}</td>
              <td v-show="userType=='user'" @click="buydia(AllIpList[index])" class="buycl">购买</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="pageclass">
        <el-pagination
          small
          layout="prev, pager, next"
          :page-size="5"
          :total="ipListLength"
          @current-change="getpagelist"
          >
        </el-pagination>
      </div>
    </div>
    <div v-show="regrist" class="dialog">
     <img class="img1" @click="close()" src="../assets/img/gb.png">
      <div class="tromcontent">
        <div class="fromtitle"><span>注册</span></div>
        <div class="conteninput">
          <input type="text" name="" v-model="signInParam.userId" placeholder="请输入用户名">
          <input type="text" name="" v-model="signInParam.email" placeholder="请输入邮箱">
          <input type="password" name="" placeholder="请输入密码">
          <button @click="signIn()">注册</button>
        </div>
      </div>
    </div>
    <div v-show="loging" class="dialog">
    <img class="img1" @click="close()" src="../assets/img/gb.png">
      <div class="tromcontent">
        <div class="fromtitle"><span>登录</span></div>
        <div class="conteninput">
          <input type="text" name="" v-model="userId" placeholder="请输入用户名">
          <input type="password" name="" placeholder="请输入密码">
          <button @click="login()">登录</button>
          <p>还没有账户?<span class="clickclass" @click="godignIn()">点我注册</span></p>
        </div>
      </div>
    </div>
    <div v-show="entryloging" class="dialog">
    <img class="img1" @click="close()" src="../assets/img/gb.png">
      <div class="tromcontent">
        <div class="fromtitle"><span>机构登录</span></div>
        <div class="conteninput">
          <input type="text" name="" v-model="orgId" placeholder="请输入用户名">
          <input type="password" name="" placeholder="请输入密码">
          <button @click="enterlogin()">登录</button>
          <!-- <p>还没有账户?<span class="clickclass">点我注册</span></p> -->
        </div>
      </div>
    </div>
    <div v-show="buy" class="dialog">
    <img class="img1" @click="close()" src="../assets/img/gb.png">
      <div class="tromcontent">
        <div class="fromtitle"><span>购买</span></div>
        <div class="conteninput">
          <div class="cnane">
            <span class="span1">原作者</span>
            <span class="span2">{{buyparam.author}}</span>
          </div>
          <div class="cnane">
            <span class="span1">备案信息</span>
            <span class="span2">{{buyparam.proposalUrl}}</span>
          </div>
          <div class="cnane spanimg">
            <span class="span1">照片</span>
            <span class="span2"><img class="" :src="apiurl+'download/file/'+buyparam.pictureUrl"></span>
          </div>
          <div class="cnane">
            <span class="span1">商品子代码</span>
            <span class="span2">{{buyparam.subId}}</span>
          </div>
          <div class="cnane">
            <span class="span1">价格</span>
            <span class="span2">{{buyparam.Price}}</span>
          </div>
          <button @click="gogetIp()">确认购买</button>
        </div>
      </div>
    </div>
    <div v-show="showimg" class="dialog dialogimg">
    <img class="img1" @click="close()" src="../assets/img/gb.png">
      <div class="tromcontent">
      <!-- <div class="fromtitle"><span>机构登录</span></div> -->
        <img class="yuimg" :src="yuanimgurl">
      </div>
      </div>
    </div>
  </div>
</template>

<script>
// import axios from 'axios'
import api from '../api/config'
export default {
  // name: 'hello',
  data () {
    return {
      apiurl:'',
      userType:'',
      userName:'',
      page:1,
      pageSize:5,
      logintue:false,
      entryloging:false,
      regrist:false,
      loging:false,
      zhezhao:false,
      buy:false,
      showimg:false,
      ipListLength:0,
      yuanimgurl:'',
      userId:'',
      orgId:'',
      AllIpList:[],
      pageIpList:[],
      buyparam:{
        author:'',
        proposalUrl:'',
        pictureUrl:'',
        Price:'',
        subId:'',
      },
      gobuy:{
        ipId:'',
        userId:'',
      },
      signInParam:{
        userId:'',
        email:'',
      },
    }
  },
  mounted: function() {
    this.$nextTick(() => {
      this.IpList();
      this.apiurl=api.curl;
      if(sessionStorage.userId){
        this.userName=sessionStorage.userId
        this.logintue=true;
        this.userType=sessionStorage.userType;
      }
    })
  },
  methods:{
    registerdialog(){
      this.close();
      this.regrist=true;
      this.zhezhao=true;
    },
    godignIn(){
      this.registerdialog();
    },
    logindialog(){
      this.close();
      this.zhezhao=true;
      this.loging=true;
    },
    entrylogindialog(){
      this.close();
      this.zhezhao=true;
      this.entryloging=true;
    },
    close(){
      this.entryloging=false;
      this.regrist=false;
      this.zhezhao=false;
      this.buy=false;
      this.loging=false;
      this.showimg=false;
    },
    buydia(content){
      this.close();
      this.zhezhao=true;
      this.buy=true;
      this.buyparam.author=content.author;
      this.buyparam.proposalUrl=content.proposalUrl;
      this.buyparam.pictureUrl=content.pictureUrl;
      this.buyparam.Price=content.Price;
      this.buyparam.subId=content.subId;
      this.gobuy.ipId=content.ipId;
    },
    showimgss(imgurl){
      console.log(imgurl);
      this.zhezhao=true;
      this.showimg=true;
      this.yuanimgurl=api.curl+'download/file/'+imgurl;
    },
    getpagelist(response){
        var start=(response-1)*this.pageSize;
        var end=response*this.pageSize;
        this.pageIpList=this.AllIpList.slice(start,end);
    },
    gogetIp(){
      var _this=this;
      this.gobuy.userId=sessionStorage.userId;
      $.ajax({
        type:"POST",
        url:api.curl+'buy',
        dataType:'json',
        data:_this.gobuy,
        success:function(data){
          if(data.code==0){
            _this.$notify.success({
              title: '成功',
               message: '提交成功'
            });
            _this.zhezhao=false;
            _this.buy=false;
            _this.IpList();
          }else{
            _this.$notify.error({
              title: '失败',
               message: '购买失败'
            });
          }
        },
        error:function(){
          console.log(1111)
        }
      });
    },
    goout(){
      sessionStorage.clear();
      this.userName='';
      this.userType='';
      this.logintue=false;
    },
    login(){
      var userid=this.userId;
      var _this=this;
      // var url=api.curl+api.Login;
     $.ajax({
      type:"POST",
      url:api.curl+'queryUser',
      dataType:'json',
      data:'userId='+userid,
      success:function(data){
        if(data.code==0){
          var response=JSON.parse(data.data);
          _this.userName=response.userName;
          _this.logintue=true;
          _this.userType='user';
          sessionStorage.setItem('userName',response.userName);
          sessionStorage.setItem('coin',response.coin);
          sessionStorage.setItem('userType','user');
          sessionStorage.setItem('userId',response.userId);
          _this.$router.push({
            path: '/user'
          });
        }else{
          _this.$notify.error({
            title: '错误',
            message: '登录失败'
          });
        }
      },
      error:function(){
        console.log(1111)
      }
     });
    },
    enterlogin(){
      var userid=this.orgId;
      var _this=this;
     $.ajax({
      type:"POST",
      url:api.curl+'queryOrg',
      dataType:'json',
      data:'orgId='+userid,
      success:function(data){
        if(data.code==0){
          var response=JSON.parse(data.data);
          _this.userName=response.userName;
          _this.logintue=true;
          _this.userType='org';
          sessionStorage.setItem('userName',response.orgName);
          sessionStorage.setItem('userType','org');
          sessionStorage.setItem('userId',response.orgId);
          sessionStorage.setItem('coin',response.coin);
         _this.$router.push({
          path: '/entery'
        });
        }else{
          _this.$notify.error({
            title: '错误',
            message: '登录失败'
          });
        }
      },
      error:function(){
        console.log(1111)
      }
     });
    },
    signIn(){
      var _this=this;
     $.ajax({
      type:"POST",
      url:api.curl+'signIn',
      dataType:'json',
      data:_this.signInParam,
      success:function(data){
        if(data.code==0){
          _this.$notify.success({
            title: '成功',
            message: '注册成功'
          });
          _this.userId=_this.signInParam.userId;
          _this.login();
        }else{
          _this.$notify.error({
            title: '错误',
            message: '注册失败'
          });
        }
        
      },
      error:function(){
        console.log(1111)
      }
     });
    },
    IpList(){
      var _this=this;
     $.ajax({
      type:"POST",
      url:api.curl+'queryIPList',
      dataType:'json',
      data:'',
      success:function(data){
        // var response=JSON.parse(data.data);
        _this.AllIpList=JSON.parse(data.data);
        _this.pageIpList=_this.AllIpList.slice(0,5);
        _this.ipListLength=_this.AllIpList.length;
      },
      error:function(){
        console.log(1111)
      }
     });
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
.tuichu a:hover{
  color: #42b983;
}
.headerindex{
  width: 100%;
  height: 100%;
  .zhezhao{
    width: 100%;
    height: 100%;
    z-index: 200;
    position: absolute;
    background: rgba(0,0,0,0.6);
  }
  .dialog{
    width: 400px;
    position: fixed;
    z-index: 2001;
    background: #fff;
    height: 400px;
    top:300px;
    left: 50%;
    margin: -200px;
    border-radius: 6px;
    .img1{
      position: absolute;
      top:10px;
      right: 10px;
    }
    .tromcontent{
      .fromtitle{
        width: 320px;
        margin: 25px auto 0;
        height: 45px;
        line-height: 45px;
        border-bottom: 2px solid #4778c7;
        color: #4778c7;
        font-size: 16px;
        text-align: left;
        font-weight: 500;
      }
      .conteninput{
        width: 320px;
        margin: 0 auto;
        padding-top: 15px;
        input{
          width: 320px;
          height: 35px;
          line-height: 35px;
          border:1px solid #eee;
          margin-top: 10px;
          border-radius: 4px;
          padding-left: 15px;
        }
        .cnane{
          width: 320px;
          height: 35px;
          line-height: 35px;
          border-bottom:1px solid #eee;
          span{
            display: inline-block;
            text-align: left;
          }
          .span1{
            width: 95px;
            color: #ccc;
          }
          .span2{
            width: 220px;
            color: #000;
            img{
              width: 200px;
              height: 70px;
            }
          }
        }
        .spanimg{
          height: 80px;
        }
        button{
          width: 200px;
          height: 35px;
          margin: 35px auto 0;
          color: #fff;
          font-size: 14px;
          background: #8faddd;
          border-radius: 4px;
          line-height: 35px;
          border:1px solid #8faddd;
        }
        p{
          color: #ccc;
          text-align: center;
          .clickclass{
            color: #8faddd;
            cursor: pointer;
          }
        }
      }
    } 
  }
  .dialogimg{
    width: 600px;
    height: 600px;
    .yuimg{
      margin-top: 40px;
      width: 600px;
      height: 550px;
    }
  }
  .index-header{
    width: 100%;
    height: 60px;
    background: #4778c7;
    .titlelist{
      float: left;
      line-height: 60px;
      color: #fff;
      padding-left: 30px;
    }
    ul{
      width: 1200px;
      margin: 0 auto!important;
      height: 60px;
      padding-top: 12px;
      li{
        float: right;
        width: 100px;
        height: 35px;
        line-height: 35px;
        border:1px solid #8faddd;
        border-radius: 3px;
        color: #fff;
        font-size: 14px;
        cursor: pointer;
        a{
          display: inline-block;
          height: 100%;
          width: 100%;
        }
      }
      li:hover{
        background: #fff;
        color: #4778c7;
      }
      .name{
        border:0;
        p{
          margin: 0;
          padding: 0;
          line-height: 35px;
        }
      }
      .name:hover{
        background: #4778c7;
        color: #fff;
      }
    }
  }
  .index-content{
    width: 1200px;
    height: 800px;
    margin:0 auto;
    .index-con1{
      width: 1200px;
      height: 40px;
      line-height: 40px;
      background: #fff;
      margin-top: 20px;
      color: #4778c7;
      text-align: left;
      span{
        padding-left: 20px;
      }
    }
    .table1{
      margin-top: 10px;
      a{
        color: #333!important;
      }
      .buycl{
        color: #4778c7;
      }
      table{
        display: table;
        thead{
          tr{
            background: #f5f8fd;
          }
        }
        tbody{
          tr{
            background: #fff;
          }
        }
        th,td{
         width: 200px;
         height: 35px;
         overflow: hidden;
         line-height: 35px;
         color: #333;
         font-size: 14px;
         font-weight: 500;
         border: 0;
         .alink{
          display: inline-block;
         }
         .showimg{
          width: 200px;
          height: 80px;
         }
        }
      }
    }
    .btn-prev{
       position: relative;
        top:20px;
    }
    .pageclass{
      width: 100%;
      background: #fff;
      .el-pager{
        li{
          margin:0!important;
        }
      }
       
    }
  }
}

</style>
