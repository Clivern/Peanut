/** @format */

import ApiService from "./api.service.js";

const getServices = (payload) => {
	let apiKey = "";

	if (payload["apiKey"]) {
		apiKey = payload["apiKey"];
	}

	return ApiService.get("/api/v1/service", apiKey);
};

const getService = (payload) => {
	let apiKey = "";

	if (payload["apiKey"]) {
		apiKey = payload["apiKey"];
	}

	return ApiService.get("/api/v1/service/" + payload["serviceId"], apiKey);
};

const deleteService = (payload) => {
	let apiKey = "";

	if (payload["apiKey"]) {
		apiKey = payload["apiKey"];
	}

	return ApiService.delete("/api/v1/service/" + payload["serviceId"], apiKey);
};

const getJob = (payload) => {
	let apiKey = "";

	if (payload["apiKey"]) {
		apiKey = payload["apiKey"];
	}

	return ApiService.get(
		"/api/v1/job/" + payload["serviceId"] + "/" + payload["jobId"],
		apiKey
	);
};

const getServiceTags = (payload) => {
	let apiKey = "";

	if (payload["apiKey"]) {
		apiKey = payload["apiKey"];
	}

	return ApiService.get(
		"/api/v1/tag/" +
			payload["org"] +
			"/" +
			payload["service"] +
			"/" +
			payload["cache"],
		apiKey
	);
};

const deployService = (payload) => {
	let apiKey = "";

	if (payload["apiKey"]) {
		apiKey = payload["apiKey"];
	}

	return ApiService.post("/api/v1/service", payload, apiKey);
};

export {
	getServices,
	getService,
	deleteService,
	getJob,
	deployService,
	getServiceTags,
};
