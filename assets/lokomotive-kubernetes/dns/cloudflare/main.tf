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
  description = "Cloudflare zone name (e.g. example.com)"
}

data "cloudflare_zones" "selected" {
  filter {
    name   = var.dns_zone
    status = "active"
    paused = false
  }
}

resource "cloudflare_record" "apiserver_public" {
  count = length(var.controllers_public_ipv4)

  zone_id = lookup(data.cloudflare_zones.selected.zones[0], "id")
  name    = format("%s.%s.", var.cluster_name, var.dns_zone)
  type    = "A"
  ttl     = 300
  value   = var.controllers_public_ipv4[count.index]
}

resource "cloudflare_record" "apiserver_private" {
  count = length(var.controllers_private_ipv4)

  zone_id = lookup(data.cloudflare_zones.selected.zones[0], "id")
  name    = format("%s-private.%s.", var.cluster_name, var.dns_zone)
  type    = "A"
  ttl     = 300
  value   = var.controllers_private_ipv4[count.index]
}

resource "cloudflare_record" "etcd" {
  count = length(var.controllers_private_ipv4)

  zone_id = lookup(data.cloudflare_zones.selected.zones[0], "id")
  name    = format("%s-etcd%d.%s.", var.cluster_name, count.index, var.dns_zone)
  type    = "A"
  ttl     = 300
  value   = var.controllers_private_ipv4[count.index]
}
