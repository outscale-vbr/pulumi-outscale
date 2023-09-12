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

func LookupVirtualGateway(ctx *pulumi.Context, args *LookupVirtualGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVirtualGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualGatewayResult
	err := ctx.Invoke("outscale:index/getVirtualGateway:getVirtualGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualGateway.
type LookupVirtualGatewayArgs struct {
	ConnectionType   *string                   `pulumi:"connectionType"`
	Filters          []GetVirtualGatewayFilter `pulumi:"filters"`
	State            *string                   `pulumi:"state"`
	VirtualGatewayId *string                   `pulumi:"virtualGatewayId"`
}

// A collection of values returned by getVirtualGateway.
type LookupVirtualGatewayResult struct {
	ConnectionType string                    `pulumi:"connectionType"`
	Filters        []GetVirtualGatewayFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string                                     `pulumi:"id"`
	NetToVirtualGatewayLinks []GetVirtualGatewayNetToVirtualGatewayLink `pulumi:"netToVirtualGatewayLinks"`
	RequestId                string                                     `pulumi:"requestId"`
	State                    string                                     `pulumi:"state"`
	Tags                     []GetVirtualGatewayTag                     `pulumi:"tags"`
	VirtualGatewayId         string                                     `pulumi:"virtualGatewayId"`
}

func LookupVirtualGatewayOutput(ctx *pulumi.Context, args LookupVirtualGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualGatewayResult, error) {
			args := v.(LookupVirtualGatewayArgs)
			r, err := LookupVirtualGateway(ctx, &args, opts...)
			var s LookupVirtualGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualGatewayResultOutput)
}

// A collection of arguments for invoking getVirtualGateway.
type LookupVirtualGatewayOutputArgs struct {
	ConnectionType   pulumi.StringPtrInput             `pulumi:"connectionType"`
	Filters          GetVirtualGatewayFilterArrayInput `pulumi:"filters"`
	State            pulumi.StringPtrInput             `pulumi:"state"`
	VirtualGatewayId pulumi.StringPtrInput             `pulumi:"virtualGatewayId"`
}

func (LookupVirtualGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualGateway.
type LookupVirtualGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualGatewayResult)(nil)).Elem()
}

func (o LookupVirtualGatewayResultOutput) ToLookupVirtualGatewayResultOutput() LookupVirtualGatewayResultOutput {
	return o
}

func (o LookupVirtualGatewayResultOutput) ToLookupVirtualGatewayResultOutputWithContext(ctx context.Context) LookupVirtualGatewayResultOutput {
	return o
}

func (o LookupVirtualGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualGatewayResult] {
	return pulumix.Output[LookupVirtualGatewayResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVirtualGatewayResultOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) string { return v.ConnectionType }).(pulumi.StringOutput)
}

func (o LookupVirtualGatewayResultOutput) Filters() GetVirtualGatewayFilterArrayOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) []GetVirtualGatewayFilter { return v.Filters }).(GetVirtualGatewayFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualGatewayResultOutput) NetToVirtualGatewayLinks() GetVirtualGatewayNetToVirtualGatewayLinkArrayOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) []GetVirtualGatewayNetToVirtualGatewayLink {
		return v.NetToVirtualGatewayLinks
	}).(GetVirtualGatewayNetToVirtualGatewayLinkArrayOutput)
}

func (o LookupVirtualGatewayResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func (o LookupVirtualGatewayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupVirtualGatewayResultOutput) Tags() GetVirtualGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) []GetVirtualGatewayTag { return v.Tags }).(GetVirtualGatewayTagArrayOutput)
}

func (o LookupVirtualGatewayResultOutput) VirtualGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualGatewayResult) string { return v.VirtualGatewayId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualGatewayResultOutput{})
}