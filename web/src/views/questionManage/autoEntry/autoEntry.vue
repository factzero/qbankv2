<template>
	<div class="gva-table-box">
		<el-upload ref="uploadRef" class="upload-demo" accept=".png,.jpg,.jpeg,.pdf,.xls,.xlsx" action="#" :auto-upload="false">
			<template #trigger>
				<el-button type="primary">选择文件</el-button>
			</template>

			<el-button class="ml-3" type="success" @click="submitUpload">上传文件</el-button>

			<template #tip>
				<div class="el-upload__tip">
					<p>只支持.png / .jpg / .jpeg / .pdf / .xls / .xlsx文件</p>
					<p>最多上传5个文件</p>
				</div>
			</template>
		</el-upload>
	</div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import type { UploadInstance } from 'element-plus';

const uploadRef = ref<UploadInstance>();

const submitUpload = () => {
	uploadRef.value!.submit();
};
</script>
