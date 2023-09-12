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

type PublicIp struct {
	pulumi.CustomResourceState

	LinkPublicIpId pulumi.StringOutput    `pulumi:"linkPublicIpId"`
	NicAccountId   pulumi.StringOutput    `pulumi:"nicAccountId"`
	NicId          pulumi.StringOutput    `pulumi:"nicId"`
	PrivateIp      pulumi.StringOutput    `pulumi:"privateIp"`
	PublicIp       pulumi.StringOutput    `pulumi:"publicIp"`
	PublicIpId     pulumi.StringOutput    `pulumi:"publicIpId"`
	RequestId      pulumi.StringOutput    `pulumi:"requestId"`
	Tags           PublicIpTagArrayOutput `pulumi:"tags"`
	VmId           pulumi.StringOutput    `pulumi:"vmId"`
}

// NewPublicIp registers a new resource with the given unique name, arguments, and options.
func NewPublicIp(ctx *pulumi.Context,
	name string, args *PublicIpArgs, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	if args == nil {
		args = &PublicIpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicIp
	err := ctx.RegisterResource("outscale:index/publicIp:PublicIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIp gets an existing PublicIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpState, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	var resource PublicIp
	err := ctx.ReadResource("outscale:index/publicIp:PublicIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIp resources.
type publicIpState struct {
	LinkPublicIpId *string       `pulumi:"linkPublicIpId"`
	NicAccountId   *string       `pulumi:"nicAccountId"`
	NicId          *string       `pulumi:"nicId"`
	PrivateIp      *string       `pulumi:"privateIp"`
	PublicIp       *string       `pulumi:"publicIp"`
	PublicIpId     *string       `pulumi:"publicIpId"`
	RequestId      *string       `pulumi:"requestId"`
	Tags           []PublicIpTag `pulumi:"tags"`
	VmId           *string       `pulumi:"vmId"`
}

type PublicIpState struct {
	LinkPublicIpId pulumi.StringPtrInput
	NicAccountId   pulumi.StringPtrInput
	NicId          pulumi.StringPtrInput
	PrivateIp      pulumi.StringPtrInput
	PublicIp       pulumi.StringPtrInput
	PublicIpId     pulumi.StringPtrInput
	RequestId      pulumi.StringPtrInput
	Tags           PublicIpTagArrayInput
	VmId           pulumi.StringPtrInput
}

func (PublicIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpState)(nil)).Elem()
}

type publicIpArgs struct {
	Tags []PublicIpTag `pulumi:"tags"`
}

// The set of arguments for constructing a PublicIp resource.
type PublicIpArgs struct {
	Tags PublicIpTagArrayInput
}

func (PublicIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpArgs)(nil)).Elem()
}

type PublicIpInput interface {
	pulumi.Input

	ToPublicIpOutput() PublicIpOutput
	ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput
}

func (*PublicIp) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIp)(nil)).Elem()
}

func (i *PublicIp) ToPublicIpOutput() PublicIpOutput {
	return i.ToPublicIpOutputWithContext(context.Background())
}

func (i *PublicIp) ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpOutput)
}

func (i *PublicIp) ToOutput(ctx context.Context) pulumix.Output[*PublicIp] {
	return pulumix.Output[*PublicIp]{
		OutputState: i.ToPublicIpOutputWithContext(ctx).OutputState,
	}
}

// PublicIpArrayInput is an input type that accepts PublicIpArray and PublicIpArrayOutput values.
// You can construct a concrete instance of `PublicIpArrayInput` via:
//
//	PublicIpArray{ PublicIpArgs{...} }
type PublicIpArrayInput interface {
	pulumi.Input

	ToPublicIpArrayOutput() PublicIpArrayOutput
	ToPublicIpArrayOutputWithContext(context.Context) PublicIpArrayOutput
}

type PublicIpArray []PublicIpInput

func (PublicIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIp)(nil)).Elem()
}

func (i PublicIpArray) ToPublicIpArrayOutput() PublicIpArrayOutput {
	return i.ToPublicIpArrayOutputWithContext(context.Background())
}

func (i PublicIpArray) ToPublicIpArrayOutputWithContext(ctx context.Context) PublicIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpArrayOutput)
}

