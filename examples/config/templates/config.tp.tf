
locals {
  # Read the root config file
  config = yamldecode(file("{{ .Root }}/config.yaml"))
  # Create the cluster config for a specific cluster based on name and environment
  cluster_config = merge(
    # Set any cluster defaults
    {
      name         = "${var.cluster}-${var.environment}",
      loadbalancer = []
    },
    # Get the cluster config from the config YAML file, by looking for the
    # cluster name and environment, e.g.
    # clusters:
    #   cluster_name:
    #     environment: {}
    lookup(lookup(local.config.clusters, var.cluster), var.environment)
  )
}
