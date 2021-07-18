/** @format */

export { getServices, getService, deleteService, getJob, deployService };

import {
	getServices,
	getService,
	deleteService,
	getJob,
	deployService,
	getServiceTags,
} from "@/common/service.api";

const state = () => ({
	getServicesResult: {},
	getServiceResult: {},
	getDeleteServiceResult: {},
	getJobResult: {},
	getDeployServiceResult: {},
	getServiceTagsResult: {},
});

const getters = {
	getServicesResult: (state) => {
		return state.getServicesResult;
	},
	getServiceResult: (state) => {
		return state.getServiceResult;
	},
	getDeleteServiceResult: (state) => {
		return state.getDeleteServiceResult;
	},
	getJobResult: (state) => {
		return state.getJobResult;
	},
	getDeployServiceResult: (state) => {
		return state.getDeployServiceResult;
	},
	getServiceTagsResult: (state) => {
		return state.getServiceTagsResult;
	},
};

const actions = {
	async getServicesAction(context, payload) {
		const result = await getServices(payload);
		context.commit("SET_GET_SERVICES_RESULT", result.data);
		return result;
	},
	async getServiceAction(context, payload) {
		const result = await getService(payload);
		context.commit("SET_GET_SERVICE_RESULT", result.data);
		return result;
	},
	async deleteServiceAction(context, payload) {
		const result = await deleteService(payload);
		context.commit("SET_DELETE_SERVICE_RESULT", result.data);
		return result;
	},
	async getJobAction(context, payload) {
		const result = await getJob(payload);
		context.commit("SET_GET_JOB_RESULT", result.data);
		return result;
	},
	async deployServiceAction(context, payload) {
		const result = await deployService(payload);
		context.commit("SET_DEPLOY_SERVICE_RESULT", result.data);
		return result;
	},
	async getServiceTagsAction(context, payload) {
		const result = await getServiceTags(payload);
		context.commit("SET_SERVICE_TAGS_RESULT", result.data);
		return result;
	},
};

const mutations = {
	SET_GET_SERVICES_RESULT(state, getServicesResult) {
		state.getServicesResult = getServicesResult;
	},
	SET_GET_SERVICE_RESULT(state, getServiceResult) {
		state.getServiceResult = getServiceResult;
	},
	SET_DELETE_SERVICE_RESULT(state, getDeleteServiceResult) {
		state.getDeleteServiceResult = getDeleteServiceResult;
	},
	SET_GET_JOB_RESULT(state, getJobResult) {
		state.getJobResult = getJobResult;
	},
	SET_DEPLOY_SERVICE_RESULT(state, getDeployServiceResult) {
		state.getDeployServiceResult = getDeployServiceResult;
	},
	SET_SERVICE_TAGS_RESULT(state, getServiceTagsResult) {
		state.getServiceTagsResult = getServiceTagsResult;
	},
};

export default {
	namespaced: true,
	state,
	getters,
	actions,
	mutations,
};
