<template>
	<div :style="{ background: '#fff' }">
		<el-scrollbar style="height: calc(100% - ${'60px'})">
			<transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
				<el-menu
					active-text-color="#ffd04b"
					background-color="#191a23"
					class="el-menu-vertical"
					default-active="2"
					text-color="#fff"
					unique-opened
					@select="selectMenuItem"
				>
					<el-menu-item index="dashboard">
						<template #title>
							<el-icon><Odometer /></el-icon>
							<span>仪表盘</span>
						</template>
					</el-menu-item>
					<el-sub-menu index="superadmin">
						<template #title>
							<el-icon><User /></el-icon>
							<span>超级管理员</span>
						</template>
						<el-menu-item index="usermanage">
							<el-icon><Coordinate /></el-icon>
							<span>用户管理</span>
						</el-menu-item>
						<el-menu-item index="sysoprecord">
							<el-icon><pie-chart /></el-icon>
							<span>操作历史</span>
						</el-menu-item>
					</el-sub-menu>
				</el-menu>
			</transition>
		</el-scrollbar>
	</div>
</template>

<script lang="ts" setup>
import { Odometer, User, Coordinate, PieChart } from '@element-plus/icons-vue';
import router from '@/router/index';

defineOptions({
	name: 'AsideIndex',
});

const selectMenuItem = (e) => {
	console.log('select:', e);
	router.push({ path: e });
};
</script>
