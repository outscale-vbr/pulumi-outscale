package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	outscale "github.com/outscale/terraform-provider-outscale/outscale"
	"github.com/outscale-vbr/pulumi-outscale/provider/pkg/version"
)

func NewProvider() tfpf.Provider {
	prov := outscale.New(version.Version)

	return prov()
}