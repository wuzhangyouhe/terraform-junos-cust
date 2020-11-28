# Configure the Junos Provider
provider "junos" {
  ip         = "192.168.1.1"
  username   = "lab"
  password   = "lab123"
}

# Configure an prefix list
resource junos_policyoptions_prefix_list "PE1_project_1" {
  name   = "import-GCP-DX"
  prefix = ["10.233.0.0/24"]
}
