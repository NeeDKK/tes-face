import Vue from 'vue'
import App from './App.vue'
import router from './router'
import 'element-ui/lib/theme-chalk/index.css';
import ElementUI from 'element-ui';
import ViewUI from 'view-design';
import axios from 'axios';
import './assets/css/main.css';
import './assets/css/theme.scss';

Vue.config.productionTip = false

Vue.use(ElementUI);
Vue.use(ViewUI);

const instance = axios.create({
  baseURL: '/api'
});

// 配置过滤response
instance.interceptors.response.use((response) => {
  if (response) {
    return response;
  }
}, error => {
  return Promise.reject(error);
});
instance.interceptors.request.use((request) => {
  return request;
});
// 跨域请求时是否需要使用凭证
instance.defaults.withCredentials = true;
instance.defaults.headers.post['Content-Type'] = 'application/json';
instance.defaults.headers.get['Content-Type'] = 'application/json';

Vue.prototype.$axios = instance;
axios.defaults.baseURL = 'http://localhost:9999';

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
