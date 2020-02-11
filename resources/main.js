import Vue from 'vue'
import router from './routes'
import App from './views/App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import VueRouter from 'vue-router';

import { library } from '@fortawesome/fontawesome-svg-core';
// internal icons
import { faEdit, faTrash } from "@fortawesome/free-solid-svg-icons";
library.add(faEdit, faTrash);
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

Vue.component('vue-fontawesome', FontAwesomeIcon);

Vue
  .use(VueRouter)
  .use(Buefy,{
    defaultIconComponent: 'vue-fontawesome',
    defaultIconPack: 'fa',
  });

new Vue({
  router,
  el: '#app',
  render: h => h(App),
});
