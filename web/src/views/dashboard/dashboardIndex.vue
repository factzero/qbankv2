<template>
	<div class="page">
		<div class="gva-card-box">
			<div class="gva-card gva-top-card">
				<div class="gva-top-card-left">
					<div class="gva-top-card-left-title">早安，管理员，请开始一天的工作吧</div>
					<div class="gva-top-card-left-dot">{{ weatherInfo }}</div>
					<el-row class="my-8 w-[500px]"> </el-row>
				</div>
				<img src="@/assets/dashboard.png" class="gva-top-card-right" alt />
			</div>
		</div>
		<div class="gva-card-box">
			<div class="gva-card">
				<div class="gva-card-title">数据统计</div>
				<div class="p-4">
					<el-row :gutter="20">
						<el-col :xs="24" :sm="18">
							<echarts-line />
						</el-col>
						<el-col :xs="24" :sm="6">
							<dashboard-table />
						</el-col>
					</el-row>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import EchartsLine from '@/views/dashboard/dashboardCharts/echartsLine.vue';
import DashboardTable from '@/views/dashboard/dashboardTable/dashboardTable.vue';
import { useWeatherInfo } from '@/views/dashboard/weather.ts';

defineOptions({
	name: 'DashboardIndex',
});

const weatherInfo = useWeatherInfo();
</script>

<style lang="scss" scoped>
.page {
	@apply p-0;
	.gva-card-box {
		@apply p-4;
		& + .gva-card-box {
			@apply pt-0;
		}
	}
	.gva-card {
		@apply box-border bg-white rounded h-auto px-6 py-8 overflow-hidden shadow-sm;
		.gva-card-title {
			@apply pb-5 border-t-0 border-l-0 border-r-0 border-b border-solid border-gray-100;
		}
	}
	.gva-top-card {
		@apply h-72 flex items-center justify-between text-gray-500;
		&-left {
			@apply h-full flex flex-col w-auto;
			&-title {
				@apply text-3xl text-gray-600;
			}
			&-dot {
				@apply mt-4 text-gray-600 text-lg;
			}
			&-item {
				+ .gva-top-card-left-item {
					margin-top: 24px;
				}
				margin-top: 14px;
			}
		}
		&-right {
			height: 600px;
			width: 600px;
			margin-top: 28px;
		}
	}
	::v-deep(.el-card__header) {
		@apply p-0  border-gray-200;
	}
	.card-header {
		@apply pb-5 border-b border-solid border-gray-200 border-t-0 border-l-0 border-r-0;
	}
	.quick-entrance-items {
		@apply flex items-center justify-center text-center text-gray-800;
		.quick-entrance-item {
			@apply px-8 py-6 flex items-center flex-col transition-all duration-100 ease-in-out rounded-lg cursor-pointer;
			&:hover {
				@apply shadow-lg;
			}
			&-icon {
				@apply flex items-center h-16 w-16 rounded-lg justify-center mx-0 my-auto text-2xl;
			}
			p {
				@apply mt-2.5;
			}
		}
	}
}
.dashboard-icon {
	@apply flex items-center text-xl mr-2 text-blue-400;
}
</style>