func (i PublicIpArray) ToOutput(ctx context.Context) pulumix.Output[[]*PublicIp] {
	return pulumix.Output[[]*PublicIp]{
		OutputState: i.ToPublicIpArrayOutputWithContext(ctx).OutputState,
	}
}

// PublicIpMapInput is an input type that accepts PublicIpMap and PublicIpMapOutput values.
// You can construct a concrete instance of `PublicIpMapInput` via:
//
//	PublicIpMap{ "key": PublicIpArgs{...} }
type PublicIpMapInput interface {
	pulumi.Input

	ToPublicIpMapOutput() PublicIpMapOutput
	ToPublicIpMapOutputWithContext(context.Context) PublicIpMapOutput
}

type PublicIpMap map[string]PublicIpInput

func (PublicIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIp)(nil)).Elem()
}

func (i PublicIpMap) ToPublicIpMapOutput() PublicIpMapOutput {
	return i.ToPublicIpMapOutputWithContext(context.Background())
}

func (i PublicIpMap) ToPublicIpMapOutputWithContext(ctx context.Context) PublicIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpMapOutput)
}

func (i PublicIpMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PublicIp] {
	return pulumix.Output[map[string]*PublicIp]{
		OutputState: i.ToPublicIpMapOutputWithContext(ctx).OutputState,
	}
}

type PublicIpOutput struct{ *pulumi.OutputState }

func (PublicIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIp)(nil)).Elem()
}

func (o PublicIpOutput) ToPublicIpOutput() PublicIpOutput {
	return o
}

func (o PublicIpOutput) ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput {
	return o
}

func (o PublicIpOutput) ToOutput(ctx context.Context) pulumix.Output[*PublicIp] {
	return pulumix.Output[*PublicIp]{
		OutputState: o.OutputState,
	}
}

func (o PublicIpOutput) LinkPublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.LinkPublicIpId }).(pulumi.StringOutput)
}

func (o PublicIpOutput) NicAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.NicAccountId }).(pulumi.StringOutput)
}

func (o PublicIpOutput) NicId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.NicId }).(pulumi.StringOutput)
}

func (o PublicIpOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o PublicIpOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

func (o PublicIpOutput) PublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.PublicIpId }).(pulumi.StringOutput)
}

func (o PublicIpOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o PublicIpOutput) Tags() PublicIpTagArrayOutput {
	return o.ApplyT(func(v *PublicIp) PublicIpTagArrayOutput { return v.Tags }).(PublicIpTagArrayOutput)
}

func (o PublicIpOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIp) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

type PublicIpArrayOutput struct{ *pulumi.OutputState }

func (PublicIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIp)(nil)).Elem()
}

func (o PublicIpArrayOutput) ToPublicIpArrayOutput() PublicIpArrayOutput {
	return o
}

func (o PublicIpArrayOutput) ToPublicIpArrayOutputWithContext(ctx context.Context) PublicIpArrayOutput {
	return o
}

func (o PublicIpArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PublicIp] {
	return pulumix.Output[[]*PublicIp]{
		OutputState: o.OutputState,
	}
}

func (o PublicIpArrayOutput) Index(i pulumi.IntInput) PublicIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicIp {
		return vs[0].([]*PublicIp)[vs[1].(int)]
	}).(PublicIpOutput)
}

type PublicIpMapOutput struct{ *pulumi.OutputState }

func (PublicIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIp)(nil)).Elem()
}

func (o PublicIpMapOutput) ToPublicIpMapOutput() PublicIpMapOutput {
	return o
}

func (o PublicIpMapOutput) ToPublicIpMapOutputWithContext(ctx context.Context) PublicIpMapOutput {
	return o
}

func (o PublicIpMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PublicIp] {
	return pulumix.Output[map[string]*PublicIp]{
		OutputState: o.OutputState,
	}
}

func (o PublicIpMapOutput) MapIndex(k pulumi.StringInput) PublicIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicIp {
		return vs[0].(map[string]*PublicIp)[vs[1].(string)]
	}).(PublicIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpInput)(nil)).Elem(), &PublicIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpArrayInput)(nil)).Elem(), PublicIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpMapInput)(nil)).Elem(), PublicIpMap{})
	pulumi.RegisterOutputType(PublicIpOutput{})
	pulumi.RegisterOutputType(PublicIpArrayOutput{})
	pulumi.RegisterOutputType(PublicIpMapOutput{})
}
