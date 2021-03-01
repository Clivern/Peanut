<!-- @format -->

<template>
	<section>
		<div class="columns is-desktop is-centered">
			<div class="column"></div>
			<div class="column is-three-quarters">
				<b-table
					:data="data"
					ref="table"
					paginated
					per-page="20"
					:opened-detailed="defaultOpenedDetails"
					detailed
					detail-key="id"
					:detail-transition="transitionName"
					:show-detail-icon="showDetailIcon"
					aria-next-label="Next page"
					aria-previous-label="Previous page"
					aria-page-label="Page"
					aria-current-label="Current page"
				>
					<b-table-column
						field="service"
						label="Service"
						centered
						v-slot="props"
					>
						{{ services[props.row.service] }}
					</b-table-column>

					<b-table-column
						field="address"
						label="Address"
						centered
						v-slot="props"
					>
						<span class="tag is-success is-light">
							{{ props.row.configs.address }}
						</span>
					</b-table-column>

					<b-table-column field="id" label="UUID" centered v-slot="props">
						<span class="tag is-warning is-light">{{ props.row.id }}</span>
					</b-table-column>

					<b-table-column
						field="deleteAfter"
						label="Delete After"
						centered
						v-slot="props"
					>
						<span class="tag is-success is-light">
							<template v-if="props.row.deleteAfter != ''">
								{{ props.row.deleteAfter }}
							</template>
							<template v-else>N/A</template>
						</span>
					</b-table-column>

					<b-table-column
						field="createdAt"
						label="Created at"
						centered
						v-slot="props"
					>
						<span class="tag is-danger is-light">
							{{ new Date(props.row.createdAt).toLocaleDateString() }}
						</span>
					</b-table-column>

					<b-table-column label="Actions" centered v-slot="props">
						<template v-if="props.row.service == 'grafana'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>

						<template v-if="props.row.service == 'graphite'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>

						<template v-if="props.row.service == 'prometheus'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>

						<template v-if="props.row.service == 'zipkin'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>

						<template v-if="props.row.service == 'mailhog'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>

						<template v-if="props.row.service == 'jaeger'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>

						<template v-if="props.row.service == 'rabbitmq'">
							<b-button
								size="is-small"
								type="is-link is-success is-light"
								@click="openServiceDashboard(props.row)"
								>Visit</b-button
							>
							-
						</template>
						<b-button
							size="is-small"
							type="is-link is-danger is-light"
							@click="deleteService(props.row.id)"
							>Delete</b-button
						>
					</b-table-column>

					<template #detail="props">
						<code>{{ props.row.configs }}</code>
					</template>

					<td slot="empty" colspan="6">No records found.</td>
				</b-table>
			</div>
			<div class="column"></div>
		</div>
	</section>
</template>

<script>
export default {
	data() {
		return {
			data: [],
			defaultOpenedDetails: [1],
			showDetailIcon: true,
			useTransition: false,

			// Loader
			loader: {
				isFullPage: true,
				ref: null,
			},

			services: {
				mysql: "MySQL",
				mariadb: "MariaDB",
				postgresql: "PostgreSQL",
				redis: "Redis",
				etcd: "Etcd",
				grafana: "Grafana",
				elasticsearch: "Elasticsearch",
				mongodb: "MongoDB",
				graphite: "Graphite",
				prometheus: "Prometheus",
				zipkin: "Zipkin",
				memcached: "Memcached",
				mailhog: "Mailhog",
				jaeger: "Jaeger",
				rabbitmq: "RabbitMQ",
			},
		};
	},
	computed: {
		transitionName() {
			if (this.useTransition) {
				return "fade";
			}
		},
	},
	methods: {
		loading() {
			this.loader.ref = this.$buefy.loading.open({
				container: this.loader.isFullPage ? null : this.$refs.element.$el,
			});
		},
		loadInitialState() {
			this.loading();
			this.$store.dispatch("service/getServicesAction", {}).then(
				() => {
					let data = this.$store.getters["service/getServicesResult"];

					if (data.services) {
						this.data = data.services;
					} else {
						this.data = [];
					}
					this.loader.ref.close();
				},
				(err) => {
					this.$buefy.toast.open({
						message: err.response.data.errorMessage,
						type: "is-danger is-light",
					});
					this.loader.ref.close();
				}
			);
		},
		openServiceDashboard(data) {
			if (data.service == "grafana") {
				window.open("//" + data.configs.address + ":" + data.configs.port);
			}

			if (data.service == "graphite") {
				window.open("//" + data.configs.address + ":" + data.configs.webPort);
			}

			if (data.service == "prometheus") {
				window.open("//" + data.configs.address + ":" + data.configs.port);
			}

			if (data.service == "zipkin") {
				window.open("//" + data.configs.address + ":" + data.configs.port);
			}

			if (data.service == "mailhog") {
				window.open("//" + data.configs.address + ":" + data.configs.httpPort);
			}

			if (data.service == "jaeger") {
				window.open("//" + data.configs.address + ":" + data.configs.httpPort2);
			}

			if (data.service == "rabbitmq") {
				window.open(
					"//" + data.configs.address + ":" + data.configs.dashboardPort
				);
			}
		},
		deleteService(serviceId) {
			this.$buefy.dialog.confirm({
				message: "Are you sure?",
				onConfirm: () => {
					this.$store
						.dispatch("service/deleteServiceAction", {
							serviceId: serviceId,
						})
						.then(
							() => {
								this.loading();

								let data =
									this.$store.getters["service/getDeleteServiceResult"];

								var timer = setInterval(() => {
									this.$store
										.dispatch("service/getJobAction", {
											serviceId: data.service,
											jobId: data.id,
										})
										.then(
											() => {},
											() => {
												// If job not found anymore, means service got deleted
												// this behavior to clear etcd cluster
												this.$buefy.toast.open({
													message: "Service delete successfully!",
													type: "is-success",
												});
												this.loader.ref.close();
												this.loadInitialState();
												clearInterval(timer);
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
							}
						);
				},
			});
		},
	},
	mounted() {
		this.loadInitialState();
	},
};
</script>
