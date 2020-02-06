import VueRouter from 'vue-router';
import Home from './views/Home.vue'
import Jobs from './views/Jobs.vue'
import Logs from './views/Logs.vue'

let routes = [
  {
    path: '/',
    component: Home
  },  {
    path: '/jobs',
    component: Jobs
  },  {
    path: '/logs',
    component: Logs
  }
];

export default new VueRouter({
  routes
});
