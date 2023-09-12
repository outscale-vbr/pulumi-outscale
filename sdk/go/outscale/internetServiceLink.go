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

type InternetServiceLink struct {
	pulumi.CustomResourceState

	InternetServiceId pulumi.StringOutput               `pulumi:"internetServiceId"`
	NetId             pulumi.StringOutput               `pulumi:"netId"`
	RequestId         pulumi.StringOutput               `pulumi:"requestId"`
	State             pulumi.StringOutput               `pulumi:"state"`
	Tags              InternetServiceLinkTagArrayOutput `pulumi:"tags"`
}

// NewInternetServiceLink registers a new resource with the given unique name, arguments, and options.
func NewInternetServiceLink(ctx *pulumi.Context,
	name string, args *InternetServiceLinkArgs, opts ...pulumi.ResourceOption) (*InternetServiceLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InternetServiceId == nil {
		return nil, errors.New("invalid value for required argument 'InternetServiceId'")
	}
	if args.NetId == nil {
		return nil, errors.New("invalid value for required argument 'NetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InternetServiceLink
	err := ctx.RegisterResource("outscale:index/internetServiceLink:InternetServiceLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetServiceLink gets an existing InternetServiceLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetServiceLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetServiceLinkState, opts ...pulumi.ResourceOption) (*InternetServiceLink, error) {
	var resource InternetServiceLink
	err := ctx.ReadResource("outscale:index/internetServiceLink:InternetServiceLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InternetServiceLink resources.
type internetServiceLinkState struct {
	InternetServiceId *string                  `pulumi:"internetServiceId"`
	NetId             *string                  `pulumi:"netId"`
	RequestId         *string                  `pulumi:"requestId"`
	State             *string                  `pulumi:"state"`
	Tags              []InternetServiceLinkTag `pulumi:"tags"`
}

type InternetServiceLinkState struct {
	InternetServiceId pulumi.StringPtrInput
	NetId             pulumi.StringPtrInput
	RequestId         pulumi.StringPtrInput
	State             pulumi.StringPtrInput
	Tags              InternetServiceLinkTagArrayInput
}

func (InternetServiceLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetServiceLinkState)(nil)).Elem()
}

type internetServiceLinkArgs struct {
	InternetServiceId string `pulumi:"internetServiceId"`
	NetId             string `pulumi:"netId"`
}

// The set of arguments for constructing a InternetServiceLink resource.
type InternetServiceLinkArgs struct {
	InternetServiceId pulumi.StringInput
	NetId             pulumi.StringInput
}

func (InternetServiceLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetServiceLinkArgs)(nil)).Elem()
}

type InternetServiceLinkInput interface {
	pulumi.Input

	ToInternetServiceLinkOutput() InternetServiceLinkOutput
	ToInternetServiceLinkOutputWithContext(ctx context.Context) InternetServiceLinkOutput
}

func (*InternetServiceLink) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetServiceLink)(nil)).Elem()
}

func (i *InternetServiceLink) ToInternetServiceLinkOutput() InternetServiceLinkOutput {
	return i.ToInternetServiceLinkOutputWithContext(context.Background())
}

func (i *InternetServiceLink) ToInternetServiceLinkOutputWithContext(ctx context.Context) InternetServiceLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetServiceLinkOutput)
}

func (i *InternetServiceLink) ToOutput(ctx context.Context) pulumix.Output[*InternetServiceLink] {
	return pulumix.Output[*InternetServiceLink]{
		OutputState: i.ToInternetServiceLinkOutputWithContext(ctx).OutputState,
	}
}

// InternetServiceLinkArrayInput is an input type that accepts InternetServiceLinkArray and InternetServiceLinkArrayOutput values.
// You can construct a concrete instance of `InternetServiceLinkArrayInput` via:
//
//	InternetServiceLinkArray{ InternetServiceLinkArgs{...} }
type InternetServiceLinkArrayInput interface {
	pulumi.Input

	ToInternetServiceLinkArrayOutput() InternetServiceLinkArrayOutput
	ToInternetServiceLinkArrayOutputWithContext(context.Context) InternetServiceLinkArrayOutput
}

type InternetServiceLinkArray []InternetServiceLinkInput

func (InternetServiceLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InternetServiceLink)(nil)).Elem()
}

func (i InternetServiceLinkArray) ToInternetServiceLinkArrayOutput() InternetServiceLinkArrayOutput {
	return i.ToInternetServiceLinkArrayOutputWithContext(context.Background())
}

func (i InternetServiceLinkArray) ToInternetServiceLinkArrayOutputWithContext(ctx context.Context) InternetServiceLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetServiceLinkArrayOutput)
}

func (i InternetServiceLinkArray) ToOutput(ctx context.Context) pulumix.Output[[]*InternetServiceLink] {
	return pulumix.Output[[]*InternetServiceLink]{
		OutputState: i.ToInternetServiceLinkArrayOutputWithContext(ctx).OutputState,
	}
}

