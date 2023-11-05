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
					<el-menu-item index="1">
						<template #title>
							<el-icon><Odometer /></el-icon>
							<span>仪表盘</span>
						</template>
					</el-menu-item>
					<el-sub-menu index="3">
						<template #title>
							<el-icon><icon-menu /></el-icon>
							<span>超级管理员</span>
						</template>
						<el-menu-item-group title="Group One">
							<el-menu-item index="3-1">item one</el-menu-item>
							<el-menu-item index="3-2">item two</el-menu-item>
						</el-menu-item-group>
						<el-menu-item-group title="Group Two">
							<el-menu-item index="3-3">item three</el-menu-item>
						</el-menu-item-group>
						<el-sub-menu index="3-4">
							<template #title>item four</template>
							<el-menu-item index="3-4-1">item one</el-menu-item>
						</el-sub-menu>
					</el-sub-menu>
				</el-menu>
			</transition>
		</el-scrollbar>
	</div>
</template>

<script lang="ts" setup>
import { Menu as IconMenu, Odometer } from '@element-plus/icons-vue';
import router from '@/router/index';

defineOptions({
	name: 'AsideIndex',
});

const selectMenuItem = (e) => {
	console.log('select:', e);
	router.push({ path: e });
};
</script>
