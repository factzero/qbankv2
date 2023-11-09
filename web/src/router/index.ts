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
		name: 'home',
		redirect: '/dashboard',
		component: () => import('@/views/layout/layoutIndex.vue'),
		meta: {
			title: 'home',
		},
		children: [
			{
				path: '/dashboard',
				name: 'dashboard',
				component: () => import('@/views/dashboard/dashboardIndex.vue'),
				meta: {
					title: 'dashboard',
				},
			},
			{
				path: '/superadmin',
				name: 'superadmin',
				redirect: '/usermanage',
				component: () => import('@/views/superAdmin/superAdminIndex.vue'),
				meta: {
					title: 'superadmin',
				},
				children: [
					{
						path: '/usermanage',
						name: 'usermanage',
						component: () => import('@/views/superAdmin/user/userManage.vue'),
						meta: {
							title: 'usermanage',
						},
					},
					{
						path: '/sysoprecord',
						name: 'sysoprecord',
						component: () => import('@/views/superAdmin/operation/sysOperationRecord.vue'),
						meta: {
							title: 'sysoprecord',
						},
					},
				],
			},
			{
				path: '/questionmanage',
				name: 'questionmanage',
				component: () => import('@/views/questionManage/questionManage.vue'),
				meta: {
					title: 'questionmanage',
				},
				children: [
					{
						path: '/manualentry',
						name: 'manualentry',
						component: () => import('@/views/questionManage/manualEntry/manualEntry.vue'),
						meta: {
							title: 'manualentry',
						},
					},
				],
			},
		],
	},
];

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes,
});

export default router;
