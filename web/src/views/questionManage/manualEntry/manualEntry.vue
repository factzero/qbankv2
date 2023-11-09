<template>
	<div>
		<div class="gva-search-box">
			<el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="120px" class="demo-ruleForm" :size="formSize" status-icon>
				<el-form-item label="题型" prop="qtype">
					<el-select v-model="ruleForm.qtype" placeholder="请选择">
						<el-option label="单选题" value="radio" />
						<el-option label="多选题" value="multiple" />
						<el-option label="判断题" value="determine" />
						<el-option label="填空题" value="blanks" />
						<el-option label="简答题" value="analytic" />
					</el-select>
				</el-form-item>
				<el-form-item label="题干" prop="stem">
					<el-input v-model="ruleForm.stem" />
				</el-form-item>
				<el-form-item label="选项" prop="option">
					<el-input v-model="ruleForm.option" />
				</el-form-item>
				<el-form-item label="答案" prop="answer">
					<el-input v-model="ruleForm.answer" />
				</el-form-item>
				<el-form-item label="解析" prop="analysis">
					<el-input v-model="ruleForm.analysis" />
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="submitForm(ruleFormRef)">提交</el-button>
					<el-button @click="resetForm(ruleFormRef)">重置</el-button>
				</el-form-item>
			</el-form>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';

defineOptions({
	name: 'ManualEntry',
});

interface RuleForm {
	qtype: string;
	stem: string;
	option: string;
	answer: string;
	analysis: string;
}

const formSize = ref('default');
const ruleFormRef = ref<FormInstance>();
const ruleForm = reactive<RuleForm>({
	qtype: '',
	stem: '',
	option: '',
	answer: '',
	analysis: '',
});

const rules = reactive<FormRules<RuleForm>>({});

const submitForm = async (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	await formEl.validate((valid, fields) => {
		if (valid) {
			console.log('submit!');
		} else {
			console.log('error submit!', fields);
		}
	});
};

const resetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.resetFields();
};
</script>
