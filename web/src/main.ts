import './style/element_visiable.scss';
import { createApp } from 'vue';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import ElementPlus from 'element-plus';
import zhCN from 'element-plus/dist/locale/zh-cn.mjs';
import App from './App.vue';
import router from './router';

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
	app.component(key, component);
}

app.use(router);
app.use(ElementPlus, { locale: zhCN });

app.mount('#app');
