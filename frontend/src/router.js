import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import Finished from './views/Finished.vue';
import Convert from './views/Convert.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  linkExactActiveClass: 'active',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/convert',
      name: 'convert',
      component: Convert
    },
    {
      path: '/finished',
      name: 'finished',
      component: Finished
    }
  ]
});
