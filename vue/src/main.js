import Vue from 'vue'
import App from './App.vue'
import './utils/firebase.js'
import VueFire from 'vuefire'
import router from './routers';

Vue.use(VueFire)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
