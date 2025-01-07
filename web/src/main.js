import './assets/css/input.css';
import './assets/css/style.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router/route';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount('#app');