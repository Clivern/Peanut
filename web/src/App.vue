<!-- @format -->

<template>
	<div id="app">
		<div id="nav">
			<router-link to="/"
				><b-icon pack="fas" icon="home" size="is-small"> </b-icon>
				Home</router-link
			>
		</div>
		<router-view @refresh-state="refreshState" />
	</div>
</template>

<style>
#app {
	text-align: center;
	color: #2c3e50;
}

#nav {
	padding: 30px;
}

#nav a {
	font-weight: bold;
	color: #2c3e50;
}

#nav a.router-link-exact-active {
	color: #42b983;
}
</style>

<script>
export default {
	data() {
		return {
			logged:
				localStorage.getItem("user_api_key") != null &&
				localStorage.getItem("user_email") != null,
		};
	},
	methods: {
		logout() {
			this.logged = false;
			localStorage.removeItem("user_api_key");
			localStorage.removeItem("user_email");
			localStorage.removeItem("user_id");
			localStorage.removeItem("user_name");
		},
		refreshState() {
			this.logged =
				localStorage.getItem("user_api_key") != null &&
				localStorage.getItem("user_email") != null;
		},
	},
	mounted() {
		this.logged =
			localStorage.getItem("user_api_key") != null &&
			localStorage.getItem("user_email") != null;
	},
};
</script>
