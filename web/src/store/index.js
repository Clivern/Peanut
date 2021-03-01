/** @format */

import Vue from "vue";
import Vuex from "vuex";
import service from "./service.module";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: { service },
});
