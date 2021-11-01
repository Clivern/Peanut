<!-- @format -->

<template>
	<div class="columns is-desktop is-centered">
		<div class="column is-4">
			<br /><br />
			<section>
				<b-field label="Service Type">
					<b-select
						placeholder="Select a service type"
						expanded
						v-model="form.type"
						@input="onServiceChange('true')"
					>
						<option value="mysql">MySQL</option>
						<option value="mariadb">MariaDB</option>
						<option value="postgresql">PostgreSQL</option>
						<option value="redis">Redis</option>
						<option value="etcd">Etcd</option>
						<option value="grafana">Grafana</option>
						<option value="elasticsearch">Elasticsearch</option>
						<option value="mongodb">MongoDB</option>
						<option value="graphite">Graphite</option>
						<option value="prometheus">Prometheus</option>
						<option value="zipkin">Zipkin</option>
						<option value="memcached">Memcached</option>
						<option value="mailhog">Mailhog</option>
						<option value="jaeger">Jaeger</option>
						<option value="rabbitmq">RabbitMQ</option>
						<option value="consul">Consul</option>
						<option value="vault">Vault</option>
						<option value="cassandra">Cassandra</option>
						<option value="minio">Minio</option>
						<option value="registry">Registry</option>
						<option value="ghost">Ghost</option>
						<option value="httpbin">Httpbin</option>
						<option value="etherpad">Etherpad</option>
					</b-select>
				</b-field>

				<b-field label="Version">
					<b-select
						placeholder="Select a service version"
						expanded
						v-model="form.version.selected"
					>
						<option
							v-for="option in form.version.options"
							v-bind:key="option.value"
						>
							{{ option.text }}
						</option>
					</b-select>
					<p class="control">
						<b-button
							@click="onServiceChange('false')"
							label="Get Latest Versions"
							type="is-warning"
						/>
					</p>
				</b-field>

				<b-field label="Delete After">
					<b-input
						type="number"
						v-model="form.delete_after_period"
						placeholder="0"
					></b-input>
					<b-select
						placeholder="Period (disabled by default)"
						expanded
						v-model="form.delete_after_type"
					>
						<option value="" selected>Disabled</option>
						<option value="sec">Seconds</option>
						<option value="min">Minutes</option>
						<option value="hours">Hours</option>
						<option value="days">Days</option>
					</b-select>
				</b-field>
				<template v-if="form.type == 'redis'">
					<b-field label="Redis Password">
						<b-input value="" v-model="form.configs.redis.password"></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'mysql'">
					<b-field label="MySQL Root Password">
						<b-input
							value=""
							v-model="form.configs.mysql.rootPassword"
						></b-input>
					</b-field>
					<b-field label="MySQL Database">
						<b-input value="" v-model="form.configs.mysql.database"></b-input>
					</b-field>
					<b-field label="MySQL Username">
						<b-input value="" v-model="form.configs.mysql.username"></b-input>
					</b-field>
					<b-field label="MySQL Password">
						<b-input value="" v-model="form.configs.mysql.password"></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'mariadb'">
					<b-field label="MariaDB Root Password">
						<b-input
							value=""
							v-model="form.configs.mariadb.rootPassword"
						></b-input>
					</b-field>
					<b-field label="MariaDB Database">
						<b-input value="" v-model="form.configs.mariadb.database"></b-input>
					</b-field>
					<b-field label="MariaDB Username">
						<b-input value="" v-model="form.configs.mariadb.username"></b-input>
					</b-field>
					<b-field label="MariaDB Password">
						<b-input value="" v-model="form.configs.mariadb.password"></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'postgresql'">
					<b-field label="PostgreSQL Database">
						<b-input
							value=""
							v-model="form.configs.postgresql.database"
						></b-input>
					</b-field>
					<b-field label="PostgreSQL Username">
						<b-input
							value=""
							v-model="form.configs.postgresql.username"
						></b-input>
					</b-field>
					<b-field label="PostgreSQL Password">
						<b-input
							value=""
							v-model="form.configs.postgresql.password"
						></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'grafana'">
					<b-field label="Login Username">
						<b-input value="" v-model="form.configs.grafana.username"></b-input>
					</b-field>
					<b-field label="Login Password">
						<b-input value="" v-model="form.configs.grafana.password"></b-input>
					</b-field>
					<b-field label="Allow Signup">
						<b-select
							placeholder="Allow Signup"
							expanded
							v-model="form.configs.grafana.allowSignup"
						>
							<option value="false">off</option>
							<option value="true">on</option>
						</b-select>
					</b-field>
					<b-field label="Anonymous Access">
						<b-select
							placeholder="Anonymous Access"
							expanded
							v-model="form.configs.grafana.anonymousAccess"
						>
							<option value="true">on</option>
							<option value="false">off</option>
						</b-select>
					</b-field>
				</template>

				<template v-if="form.type == 'mongodb'">
					<b-field label="MongoDB Database">
						<b-input value="" v-model="form.configs.mongodb.database"></b-input>
					</b-field>
					<b-field label="MongoDB Username">
						<b-input value="" v-model="form.configs.mongodb.username"></b-input>
					</b-field>
					<b-field label="MongoDB Password">
						<b-input value="" v-model="form.configs.mongodb.password"></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'vault'">
					<b-field label="Vault Root Token">
						<b-input value="" v-model="form.configs.vault.token"></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'prometheus'">
					<b-field label="YAML Configs (please ensure it is valid)">
						<b-input
							type="textarea"
							rows="18"
							v-model="form.configs.prometheus.configsBase64Decoded"
						></b-input>
					</b-field>
				</template>

				<template v-if="form.type == 'minio'">
					<b-field label="Minio Username">
						<b-input value="" v-model="form.configs.minio.username"></b-input>
					</b-field>
					<b-field label="Minio Password">
						<b-input value="" v-model="form.configs.minio.password"></b-input>
					</b-field>
				</template>

				<br />
				<div class="field">
					<p class="control">
						<b-button
							type="is-danger is-light"
							v-bind:disabled="form.button_disabled"
							@click="deployEvent"
							>Deploy</b-button
						>
					</p>
				</div>
			</section>
		</div>
	</div>
