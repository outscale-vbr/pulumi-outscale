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

type ApiAccessRule struct {
	pulumi.CustomResourceState

	ApiAccessRuleId pulumi.StringOutput      `pulumi:"apiAccessRuleId"`
	CaIds           pulumi.StringArrayOutput `pulumi:"caIds"`
	Cns             pulumi.StringArrayOutput `pulumi:"cns"`
	Description     pulumi.StringPtrOutput   `pulumi:"description"`
	IpRanges        pulumi.StringArrayOutput `pulumi:"ipRanges"`
	RequestId       pulumi.StringOutput      `pulumi:"requestId"`
}

// NewApiAccessRule registers a new resource with the given unique name, arguments, and options.
func NewApiAccessRule(ctx *pulumi.Context,
	name string, args *ApiAccessRuleArgs, opts ...pulumi.ResourceOption) (*ApiAccessRule, error) {
	if args == nil {
		args = &ApiAccessRuleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiAccessRule
	err := ctx.RegisterResource("outscale:index/apiAccessRule:ApiAccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiAccessRule gets an existing ApiAccessRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiAccessRuleState, opts ...pulumi.ResourceOption) (*ApiAccessRule, error) {
	var resource ApiAccessRule
	err := ctx.ReadResource("outscale:index/apiAccessRule:ApiAccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiAccessRule resources.
type apiAccessRuleState struct {
	ApiAccessRuleId *string  `pulumi:"apiAccessRuleId"`
	CaIds           []string `pulumi:"caIds"`
	Cns             []string `pulumi:"cns"`
	Description     *string  `pulumi:"description"`
	IpRanges        []string `pulumi:"ipRanges"`
	RequestId       *string  `pulumi:"requestId"`
}

type ApiAccessRuleState struct {
	ApiAccessRuleId pulumi.StringPtrInput
	CaIds           pulumi.StringArrayInput
	Cns             pulumi.StringArrayInput
	Description     pulumi.StringPtrInput
	IpRanges        pulumi.StringArrayInput
	RequestId       pulumi.StringPtrInput
}

func (ApiAccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiAccessRuleState)(nil)).Elem()
}

type apiAccessRuleArgs struct {
	CaIds       []string `pulumi:"caIds"`
	Cns         []string `pulumi:"cns"`
	Description *string  `pulumi:"description"`
	IpRanges    []string `pulumi:"ipRanges"`
}

// The set of arguments for constructing a ApiAccessRule resource.
type ApiAccessRuleArgs struct {
	CaIds       pulumi.StringArrayInput
	Cns         pulumi.StringArrayInput
	Description pulumi.StringPtrInput
	IpRanges    pulumi.StringArrayInput
}

func (ApiAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiAccessRuleArgs)(nil)).Elem()
}

type ApiAccessRuleInput interface {
	pulumi.Input

	ToApiAccessRuleOutput() ApiAccessRuleOutput
	ToApiAccessRuleOutputWithContext(ctx context.Context) ApiAccessRuleOutput
}

func (*ApiAccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiAccessRule)(nil)).Elem()
}

func (i *ApiAccessRule) ToApiAccessRuleOutput() ApiAccessRuleOutput {
	return i.ToApiAccessRuleOutputWithContext(context.Background())
}

func (i *ApiAccessRule) ToApiAccessRuleOutputWithContext(ctx context.Context) ApiAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiAccessRuleOutput)
}

func (i *ApiAccessRule) ToOutput(ctx context.Context) pulumix.Output[*ApiAccessRule] {
	return pulumix.Output[*ApiAccessRule]{
		OutputState: i.ToApiAccessRuleOutputWithContext(ctx).OutputState,
	}
}

// ApiAccessRuleArrayInput is an input type that accepts ApiAccessRuleArray and ApiAccessRuleArrayOutput values.
// You can construct a concrete instance of `ApiAccessRuleArrayInput` via:
//
//	ApiAccessRuleArray{ ApiAccessRuleArgs{...} }
type ApiAccessRuleArrayInput interface {
	pulumi.Input

	ToApiAccessRuleArrayOutput() ApiAccessRuleArrayOutput
	ToApiAccessRuleArrayOutputWithContext(context.Context) ApiAccessRuleArrayOutput
}

type ApiAccessRuleArray []ApiAccessRuleInput

func (ApiAccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiAccessRule)(nil)).Elem()
}

func (i ApiAccessRuleArray) ToApiAccessRuleArrayOutput() ApiAccessRuleArrayOutput {
	return i.ToApiAccessRuleArrayOutputWithContext(context.Background())
}

func (i ApiAccessRuleArray) ToApiAccessRuleArrayOutputWithContext(ctx context.Context) ApiAccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiAccessRuleArrayOutput)
}

func (i ApiAccessRuleArray) ToOutput(ctx context.Context) pulumix.Output[[]*ApiAccessRule] {
	return pulumix.Output[[]*ApiAccessRule]{
		OutputState: i.ToApiAccessRuleArrayOutputWithContext(ctx).OutputState,
	}
}

