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

func GetQuota(ctx *pulumi.Context, args *GetQuotaArgs, opts ...pulumi.InvokeOption) (*GetQuotaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetQuotaResult
	err := ctx.Invoke("outscale:index/getQuota:getQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQuota.
type GetQuotaArgs struct {
	Filters []GetQuotaFilter `pulumi:"filters"`
}

// A collection of values returned by getQuota.
type GetQuotaResult struct {
	AccountId   string           `pulumi:"accountId"`
	Description string           `pulumi:"description"`
	Filters     []GetQuotaFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	MaxValue         int    `pulumi:"maxValue"`
	Name             string `pulumi:"name"`
	QuotaCollection  string `pulumi:"quotaCollection"`
	QuotaType        string `pulumi:"quotaType"`
	RequestId        string `pulumi:"requestId"`
	ShortDescription string `pulumi:"shortDescription"`
	UsedValue        int    `pulumi:"usedValue"`
}

func GetQuotaOutput(ctx *pulumi.Context, args GetQuotaOutputArgs, opts ...pulumi.InvokeOption) GetQuotaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetQuotaResult, error) {
			args := v.(GetQuotaArgs)
			r, err := GetQuota(ctx, &args, opts...)
			var s GetQuotaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetQuotaResultOutput)
}

// A collection of arguments for invoking getQuota.
type GetQuotaOutputArgs struct {
	Filters GetQuotaFilterArrayInput `pulumi:"filters"`
}

func (GetQuotaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQuotaArgs)(nil)).Elem()
}

// A collection of values returned by getQuota.
type GetQuotaResultOutput struct{ *pulumi.OutputState }

func (GetQuotaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQuotaResult)(nil)).Elem()
}

func (o GetQuotaResultOutput) ToGetQuotaResultOutput() GetQuotaResultOutput {
	return o
}

func (o GetQuotaResultOutput) ToGetQuotaResultOutputWithContext(ctx context.Context) GetQuotaResultOutput {
	return o
}

func (o GetQuotaResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetQuotaResult] {
	return pulumix.Output[GetQuotaResult]{
		OutputState: o.OutputState,
	}
}

func (o GetQuotaResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) Filters() GetQuotaFilterArrayOutput {
	return o.ApplyT(func(v GetQuotaResult) []GetQuotaFilter { return v.Filters }).(GetQuotaFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetQuotaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) MaxValue() pulumi.IntOutput {
	return o.ApplyT(func(v GetQuotaResult) int { return v.MaxValue }).(pulumi.IntOutput)
}

func (o GetQuotaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) QuotaCollection() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.QuotaCollection }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) QuotaType() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.QuotaType }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaResult) string { return v.ShortDescription }).(pulumi.StringOutput)
}

func (o GetQuotaResultOutput) UsedValue() pulumi.IntOutput {
	return o.ApplyT(func(v GetQuotaResult) int { return v.UsedValue }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetQuotaResultOutput{})
}
