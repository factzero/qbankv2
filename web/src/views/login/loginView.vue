<template>
	<div id="userLayout" class="w-full h-full relative">
		<div class="rounded-lg flex items-center justify-evenly w-full h-full bg-white md:w-screen md:h-screen md:bg-[#194bfb]">
			<div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
				<div class="oblique h-[130%] w-3/5 bg-white transform -rotate-12 absolute -ml-52" />
				<div class="z-[999] pt-12 pb-10 md:w-96 w-full rounded-lg flex flex-col justify-between box-border">
					<div>
						<div class="flex items-center justify-center">
							<img class="w-24" src="../../assets/logo.jpg" />
						</div>
						<div class="mb-9">
							<p class="text-center text-4xl font-bold">试题管理系统</p>
						</div>
						<el-form ref="loginForm" :model="loginFormData" :rules="rules" Q :validate-on-rule-change="false" @keyup.enter="submitForm">
							<el-form-item class="mb-6" prop="username">
								<el-input size="large" v-model="loginFormData.username" placeholder="请输入用户名" suffix-icon="user" />
							</el-form-item>
							<el-form-item class="mb-6" prop="password">
								<el-input size="large" v-model="loginFormData.password" show-password placeholder="请输入密码" />
							</el-form-item>
							<el-form-item class="mb-6">
								<el-button class="shadow shadow-blue-600 h-11 w-full" size="large" type="primary" @click="submitForm">登录</el-button>
							</el-form-item>
						</el-form>
					</div>
				</div>
			</div>
			<div class="hidden md:block w-full h-full float-right bg-[#194bfb]">
				<img class="h-full" src="../../assets/login_right_banner.jpg" alt="banner" />
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';
import type { FormRules, ElMessage } from 'element-plus';
import { useUserStore } from '../../stores/user';
import router from '@/router/index';

const checkUsername = (rule: any, value: any, callback: any) => {
	if (value.length < 4) {
		return callback(new Error('请输入正确的用户名'));
	} else {
		callback();
	}
};

const validatePassword = (rule: any, value: any, callback: any) => {
	if (value.length < 4) {
		return callback(new Error('请输入正确的密码'));
	} else {
		callback();
	}
};

const loginForm = ref(null);
const loginFormData = reactive({
	username: '',
	password: '',
});

const rules = reactive<FormRules<typeof loginFormData>>({
	username: [{ validator: checkUsername, trigger: 'blur' }],
	password: [{ validator: validatePassword, trigger: 'blur' }],
});

const userStore = useUserStore();
const submitForm = () => {
	loginForm.value.validate(async (v) => {
		if (v) {
			console.log('submit!');
			userStore.LoginIn(loginFormData).then(() => {
				router.push({ path: '/dashboard' });
			});
		} else {
			ElMessage({
				type: 'error',
				message: '请正确填写登录信息',
				showClose: true,
			});
			return false;
		}
	});
};
</script>
