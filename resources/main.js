import Vue from 'vue'
import router from './routes'
import App from './views/App.vue'
import 'bulma/css/bulma.css'
import VueRouter from 'vue-router';

Vue.use(VueRouter);

new Vue({
  router,
  el: '#app',
  render: h => h(App),
});
