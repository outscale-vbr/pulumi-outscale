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

func LookupApiAccessRule(ctx *pulumi.Context, args *LookupApiAccessRuleArgs, opts ...pulumi.InvokeOption) (*LookupApiAccessRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiAccessRuleResult
	err := ctx.Invoke("outscale:index/getApiAccessRule:getApiAccessRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiAccessRule.
type LookupApiAccessRuleArgs struct {
	Filters []GetApiAccessRuleFilter `pulumi:"filters"`
}

// A collection of values returned by getApiAccessRule.
type LookupApiAccessRuleResult struct {
	ApiAccessRuleId string                   `pulumi:"apiAccessRuleId"`
	CaIds           []string                 `pulumi:"caIds"`
	Cns             []string                 `pulumi:"cns"`
	Description     string                   `pulumi:"description"`
	Filters         []GetApiAccessRuleFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	IpRanges  []string `pulumi:"ipRanges"`
	RequestId string   `pulumi:"requestId"`
}

func LookupApiAccessRuleOutput(ctx *pulumi.Context, args LookupApiAccessRuleOutputArgs, opts ...pulumi.InvokeOption) LookupApiAccessRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiAccessRuleResult, error) {
			args := v.(LookupApiAccessRuleArgs)
			r, err := LookupApiAccessRule(ctx, &args, opts...)
			var s LookupApiAccessRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiAccessRuleResultOutput)
}

// A collection of arguments for invoking getApiAccessRule.
type LookupApiAccessRuleOutputArgs struct {
	Filters GetApiAccessRuleFilterArrayInput `pulumi:"filters"`
}

func (LookupApiAccessRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiAccessRuleArgs)(nil)).Elem()
}

// A collection of values returned by getApiAccessRule.
type LookupApiAccessRuleResultOutput struct{ *pulumi.OutputState }

func (LookupApiAccessRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiAccessRuleResult)(nil)).Elem()
}

func (o LookupApiAccessRuleResultOutput) ToLookupApiAccessRuleResultOutput() LookupApiAccessRuleResultOutput {
	return o
}

func (o LookupApiAccessRuleResultOutput) ToLookupApiAccessRuleResultOutputWithContext(ctx context.Context) LookupApiAccessRuleResultOutput {
	return o
}

func (o LookupApiAccessRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiAccessRuleResult] {
	return pulumix.Output[LookupApiAccessRuleResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupApiAccessRuleResultOutput) ApiAccessRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) string { return v.ApiAccessRuleId }).(pulumi.StringOutput)
}

func (o LookupApiAccessRuleResultOutput) CaIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) []string { return v.CaIds }).(pulumi.StringArrayOutput)
}

func (o LookupApiAccessRuleResultOutput) Cns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) []string { return v.Cns }).(pulumi.StringArrayOutput)
}

func (o LookupApiAccessRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupApiAccessRuleResultOutput) Filters() GetApiAccessRuleFilterArrayOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) []GetApiAccessRuleFilter { return v.Filters }).(GetApiAccessRuleFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApiAccessRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiAccessRuleResultOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

func (o LookupApiAccessRuleResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiAccessRuleResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiAccessRuleResultOutput{})
}
