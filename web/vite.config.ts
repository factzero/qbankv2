import * as path from 'path';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import vue from '@vitejs/plugin-vue';
import eslintPlugin from 'vite-plugin-eslint';
import AutoImport from 'unplugin-auto-import/vite';
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';

// https://vitejs.dev/config/
export default ({ command, mode }) => {
	const NODE_ENV = mode || 'development';
	const envFiles = [`.env.${NODE_ENV}`];
	for (const file of envFiles) {
		const envConfig = dotenv.parse(fs.readFileSync(file));
		for (const k in envConfig) {
			process.env[k] = envConfig[k];
		}
	}

	const alias = {
		'@': path.resolve(__dirname, './src'),
		vue$: 'vue/dist/vue.runtime.esm-bundler.js',
	};

	const config = {
		base: './', // index.html文件所在位置
		root: './', // js导入的资源路径，src
		resolve: {
			alias,
		},
		define: {
			'process.env': {},
		},
		server: {
			// 如果使用docker-compose开发模式，设置为false
			open: true,
			port: process.env.VITE_APP_PORT,
			proxy: {
				// 把key的路径代理到target位置
				// detail: https://cli.vuejs.org/config/#devserver-proxy
				[process.env.VITE_BASE_API]: {
					// 需要代理的路径   例如 '/api'
					target: `${process.env.VITE_SERVER_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
					changeOrigin: true,
					rewrite: (path) => path.replace(new RegExp('^' + process.env.VITE_BASE_API), ''),
				},
			},
		},
		plugins: [
			vue(),
			eslintPlugin({
				include: ['src/**/*.ts', 'src/**/*.vue', 'src/*.ts', 'src/*.vue'],
			}),
			AutoImport({
				resolvers: [ElementPlusResolver()],
			}),
			Components({
				resolvers: [ElementPlusResolver()],
			}),
		],
	};
	return config;
};
