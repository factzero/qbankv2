import { service } from '../utils/request';

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /admin/user/login [post]
export const login = (data) => {
	return service({
		url: '/admin/user/login',
		method: 'post',
		data: data,
	});
};
