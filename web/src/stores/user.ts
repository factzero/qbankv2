import { ref } from 'vue';
import { defineStore } from 'pinia';
import { login } from '../api/user';

export const useUserStore = defineStore('user', () => {
	const userInfo = ref({
		uuid: '',
		nickName: '',
		headerImg: '',
		authority: {},
	});

	const setUserInfo = (val) => {
		userInfo.value = val;
	};

	const token = ref(window.localStorage.getItem('token') || '');
	const setToken = (val) => {
		token.value = val;
	};

	/* 登录*/
	const LoginIn = (loginData) => {
		return new Promise((resolve, reject) => {
			login(loginData)
				.then((respone) => {
					setUserInfo(respone.user);
					console.log('logintoken:', respone.data);
					setToken(respone.data);
					console.log('settoken:', token);
					resolve(respone.token);
				})
				.catch((err) => {
					reject(err);
				});
		});
	};

	return { LoginIn, token };
});
