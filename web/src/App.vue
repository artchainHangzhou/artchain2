<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
html,body{
  width: 100%;
  height: 100%;
}
body{
  margin: 0;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100%;
  a{
    text-decoration: none;
    /*color: #fff;*/
  }
  input:active{
    border:0;
  }
  th {
      text-align:inherit;
  }
  table {
    border-collapse:collapse;
    border-spacing:0;
  }
  ul {
    list-style-type: none;
    padding: 0;
    margin:0;
    -webkit-margin-before: 0;
    -webkit-margin-after: 0;
    -webkit-margin-start: 0px;
    -webkit-margin-end: 0px;
    -webkit-padding-start: 0;
  }

  li {
    display: inline-block;
    
  }
  .index-header{
    li{
      margin: 12px 10px;
    }
  }
}
</style>
