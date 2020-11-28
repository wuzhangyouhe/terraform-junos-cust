#!/usr/bin/env bash

tfPath="${HOME}/.terraform.d/plugins/registry.local/jeremmfr/junos/twitter"
mkdir -p $tfPath
go build -o $tfPath/terraform-provider-junos_twitter