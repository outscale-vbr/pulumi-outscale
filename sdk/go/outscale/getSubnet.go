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

func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSubnetResult
	err := ctx.Invoke("outscale:index/getSubnet:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetArgs struct {
	Filters  []GetSubnetFilter `pulumi:"filters"`
	SubnetId *string           `pulumi:"subnetId"`
}

// A collection of values returned by getSubnet.
type LookupSubnetResult struct {
	AvailableIpsCount int               `pulumi:"availableIpsCount"`
	Filters           []GetSubnetFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string         `pulumi:"id"`
	IpRange             string         `pulumi:"ipRange"`
	MapPublicIpOnLaunch bool           `pulumi:"mapPublicIpOnLaunch"`
	NetId               string         `pulumi:"netId"`
	RequestId           string         `pulumi:"requestId"`
	State               string         `pulumi:"state"`
	SubnetId            string         `pulumi:"subnetId"`
	SubregionName       string         `pulumi:"subregionName"`
	Tags                []GetSubnetTag `pulumi:"tags"`
}

func LookupSubnetOutput(ctx *pulumi.Context, args LookupSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetResult, error) {
			args := v.(LookupSubnetArgs)
			r, err := LookupSubnet(ctx, &args, opts...)
			var s LookupSubnetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetResultOutput)
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetOutputArgs struct {
	Filters  GetSubnetFilterArrayInput `pulumi:"filters"`
	SubnetId pulumi.StringPtrInput     `pulumi:"subnetId"`
}

func (LookupSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}

// A collection of values returned by getSubnet.
type LookupSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetResult)(nil)).Elem()
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutput() LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutputWithContext(ctx context.Context) LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSubnetResult] {
	return pulumix.Output[LookupSubnetResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSubnetResultOutput) AvailableIpsCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetResult) int { return v.AvailableIpsCount }).(pulumi.IntOutput)
}

func (o LookupSubnetResultOutput) Filters() GetSubnetFilterArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []GetSubnetFilter { return v.Filters }).(GetSubnetFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSubnetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) IpRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.IpRange }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) MapPublicIpOnLaunch() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetResult) bool { return v.MapPublicIpOnLaunch }).(pulumi.BoolOutput)
}

func (o LookupSubnetResultOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.NetId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) SubregionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.SubregionName }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Tags() GetSubnetTagArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []GetSubnetTag { return v.Tags }).(GetSubnetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetResultOutput{})
}
