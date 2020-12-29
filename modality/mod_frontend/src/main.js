import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import router from './routers/router'
import MuseUI from 'muse-ui'
import 'muse-ui/dist/muse-ui.css'
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome"
import {library} from '@fortawesome/fontawesome-svg-core'
import {faTimesCircle, faCheckCircle} from '@fortawesome/free-solid-svg-icons'

Vue.use(MuseUI)
Vue.use(VueRouter)
library.add(faCheckCircle, faTimesCircle)
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
