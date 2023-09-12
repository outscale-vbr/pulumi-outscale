// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/outscale-vbr/pulumi-outscale/sdk/go/outscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The Access Key ID for API operations
func GetAccessKeyId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "outscale:accessKeyId")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "OUTSCALE_ACCESSKEYID"); d != nil {
		value = d.(string)
	}
	return value
}
func GetEndpoints(ctx *pulumi.Context) string {
	return config.Get(ctx, "outscale:endpoints")
}

// The Region for API operations.
func GetRegion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "outscale:region")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "OUTSCALE_REGION"); d != nil {
		value = d.(string)
	}
	return value
}

// The Secret Key ID for API operations.
func GetSecretKeyId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "outscale:secretKeyId")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "OUTSCALE_SECRETKEYID"); d != nil {
		value = d.(string)
	}
	return value
}

// The path to your x509 cert
func GetX509CertPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "outscale:x509CertPath")
}

// The path to your x509 key
func GetX509KeyPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "outscale:x509KeyPath")
}