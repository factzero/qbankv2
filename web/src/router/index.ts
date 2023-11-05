import { createRouter, createWebHistory } from 'vue-router';

// 静态路由
const routes = [
	{
		path: '/',
		redirect: '/login',
	},
	{
		path: '/login',
		name: 'Login',
		component: () => import('@/views/login/loginView.vue'),
	},
	{
		path: '/home',
		name: 'Home',
		component: () => import('@/views/layout/layoutIndex.vue'),
		children: [
			{
				path: '/1',
				name: '1',
				meta: {
					title: '首页',
					icon: 'ios-list-box',
				},
				component: () => import('@/views/dashboard/dashboardIndex.vue'),
			},
		],
	},
];

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes,
});

export default router;
