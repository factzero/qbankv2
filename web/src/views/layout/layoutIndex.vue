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
				<div :style="{ width: `calc(100% - ${'220px'})` }" class="fixed top-0 box-border z-50">
					<el-row>
						<el-col>
							<el-header class="header-cont">
								<el-row class="p-0 h-full">
									<el-col :xs="2" :lg="1" :md="1" :sm="1" :xl="1" class="z-50 flex items-center pl-3">
										<div class="text-black cursor-pointer text-lg leading-5" @click="totalCollapse">
											<div class="gvaIcon gvaIcon-arrow-double-right" />
										</div>
									</el-col>
									<el-col :xs="10" :lg="14" :md="14" :sm="9" :xl="14" :pull="1" class="flex items-center">
										<!-- 修改为手机端不显示顶部标签 -->
										<el-breadcrumb class="breadcrumb">
											<el-breadcrumb-item v-for="(item, index) in breadcrumbData" :key="item.path">
												<!-- 不可点击项 -->
												<span v-if="index === breadcrumbData.length - 1" class="no-redirect">{{ item.meta.title }}</span>
												<!-- 可点击项 -->
												<a v-else class="redirect" @click.prevent="onLinkClick(item)">{{ item.meta.title }}</a>
											</el-breadcrumb-item>
										</el-breadcrumb>
									</el-col>
								</el-row>
							</el-header>
						</el-col>
					</el-row>
				</div>
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
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import AsideIndex from '@/views/layout/aside/asideIndex.vue';

const route = useRoute();

// 生成数组数据
const breadcrumbData = ref([]);
const getBreadcrumbData = () => {
	breadcrumbData.value = route.matched.filter((item) => item.meta && item.meta.title);
	console.log(breadcrumbData.value);
};

// 监听路由变化时触发
watch(
	route,
	() => {
		getBreadcrumbData();
	},
	{
		immediate: true,
	}
);

// 处理点击事件
const router = useRouter();
const onLinkClick = (item) => {
	// console.log(item)
	router.push(item.path);
};
</script>
