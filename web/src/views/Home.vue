<!-- @format -->

<template>
	<div class="home">
		<img alt="logo" src="../assets/logo.png" width="300" />
		<div class="hello">
			<br />
			<strong>Welcome to Peanut</strong>
			<p>
				If you have any suggestions, bug reports, or annoyances
				<br />please report them to our
				<a
					href="https://github.com/clivern/peanut/issues"
					target="_blank"
					rel="noopener"
					>issue tracker</a
				>.
			</p>
			<br /><br />
			<small>
				<br />
				Made with
				<span class="icon has-text-danger"><i class="fas fa-heart"></i></span>
				by
				<a href="https://github.com/clivern" target="_blank" rel="noopener"
					>Clivern</a
				><br />
			</small>
		</div>
	</div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
	color: #42b983;
}
</style>

<script>
export default {
	name: "home",

	data() {
		return {};
	},

	methods: {},

	mounted() {
		this.$emit("refresh-state");
	},
};
</script>
