import Vue from 'vue'
import App from './App.vue'
import './firebase.js'
import VueFire from 'vuefire'

Vue.use(VueFire)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  beforeCreate: ()=>{
    console.log('beforeCreate');
  }
}).$mount('#app')
