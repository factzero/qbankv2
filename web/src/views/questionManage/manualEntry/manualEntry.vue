<template>
	<div>
		<div class="gva-table-box">
			<el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="120px" class="demo-ruleForm" :size="formSize" status-icon>
				<el-form-item label="题型" prop="qtype">
					<el-select v-model="ruleForm.qtype" placeholder="填空题">
						<el-option v-for="item in questionType" :label="item.label" :key="item.value" :value="item.value"></el-option>
					</el-select>
				</el-form-item>
				<el-form-item label="题干" prop="stem">
					<el-input v-model="ruleForm.stem" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="Please input" />
				</el-form-item>
				<template v-if="['radio', 'multiple'].includes(ruleForm.qtype)">
					<el-row :gutter="20">
						<el-col :span="4">
							<el-form-item>
								<el-button type="primary" @click="addOptionConfig">新增行</el-button>
							</el-form-item>
						</el-col>
					</el-row>
					<el-row :gutter="20" v-for="(item, index) in ruleForm.optionConfig" :key="item.index">
						<el-col :span="1">
							<el-form-item label="选项" prop="'option' + index">
								{{ setOptionCfgID(item, index) }}
							</el-form-item>
						</el-col>
						<el-col :span="15">
							<el-form-item label="" prop="'context' + index">
								<el-input v-model="item.context" :autosize="{ minRows: 1, maxRows: 4 }" type="textarea" placeholder="Please input" />
							</el-form-item>
						</el-col>
						<el-col :span="4">
							<el-button @click.prevent="removeOptionConfig(item)">删除行</el-button>
						</el-col>
					</el-row>
					<el-form-item label="答案" prop="answer" v-if="ruleForm.qtype === 'radio'">
						<el-select v-model="value1" placeholder="Select" style="width: 240px">
							<el-option v-for="item in ruleForm.optionConfig" :key="item.index" :label="item.index" :value="item.index" />
						</el-select>
					</el-form-item>
					<el-form-item label="答案" prop="answer" v-else>
						<el-select v-model="value1" multiple placeholder="Select" style="width: 240px">
							<el-option v-for="item in ruleForm.optionConfig" :key="item.index" :label="item.index" :value="item.index" />
						</el-select>
					</el-form-item>
				</template>
				<template v-else-if="['determine'].includes(ruleForm.qtype)">
					<el-form-item label="答案" prop="answer">
						<el-radio-group v-model="radio1" class="ml-4">
							<el-radio label="1" size="large">对</el-radio>
							<el-radio label="2" size="large">错</el-radio>
						</el-radio-group>
					</el-form-item>
				</template>
				<template v-else>
					<el-form-item label="答案" prop="answer">
						<el-input v-model="ruleForm.answer" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="Please input" />
					</el-form-item>
				</template>
				<el-form-item label="解析" prop="analysis">
					<el-input v-model="ruleForm.analysis" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="Please input" />
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
import { useManageQStore } from '@/stores/managequestion';

defineOptions({
	name: 'ManualEntry',
});

const questionType = [
	{
		value: 'radio',
		label: '单选题',
	},
	{
		value: 'multiple',
		label: '多选题',
	},
	{
		value: 'determine',
		label: '判断题',
	},
	{
		value: 'blanks',
		label: '填空题',
	},
	{
		value: 'analytic',
		label: '简答题',
	},
];

interface RuleForm {
	qtype: string;
	stem: string;
	option: string;
	answer: string;
	analysis: string;
	optionConfig: [
		{
			index: string;
			context: string;
		},
	];
}

const formSize = ref('default');
const ruleFormRef = ref<FormInstance>();
const ruleForm = reactive<RuleForm>({
	qtype: 'blanks',
	stem: '',
	option: '',
	answer: '',
	analysis: '',
	optionConfig: [
		{
			index: '',
			context: '',
		},
	],
});

const rules = reactive<FormRules<RuleForm>>({});

const setOptionCfgID = (item, n) => {
	const index = ruleForm.optionConfig.indexOf(item);
	const c = String.fromCharCode(65 + n).toUpperCase();
	ruleForm.optionConfig[index].index = c;
	return c;
};

// const num2char = (n) => {
// 	return String.fromCharCode(65 + n).toUpperCase();
// };

const addOptionConfig = () => {
	ruleForm.optionConfig.push({
		index: '',
		context: '',
	});
};

const removeOptionConfig = (item) => {
	const index = ruleForm.optionConfig.indexOf(item);
	if (index !== -1) {
		ruleForm.optionConfig.splice(index, 1);
	}
};

const value1 = ref([]);
const radio1 = ref('1');

const restFormData = () => {
	ruleForm.stem = '';
	ruleForm.answer = '';
	ruleForm.analysis = '';
	ruleForm.optionConfig = [
		{
			index: '',
			context: '',
		},
	];
};

const submitForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	formEl.validate((valid, fields) => {
		if (valid) {
			useManageQStore()
				.ManualEntry(ruleForm)
				.then(() => {
					console.log('submitFormdddd submit!');
					restFormData();
				});
			console.log('submitForm submit!');
		} else {
			console.log('submitForm error submit!', fields);
		}
	});
};

const resetForm = (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	restFormData();
};
</script>
