import Vue from 'vue'
import App from './App.vue'
import Pulse from 'pulse-framework';
import BootstrapVue from 'bootstrap-vue';

const pulse = new Pulse.Library({
  collections: {
    channels: {},
    posts: {}
  }
});


Vue.use(BootstrapVue);
Vue.use(pulse);

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

export default pulse;