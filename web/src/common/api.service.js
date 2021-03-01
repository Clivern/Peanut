/** @format */

import axios from "axios";

const ApiService = {
	getURL(endpoint) {
		let apiURL = "";

		if (process.env.VUE_APP_API_URL) {
			apiURL = process.env.VUE_APP_API_URL.replace(/\/$/, "");
		}

		return apiURL + endpoint;
	},

	getHeaders(apiKey = "") {
		if (localStorage.getItem("x_api_key") != null) {
			apiKey = localStorage.getItem("x_api_key");
		}

		return {
			crossdomain: true,

			headers: {
				"X-API-Key": apiKey,
				"X-Client-ID": "dashboard",
				"X-Requested-With": "XMLHttpRequest",
				"Content-Type": "application/json",
			},
		};
	},

	get(endpoint, apiKey = "") {
		return axios.get(this.getURL(endpoint), this.getHeaders(apiKey));
	},

	delete(endpoint, apiKey = "") {
		return axios.delete(this.getURL(endpoint), this.getHeaders(apiKey));
	},

	post(endpoint, data = {}, apiKey = "") {
		return axios.post(this.getURL(endpoint), data, this.getHeaders(apiKey));
	},

	put(endpoint, data = {}, apiKey = "") {
		return axios.put(this.getURL(endpoint), data, this.getHeaders(apiKey));
	},
};

export default ApiService;
