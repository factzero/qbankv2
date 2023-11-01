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
		base: './', // index.html�ļ�����λ��
		root: './', // js�������Դ·����src
		resolve: {
			alias,
		},
		define: {
			'process.env': {},
		},
		server: {
			// ���ʹ��docker-compose����ģʽ������Ϊfalse
			open: true,
			port: process.env.VITE_APP_PORT,
			proxy: {
				// ��key��·������targetλ��
				// detail: https://cli.vuejs.org/config/#devserver-proxy
				[process.env.VITE_BASE_API]: {
					// ��Ҫ�����·��   ���� '/api'
					target: `${process.env.VITE_SERVER_PATH}:${process.env.VITE_SERVER_PORT}/`, // ���� Ŀ��·��
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
