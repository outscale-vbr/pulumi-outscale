// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package outscale

import (
	"github.com/outscale-vbr/pulumi-outscale/sdk/go/outscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiAccessPolicy(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupApiAccessPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiAccessPolicyResult
	err := ctx.Invoke("outscale:index/getApiAccessPolicy:getApiAccessPolicy", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getApiAccessPolicy.
type LookupApiAccessPolicyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                            string `pulumi:"id"`
	MaxAccessKeyExpirationSeconds int    `pulumi:"maxAccessKeyExpirationSeconds"`
	RequestId                     string `pulumi:"requestId"`
	RequireTrustedEnv             bool   `pulumi:"requireTrustedEnv"`
}
