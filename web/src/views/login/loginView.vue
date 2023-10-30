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
						<el-form ref="ruleFormRef" :model="ruleForm" :rules="rules">
							<el-form-item class="mb-6" prop="username">
								<el-input size="large" v-model="ruleForm.username" placeholder="请输入用户名" suffix-icon="user" />
							</el-form-item>
							<el-form-item class="mb-6" prop="password">
								<el-input size="large" v-model="ruleForm.password" show-password placeholder="请输入密码" />
							</el-form-item>
							<el-form-item class="mb-6">
								<el-button class="shadow shadow-blue-600 h-11 w-full" size="large" type="primary" @click="submitForm(ruleFormRef)">登录</el-button>
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
import type { FormInstance, FormRules } from 'element-plus';
import { useUserStore } from '../../stores/user';

const ruleFormRef = ref<FormInstance>();

const checkUsername = (rule: any, value: any, callback: any) => {
	if (value.length < 5) {
		return callback(new Error('请输入正确的用户名'));
	} else {
		callback();
	}
};

const validatePassword = (rule: any, value: any, callback: any) => {
	if (value.length < 6) {
		return callback(new Error('请输入正确的密码'));
	} else {
		callback();
	}
};

const ruleForm = reactive({
	username: '',
	password: '',
});

const rules = reactive<FormRules<typeof ruleForm>>({
	username: [{ validator: checkUsername, trigger: 'blur' }],
	password: [{ validator: validatePassword, trigger: 'blur' }],
});

const userStore = useUserStore();
const submitForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate((valid) => {
		if (valid) {
			console.log('submit!');
			const flag = userStore.LoginIn(ruleForm);
			console.log('submit!', flag);
		} else {
			console.log('error submit!');
			return false;
		}
	});
};
</script>
../../stores/user
