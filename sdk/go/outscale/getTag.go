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

func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTagResult
	err := ctx.Invoke("outscale:index/getTag:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTag.
type LookupTagArgs struct {
	Filters []GetTagFilter `pulumi:"filters"`
}

// A collection of values returned by getTag.
type LookupTagResult struct {
	Filters []GetTagFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	Key          string `pulumi:"key"`
	ResourceId   string `pulumi:"resourceId"`
	ResourceType string `pulumi:"resourceType"`
	Value        string `pulumi:"value"`
}

func LookupTagOutput(ctx *pulumi.Context, args LookupTagOutputArgs, opts ...pulumi.InvokeOption) LookupTagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagResult, error) {
			args := v.(LookupTagArgs)
			r, err := LookupTag(ctx, &args, opts...)
			var s LookupTagResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagResultOutput)
}

// A collection of arguments for invoking getTag.
type LookupTagOutputArgs struct {
	Filters GetTagFilterArrayInput `pulumi:"filters"`
}

func (LookupTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagArgs)(nil)).Elem()
}

// A collection of values returned by getTag.
type LookupTagResultOutput struct{ *pulumi.OutputState }

func (LookupTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagResult)(nil)).Elem()
}

func (o LookupTagResultOutput) ToLookupTagResultOutput() LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToLookupTagResultOutputWithContext(ctx context.Context) LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTagResult] {
	return pulumix.Output[LookupTagResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupTagResultOutput) Filters() GetTagFilterArrayOutput {
	return o.ApplyT(func(v LookupTagResult) []GetTagFilter { return v.Filters }).(GetTagFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTagResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagResultOutput{})
}
