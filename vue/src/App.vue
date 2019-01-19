<template>
  <div id="app">
    <div>
      <label for="name">
        name
        <input type="text" v-model="name">
      </label>
      <button @click="submitName()">add</button>
    </div>
  </div>
</template>

<script>
import {
  db,
  namesRef
} from './firebase';
export default {
  name: 'app',
  data(){
    return {
      name: 'pengliheng',
    }
  },
  methods: {
    submitName(){
      namesRef.push({name: this.name, edit: false})
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
