<template>
	<div>
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button type="primary" icon="plus" @click="addUser">新增用户</el-button>
			</div>
			<el-table :data="tableData" row-key="ID">
				<el-table-column align="left" label="头像" min-width="75" />
				<el-table-column align="left" label="ID" min-width="50" prop="ID" />
				<el-table-column align="left" label="用户名" min-width="150" prop="userName" />
				<el-table-column align="left" label="昵称" min-width="150" prop="nickName" />
				<el-table-column align="left" label="手机号" min-width="180" prop="phone" />
				<el-table-column align="left" label="邮箱" min-width="180" prop="email" />
				<el-table-column label="操作" min-width="250" fixed="right">
					<template #default="scope">
						<el-popover v-model="scope.row.visible" placement="top" width="160">
							<p>确定要删除此用户吗</p>
							<div style="text-align: right; margin-top: 8px">
								<el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
								<el-button type="primary" @click="deleteUserFunc(scope.row)">确定</el-button>
							</div>
							<template #reference>
								<el-button type="primary" link icon="delete">删除</el-button>
							</template>
						</el-popover>
						<el-button type="primary" link icon="edit" @click="openEdit(scope.row)">编辑</el-button>
						<el-button type="primary" link icon="magic-stick" @click="resetPasswordFunc(scope.row)">重置密码</el-button>
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
		<el-dialog v-model="addUserDialog" title="用户" :show-close="false" :close-on-press-escape="false" :close-on-click-modal="false">
			<div style="height: 60vh; overflow: auto; padding: 0 12px">
				<el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
					<el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
						<el-input v-model="userInfo.userName" />
					</el-form-item>
					<el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
						<el-input v-model="userInfo.password" />
					</el-form-item>
					<el-form-item label="昵称" prop="nickName">
						<el-input v-model="userInfo.nickName" />
					</el-form-item>
					<el-form-item label="手机号" prop="phone">
						<el-input v-model="userInfo.phone" />
					</el-form-item>
					<el-form-item label="邮箱" prop="email">
						<el-input v-model="userInfo.email" />
					</el-form-item>
					<el-form-item label="用户角色" prop="authorityId">
						<el-cascader
							v-model="userInfo.authorityIds"
							style="width: 100%"
							:options="authOptions"
							:show-all-levels="false"
							:props="{ multiple: true, checkStrictly: true, label: 'authorityName', value: 'authorityId', disabled: 'disabled', emitPath: false }"
							:clearable="false"
						/>
					</el-form-item>
					<el-form-item label="启用" prop="disabled">
						<el-switch v-model="userInfo.enable" inline-prompt :active-value="1" :inactive-value="2" />
					</el-form-item>
					<el-form-item label="头像" label-width="80px">
						<div style="display: inline-block" @click="openHeaderChange">
							<img
								v-if="userInfo.headerImg"
								alt="头像"
								class="header-img-box"
								:src="userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http' ? path + userInfo.headerImg : userInfo.headerImg"
							/>
							<div v-else class="header-img-box">从媒体库选择</div>
						</div>
					</el-form-item>
				</el-form>
			</div>

			<template #footer>
				<div class="dialog-footer">
					<el-button @click="closeAddUserDialog">取 消</el-button>
					<el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
				</div>
			</template>
		</el-dialog>
	</div>
</template>

<script setup>
import { ref, watch } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';

defineOptions({
	name: 'UserManage',
});

const path = ref(import.meta.env.VITE_BASE_API + '/');
// 初始化相关

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
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

watch(
	() => tableData.value,
	() => {
		setAuthorityIds();
	}
);

const initPage = async () => {};

initPage();

const resetPasswordFunc = (row) => {
	ElMessageBox.confirm('是否将此用户密码重置为123456?', '警告', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	});
	console.log(row);
};
const setAuthorityIds = () => {
	tableData.value &&
		tableData.value.forEach((user) => {
			user.authorityIds =
				user.authorities &&
				user.authorities.map((i) => {
					return i.authorityId;
				});
		});
};

const chooseImg = ref(null);
const openHeaderChange = () => {
	chooseImg.value.open();
};

const authOptions = ref([]);

const deleteUserFunc = async (row) => {
	ElMessage.success('删除成功');
	console.log(row);
};

// 弹窗相关
const userInfo = ref({
	username: '',
	password: '',
	nickName: '',
	headerImg: '',
	authorityId: '',
	authorityIds: [],
	enable: 1,
});

const rules = ref({
	userName: [
		{ required: true, message: '请输入用户名', trigger: 'blur' },
		{ min: 5, message: '最低5位字符', trigger: 'blur' },
	],
	password: [
		{ required: true, message: '请输入用户密码', trigger: 'blur' },
		{ min: 6, message: '最低6位字符', trigger: 'blur' },
	],
	nickName: [{ required: true, message: '请输入用户昵称', trigger: 'blur' }],
	phone: [{ pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' }],
	email: [{ pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur' }],
	authorityId: [{ required: true, message: '请选择用户角色', trigger: 'blur' }],
});
const userForm = ref(null);
const enterAddUserDialog = async () => {
	userInfo.value.authorityId = userInfo.value.authorityIds[0];
	userForm.value.validate(async (valid) => {
		if (valid) {
			if (dialogFlag.value === 'add') {
				ElMessage({ type: 'success', message: '创建成功' });
				closeAddUserDialog();
			}
			if (dialogFlag.value === 'edit') {
				ElMessage({ type: 'success', message: '编辑成功' });
				closeAddUserDialog();
			}
		}
	});
};

const addUserDialog = ref(false);
const closeAddUserDialog = () => {
	userForm.value.resetFields();
	userInfo.value.headerImg = '';
	userInfo.value.authorityIds = [];
	addUserDialog.value = false;
};

const dialogFlag = ref('add');

const addUser = () => {
	dialogFlag.value = 'add';
	addUserDialog.value = true;
};

const openEdit = (row) => {
	dialogFlag.value = 'edit';
	userInfo.value = JSON.parse(JSON.stringify(row));
	addUserDialog.value = true;
};
</script>

<style lang="scss">
.header-img-box {
	@apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
</style>
