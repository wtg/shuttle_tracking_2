import Vue from 'vue';
import 'leaflet/dist/leaflet.css';
import 'typeface-open-sans';
import Buefy from 'buefy';
import 'buefy/dist/buefy.css';

// @ts-ignore
import VueAnalytics from 'vue-analytics';

// only include the specific fontawesome icons that we use
import '@fortawesome/fontawesome-free/css/fontawesome.css';
import '@fortawesome/fontawesome-free/css/solid.css';

import App from '@/App.vue';
import store from '@/store';
import router from '@/index';
import '@/assets/styles.scss';

Vue.config.productionTip = false;

Vue.use(Buefy);
Vue.use(VueAnalytics, {
  id: 'UA-28203673-6',
  autoTracking: {
    exception: true,
  },
  router,
});

/**
 * Declare the main Vue instance with components and vuex store.
 */
new Vue({
  store,
  router,
  render: (h) => h(App),
  beforeCreate() {
    this.$store.commit('initializeSettings');
  },
}).$mount('#app');

