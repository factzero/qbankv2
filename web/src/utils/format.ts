import { formatTimeToStr } from '@/utils/date';

export const formatDate = (time) => {
	if (time !== null && time !== '') {
		const date = new Date(time);
		return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss');
	} else {
		return '';
	}
};
