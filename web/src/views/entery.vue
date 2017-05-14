<template>
  <div class="usercontwe">
    <header class="index-header">
    <div class="user-con">
      <ul class="left">
        <!-- <li><router-link to="/user">我的账户</router-link></li> -->
        <!-- <li ><router-link to="/tradefrom">提交IP</router-link></li> -->
        <li><router-link to="/index">首页</router-link></li>
        <li class="mybutton"><router-link to="/entery">机构管理</router-link></li>
      </ul>
      <ul class="right">
        <li class="tuichu"  @click="goout()">退出</li>
        <li class="name">
          <p>你好:{{userName}}</p>
        </li>
      </ul>
    </div>
    </header>
    <div class="user-ssec"><span class="mymis">交易明细</span><span class="mycoin">机构余额: <b class="coincolor">{{coin}}</b></span></div>
    <div class="table1">
      <table>
        <thead>
          <tr>
            <th>原作名称</th>
            <th>商品子代码</th>
            <th>成交时间</th>
            <th>价格</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in caozuoList">
            <td>{{item.ipName}}</td>
            <td>{{item.ipId}}</td>
            <td>{{item.updateTime}}</td>
            <td>{{item.price}}</td>
          </tr>
        </tbody>
      </table>
      <el-pagination
        small
        layout="prev, pager, next"
        :page-size="10"
        :total="pagecaozuoListLeng"
        @current-change="getpage3list"
        >
      </el-pagination>
    </div>
     <div class="user-ssec"><span class="mymis">用户查询</span></div>
    <div class="table2">
      <table>
        <thead>
          <tr>
            <th class="fristthtd">用户名</th>
            <th class="twothtd">账户余额</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in userList">
            <td class="fristthtd">{{item.userName}}</td>
            <td class="twothtd">{{item.coin}}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
import api from '../api/config'
  export default {
  // name: 'hello',
    data () {
      return {
        coin:'0.00',
        userName:'',
        userList:[],
        caozuoList:[],
        pageSize3:5,
        page:0,
        pagecaozuoList:[],
        pagecaozuoListLeng:0,
      }
    },
    mounted: function() {
      this.$nextTick(() => {
        this.userName=sessionStorage.userName;
        this.coin=sessionStorage.coin;
        this.UserList();
        this.getmyhistoryList();
      })
    },
    methods:{
      goout(){
        sessionStorage.clear();
        this.$router.push({
          path: '/index'
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
          _this.coin=response.coin;
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
      getpage3list(response){
        var start=(response-1)*this.pageSize3;
        var end=response*this.pageSize3;
        this.pagecaozuoList=this.caozuoList.slice(start,end);
      },
      UserList(){
        var _this=this;
        $.ajax({
          type:"POST",
          url:api.curl+'queryUserList',
          dataType:'json',
          data:'',
          success:function(data){
            _this.userList=JSON.parse(data.data)
          },
          error:function(){
            console.log(1111)
          }
        });
      },
      getmyhistoryList(){
        var _this=this;
        var userid=sessionStorage.userId;
        $.ajax({
          type:"POST",
          url:api.curl+'queryTransaction',
          dataType:'json',
          data:'userId='+userid,
          success:function(data){
            _this.caozuoList=JSON.parse(data.data)
            _this.pagecaozuoList=_this.caozuoList.slice(0,5);
            _this.pagecaozuoListLeng=_this.caozuoList.length;
          },
          error:function(){
            console.log(1111)
          }
        });
      },
    },
    
  }
</script>
<style>
  .usercontwe{
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
      ul{
        width: 600px;
        margin: 0 auto;
        height: 60px;
        /*padding-top: 20px;*/
        li{
          /*float: right;*/
          width: 100px;
          height: 35px;
          margin-top: 12px;
          line-height: 35px;
          border:1px solid #8faddd;
          border-radius: 3px;
          color: #fff;
          font-size: 14px;
          cursor: pointer;
          a{
            display: inline-block;
            width: 100%;
            height: 100%;
          }
        }
        .tuichu:hover{
          background: #fff;
          color: #4778c7!important;
        }
      }
      .left{
        float: left;
        li{
          float: left;
        }
        .mybutton{
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
    }
    .user-ssec{
      width: 1200px;
      height: 40px;
      margin: 20px auto 10px; 
      background: #fff;
      line-height: 40px;
      /*margin-top: 20px;*/
      text-align: left;
      .mymis{
        color: #4778c7;
        padding-left: 20px;
      }
      .mycoin{
        float: right;
        padding-right: 20px;
        .coincolor{
          color: red;
        }
      }
    }
    .table1{
      width: 1200px;
      margin: 0px auto;
      background: #fff;
      table{
        padding: 10px;
        thead{
          tr{
            background: #f5f8fd;
          }
        }
        th,td{
         width: 300px;
         height: 35px;
         line-height: 35px;
         color: #333;
         font-size: 14px;
         font-weight: 500;
         border: 0;
        }
      }
    }
    .table2{
      width: 1200px;
      margin: 0px auto;
      background: #fff;
      table{
        padding: 10px;
        thead{
          tr{
            background: #f5f8fd;
          }
        }
        th,td{
         height: 35px;
         line-height: 35px;
         color: #333;
         font-size: 14px;
         font-weight: 500;
         border: 0;
        }
        .fristthtd{
          width: 300px;
        }
        .twothtd{
          width: 900px;
          text-align: left!important;
        }
      }
    }
  }
</style>
