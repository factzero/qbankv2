import { service } from '../utils/request';

// @Summary 手工录入试题
// @Produce  application/json
// @Param data body {}
// @Router /admin/managequestion/managequestion/manualEntry [post]
export const manualentry = (data) => {
	return service({
		url: '/admin/managequestion/managequestion/manualEntry',
		method: 'post',
		data: data,
	});
};

// @Summary 查询试题
// @Produce  application/json
// @Param data body {}
// @Router /admin/managequestion/managequestion/manualEntry [post]
export const getquestions = (data) => {
	return service({
		url: '/admin/managequestion/managequestion/get_questions',
		method: 'get',
		data: data,
	});
};
