import VueRouter from 'vue-router';
import mainComponent from "../components/mainComponent";
import modalityComponent from "../components/modalityComponent";
import chartComponent from "../components/chartComponent";

/* eslint-disable no-console */
// eslint-disable-next-line no-unused-vars
const router =  new VueRouter({
  routes: [
    {
      path: '/',
      name: 'mainComponent',
      component: mainComponent,
      children: [
        {
          path: 'modality',
          name: 'Модальности',
          component: modalityComponent
        },
        {
          path: 'charts',
          name: 'Графики',
          component: chartComponent
        }
      ]
    },
  ],
});

export default router;
