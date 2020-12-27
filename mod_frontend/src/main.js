import Vue from 'vue'
import VueRouter from 'vue-router'
import router from './routers/router'
import MuseUI from 'muse-ui'
import 'muse-ui/dist/muse-ui.css'
import App from './App.vue'

Vue.use(MuseUI)
Vue.use(VueRouter)
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
