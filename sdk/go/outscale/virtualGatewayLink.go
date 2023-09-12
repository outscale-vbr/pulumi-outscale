// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package outscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/outscale-vbr/pulumi-outscale/sdk/go/outscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type VirtualGatewayLink struct {
	pulumi.CustomResourceState

	DryRun                   pulumi.StringOutput                                  `pulumi:"dryRun"`
	NetId                    pulumi.StringOutput                                  `pulumi:"netId"`
	NetToVirtualGatewayLinks VirtualGatewayLinkNetToVirtualGatewayLinkArrayOutput `pulumi:"netToVirtualGatewayLinks"`
	RequestId                pulumi.StringOutput                                  `pulumi:"requestId"`
	VirtualGatewayId         pulumi.StringOutput                                  `pulumi:"virtualGatewayId"`
}

// NewVirtualGatewayLink registers a new resource with the given unique name, arguments, and options.
func NewVirtualGatewayLink(ctx *pulumi.Context,
	name string, args *VirtualGatewayLinkArgs, opts ...pulumi.ResourceOption) (*VirtualGatewayLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetId == nil {
		return nil, errors.New("invalid value for required argument 'NetId'")
	}
	if args.VirtualGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualGatewayId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualGatewayLink
	err := ctx.RegisterResource("outscale:index/virtualGatewayLink:VirtualGatewayLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGatewayLink gets an existing VirtualGatewayLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGatewayLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGatewayLinkState, opts ...pulumi.ResourceOption) (*VirtualGatewayLink, error) {
	var resource VirtualGatewayLink
	err := ctx.ReadResource("outscale:index/virtualGatewayLink:VirtualGatewayLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGatewayLink resources.
type virtualGatewayLinkState struct {
	DryRun                   *string                                     `pulumi:"dryRun"`
	NetId                    *string                                     `pulumi:"netId"`
	NetToVirtualGatewayLinks []VirtualGatewayLinkNetToVirtualGatewayLink `pulumi:"netToVirtualGatewayLinks"`
	RequestId                *string                                     `pulumi:"requestId"`
	VirtualGatewayId         *string                                     `pulumi:"virtualGatewayId"`
}

type VirtualGatewayLinkState struct {
	DryRun                   pulumi.StringPtrInput
	NetId                    pulumi.StringPtrInput
	NetToVirtualGatewayLinks VirtualGatewayLinkNetToVirtualGatewayLinkArrayInput
	RequestId                pulumi.StringPtrInput
	VirtualGatewayId         pulumi.StringPtrInput
}

func (VirtualGatewayLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGatewayLinkState)(nil)).Elem()
}

type virtualGatewayLinkArgs struct {
	DryRun           *string `pulumi:"dryRun"`
	NetId            string  `pulumi:"netId"`
	VirtualGatewayId string  `pulumi:"virtualGatewayId"`
}

// The set of arguments for constructing a VirtualGatewayLink resource.
type VirtualGatewayLinkArgs struct {
	DryRun           pulumi.StringPtrInput
	NetId            pulumi.StringInput
	VirtualGatewayId pulumi.StringInput
}

func (VirtualGatewayLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGatewayLinkArgs)(nil)).Elem()
}

type VirtualGatewayLinkInput interface {
	pulumi.Input

	ToVirtualGatewayLinkOutput() VirtualGatewayLinkOutput
	ToVirtualGatewayLinkOutputWithContext(ctx context.Context) VirtualGatewayLinkOutput
}

func (*VirtualGatewayLink) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGatewayLink)(nil)).Elem()
}

func (i *VirtualGatewayLink) ToVirtualGatewayLinkOutput() VirtualGatewayLinkOutput {
	return i.ToVirtualGatewayLinkOutputWithContext(context.Background())
}

func (i *VirtualGatewayLink) ToVirtualGatewayLinkOutputWithContext(ctx context.Context) VirtualGatewayLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGatewayLinkOutput)
}

func (i *VirtualGatewayLink) ToOutput(ctx context.Context) pulumix.Output[*VirtualGatewayLink] {
	return pulumix.Output[*VirtualGatewayLink]{
		OutputState: i.ToVirtualGatewayLinkOutputWithContext(ctx).OutputState,
	}
}

// VirtualGatewayLinkArrayInput is an input type that accepts VirtualGatewayLinkArray and VirtualGatewayLinkArrayOutput values.
// You can construct a concrete instance of `VirtualGatewayLinkArrayInput` via:
//
//	VirtualGatewayLinkArray{ VirtualGatewayLinkArgs{...} }
type VirtualGatewayLinkArrayInput interface {
	pulumi.Input

	ToVirtualGatewayLinkArrayOutput() VirtualGatewayLinkArrayOutput
	ToVirtualGatewayLinkArrayOutputWithContext(context.Context) VirtualGatewayLinkArrayOutput
}

type VirtualGatewayLinkArray []VirtualGatewayLinkInput

func (VirtualGatewayLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGatewayLink)(nil)).Elem()
}

func (i VirtualGatewayLinkArray) ToVirtualGatewayLinkArrayOutput() VirtualGatewayLinkArrayOutput {
	return i.ToVirtualGatewayLinkArrayOutputWithContext(context.Background())
}

func (i VirtualGatewayLinkArray) ToVirtualGatewayLinkArrayOutputWithContext(ctx context.Context) VirtualGatewayLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGatewayLinkArrayOutput)
}

func (i VirtualGatewayLinkArray) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualGatewayLink] {
	return pulumix.Output[[]*VirtualGatewayLink]{
		OutputState: i.ToVirtualGatewayLinkArrayOutputWithContext(ctx).OutputState,
	}
}

