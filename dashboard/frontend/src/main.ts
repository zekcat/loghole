import Vue from 'vue';
import Buefy from 'buefy';
import axios from 'axios';
import VueAxios from 'vue-axios';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/scss/app.scss';

Vue.use(VueAxios, axios);
Vue.use(Buefy);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
