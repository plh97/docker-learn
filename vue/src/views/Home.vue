<template>
  <div id="app">
    <div>
      <label for="name">
        name
        <input class="content" ref="input" @keyup.enter="submitName" type="text" v-model="name">
      </label>
      <button @click="submitName">add</button>
    </div>
    <div>
      <ul>
        <li v-for="personName of names" :key="personName['.key']">
          <div v-if="personName.edit" class="item">
            <!-- 可修改 -->
            Name: <input autofocus :ref="`input-${personName['.key']}`" class="content" type="text" @keyup.enter="save(personName)" v-model="personName.name">
            <button @click="save(personName)">保存</button>
            <button @click="cancel(personName['.key'])">取消</button>
            <!-- 可修改 -->
          </div>
          <div v-else class="item">
            <!-- 不可修改 -->
            Name: 
            <span class="content">{{personName.name}}</span>
            <button @click="edit(personName['.key'])">修改</button>
            <button @click="clear(personName['.key'])">清空</button>
            <button @click="remove(personName['.key'])">删除</button>
            <!-- 不可修改 -->
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import {
  // db,
  namesRef
} from '../utils/firebase';
export default {
  name: 'Home',
  data(){
    return {
      name: '',
    }
  },
  methods: {
    submitName(){
      if(!this.name)return;
      namesRef.push({name: this.name, edit: false})
      this.name = '';
    },
    remove(key){
      namesRef.child(key).remove();
    },
    clear(key){
      namesRef.child(key).update({key,name:' '});
    },
    edit(key){
      // this.$nextTick(()=>{
        setTimeout(() => {
          this.$refs[`input-${key}`][0].focus()
        }, 100);
      // })
      namesRef.child(key).update({edit: true});
    },
    cancel(key){
      namesRef.child(key).update({edit: false});
    },
    save(person){
      namesRef.child(person['.key']).update({name:person.name,edit:false});
    }
  },
  firebase:{
    names: namesRef
  }
}
</script>
<style lang="less" scoped>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;

  .item {
    display: inline-flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    height: 30px;
    width: 300px;
    button {
      border-radius: 0;
      border: 2px solid #000;
      color: #000;
    }
    button+button{
      margin-left: 4px;
    }
    .content {
      height: 20px;
      margin: 0 6px 0px 1px;
      width: 100px;
      font-size: 12px;
      text-align: left;
      color: #2c3e50;
      box-sizing: border-box;
      display: inline-flex;
      align-items: center;
      padding-left: 3px;
      border: 1px solid #eee;
    }
  }

}
</style>
