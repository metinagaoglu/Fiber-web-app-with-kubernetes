import { createApp } from 'vue';
import App from './App.vue';
import 'bootstrap/dist/css/bootstrap.css';
import axios from 'axios';
import { createRouter, createWebHistory } from 'vue-router';

import LoginPage from './components/Pages/Login.vue';
import RegisterPage from './components/Pages/Register.vue';
import DashboardPage from './components/Pages/Home.vue';

console.log("VUE_APP_API_URL:",import.meta.env.VUE_APP_API_URL);
axios.defaults.baseURL = "http://172.19.0.2:3000";
/*
axios.interceptors.request.use(function (config) {
  config.headers['X-Binarybox-Api-Key'] = process.env.VUE_APP_API_KEY;
  return config;
});
  */
  
const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: LoginPage },
    { path: '/register', component: RegisterPage },
    { path: '/dashboard', component: DashboardPage },
  ],
});
  
createApp(App).use(router).mount('#app');