// InternetServiceLinkMapInput is an input type that accepts InternetServiceLinkMap and InternetServiceLinkMapOutput values.
// You can construct a concrete instance of `InternetServiceLinkMapInput` via:
//
//	InternetServiceLinkMap{ "key": InternetServiceLinkArgs{...} }
type InternetServiceLinkMapInput interface {
	pulumi.Input

	ToInternetServiceLinkMapOutput() InternetServiceLinkMapOutput
	ToInternetServiceLinkMapOutputWithContext(context.Context) InternetServiceLinkMapOutput
}

type InternetServiceLinkMap map[string]InternetServiceLinkInput

func (InternetServiceLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InternetServiceLink)(nil)).Elem()
}

func (i InternetServiceLinkMap) ToInternetServiceLinkMapOutput() InternetServiceLinkMapOutput {
	return i.ToInternetServiceLinkMapOutputWithContext(context.Background())
}

func (i InternetServiceLinkMap) ToInternetServiceLinkMapOutputWithContext(ctx context.Context) InternetServiceLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetServiceLinkMapOutput)
}

func (i InternetServiceLinkMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*InternetServiceLink] {
	return pulumix.Output[map[string]*InternetServiceLink]{
		OutputState: i.ToInternetServiceLinkMapOutputWithContext(ctx).OutputState,
	}
}

type InternetServiceLinkOutput struct{ *pulumi.OutputState }

func (InternetServiceLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetServiceLink)(nil)).Elem()
}

func (o InternetServiceLinkOutput) ToInternetServiceLinkOutput() InternetServiceLinkOutput {
	return o
}

func (o InternetServiceLinkOutput) ToInternetServiceLinkOutputWithContext(ctx context.Context) InternetServiceLinkOutput {
	return o
}

func (o InternetServiceLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*InternetServiceLink] {
	return pulumix.Output[*InternetServiceLink]{
		OutputState: o.OutputState,
	}
}

func (o InternetServiceLinkOutput) InternetServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetServiceLink) pulumi.StringOutput { return v.InternetServiceId }).(pulumi.StringOutput)
}

func (o InternetServiceLinkOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetServiceLink) pulumi.StringOutput { return v.NetId }).(pulumi.StringOutput)
}

func (o InternetServiceLinkOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetServiceLink) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o InternetServiceLinkOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetServiceLink) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o InternetServiceLinkOutput) Tags() InternetServiceLinkTagArrayOutput {
	return o.ApplyT(func(v *InternetServiceLink) InternetServiceLinkTagArrayOutput { return v.Tags }).(InternetServiceLinkTagArrayOutput)
}

type InternetServiceLinkArrayOutput struct{ *pulumi.OutputState }

func (InternetServiceLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InternetServiceLink)(nil)).Elem()
}

func (o InternetServiceLinkArrayOutput) ToInternetServiceLinkArrayOutput() InternetServiceLinkArrayOutput {
	return o
}

func (o InternetServiceLinkArrayOutput) ToInternetServiceLinkArrayOutputWithContext(ctx context.Context) InternetServiceLinkArrayOutput {
	return o
}

func (o InternetServiceLinkArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*InternetServiceLink] {
	return pulumix.Output[[]*InternetServiceLink]{
		OutputState: o.OutputState,
	}
}

func (o InternetServiceLinkArrayOutput) Index(i pulumi.IntInput) InternetServiceLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InternetServiceLink {
		return vs[0].([]*InternetServiceLink)[vs[1].(int)]
	}).(InternetServiceLinkOutput)
}

type InternetServiceLinkMapOutput struct{ *pulumi.OutputState }

func (InternetServiceLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InternetServiceLink)(nil)).Elem()
}

func (o InternetServiceLinkMapOutput) ToInternetServiceLinkMapOutput() InternetServiceLinkMapOutput {
	return o
}

func (o InternetServiceLinkMapOutput) ToInternetServiceLinkMapOutputWithContext(ctx context.Context) InternetServiceLinkMapOutput {
	return o
}

func (o InternetServiceLinkMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*InternetServiceLink] {
	return pulumix.Output[map[string]*InternetServiceLink]{
		OutputState: o.OutputState,
	}
}

func (o InternetServiceLinkMapOutput) MapIndex(k pulumi.StringInput) InternetServiceLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InternetServiceLink {
		return vs[0].(map[string]*InternetServiceLink)[vs[1].(string)]
	}).(InternetServiceLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetServiceLinkInput)(nil)).Elem(), &InternetServiceLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetServiceLinkArrayInput)(nil)).Elem(), InternetServiceLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetServiceLinkMapInput)(nil)).Elem(), InternetServiceLinkMap{})
	pulumi.RegisterOutputType(InternetServiceLinkOutput{})
	pulumi.RegisterOutputType(InternetServiceLinkArrayOutput{})
	pulumi.RegisterOutputType(InternetServiceLinkMapOutput{})
}