<!-- @format -->

<template>
	<div class="columns is-desktop is-centered">
		<div class="column is-4">
			<b-field label="API Key">
				<b-input
					v-model="form.api_key"
					placeholder="76a97318-2560-4451-a90d-5ba63126d055"
					rounded
				></b-input>
			</b-field>

			<div class="field">
				<p class="control">
					<b-button
						type="submit is-danger is-light"
						v-bind:disabled="form.button_disabled"
						@click="loginEvent"
						>Login</b-button
					>
				</p>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: "LoginPage",
	data() {
		return {
			form: {
				api_key: "",
				button_disabled: false,
			},
		};
	},
	methods: {
		loginEvent() {
			this.form.button_disabled = true;
			this.$store
				.dispatch("service/getServicesAction", {
					apiKey: this.form.api_key,
				})
				.then(
					() => {
						this.$buefy.toast.open({
							message: "You logged in successfully",
							type: "is-success",
						});
						localStorage.setItem("x_api_key", this.form.api_key);
						this.$router.push("/");
					},
					(err) => {
						if (err.response.data.errorMessage) {
							this.$buefy.toast.open({
								message: err.response.data.errorMessage,
								type: "is-danger",
							});
						} else {
							this.$buefy.toast.open({
								message: "Oops! invalid api key",
								type: "is-danger",
							});
						}
						this.form.button_disabled = false;
					}
				);
		},
	},
	mounted() {
		this.loading();
	},
};
</script>
