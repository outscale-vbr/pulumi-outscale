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

func GetLoadBalancerVmHealth(ctx *pulumi.Context, args *GetLoadBalancerVmHealthArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancerVmHealthResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLoadBalancerVmHealthResult
	err := ctx.Invoke("outscale:index/getLoadBalancerVmHealth:getLoadBalancerVmHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancerVmHealth.
type GetLoadBalancerVmHealthArgs struct {
	BackendVmIds     []string                        `pulumi:"backendVmIds"`
	Filters          []GetLoadBalancerVmHealthFilter `pulumi:"filters"`
	LoadBalancerName string                          `pulumi:"loadBalancerName"`
}

// A collection of values returned by getLoadBalancerVmHealth.
type GetLoadBalancerVmHealthResult struct {
	BackendVmHealths []GetLoadBalancerVmHealthBackendVmHealth `pulumi:"backendVmHealths"`
	BackendVmIds     []string                                 `pulumi:"backendVmIds"`
	Filters          []GetLoadBalancerVmHealthFilter          `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	LoadBalancerName string `pulumi:"loadBalancerName"`
	RequestId        string `pulumi:"requestId"`
}

func GetLoadBalancerVmHealthOutput(ctx *pulumi.Context, args GetLoadBalancerVmHealthOutputArgs, opts ...pulumi.InvokeOption) GetLoadBalancerVmHealthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLoadBalancerVmHealthResult, error) {
			args := v.(GetLoadBalancerVmHealthArgs)
			r, err := GetLoadBalancerVmHealth(ctx, &args, opts...)
			var s GetLoadBalancerVmHealthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLoadBalancerVmHealthResultOutput)
}

// A collection of arguments for invoking getLoadBalancerVmHealth.
type GetLoadBalancerVmHealthOutputArgs struct {
	BackendVmIds     pulumi.StringArrayInput                 `pulumi:"backendVmIds"`
	Filters          GetLoadBalancerVmHealthFilterArrayInput `pulumi:"filters"`
	LoadBalancerName pulumi.StringInput                      `pulumi:"loadBalancerName"`
}

func (GetLoadBalancerVmHealthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancerVmHealthArgs)(nil)).Elem()
}

// A collection of values returned by getLoadBalancerVmHealth.
type GetLoadBalancerVmHealthResultOutput struct{ *pulumi.OutputState }

func (GetLoadBalancerVmHealthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancerVmHealthResult)(nil)).Elem()
}

func (o GetLoadBalancerVmHealthResultOutput) ToGetLoadBalancerVmHealthResultOutput() GetLoadBalancerVmHealthResultOutput {
	return o
}

func (o GetLoadBalancerVmHealthResultOutput) ToGetLoadBalancerVmHealthResultOutputWithContext(ctx context.Context) GetLoadBalancerVmHealthResultOutput {
	return o
}

func (o GetLoadBalancerVmHealthResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLoadBalancerVmHealthResult] {
	return pulumix.Output[GetLoadBalancerVmHealthResult]{
		OutputState: o.OutputState,
	}
}

func (o GetLoadBalancerVmHealthResultOutput) BackendVmHealths() GetLoadBalancerVmHealthBackendVmHealthArrayOutput {
	return o.ApplyT(func(v GetLoadBalancerVmHealthResult) []GetLoadBalancerVmHealthBackendVmHealth {
		return v.BackendVmHealths
	}).(GetLoadBalancerVmHealthBackendVmHealthArrayOutput)
}

func (o GetLoadBalancerVmHealthResultOutput) BackendVmIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancerVmHealthResult) []string { return v.BackendVmIds }).(pulumi.StringArrayOutput)
}

func (o GetLoadBalancerVmHealthResultOutput) Filters() GetLoadBalancerVmHealthFilterArrayOutput {
	return o.ApplyT(func(v GetLoadBalancerVmHealthResult) []GetLoadBalancerVmHealthFilter { return v.Filters }).(GetLoadBalancerVmHealthFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLoadBalancerVmHealthResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancerVmHealthResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLoadBalancerVmHealthResultOutput) LoadBalancerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancerVmHealthResult) string { return v.LoadBalancerName }).(pulumi.StringOutput)
}

func (o GetLoadBalancerVmHealthResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancerVmHealthResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoadBalancerVmHealthResultOutput{})
}
