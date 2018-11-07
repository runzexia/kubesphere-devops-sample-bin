package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/runzexia/kubesphere-devops-sample-bin/qingcloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: qingcloud.Provider,
	})
}
