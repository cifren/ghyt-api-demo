import Vue from 'vue'
import router from './routes'
import App from './views/App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import './assets/ghyt-admin.css'
import VueRouter from 'vue-router';
import axios from 'axios'
import VueAxios from 'vue-axios'

import { library } from '@fortawesome/fontawesome-svg-core';
// internal icons
import { faEdit, faTrash, faPlus, faBriefcase, faAngleDoubleRight, faCaretDown, faInfo,
  faSave } from "@fortawesome/free-solid-svg-icons";
library.add(faEdit, faTrash, faPlus, faBriefcase, faAngleDoubleRight, faCaretDown, faInfo, faSave);
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

Vue.component('vue-fontawesome', FontAwesomeIcon);

Vue
  .use(VueRouter)
  .use(VueAxios, axios)
  .use(Buefy,{
    defaultIconComponent: 'vue-fontawesome',
    defaultIconPack: 'fa',
  });

new Vue({
  router,
  el: '#app',
  render: h => h(App),
});
