# Configure the Junos Provider
provider "junos" {
  ip         = "192.168.1.1"
  username   = "lab"
  password   = "lab123"
}

# Configure an prefix list
resource junos_policyoptions_prefix_list "PE1_plist" {
  name   = "import-GCP-DX"
  prefix = [
            # The prefixes from project Interana
            "10.247.0.0/24", 
            "10.248.0.0/24",
            "11.123.0.0/28",
            # The prefixes from PoC airflow
            "10.128.0.0/29",
            # The prefixes from GKE project
            "10.100.0.0/16"
           ]
}
