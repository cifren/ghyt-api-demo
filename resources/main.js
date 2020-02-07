import Vue from 'vue'
import router from './routes'
import App from './views/App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import VueRouter from 'vue-router';

Vue
  .use(VueRouter)
  .use(Buefy);

new Vue({
  router,
  el: '#app',
  render: h => h(App),
});
