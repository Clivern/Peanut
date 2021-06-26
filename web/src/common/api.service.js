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

	getHeaders() {
		let apiKey = "";
		let email = "";
		let id = "";

		if (localStorage.getItem("user_api_key") != null) {
			apiKey = localStorage.getItem("user_api_key");
		}

		if (localStorage.getItem("user_email") != null) {
			email = localStorage.getItem("user_email");
		}

		if (localStorage.getItem("user_id") != null) {
			id = localStorage.getItem("user_id");
		}

		return {
			crossdomain: true,

			headers: {
				"X-API-Key": apiKey,
				"X-User-Email": email,
				"X-User-ID": id,
				"X-Client-ID": "dashboard",
				"X-Requested-With": "XMLHttpRequest",
				"Content-Type": "application/json",
			},
		};
	},

	get(endpoint) {
		return axios.get(this.getURL(endpoint), this.getHeaders());
	},

	delete(endpoint) {
		return axios.delete(this.getURL(endpoint), this.getHeaders());
	},

	post(endpoint, data = {}) {
		return axios.post(this.getURL(endpoint), data, this.getHeaders());
	},

	put(endpoint, data = {}) {
		return axios.put(this.getURL(endpoint), data, this.getHeaders());
	},
};

export default ApiService;
