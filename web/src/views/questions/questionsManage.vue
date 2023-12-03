<template>
	<div>
		<div class="gva-search-box">
			<el-form :inline="true" :model="searchInfo">
				<el-form-item>
					<el-select v-model="value" clearable placeholder="题型">
						<el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
					</el-select>
				</el-form-item>
				<el-form-item>
					<el-input v-model="searchInfo.path" placeholder="关键字" />
				</el-form-item>
				<el-form-item>
					<el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
					<el-button icon="refresh" @click="onReset">重置</el-button>
				</el-form-item>

				<el-form-item>
					<el-radio-group v-model="labelPosition" label="label position">
						<el-radio-button label="left">Left</el-radio-button>
						<el-radio-button label="right">Right</el-radio-button>
						<el-radio-button label="top">Top</el-radio-button>
					</el-radio-group>
				</el-form-item>
			</el-form>
		</div>
		<div class="gva-table-box">
			<el-table
				ref="multipleTable"
				:data="tableData"
				style="width: 100%"
				tooltip-effect="dark"
				row-key="ID"
				@selection-change="handleSelectionChange"
			>
				<el-table-column align="left" type="selection" width="55" />
				<el-table-column align="left" label="题型" prop="qtype" width="120" />
				<el-table-column align="left" label="题干" prop="stem" width="420" />
				<el-table-column align="left" label="创建时间" width="180">
					<template #default="scope">{{ formatDate(scope.row.createtime) }}</template>
				</el-table-column>
				<el-table-column align="left" label="修改时间" width="180">
					<template #default="scope">{{ formatDate(scope.row.updatetime) }}</template>
				</el-table-column>
				<el-table-column align="left" label="编辑">
					<el-button icon="delete" type="primary" link @click="onDelete" />
					<el-button icon="edit" type="primary" link @click="onDelete" />
				</el-table-column>
			</el-table>
			<div class="gva-pagination">
				<el-pagination
					:current-page="page"
					:page-size="pageSize"
					:page-sizes="[10, 30, 50, 100]"
					:total="total"
					layout="total, sizes, prev, pager, next, jumper"
					@current-change="handleCurrentChange"
					@size-change="handleSizeChange"
				/>
			</div>
		</div>
	</div>
</template>

<script setup>
import { formatDate } from '@/utils/format';
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { useManageQStore } from '@/stores/managequestion';

defineOptions({
	name: 'QuestionRecord',
});

const value = ref('');
const options = [
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

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const onReset = () => {
	searchInfo.value = {};
};
// 条件搜索前端看此方法
const onSubmit = () => {
	page.value = 1;
	pageSize.value = 10;
	if (searchInfo.value.status === '') {
		searchInfo.value.status = null;
	}
	getTableData();
};

// 分页
const handleSizeChange = (val) => {
	pageSize.value = val;
	getTableData();
};

const handleCurrentChange = (val) => {
	page.value = val;
	getTableData();
};

// 查询
const getTableData = async () => {
	useManageQStore()
		.GetQuestions({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
		.then((res) => {
			tableData.value = res.data.item;
			console.log('getTableData:', res);
		});
};

getTableData();

const multipleSelection = ref([]);
const handleSelectionChange = (val) => {
	multipleSelection.value = val;
};

const onDelete = async () => {
	ElMessage({
		type: 'success',
		message: '删除成功',
	});
};
</script>

<style lang="scss">
.table-expand {
	padding-left: 60px;
	font-size: 0;
	label {
		width: 90px;
		color: #99a9bf;
		.el-form-item {
			margin-right: 0;
			margin-bottom: 0;
			width: 50%;
		}
	}
}
.popover-box {
	background: #112435;
	color: #f08047;
	height: 600px;
	width: 420px;
	overflow: auto;
}
.popover-box::-webkit-scrollbar {
	display: none; /* Chrome Safari */
}
</style>
