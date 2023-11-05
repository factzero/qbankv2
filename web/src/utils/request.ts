import axios from 'axios';
import { ElMessage } from 'element-plus';
import router from '@/router/index';

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
			console.log('22222');
			return response.data;
		} else {
			ElMessage({
				showClose: true,
				message: response.data.msg || decodeURI(response.headers.msg),
				type: 'error',
			});
			router.push({ name: 'Login', replace: true });
			return Promise.reject(new Error(response.data.msg || 'Error'));
		}
	},
	(error) => {
		return Promise.reject(error);
	}
);