// VirtualGatewayLinkMapInput is an input type that accepts VirtualGatewayLinkMap and VirtualGatewayLinkMapOutput values.
// You can construct a concrete instance of `VirtualGatewayLinkMapInput` via:
//
//	VirtualGatewayLinkMap{ "key": VirtualGatewayLinkArgs{...} }
type VirtualGatewayLinkMapInput interface {
	pulumi.Input

	ToVirtualGatewayLinkMapOutput() VirtualGatewayLinkMapOutput
	ToVirtualGatewayLinkMapOutputWithContext(context.Context) VirtualGatewayLinkMapOutput
}

type VirtualGatewayLinkMap map[string]VirtualGatewayLinkInput

func (VirtualGatewayLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGatewayLink)(nil)).Elem()
}

func (i VirtualGatewayLinkMap) ToVirtualGatewayLinkMapOutput() VirtualGatewayLinkMapOutput {
	return i.ToVirtualGatewayLinkMapOutputWithContext(context.Background())
}

func (i VirtualGatewayLinkMap) ToVirtualGatewayLinkMapOutputWithContext(ctx context.Context) VirtualGatewayLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGatewayLinkMapOutput)
}

func (i VirtualGatewayLinkMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualGatewayLink] {
	return pulumix.Output[map[string]*VirtualGatewayLink]{
		OutputState: i.ToVirtualGatewayLinkMapOutputWithContext(ctx).OutputState,
	}
}

type VirtualGatewayLinkOutput struct{ *pulumi.OutputState }

func (VirtualGatewayLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGatewayLink)(nil)).Elem()
}

func (o VirtualGatewayLinkOutput) ToVirtualGatewayLinkOutput() VirtualGatewayLinkOutput {
	return o
}

func (o VirtualGatewayLinkOutput) ToVirtualGatewayLinkOutputWithContext(ctx context.Context) VirtualGatewayLinkOutput {
	return o
}

func (o VirtualGatewayLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualGatewayLink] {
	return pulumix.Output[*VirtualGatewayLink]{
		OutputState: o.OutputState,
	}
}

func (o VirtualGatewayLinkOutput) DryRun() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGatewayLink) pulumi.StringOutput { return v.DryRun }).(pulumi.StringOutput)
}

func (o VirtualGatewayLinkOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGatewayLink) pulumi.StringOutput { return v.NetId }).(pulumi.StringOutput)
}

func (o VirtualGatewayLinkOutput) NetToVirtualGatewayLinks() VirtualGatewayLinkNetToVirtualGatewayLinkArrayOutput {
	return o.ApplyT(func(v *VirtualGatewayLink) VirtualGatewayLinkNetToVirtualGatewayLinkArrayOutput {
		return v.NetToVirtualGatewayLinks
	}).(VirtualGatewayLinkNetToVirtualGatewayLinkArrayOutput)
}

func (o VirtualGatewayLinkOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGatewayLink) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o VirtualGatewayLinkOutput) VirtualGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGatewayLink) pulumi.StringOutput { return v.VirtualGatewayId }).(pulumi.StringOutput)
}

type VirtualGatewayLinkArrayOutput struct{ *pulumi.OutputState }

func (VirtualGatewayLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGatewayLink)(nil)).Elem()
}

func (o VirtualGatewayLinkArrayOutput) ToVirtualGatewayLinkArrayOutput() VirtualGatewayLinkArrayOutput {
	return o
}

func (o VirtualGatewayLinkArrayOutput) ToVirtualGatewayLinkArrayOutputWithContext(ctx context.Context) VirtualGatewayLinkArrayOutput {
	return o
}

func (o VirtualGatewayLinkArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualGatewayLink] {
	return pulumix.Output[[]*VirtualGatewayLink]{
		OutputState: o.OutputState,
	}
}

func (o VirtualGatewayLinkArrayOutput) Index(i pulumi.IntInput) VirtualGatewayLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualGatewayLink {
		return vs[0].([]*VirtualGatewayLink)[vs[1].(int)]
	}).(VirtualGatewayLinkOutput)
}

type VirtualGatewayLinkMapOutput struct{ *pulumi.OutputState }

func (VirtualGatewayLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGatewayLink)(nil)).Elem()
}

func (o VirtualGatewayLinkMapOutput) ToVirtualGatewayLinkMapOutput() VirtualGatewayLinkMapOutput {
	return o
}

func (o VirtualGatewayLinkMapOutput) ToVirtualGatewayLinkMapOutputWithContext(ctx context.Context) VirtualGatewayLinkMapOutput {
	return o
}

func (o VirtualGatewayLinkMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualGatewayLink] {
	return pulumix.Output[map[string]*VirtualGatewayLink]{
		OutputState: o.OutputState,
	}
}

func (o VirtualGatewayLinkMapOutput) MapIndex(k pulumi.StringInput) VirtualGatewayLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualGatewayLink {
		return vs[0].(map[string]*VirtualGatewayLink)[vs[1].(string)]
	}).(VirtualGatewayLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGatewayLinkInput)(nil)).Elem(), &VirtualGatewayLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGatewayLinkArrayInput)(nil)).Elem(), VirtualGatewayLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGatewayLinkMapInput)(nil)).Elem(), VirtualGatewayLinkMap{})
	pulumi.RegisterOutputType(VirtualGatewayLinkOutput{})
	pulumi.RegisterOutputType(VirtualGatewayLinkArrayOutput{})
	pulumi.RegisterOutputType(VirtualGatewayLinkMapOutput{})
}