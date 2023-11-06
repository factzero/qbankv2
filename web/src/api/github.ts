import axios from 'axios';

const service = axios.create();

export function Commits(page) {
	return service({
		url: 'https://api.github.com/repos/factzero/qbankv2/commits?page=' + page,
		method: 'get',
	});
}

export function Members() {
	return service({
		url: 'https://api.github.com/repos/factzero/members',
		method: 'get',
	});
}
