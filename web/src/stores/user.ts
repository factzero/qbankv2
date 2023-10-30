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
	const LoginIn = (loginInfo) => {
		try {
			const res = login(loginInfo);
			if (res.code === 0) {
				setUserInfo(res.data.user);
				return true;
			}
		} catch (e) {
			console.log(e);
		}
	};

	return { LoginIn };
});
