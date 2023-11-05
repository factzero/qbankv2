<template>
	<el-container class="layout-cont">
		<el-container>
			<el-aside class="main-cont gva-aside" :style="{ width: '220px' }">
				<div class="min-h-[60px] text-center transition-all duration-300 flex items-center justify-center gap-2" :style="{ background: '#191a23' }">
					<div class="inline-flex font-bold text-2xl" :style="{ color: '#fff' }">试题管理系统</div>
				</div>
				<AsideIndex />
			</el-aside>
			<el-main class="main-cont main-right">
				<span> main</span>
				<router-view v-slot="{ Component }" element-loading-text="正在加载中" class="admin-box">
					<div>
						<transition mode="out-in" name="el-fade-in-linear">
							<keep-alive>
								<component :is="Component" />
							</keep-alive>
						</transition>
					</div>
				</router-view>
			</el-main>
		</el-container>
	</el-container>
</template>

<script setup>
import AsideIndex from '@/views/layout/aside/asideIndex.vue';
</script>
