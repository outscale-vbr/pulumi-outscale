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

func GetFlexibleGpus(ctx *pulumi.Context, args *GetFlexibleGpusArgs, opts ...pulumi.InvokeOption) (*GetFlexibleGpusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFlexibleGpusResult
	err := ctx.Invoke("outscale:index/getFlexibleGpus:getFlexibleGpus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlexibleGpus.
type GetFlexibleGpusArgs struct {
	Filters []GetFlexibleGpusFilter `pulumi:"filters"`
}

// A collection of values returned by getFlexibleGpus.
type GetFlexibleGpusResult struct {
	Filters        []GetFlexibleGpusFilter       `pulumi:"filters"`
	FlexibleGpuses []GetFlexibleGpusFlexibleGpus `pulumi:"flexibleGpuses"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	RequestId string `pulumi:"requestId"`
}

func GetFlexibleGpusOutput(ctx *pulumi.Context, args GetFlexibleGpusOutputArgs, opts ...pulumi.InvokeOption) GetFlexibleGpusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFlexibleGpusResult, error) {
			args := v.(GetFlexibleGpusArgs)
			r, err := GetFlexibleGpus(ctx, &args, opts...)
			var s GetFlexibleGpusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFlexibleGpusResultOutput)
}

// A collection of arguments for invoking getFlexibleGpus.
type GetFlexibleGpusOutputArgs struct {
	Filters GetFlexibleGpusFilterArrayInput `pulumi:"filters"`
}

func (GetFlexibleGpusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlexibleGpusArgs)(nil)).Elem()
}

// A collection of values returned by getFlexibleGpus.
type GetFlexibleGpusResultOutput struct{ *pulumi.OutputState }

func (GetFlexibleGpusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlexibleGpusResult)(nil)).Elem()
}

func (o GetFlexibleGpusResultOutput) ToGetFlexibleGpusResultOutput() GetFlexibleGpusResultOutput {
	return o
}

func (o GetFlexibleGpusResultOutput) ToGetFlexibleGpusResultOutputWithContext(ctx context.Context) GetFlexibleGpusResultOutput {
	return o
}

func (o GetFlexibleGpusResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetFlexibleGpusResult] {
	return pulumix.Output[GetFlexibleGpusResult]{
		OutputState: o.OutputState,
	}
}

func (o GetFlexibleGpusResultOutput) Filters() GetFlexibleGpusFilterArrayOutput {
	return o.ApplyT(func(v GetFlexibleGpusResult) []GetFlexibleGpusFilter { return v.Filters }).(GetFlexibleGpusFilterArrayOutput)
}

func (o GetFlexibleGpusResultOutput) FlexibleGpuses() GetFlexibleGpusFlexibleGpusArrayOutput {
	return o.ApplyT(func(v GetFlexibleGpusResult) []GetFlexibleGpusFlexibleGpus { return v.FlexibleGpuses }).(GetFlexibleGpusFlexibleGpusArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlexibleGpusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlexibleGpusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFlexibleGpusResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlexibleGpusResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlexibleGpusResultOutput{})
}