// ApiAccessRuleMapInput is an input type that accepts ApiAccessRuleMap and ApiAccessRuleMapOutput values.
// You can construct a concrete instance of `ApiAccessRuleMapInput` via:
//
//	ApiAccessRuleMap{ "key": ApiAccessRuleArgs{...} }
type ApiAccessRuleMapInput interface {
	pulumi.Input

	ToApiAccessRuleMapOutput() ApiAccessRuleMapOutput
	ToApiAccessRuleMapOutputWithContext(context.Context) ApiAccessRuleMapOutput
}

type ApiAccessRuleMap map[string]ApiAccessRuleInput

func (ApiAccessRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiAccessRule)(nil)).Elem()
}

func (i ApiAccessRuleMap) ToApiAccessRuleMapOutput() ApiAccessRuleMapOutput {
	return i.ToApiAccessRuleMapOutputWithContext(context.Background())
}

func (i ApiAccessRuleMap) ToApiAccessRuleMapOutputWithContext(ctx context.Context) ApiAccessRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiAccessRuleMapOutput)
}

func (i ApiAccessRuleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApiAccessRule] {
	return pulumix.Output[map[string]*ApiAccessRule]{
		OutputState: i.ToApiAccessRuleMapOutputWithContext(ctx).OutputState,
	}
}

type ApiAccessRuleOutput struct{ *pulumi.OutputState }

func (ApiAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiAccessRule)(nil)).Elem()
}

func (o ApiAccessRuleOutput) ToApiAccessRuleOutput() ApiAccessRuleOutput {
	return o
}

func (o ApiAccessRuleOutput) ToApiAccessRuleOutputWithContext(ctx context.Context) ApiAccessRuleOutput {
	return o
}

func (o ApiAccessRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*ApiAccessRule] {
	return pulumix.Output[*ApiAccessRule]{
		OutputState: o.OutputState,
	}
}

func (o ApiAccessRuleOutput) ApiAccessRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiAccessRule) pulumi.StringOutput { return v.ApiAccessRuleId }).(pulumi.StringOutput)
}

func (o ApiAccessRuleOutput) CaIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiAccessRule) pulumi.StringArrayOutput { return v.CaIds }).(pulumi.StringArrayOutput)
}

func (o ApiAccessRuleOutput) Cns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiAccessRule) pulumi.StringArrayOutput { return v.Cns }).(pulumi.StringArrayOutput)
}

func (o ApiAccessRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiAccessRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiAccessRuleOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApiAccessRule) pulumi.StringArrayOutput { return v.IpRanges }).(pulumi.StringArrayOutput)
}

func (o ApiAccessRuleOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiAccessRule) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

type ApiAccessRuleArrayOutput struct{ *pulumi.OutputState }

func (ApiAccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiAccessRule)(nil)).Elem()
}

func (o ApiAccessRuleArrayOutput) ToApiAccessRuleArrayOutput() ApiAccessRuleArrayOutput {
	return o
}

func (o ApiAccessRuleArrayOutput) ToApiAccessRuleArrayOutputWithContext(ctx context.Context) ApiAccessRuleArrayOutput {
	return o
}

func (o ApiAccessRuleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ApiAccessRule] {
	return pulumix.Output[[]*ApiAccessRule]{
		OutputState: o.OutputState,
	}
}

func (o ApiAccessRuleArrayOutput) Index(i pulumi.IntInput) ApiAccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiAccessRule {
		return vs[0].([]*ApiAccessRule)[vs[1].(int)]
	}).(ApiAccessRuleOutput)
}

type ApiAccessRuleMapOutput struct{ *pulumi.OutputState }

func (ApiAccessRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiAccessRule)(nil)).Elem()
}

func (o ApiAccessRuleMapOutput) ToApiAccessRuleMapOutput() ApiAccessRuleMapOutput {
	return o
}

func (o ApiAccessRuleMapOutput) ToApiAccessRuleMapOutputWithContext(ctx context.Context) ApiAccessRuleMapOutput {
	return o
}

func (o ApiAccessRuleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ApiAccessRule] {
	return pulumix.Output[map[string]*ApiAccessRule]{
		OutputState: o.OutputState,
	}
}

func (o ApiAccessRuleMapOutput) MapIndex(k pulumi.StringInput) ApiAccessRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiAccessRule {
		return vs[0].(map[string]*ApiAccessRule)[vs[1].(string)]
	}).(ApiAccessRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiAccessRuleInput)(nil)).Elem(), &ApiAccessRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiAccessRuleArrayInput)(nil)).Elem(), ApiAccessRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiAccessRuleMapInput)(nil)).Elem(), ApiAccessRuleMap{})
	pulumi.RegisterOutputType(ApiAccessRuleOutput{})
	pulumi.RegisterOutputType(ApiAccessRuleArrayOutput{})
	pulumi.RegisterOutputType(ApiAccessRuleMapOutput{})
}
