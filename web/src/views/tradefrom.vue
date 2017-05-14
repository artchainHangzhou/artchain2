<template>
  <div class="tradefrom">
    <header class="index-header">
      <div class="user-con">
        <ul class="left">
          <li v-if="userType=='user'"><router-link to="/user">我的账户</router-link></li>
          <li class="tradeIP"><router-link to="/tradefrom">提交IP</router-link></li>
          <li><router-link to="/index">购买</router-link></li>
          <li v-if="userType=='org'"><router-link to="/entery">机构管理</router-link></li>
          <li><router-link to="/banquan">版权查询</router-link></li>
        </ul>
        <ul class="right">
          <li class="tuichu" @click="goout()">退出</li>
          <li class="name">
            <p>你好:{{userName}}</p>
          </li>
        </ul>
      </div>
    </header>
    <div class="tromcontent">
      <div class="fromtitle"><span>提交IP许可证</span></div>
      <div class="conteninput">
        <div class="group"> 
          <span>原作名称</span><input type="text" name="" v-model="tradefrom.ipName">
        </div>
        <div class="group"> 
          <span>作者</span><input type="text" name="" v-model="tradefrom.author">
        </div>
        <!-- <input type="text" name="" v-model="tradefrom.author" placeholder="作者"> -->
        <div class="group"> 
          <span>物理信息</span><input type="text" name="" v-model="tradefrom.description">
        </div>
        <div class="group"> 
          <span>总量</span><input type="text" name="" v-model="tradefrom.total">
        </div>
        <div class="group"> 
          <span>单价</span><input type="text" name="" v-model="tradefrom.price">
        </div>
         <el-upload class='wenjian'
            :action='imgurl+"upload"'
            type="drag"
            :thumbnail-mode="false"
            name="uploadfile"
            :on-preview="handlePreview"
            :on-remove="handleRemove1"
            :on-success="uploadsuccess"
            :default-file-list="fileList"
          >
            <i class="el-icon-upload"></i>
            <div class="el-dragger__text">将律师意见书拖到此处，或<em>点击上传</em></div>
        </el-upload>
        <el-upload class="imgs"
            :action='imgurl+"upload"'
            type="drag"
            :thumbnail-mode="false"
            name="uploadfile"
            :on-preview="handlePreview"
            :on-remove="handleRemove2"
            :on-success="imgupload"
            :default-file-list="fileListss"
          >
            <i class="el-icon-upload"></i>
            <div class="el-dragger__text">将文件作品拖到此处，或<em>点击上传</em></div>
        </el-upload>
        <!-- <input type="text" name="" v-model="tradefrom.ipName" placeholder="作品文件"> -->
        <!-- <input type="text" name="" v-model="tradefrom.description" placeholder="物理信息">
        <input type="text" name="" v-model="tradefrom.total" placeholder="总量">
        <input type="text" name="" v-model="tradefrom.Price" placeholder="单价"> -->
        <button @click="tradeIPP()">提交审核</button>
      </div>
    </div>
  </div>
</template>
<script>
import api from '../api/config'
  export default {
  // name: 'hello',
    data () {
      return {
        imgurl:'',
        userName:'',
        userType:'',
        wenjian:1,
        fileList: [],
        fileListss:[],
        tradefrom:{
          userId:'',
          ipName:'',
          author:'',
          description:'',
          proposalUrl:'',
          pictureUrl:'',
          price:'',
          total:'',
        },
      }
    },
    mounted: function() {
      this.$nextTick(() => {
        this.imgurl=api.curl;
        this.userName=sessionStorage.userName;
        this.userType=sessionStorage.userType;
      })
    },
    methods:{
      handleRemove1(file, fileList) {
        $('.wenjian .el-dragger').show();
      },
       handleRemove2(file, fileList) {
        $('.imgs .el-dragger').show();
      },
      handlePreview(file) {
        console.log(file);
      },
      uploadsuccess(response){
        console.log(response.data);
        this.tradefrom.proposalUrl=response.data;
        $('.wenjian .el-dragger').hide();
      },
      imgupload(response){
        this.tradefrom.pictureUrl=response.data;
        $('.imgs .el-dragger').hide();
      },
      goout(){
        sessionStorage.clear();
        this.$router.push({
          path: '/index'
        });
      },
      tradeIPP(){
        var _this=this;
        this.tradefrom.userId=sessionStorage.userId;
        var tdata=this.tradefrom;
        $.ajax({
          type:"POST",
          url:api.curl+'apply',
          dataType:'json',
          data:tdata,
          success:function(data){
            if(data.code==0){
              _this.$notify.success({
                title: '成功',
                 message: '提交成功'
              });
              _this.tradefrom.ipName='';
              _this.tradefrom.author='';
              _this.tradefrom.description='';
              _this.tradefrom.proposalUrl='';
              _this.tradefrom.pictureUrl='';
              _this.tradefrom.price='';
              _this.tradefrom.total='';
              $('.wenjian .el-dragger').show();
              $('.imgs .el-dragger').show();
              $('.el-upload__files').hide();
            }else{
              _this.$notify.error({
                  title: '错误',
                  message: '审核失败'
                });
              }
          },
          error:function(){
            console.log(1111)
          }
        });
      }
    },
  }
</script>
<style>
    .el-upload__file__name{
      color: #48576a!important;
    }
  .tradefrom{
    a{
      color: #fff;
    }
    .index-header{
      width: 100%;
      height: 60px;
      background: #4778c7;
      .user-con{
        width: 1200px;
        margin: 0 auto;
      }
      .left{
        float: left;
        li{
          float: left;
        }
        .tradeIP{
          background: #fff;
          a{
            color: #4778c7!important;
          }
        }
      }
      .right{
        float: right;
        li{
          float: right;
        }
        .name{
          border:0;
          p{
            margin: 0;
            padding: 0;
            line-height: 35px;
          }
        }
      }
      ul{
        width: 600px;
        margin: 0 auto;
        height: 60px;
        /*padding-top: 20px;*/
        li{
          float: right;
          width: 100px;
          height: 35px;
          margin-top: 12px;
          line-height: 35px;
          border:1px solid #8faddd;
          border-radius: 3px;
          color: #fff;
          font-size: 14px;
          cursor: pointer;
        }
        .tuichu:hover{
          background: #fff;
          color: #4778c7!important; 
        }
      }
    }
    .tromcontent{
      width: 1200px;
      height: 640px;
      margin: 20px auto;
      border-radius: 5px;
      background: #fff;
      .fromtitle{
        width: 920px;
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
        .group{
          width: 320px;
          height: 35px;
          line-height: 35px;
          /*border:1px solid #eee;*/
          margin-top: 10px;
          border-radius: 4px;
          padding-left: 15px;
          text-align: left;
          color: #aaa;
          position: relative;
          span{
            display: inline-block;
            width: 80px;
            text-align: right;
            padding-right: 10px;
          }
          input{
            border:1px solid #eee;
            width: 210px;
            height: 35px;
            padding-left: 10px;
            line-height: 35px;
            /*border:0;*/
          }
        }
        .el-upload{
          width: 335px;
        }
        .el-upload__files,.el-dragger{
          width: 335px;
          height: 70px;
          /*line-height: 35px;*/
          border:1px solid #eee;
          margin-top: 10px!important;
          border-radius: 4px;
          padding-left: 15px;
          text-align: left;
          color: #c6c6c6;
          /*position: relative;*/
          .el-icon-upload{
            margin:0;
          }
        }
        .el-upload__file{
          width: 80%;
        }
        .el-icon-upload{
          font-size: 30px;
          padding-left: 100px;
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
      }
    }
  }
  
</style>