</template>

<script>
export default {
	name: "deploy",
	data() {
		return {
			form: {
				button_disabled: false,
				type: "",
				delete_after_type: "",
				delete_after_period: "",

				version: {
					selected: "",
					options: [{ text: "Default", value: "" }],
				},

				configs: {
					redis: {
						password: "",
					},
					consul: {},
					registry: {},
					ghost: {},
					httpbin: {},
					etherpad: {},
					vault: {
						token: "peanut",
					},
					mysql: {
						rootPassword: "root",
						database: "peanut",
						username: "peanut",
						password: "secret",
					},
					mariadb: {
						rootPassword: "root",
						database: "peanut",
						username: "peanut",
						password: "secret",
					},
					postgresql: {
						database: "peanut",
						username: "peanut",
						password: "secret",
					},
					grafana: {
						username: "admin",
						password: "secret",
						anonymousAccess: "true",
						allowSignup: "false",
					},
					mongodb: {
						database: "peanut",
						username: "peanut",
						password: "secret",
					},
					cassandra: {},
					prometheus: {
						configsBase64Encoded: "",
						configsBase64Decoded: this.b64_to_utf8(
							"Z2xvYmFsOgogIGV2YWx1YXRpb25faW50ZXJ2YWw6IDE1cwogIHNjcmFwZV9pbnRlcnZhbDogMTVzCnJ1bGVfZmlsZXM6IH4Kc2NyYXBlX2NvbmZpZ3M6CiAgLQogICAgam9iX25hbWU6IHByb21ldGhldXMKICAgIHNjcmFwZV9pbnRlcnZhbDogNXMKICAgIHN0YXRpY19jb25maWdzOgogICAgICAtCiAgICAgICAgdGFyZ2V0czoKICAgICAgICAgIC0gImxvY2FsaG9zdDo5MDkwIg=="
						),
					},
					minio: {
						username: "admin12345678",
						password: "admin12345678",
					},
				},
			},
			// Loader
			loader: {
				isFullPage: true,
				ref: null,
			},
		};
	},
	methods: {
		loading() {
			this.loader.ref = this.$buefy.loading.open({
				container: this.loader.isFullPage ? null : this.$refs.element.$el,
			});
		},

		utf8_to_b64(str) {
			return window.btoa(unescape(encodeURIComponent(str)));
		},

		b64_to_utf8(str) {
			return decodeURIComponent(escape(window.atob(str)));
		},

		onServiceChange(cache = "true") {
			if (this.form.type == "") {
				this.form.version = {
					selected: "",
					options: [{ text: "Default", value: "" }],
				};
				return;
			}

			this.loading();

			this.$store
				.dispatch("service/getServiceTagsAction", {
					serviceType: this.form.type,
					cache: cache,
				})
				.then(
					() => {
						let data = this.$store.getters["service/getServiceTagsResult"];

						this.form.version = {
							selected: "",
							options: [{ text: "Default", value: "" }],
						};

						for (var i = 0; i < data.tags.length; i++) {
							this.form.version.options.push({
								text: data.tags[i],
								value: data.tags[i],
							});
						}

						this.loader.ref.close();
					},
					(err) => {
						if (err.response.data.errorMessage) {
							this.$buefy.toast.open({
								message: err.response.data.errorMessage,
								type: "is-danger is-light",
							});
						} else {
							this.$buefy.toast.open({
								message: "Error status code: " + err.response.status,
								type: "is-danger is-light",
							});
						}
						this.loader.ref.close();
					}
				);
		},

		deployEvent() {
			this.loading();
			this.form.button_disabled = true;

			let deleteAfter = "";

			if (
				this.form.delete_after_type != "" &&
				this.form.delete_after_period > 0
			) {
				deleteAfter =
					this.form.delete_after_period + this.form.delete_after_type;
			}

			let configs = {};

			if (this.form.type == "redis") {
				configs = this.form.configs.redis;
			} else if (this.form.type == "mysql") {
				configs = this.form.configs.mysql;
			} else if (this.form.type == "mariadb") {
				configs = this.form.configs.mariadb;
			} else if (this.form.type == "postgresql") {
				configs = this.form.configs.postgresql;
			} else if (this.form.type == "grafana") {
				configs = this.form.configs.grafana;
			} else if (this.form.type == "mongodb") {
				configs = this.form.configs.mongodb;
			} else if (this.form.type == "prometheus") {
				configs = {
					configsBase64Encoded: this.utf8_to_b64(
						this.form.configs.prometheus.configsBase64Decoded
					),
				};
			} else if (this.form.type == "consul") {
				configs = this.form.configs.consul;
			} else if (this.form.type == "vault") {
				configs = this.form.configs.vault;
			} else if (this.form.type == "cassandra") {
				configs = this.form.configs.cassandra;
			} else if (this.form.type == "minio") {
				configs = this.form.configs.minio;
			} else if (this.form.type == "registry") {
				configs = this.form.configs.registry;
			} else if (this.form.type == "ghost") {
				configs = this.form.configs.ghost;
			} else if (this.form.type == "httpbin") {
				configs = this.form.configs.httpbin;
			} else if (this.form.type == "etherpad") {
				configs = this.form.configs.etherpad;
			}

			this.$store
				.dispatch("service/deployServiceAction", {
					service: this.form.type,
					deleteAfter: deleteAfter,
					configs: configs,
					version: this.form.version.selected,
				})
				.then(
					() => {
						let data = this.$store.getters["service/getDeployServiceResult"];

						var timer = setInterval(() => {
							this.$store
								.dispatch("service/getJobAction", {
									serviceId: data.service,
									jobId: data.id,
								})
								.then(
									() => {
										let response = this.$store.getters["service/getJobResult"];

										if (response.status == "SUCCESS") {
											this.$buefy.toast.open({
												message: "Service deployed successfully!",
												type: "is-success",
											});

											this.form.button_disabled = false;
											this.loader.ref.close();
											clearInterval(timer);
											this.$router.push("/services");
										} else if (response.status == "FAILED") {
											this.$buefy.toast.open({
												message: "Service deployment failed",
												type: "is-danger",
											});

											this.form.button_disabled = false;
											this.loader.ref.close();
											clearInterval(timer);
										}
									},
									(err) => {
										if (err.response.data.errorMessage) {
											this.$buefy.toast.open({
												message: err.response.data.errorMessage,
												type: "is-danger",
											});
										} else {
											this.$buefy.toast.open({
												message: "Error occurred: " + err.response.status,
												type: "is-danger",
											});
										}
										this.form.button_disabled = false;
									}
								);
						}, 3000);
					},
					(err) => {
						if (err.response.data.errorMessage) {
							this.$buefy.toast.open({
								message: err.response.data.errorMessage,
								type: "is-danger is-light",
							});
						} else {
							this.$buefy.toast.open({
								message: "Error status code: " + err.response.status,
								type: "is-danger is-light",
							});
						}
						this.form.button_disabled = false;
						this.loader.ref.close();
					}
				);
		},
	},
	mounted() {},
};
</script>
