import VueRouter from 'vue-router';
import HomePage from '../views/Home'
import LoginPage from '../views/Login'
import vue from 'vue'

vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    // dynamic segments start with a colon
    { path: '/', component: HomePage },
    { path: '/login', component: LoginPage },
  ],
  mode: 'history'
})

export default router;