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

type Ca struct {
	pulumi.CustomResourceState

	CaFingerprint pulumi.StringOutput    `pulumi:"caFingerprint"`
	CaId          pulumi.StringOutput    `pulumi:"caId"`
	CaPem         pulumi.StringPtrOutput `pulumi:"caPem"`
	Description   pulumi.StringPtrOutput `pulumi:"description"`
	RequestId     pulumi.StringOutput    `pulumi:"requestId"`
}

// NewCa registers a new resource with the given unique name, arguments, and options.
func NewCa(ctx *pulumi.Context,
	name string, args *CaArgs, opts ...pulumi.ResourceOption) (*Ca, error) {
	if args == nil {
		args = &CaArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ca
	err := ctx.RegisterResource("outscale:index/ca:Ca", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCa gets an existing Ca resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaState, opts ...pulumi.ResourceOption) (*Ca, error) {
	var resource Ca
	err := ctx.ReadResource("outscale:index/ca:Ca", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ca resources.
type caState struct {
	CaFingerprint *string `pulumi:"caFingerprint"`
	CaId          *string `pulumi:"caId"`
	CaPem         *string `pulumi:"caPem"`
	Description   *string `pulumi:"description"`
	RequestId     *string `pulumi:"requestId"`
}

type CaState struct {
	CaFingerprint pulumi.StringPtrInput
	CaId          pulumi.StringPtrInput
	CaPem         pulumi.StringPtrInput
	Description   pulumi.StringPtrInput
	RequestId     pulumi.StringPtrInput
}

func (CaState) ElementType() reflect.Type {
	return reflect.TypeOf((*caState)(nil)).Elem()
}

type caArgs struct {
	CaPem       *string `pulumi:"caPem"`
	Description *string `pulumi:"description"`
}

// The set of arguments for constructing a Ca resource.
type CaArgs struct {
	CaPem       pulumi.StringPtrInput
	Description pulumi.StringPtrInput
}

func (CaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caArgs)(nil)).Elem()
}

type CaInput interface {
	pulumi.Input

	ToCaOutput() CaOutput
	ToCaOutputWithContext(ctx context.Context) CaOutput
}

func (*Ca) ElementType() reflect.Type {
	return reflect.TypeOf((**Ca)(nil)).Elem()
}

func (i *Ca) ToCaOutput() CaOutput {
	return i.ToCaOutputWithContext(context.Background())
}

func (i *Ca) ToCaOutputWithContext(ctx context.Context) CaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaOutput)
}

func (i *Ca) ToOutput(ctx context.Context) pulumix.Output[*Ca] {
	return pulumix.Output[*Ca]{
		OutputState: i.ToCaOutputWithContext(ctx).OutputState,
	}
}

// CaArrayInput is an input type that accepts CaArray and CaArrayOutput values.
// You can construct a concrete instance of `CaArrayInput` via:
//
//	CaArray{ CaArgs{...} }
type CaArrayInput interface {
	pulumi.Input

	ToCaArrayOutput() CaArrayOutput
	ToCaArrayOutputWithContext(context.Context) CaArrayOutput
}

type CaArray []CaInput

func (CaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ca)(nil)).Elem()
}

func (i CaArray) ToCaArrayOutput() CaArrayOutput {
	return i.ToCaArrayOutputWithContext(context.Background())
}

func (i CaArray) ToCaArrayOutputWithContext(ctx context.Context) CaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaArrayOutput)
}

func (i CaArray) ToOutput(ctx context.Context) pulumix.Output[[]*Ca] {
	return pulumix.Output[[]*Ca]{
		OutputState: i.ToCaArrayOutputWithContext(ctx).OutputState,
	}
}

// CaMapInput is an input type that accepts CaMap and CaMapOutput values.
// You can construct a concrete instance of `CaMapInput` via:
//
//	CaMap{ "key": CaArgs{...} }
type CaMapInput interface {
	pulumi.Input

	ToCaMapOutput() CaMapOutput
	ToCaMapOutputWithContext(context.Context) CaMapOutput
}

type CaMap map[string]CaInput

func (CaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ca)(nil)).Elem()
}

func (i CaMap) ToCaMapOutput() CaMapOutput {
	return i.ToCaMapOutputWithContext(context.Background())
}

func (i CaMap) ToCaMapOutputWithContext(ctx context.Context) CaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaMapOutput)
}

func (i CaMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ca] {
	return pulumix.Output[map[string]*Ca]{
		OutputState: i.ToCaMapOutputWithContext(ctx).OutputState,
	}
}

type CaOutput struct{ *pulumi.OutputState }

func (CaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ca)(nil)).Elem()
}

func (o CaOutput) ToCaOutput() CaOutput {
	return o
}

func (o CaOutput) ToCaOutputWithContext(ctx context.Context) CaOutput {
	return o
}

func (o CaOutput) ToOutput(ctx context.Context) pulumix.Output[*Ca] {
	return pulumix.Output[*Ca]{
		OutputState: o.OutputState,
	}
}

func (o CaOutput) CaFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.CaFingerprint }).(pulumi.StringOutput)
}

func (o CaOutput) CaId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.CaId }).(pulumi.StringOutput)
}

func (o CaOutput) CaPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringPtrOutput { return v.CaPem }).(pulumi.StringPtrOutput)
}

func (o CaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CaOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ca) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

type CaArrayOutput struct{ *pulumi.OutputState }

func (CaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ca)(nil)).Elem()
}

func (o CaArrayOutput) ToCaArrayOutput() CaArrayOutput {
	return o
}

func (o CaArrayOutput) ToCaArrayOutputWithContext(ctx context.Context) CaArrayOutput {
	return o
}

func (o CaArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Ca] {
	return pulumix.Output[[]*Ca]{
		OutputState: o.OutputState,
	}
}

func (o CaArrayOutput) Index(i pulumi.IntInput) CaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ca {
		return vs[0].([]*Ca)[vs[1].(int)]
	}).(CaOutput)
}

type CaMapOutput struct{ *pulumi.OutputState }

func (CaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ca)(nil)).Elem()
}

func (o CaMapOutput) ToCaMapOutput() CaMapOutput {
	return o
}

func (o CaMapOutput) ToCaMapOutputWithContext(ctx context.Context) CaMapOutput {
	return o
}

func (o CaMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ca] {
	return pulumix.Output[map[string]*Ca]{
		OutputState: o.OutputState,
	}
}

func (o CaMapOutput) MapIndex(k pulumi.StringInput) CaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ca {
		return vs[0].(map[string]*Ca)[vs[1].(string)]
	}).(CaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaInput)(nil)).Elem(), &Ca{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaArrayInput)(nil)).Elem(), CaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaMapInput)(nil)).Elem(), CaMap{})
	pulumi.RegisterOutputType(CaOutput{})
	pulumi.RegisterOutputType(CaArrayOutput{})
	pulumi.RegisterOutputType(CaMapOutput{})
}