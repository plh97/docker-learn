<template>

  <form action="" v-if="mode==='login'">
    <div id="firebaseui-auth-container">

    </div>
    <div class="item">
      name: <input type="text">
    </div>
    <div class="item">
      password: <input type="password">
    </div>
    <div class="item">
      <input type="submit" value="登录">
      <input type="register" value="注册">
    </div>
  </form>
  <form action="" v-else-if="mode==='register'">
    <div class="item">
      name: <input type="text">
    </div>
    <div class="item">
      email: <input type="email">
    </div>
    <div class="item">
      password: <input type="password">
    </div>
    <div class="item">
      repeat password: <input type="password">
    </div>
    <div class="item">
      <input type="submit" value="提交">
      <input type="register" value="登录">
    </div>
  </form>
</template>

<style lang="less" scoped>
  @import 'https://cdn.firebase.com/libs/firebaseui/3.5.2/firebaseui.css';

</style>


<script lang="ts" >
  import Vue from 'vue'
  import firebase from 'firebase'
  // or for ES6 imports.
  import * as firebaseui from 'firebaseui'
  import uiConfig from '../utils/auth'
  export default Vue.extend({
    data(){
      return {
        mode: 'login',
      }
    },
    mounted(){
      // uiConfig
      // Initialize the FirebaseUI Widget using Firebase.
      var ui = new firebaseui.auth.AuthUI(firebase.auth());
      // The start method will wait until the DOM is loaded.
      ui.start('#firebaseui-auth-container', uiConfig);
    },
    methods: {
      //
    }
  })
</script>
