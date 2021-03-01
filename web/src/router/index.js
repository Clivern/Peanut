/** @format */

import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
	{
		path: "/",
		name: "Home",
		component: () => import("../views/Home.vue"),
		meta: {
			requiresAuth: false,
		},
	},
	{
		path: "/login",
		name: "Login",
		component: () => import("../views/Login.vue"),
		meta: {
			requiresAuth: false,
		},
	},
	{
		path: "/deploy",
		name: "Deploy",
		component: () => import("../views/Deploy.vue"),
		meta: {
			requiresAuth: true,
		},
	},
	{
		path: "/services",
		name: "Services",
		component: () => import("../views/Services.vue"),
		meta: {
			requiresAuth: true,
		},
	},
	{
		path: "/404",
		name: "NotFound",
		component: () => import("../views/NotFound.vue"),
	},
	{
		path: "*",
		redirect: "/404",
	},
];

const router = new VueRouter({
	routes,
});

// Auth Middleware
router.beforeEach((to, from, next) => {
	if (to.matched.some((record) => record.meta.requiresAuth)) {
		if (localStorage.getItem("x_api_key") == null) {
			next({
				path: "/login",
				params: { nextUrl: to.fullPath },
			});
		}
	} else if (to.name == "Login") {
		localStorage.removeItem("x_api_key");
	}
	next();
});

export default router;
