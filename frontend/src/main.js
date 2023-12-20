import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import router from './router/router.js';
import naive from 'naive-ui'
const app = createApp(App);

app.use(router)
app.use(naive)
app.mount('#app');
