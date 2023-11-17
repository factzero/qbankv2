<template>
	<div>
		<div class="gva-search-box">
			<el-form :inline="true" :model="searchInfo">
				<el-form-item label="题型">
					<el-input v-model="searchInfo.method" placeholder="搜索条件" />
				</el-form-item>
				<el-form-item label="关键字">
					<el-input v-model="searchInfo.path" placeholder="搜索条件" />
				</el-form-item>
				<el-form-item>
					<el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
					<el-button icon="refresh" @click="onReset">重置</el-button>
				</el-form-item>
			</el-form>
		</div>
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-popover v-model="deleteVisible" placement="top" width="160">
					<p>确定要删除吗？</p>
					<div style="text-align: right; margin-top: 8px">
						<el-button type="primary" link @click="deleteVisible = false">取消</el-button>
						<el-button type="primary" @click="onDelete">确定</el-button>
					</div>
					<template #reference>
						<el-button icon="delete" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
					</template>
				</el-popover>
			</div>
			<el-table
				ref="multipleTable"
				:data="tableData"
				style="width: 100%"
				tooltip-effect="dark"
				row-key="ID"
				@selection-change="handleSelectionChange"
			>
				<el-table-column align="left" type="selection" width="55" />
				<el-table-column align="left" label="题型" prop="qtype" width="80" />
				<el-table-column align="left" label="题干" prop="stem" width="120" />
				<el-table-column align="left" label="选项" prop="option" width="120" />
				<el-table-column align="left" label="答案" prop="answer" width="240" />
				<el-table-column align="left" label="解析" prop="analysis" width="240" />
				<el-table-column align="left" label="日期" width="120">
					<template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
				</el-table-column>
				<el-table-column align="left" label="操作">
					<template #default="scope">
						<el-popover v-model="scope.row.visible" placement="top" width="160">
							<p>确定要删除吗？</p>
							<div style="text-align: right; margin-top: 8px">
								<el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
								<el-button type="primary" @click="deleteSysOperationRecordFunc(scope.row)">确定</el-button>
							</div>
							<template #reference>
								<el-button icon="delete" type="primary" link @click="scope.row.visible = true">删除</el-button>
							</template>
						</el-popover>
					</template>
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

defineOptions({
	name: 'SysOperationRecord',
});

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
const getTableData = async () => {};

getTableData();

const deleteVisible = ref(false);
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

const deleteSysOperationRecordFunc = async (row) => {
	console.log(row);
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
