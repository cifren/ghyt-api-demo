import VueRouter from 'vue-router';
import Home from './views/Home.vue'
import Jobs from './views/job/Jobs.vue'
import JobForm from './views/job/JobForm.vue'
import ActionForm from './views/job/ActionForm.vue'
import ConditionForm from './views/job/ConditionForm.vue'
import Logs from './views/Logs.vue'

let routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/jobs',
    component: Jobs,
  },
  {
    path: '/job',
    component: JobForm,
    children: [
      {
        path: '/action/:id',
        component: ActionForm
      },
      {
        path: '/condition/:id',
        component: ConditionForm
      }
    ]
  },
  {
    path: '/logs',
    component: Logs
  }
];

export default new VueRouter({
  routes
});
