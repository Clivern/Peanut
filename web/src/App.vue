<!-- @format -->

<template>
	<div id="app">
		<div id="nav">
			<router-link to="/"
				><b-icon pack="fas" icon="home" size="is-small"> </b-icon>
				Home</router-link
			>
			<template v-if="logged">
				|
				<router-link to="/deploy">
					<b-icon pack="fas" icon="rocket" size="is-small"> </b-icon>
					Deploy</router-link
				>
				|
				<router-link to="/services">
					<b-icon pack="fas" icon="server" size="is-small"> </b-icon>
					Services</router-link
				>
				|
				<a href="#" @click="logout">
					<b-icon pack="fas" icon="sign-out-alt" size="is-small"> </b-icon>
					Logout</a
				>
			</template>
			<template v-else>
				|
				<router-link to="/login">
					<b-icon pack="fas" icon="sign-in-alt" size="is-small"> </b-icon>
					Login</router-link
				>
			</template>
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
			logged: localStorage.getItem("x_api_key") != null,
		};
	},
	methods: {
		logout() {
			this.logged = false;
			localStorage.removeItem("x_api_key");
			this.$router.push("/login");
		},
		refreshState() {
			this.logged = localStorage.getItem("x_api_key") != null;
		},
	},
	mounted() {
		this.logged = localStorage.getItem("x_api_key") != null;
	},
};
</script>
