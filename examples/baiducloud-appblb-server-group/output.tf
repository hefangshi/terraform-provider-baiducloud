output "appblb-sg" {
  value = "${baiducloud_appblb_server_group.my-appblb-sg}"
}

output "appblb-sgs" {
  value = "${data.baiducloud_appblb_server_groups.default.server_groups}"
}