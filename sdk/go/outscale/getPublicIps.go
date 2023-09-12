// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package outscale

import (
	"context"
	"reflect"

	"github.com/outscale-vbr/pulumi-outscale/sdk/go/outscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func GetPublicIps(ctx *pulumi.Context, args *GetPublicIpsArgs, opts ...pulumi.InvokeOption) (*GetPublicIpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPublicIpsResult
	err := ctx.Invoke("outscale:index/getPublicIps:getPublicIps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicIps.
type GetPublicIpsArgs struct {
	Filters []GetPublicIpsFilter `pulumi:"filters"`
}

// A collection of values returned by getPublicIps.
type GetPublicIpsResult struct {
	Filters []GetPublicIpsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                 `pulumi:"id"`
	PublicIps []GetPublicIpsPublicIp `pulumi:"publicIps"`
	RequestId string                 `pulumi:"requestId"`
}

func GetPublicIpsOutput(ctx *pulumi.Context, args GetPublicIpsOutputArgs, opts ...pulumi.InvokeOption) GetPublicIpsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPublicIpsResult, error) {
			args := v.(GetPublicIpsArgs)
			r, err := GetPublicIps(ctx, &args, opts...)
			var s GetPublicIpsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPublicIpsResultOutput)
}

// A collection of arguments for invoking getPublicIps.
type GetPublicIpsOutputArgs struct {
	Filters GetPublicIpsFilterArrayInput `pulumi:"filters"`
}

func (GetPublicIpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIpsArgs)(nil)).Elem()
}

// A collection of values returned by getPublicIps.
type GetPublicIpsResultOutput struct{ *pulumi.OutputState }

func (GetPublicIpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIpsResult)(nil)).Elem()
}

func (o GetPublicIpsResultOutput) ToGetPublicIpsResultOutput() GetPublicIpsResultOutput {
	return o
}

func (o GetPublicIpsResultOutput) ToGetPublicIpsResultOutputWithContext(ctx context.Context) GetPublicIpsResultOutput {
	return o
}

func (o GetPublicIpsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetPublicIpsResult] {
	return pulumix.Output[GetPublicIpsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetPublicIpsResultOutput) Filters() GetPublicIpsFilterArrayOutput {
	return o.ApplyT(func(v GetPublicIpsResult) []GetPublicIpsFilter { return v.Filters }).(GetPublicIpsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPublicIpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPublicIpsResultOutput) PublicIps() GetPublicIpsPublicIpArrayOutput {
	return o.ApplyT(func(v GetPublicIpsResult) []GetPublicIpsPublicIp { return v.PublicIps }).(GetPublicIpsPublicIpArrayOutput)
}

func (o GetPublicIpsResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpsResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicIpsResultOutput{})
}