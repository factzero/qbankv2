import { createRouter, createWebHistory } from 'vue-router';

// 静态路由
const routes = [
	{
		path: '/',
		name: 'Home',
		component: () => import('../components/HelloWorld.vue'),
	},
	{
		path: '/login',
		name: 'Login',
		component: () => import('../views/login/loginView.vue'),
	},
];

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes,
});

export default router;
