import axios from 'axios';
import { ElMessage } from 'element-plus';
import router from '@/router/index';
import { useUserStore } from '@/stores/user';

export const service = axios.create({
	baseURL: '/api',
	timeout: 99999,
});

// http response 拦截器
service.interceptors.response.use(
	(response) => {
		console.log('111');
		console.log(response);
		if (response.data.code === 0) {
			console.log('22222', response.data);
			return response.data;
		} else {
			ElMessage({
				showClose: true,
				message: response.data.message || decodeURI(response.headers.msg),
				type: 'error',
			});
			router.push({ name: 'Login', replace: true });
			return Promise.reject(new Error(response.data.message || 'Error'));
		}
	},
	(error) => {
		return Promise.reject(error);
	}
);

// request拦截器
service.interceptors.request.use(
	(config) => {
		// 在发送请求之前做些什么
		const userStore = useUserStore();
		console.log('token:', userStore.token);
		config.headers['Authorization'] = userStore.token;
		return config;
	},
	(error) => {
		// 处理请求错误
		console.log(error);
		return Promise.reject(error);
	}
);
