// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package outscale

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/outscale-vbr/pulumi-outscale/provider/pkg/version"
	//outscale "github.com/outscale/terraform-provider-outscale/outscale"
	tfpfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	//"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	shim 	"github.com/outscale/terraform-provider-outscale/shim"
	//shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	//"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	outscalePkg = "outscale"
	// modules:
	outscaleMod = "index" // the outscale module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// outscaleMember manufactures a type token for the Scaleway package and the given module and type.
func outscaleMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(outscalePkg + ":" + mod + ":" + mem)
}

// outscaleType manufactures a type token for the Scaleway package and the given module and type.
func outscaleType(mod string, typ string) tokens.Type {
	return tokens.Type(outscaleMember(mod, typ))
}

// outscaleDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Scaleway package and names the file by simply lower casing the data
// source's first character.
func outscaleDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return outscaleMember(mod+"/"+fn, res)
}

// outscaleResource manufactures a standard resource token given a module and resource name.
// It automatically uses the Scaleway package and names the file by simply lower casing the resource's
// first character.
func outscaleResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return outscaleType(mod+"/"+fn, res)
}

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider

	
	p := shim.ProviderShim()

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "outscale",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing outscale cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "outscale", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/outscale-vbr/pulumi-outscale",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "outscale-vbr",
		Config:    map[string]*tfbridge.SchemaInfo{
			"access_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OSC_ACCESS_KEY"},
				},
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OSC_SECRET_KEY"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OSC_DEFAULT_REGION"},
				},
			},

		},

		Resources:            map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			"outscale_access_key": {Tok: outscaleResource(outscaleMod, "AccessKey")},
			"outscale_api_access_policy": {Tok: outscaleResource(outscaleMod, "ApiAccessPolicy")},
			"outscale_api_access_rule": {Tok: outscaleResource(outscaleMod, "ApiAccessRule")},
			"outscale_ca": {Tok: outscaleResource(outscaleMod, "Ca")},
			"outscale_client_gateway": {Tok: outscaleResource(outscaleMod, "ClientGateway")},
			"outscale_dhcp_option": {Tok: outscaleResource(outscaleMod, "DhcpOption")},
			"outscale_flexible_gpu": {Tok: outscaleResource(outscaleMod, "FlexibleGpu")},
			"outscale_flexible_gpu_link": {Tok: outscaleResource(outscaleMod, "FlexibleGpuLink")},
			"outscale_image": {Tok: outscaleResource(outscaleMod, "Image")},
			"outscale_image_export_task": {Tok: outscaleResource(outscaleMod, "ImageExportTask")},
			"outscale_image_launch_permission": {Tok: outscaleResource(outscaleMod, "ImageLaunchPermission")},
			"outscale_internet_service": {Tok: outscaleResource(outscaleMod, "InternetService")},
			"outscale_internet_service_link": {Tok: outscaleResource(outscaleMod, "InternetServiceLink")},
			"outscale_keypair": {Tok: outscaleResource(outscaleMod, "Keypair" )},
			"outscale_load_balancer": {Tok: outscaleResource(outscaleMod, "LoadBalancer")},
			"outscale_load_balancer_attributes": {Tok: outscaleResource(outscaleMod, "LoadBalancerAttributes")},
			"outscale_load_balancer_lister_rules": {Tok: outscaleResource(outscaleMod, "LoadBalancerListerRules")},
			"outscale_load_balancer_policy": {Tok: outscaleResource(outscaleMod, "LoadBalancerPolicy")},
			"outscale_load_balancer_vms": {Tok: outscaleResource(outscaleMod, "LoadBalancerVms" )},
			"outscale_nat_service": {Tok: outscaleResource(outscaleMod, "NatService")},
			"outscale_net": {Tok: outscaleResource(outscaleMod, "Net")},
			"outscale_net_access_point": {Tok: outscaleResource(outscaleMod, "NetAccessPoint")},
			"outscale_net_attributes": {Tok: outscaleResource(outscaleMod, "NetAttributes")},
			"outscale_net_peering": {Tok: outscaleResource(outscaleMod, "NetPeering")},
			"outscale_net_peering_acceptation": {Tok: outscaleResource(outscaleMod, "NetPeeringAcception")},
			"outscale_nic": {Tok: outscaleResource(outscaleMod, "Nic")},
			"outscale_nic_link": {Tok: outscaleResource(outscaleMod, "NicLink")},
			"outscale_nic_private_ip": {Tok: outscaleResource(outscaleMod, "NicPrivateIp")},
			"outscale_public_ip": {Tok: outscaleResource(outscaleMod, "PublicIp")},
			"outscale_public_ip_link": {Tok: outscaleResource(outscaleMod, "PublicIpLink")},
			"outscale_route": {Tok: outscaleResource(outscaleMod, "Route")},
			"outscale_route_table": {Tok: outscaleResource(outscaleMod, "RouteTable")},
			"outscale_route_table_link": {Tok: outscaleResource(outscaleMod, "RouteTableLink")},
			"outscale_security_group": {Tok: outscaleResource(outscaleMod, "SecurityGroup")},
			"outscale_security_group_rule": {Tok: outscaleResource(outscaleMod, "SecurityGroupRule")},
			"outscale_server_certificate": {Tok: outscaleResource(outscaleMod, "ServerCertificate")},
			"outscale_snapshot": {Tok: outscaleResource(outscaleMod, "Snapshot")},
			"outscale_snapshot_attributes": {Tok: outscaleResource(outscaleMod, "SnapshotAttributes")},
			"outscale_snapshot_export_task": {Tok: outscaleResource(outscaleMod, "SnapshotExportTask")},
			"outscale_subnet": {Tok: outscaleResource(outscaleMod, "Subnet")},
			"outscale_virtual_gateway": {Tok: outscaleResource(outscaleMod, "VirtualGateway")},
			"outscale_virtual_gateway_link": {Tok: outscaleResource(outscaleMod, "VirtualGatewayLink")},
			"outscale_virtual_gateway_route_propagation": {Tok: outscaleResource(outscaleMod, "VirtualGatewayRoutePropagation")},
			"outscale_vm": {Tok: outscaleResource(outscaleMod, "Vm")},
			"outscale_volume": {Tok: outscaleResource(outscaleMod, "Volume")},
			"outscale_volumes_link": {Tok: outscaleResource(outscaleMod, "VolumesLink")},
			"outscale_vpn_connection": {Tok: outscaleResource(outscaleMod, "VpnConnection")},
			"outscale_vpn_connection_route": {Tok: outscaleResource(outscaleMod, "VpnConnectionRoute")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"outscale_access_key": {Tok: outscaleDataSource(outscaleMod, "getAccessKey")},
			"outscale_api_access_policy": {Tok: outscaleDataSource(outscaleMod, "getApiAccessPolicy")},
			"outscale_api_access_rule": {Tok: outscaleDataSource(outscaleMod, "getApiAccessRule")},
			"outscale_ca": {Tok: outscaleDataSource(outscaleMod, "getCa")},
			"outscale_client_gateway": {Tok: outscaleDataSource(outscaleMod, "getClientGateway")},
			"outscale_dhcp_option": {Tok: outscaleDataSource(outscaleMod, "getDhcpOption")},
			"outscale_flexible_gpu": {Tok: outscaleDataSource(outscaleMod, "getFlexibleGpu")},
			"outscale_image": {Tok: outscaleDataSource(outscaleMod, "getImage")},
			"outscale_image_export_task": {Tok: outscaleDataSource(outscaleMod, "getImageExportTask")},
			"outscale_image_launch_permission": {Tok: outscaleDataSource(outscaleMod, "getImageLaunchPermission")},
			"outscale_internet_service": {Tok: outscaleDataSource(outscaleMod, "getInternetService")},
			"outscale_internet_service_link": {Tok: outscaleDataSource(outscaleMod, "getInternetServiceLink")},
			"outscale_keypair": {Tok: outscaleDataSource(outscaleMod, "getKeypair")},
			"outscale_load_balancer": {Tok: outscaleDataSource(outscaleMod, "getLoadBalancer")},
			"outscale_load_balancer_attributes": {Tok: outscaleDataSource(outscaleMod, "getLoadBalancerAttributes")},
			"outscale_load_balancer_listener_rule": {Tok: outscaleDataSource(outscaleMod, "getLoadBalancerListenerRule")},
			"ourscale_load_balancer_vms": {Tok: outscaleDataSource(outscaleMod, "getLoadBalancerVms")},
			"outscale_nat_service": {Tok: outscaleDataSource(outscaleMod, "getNatService")},
			"outscale_net": {Tok: outscaleDataSource(outscaleMod, "getNet")},
			"outscale_net_access_point": {Tok: outscaleDataSource(outscaleMod, "getNetAccessPoint")},
			"outscale_net_attributes": {Tok: outscaleDataSource(outscaleMod, "getNetAttributes")},
			"outscale_net_peering": {Tok: outscaleDataSource(outscaleMod, "getNetPeering")},
			"outscale_net_peering_acceptation": {Tok: outscaleDataSource(outscaleMod, "getNetPeeringAcceptation")},
			"outscale_nic": {Tok: outscaleDataSource(outscaleMod, "getNic")},
			"outscale_nic_link": {Tok: outscaleDataSource(outscaleMod, "getNicLink")},
			"outscale_nic_private_ip": {Tok: outscaleDataSource(outscaleMod, "getNicPrivateIp")},
			"outscale_public_ip": {Tok: outscaleDataSource(outscaleMod, "getPublicIp")},
			"outscale_public_ip_link": {Tok: outscaleDataSource(outscaleMod, "getPublicIpLink")},
			"outscale_route": {Tok: outscaleDataSource(outscaleMod, "getRoute" )},
			"outscale_route_table": {Tok: outscaleDataSource(outscaleMod, "getRouteTable")},
			"outscale_route_table_link": {Tok: outscaleDataSource(outscaleMod, "getRouteTableLink")},
			"outscale_security_group": {Tok: outscaleDataSource(outscaleMod, "getSecurityGroup")},
			"outscale_security_group_rule": {Tok: outscaleDataSource(outscaleMod, "getSecurityGroupRule")},
			"outscale_server_certificate": {Tok: outscaleDataSource(outscaleMod, "getServerCertificate")},
			"outscale_snapshot": {Tok: outscaleDataSource(outscaleMod, "getSnapshot")},
			"outscale_snapshot_attributes": {Tok: outscaleDataSource(outscaleMod,"getSnapshotAttributes")},
			"outscale_subnet": {Tok: outscaleDataSource(outscaleMod, "getSubnet")},
			"outscale_virtual_gateway": {Tok: outscaleDataSource(outscaleMod, "getVirtualGateway")},
			"outscale_virtual_gateway_link": {Tok: outscaleDataSource(outscaleMod, "getVirtualGatewayLink")},
			"outscale_virtual_gateway_route_propagation": {Tok: outscaleDataSource(outscaleMod, "getVirtualGatewayRoutePropagation")},
			"outscale_vm": {Tok: outscaleDataSource(outscaleMod, "getVm")},
			"outscale_volume": {Tok: outscaleDataSource(outscaleMod, "getVolume")},
			"outscale_volumes_link": {Tok: outscaleDataSource(outscaleMod, "getVolumesLink")},
			"outscale_vpn_connection": {Tok: outscaleDataSource(outscaleMod, "getVpnConnection")},
			"outscale_vpn_connection_route": {Tok: outscaleDataSource(outscaleMod, "getVpnConnectionRoute")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", outscalePkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				outscalePkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference: https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	return tfpfbridge.ProviderInfo{
		ProviderInfo: prov,
		NewProvider:  shim.NewProvider,
	}
}
