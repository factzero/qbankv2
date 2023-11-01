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

	/* 登录*/
	const LoginIn = (loginData) => {
		return new Promise((resolve, reject) => {
			login(loginData)
				.then((respone) => {
					setUserInfo(respone.data.user);
					resolve(respone.data.token);
				})
				.catch((err) => {
					reject(err);
				});
		});
	};

	return { LoginIn };
});
