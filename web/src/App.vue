<script setup lang="ts">
import { RouterView } from 'vue-router';
</script>

<template>
	<div id="app">
		<RouterView />
	</div>
</template>

<style lang="scss">
@tailwind base;
@tailwind components;
@tailwind utilities;

// 引入初始化样式
#app {
	background: #eee;
	height: 100vh;
	overflow: hidden;
	font-weight: 400 !important;
}
.el-button {
	font-weight: 400 !important;
}
</style>
