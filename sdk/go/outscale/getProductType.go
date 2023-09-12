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

func GetProductType(ctx *pulumi.Context, args *GetProductTypeArgs, opts ...pulumi.InvokeOption) (*GetProductTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProductTypeResult
	err := ctx.Invoke("outscale:index/getProductType:getProductType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProductType.
type GetProductTypeArgs struct {
	Filters []GetProductTypeFilter `pulumi:"filters"`
}

// A collection of values returned by getProductType.
type GetProductTypeResult struct {
	Description string                 `pulumi:"description"`
	Filters     []GetProductTypeFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	ProductTypeId string `pulumi:"productTypeId"`
	RequestId     string `pulumi:"requestId"`
	Vendor        string `pulumi:"vendor"`
}

func GetProductTypeOutput(ctx *pulumi.Context, args GetProductTypeOutputArgs, opts ...pulumi.InvokeOption) GetProductTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductTypeResult, error) {
			args := v.(GetProductTypeArgs)
			r, err := GetProductType(ctx, &args, opts...)
			var s GetProductTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductTypeResultOutput)
}

// A collection of arguments for invoking getProductType.
type GetProductTypeOutputArgs struct {
	Filters GetProductTypeFilterArrayInput `pulumi:"filters"`
}

func (GetProductTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductTypeArgs)(nil)).Elem()
}

// A collection of values returned by getProductType.
type GetProductTypeResultOutput struct{ *pulumi.OutputState }

func (GetProductTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductTypeResult)(nil)).Elem()
}

func (o GetProductTypeResultOutput) ToGetProductTypeResultOutput() GetProductTypeResultOutput {
	return o
}

func (o GetProductTypeResultOutput) ToGetProductTypeResultOutputWithContext(ctx context.Context) GetProductTypeResultOutput {
	return o
}

func (o GetProductTypeResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetProductTypeResult] {
	return pulumix.Output[GetProductTypeResult]{
		OutputState: o.OutputState,
	}
}

func (o GetProductTypeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductTypeResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetProductTypeResultOutput) Filters() GetProductTypeFilterArrayOutput {
	return o.ApplyT(func(v GetProductTypeResult) []GetProductTypeFilter { return v.Filters }).(GetProductTypeFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProductTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProductTypeResultOutput) ProductTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductTypeResult) string { return v.ProductTypeId }).(pulumi.StringOutput)
}

func (o GetProductTypeResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductTypeResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func (o GetProductTypeResultOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductTypeResult) string { return v.Vendor }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductTypeResultOutput{})
}