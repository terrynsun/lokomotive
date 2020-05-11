variable "cluster_name" {
  type        = string
  description = "Unique cluster name (prepended to dns_zone)"
}

variable "controllers_public_ipv4" {
  type        = list(string)
  description = "Public IPv4 addresses of all the controllers in the cluster"
}

variable "controllers_private_ipv4" {
  type        = list(string)
  description = "Private IPv4 addresses of all the controllers in the cluster"
}

variable "dns_zone" {
  type        = string
  description = "Route 53 zone name (e.g. example.com)"
}

locals {
  api_external_fqdn = format("%s.%s.", var.cluster_name, var.dns_zone)
  api_fqdn          = format("%s-private.%s.", var.cluster_name, var.dns_zone)
  etcd_fqdn         = [for i, d in var.controllers_private_ipv4 : format("%s-etcd%d.%s.", var.cluster_name, i, var.dns_zone)]

  dns_entries = concat(
    [
      # apiserver public
      {
        name    = local.api_external_fqdn,
        type    = "A",
        ttl     = 300,
        records = var.controllers_public_ipv4
      },
      # apiserver private
      {
        name    = local.api_fqdn,
        type    = "A",
        ttl     = 300,
        records = var.controllers_private_ipv4
      },
    ],
    # etcd
    [
      for index, i in var.controllers_private_ipv4 :
      {
        name    = local.etcd_fqdn[index],
        type    = "A",
        ttl     = 300,
        records = [i],
      }
    ],
  )
}

output "entries" {
  value = local.dns_entries
}
