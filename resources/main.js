import Vue from 'vue'
import Buefy from 'buefy'
import VueRouter from 'vue-router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import _ from 'lodash'
import 'buefy/dist/buefy.css'
import moment from 'vue-moment'

import './assets/ghyt-admin.css'
import router from './routes'
import App from './views/App.vue'

import { library } from '@fortawesome/fontawesome-svg-core';
// internal icons
import { faEdit, faTrash, faPlus, faBriefcase, faAngleDoubleRight, faCaretDown, faInfo,
  faSave } from "@fortawesome/free-solid-svg-icons";
library.add(faEdit, faTrash, faPlus, faBriefcase, faAngleDoubleRight, faCaretDown, faInfo, faSave);
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

Vue.component('vue-fontawesome', FontAwesomeIcon);

Vue.prototype._ = _
Vue
  .use(VueRouter)
  .use(VueAxios, axios)
  .use(Buefy,{
    defaultIconComponent: 'vue-fontawesome',
    defaultIconPack: 'fa',
  })
  .use(moment)

new Vue({
  router,
  el: '#app',
  render: h => h(App),
});
