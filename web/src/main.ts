import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import Antd from 'ant-design-vue';
import * as Icons from "@ant-design/icons-vue";
import 'ant-design-vue/dist/antd.css';


const app = createApp(App)
app.use(store)
app.use(router)
app.use(Antd)

const icons: any = Icons;
for (const i in icons) {
    app.component(i, icons[i]);
}

app.mount('#app')
