<template>
	<el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="120px" class="demo-ruleForm">
		<el-form-item label="账号" prop="username">
			<el-input v-model="ruleForm.username" placeholder="请输入用户名" suffix-icon="user" />
		</el-form-item>
		<el-form-item label="密码" prop="password">
			<el-input v-model="ruleForm.password" show-password placeholder="请输入密码" />
		</el-form-item>
		<el-form-item>
			<el-button type="primary" @click="submitForm(ruleFormRef)">登录</el-button>
		</el-form-item>
	</el-form>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';

const ruleFormRef = ref<FormInstance>();

const checkUsername = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('Please input the username'));
	} else {
		callback();
	}
};

const validatePassword = (rule: any, value: any, callback: any) => {
	if (value === '') {
		callback(new Error('Please input the password'));
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

const submitForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate((valid) => {
		if (valid) {
			console.log('submit!');
		} else {
			console.log('error submit!');
			return false;
		}
	});
};
</script>
