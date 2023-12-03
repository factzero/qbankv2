import { defineStore } from 'pinia';
import { manualentry, getquestions } from '../api/managequestion';

export const useManageQStore = defineStore('manageQ', () => {
	/* 手工录入试题 */
	const ManualEntry = (Data) => {
		return new Promise((resolve, reject) => {
			manualentry(Data)
				.then(() => {
					resolve();
				})
				.catch((err) => {
					reject(err);
				});
		});
	};

	/* 查询试题 */
	const GetQuestions = (Data) => {
		return new Promise((resolve, reject) => {
			getquestions(Data)
				.then((Data) => {
					resolve(Data);
				})
				.catch((err) => {
					reject(err);
				});
		});
	};

	return { ManualEntry, GetQuestions };
});
