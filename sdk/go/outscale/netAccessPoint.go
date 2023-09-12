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

type NetAccessPoint struct {
	pulumi.CustomResourceState

	NetAccessPointId pulumi.StringOutput          `pulumi:"netAccessPointId"`
	NetId            pulumi.StringOutput          `pulumi:"netId"`
	RequestId        pulumi.StringOutput          `pulumi:"requestId"`
	RouteTableIds    pulumi.StringArrayOutput     `pulumi:"routeTableIds"`
	ServiceName      pulumi.StringOutput          `pulumi:"serviceName"`
	State            pulumi.StringOutput          `pulumi:"state"`
	Tags             NetAccessPointTagArrayOutput `pulumi:"tags"`
}

// NewNetAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewNetAccessPoint(ctx *pulumi.Context,
	name string, args *NetAccessPointArgs, opts ...pulumi.ResourceOption) (*NetAccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetId == nil {
		return nil, errors.New("invalid value for required argument 'NetId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetAccessPoint
	err := ctx.RegisterResource("outscale:index/netAccessPoint:NetAccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetAccessPoint gets an existing NetAccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetAccessPointState, opts ...pulumi.ResourceOption) (*NetAccessPoint, error) {
	var resource NetAccessPoint
	err := ctx.ReadResource("outscale:index/netAccessPoint:NetAccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetAccessPoint resources.
type netAccessPointState struct {
	NetAccessPointId *string             `pulumi:"netAccessPointId"`
	NetId            *string             `pulumi:"netId"`
	RequestId        *string             `pulumi:"requestId"`
	RouteTableIds    []string            `pulumi:"routeTableIds"`
	ServiceName      *string             `pulumi:"serviceName"`
	State            *string             `pulumi:"state"`
	Tags             []NetAccessPointTag `pulumi:"tags"`
}

type NetAccessPointState struct {
	NetAccessPointId pulumi.StringPtrInput
	NetId            pulumi.StringPtrInput
	RequestId        pulumi.StringPtrInput
	RouteTableIds    pulumi.StringArrayInput
	ServiceName      pulumi.StringPtrInput
	State            pulumi.StringPtrInput
	Tags             NetAccessPointTagArrayInput
}

func (NetAccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*netAccessPointState)(nil)).Elem()
}

type netAccessPointArgs struct {
	NetId         string              `pulumi:"netId"`
	RouteTableIds []string            `pulumi:"routeTableIds"`
	ServiceName   string              `pulumi:"serviceName"`
	Tags          []NetAccessPointTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetAccessPoint resource.
type NetAccessPointArgs struct {
	NetId         pulumi.StringInput
	RouteTableIds pulumi.StringArrayInput
	ServiceName   pulumi.StringInput
	Tags          NetAccessPointTagArrayInput
}

func (NetAccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*netAccessPointArgs)(nil)).Elem()
}

type NetAccessPointInput interface {
	pulumi.Input

	ToNetAccessPointOutput() NetAccessPointOutput
	ToNetAccessPointOutputWithContext(ctx context.Context) NetAccessPointOutput
}

func (*NetAccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAccessPoint)(nil)).Elem()
}

func (i *NetAccessPoint) ToNetAccessPointOutput() NetAccessPointOutput {
	return i.ToNetAccessPointOutputWithContext(context.Background())
}

func (i *NetAccessPoint) ToNetAccessPointOutputWithContext(ctx context.Context) NetAccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAccessPointOutput)
}

func (i *NetAccessPoint) ToOutput(ctx context.Context) pulumix.Output[*NetAccessPoint] {
	return pulumix.Output[*NetAccessPoint]{
		OutputState: i.ToNetAccessPointOutputWithContext(ctx).OutputState,
	}
}

// NetAccessPointArrayInput is an input type that accepts NetAccessPointArray and NetAccessPointArrayOutput values.
// You can construct a concrete instance of `NetAccessPointArrayInput` via:
//
//	NetAccessPointArray{ NetAccessPointArgs{...} }
type NetAccessPointArrayInput interface {
	pulumi.Input

	ToNetAccessPointArrayOutput() NetAccessPointArrayOutput
	ToNetAccessPointArrayOutputWithContext(context.Context) NetAccessPointArrayOutput
}

type NetAccessPointArray []NetAccessPointInput

func (NetAccessPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetAccessPoint)(nil)).Elem()
}

func (i NetAccessPointArray) ToNetAccessPointArrayOutput() NetAccessPointArrayOutput {
	return i.ToNetAccessPointArrayOutputWithContext(context.Background())
}

func (i NetAccessPointArray) ToNetAccessPointArrayOutputWithContext(ctx context.Context) NetAccessPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAccessPointArrayOutput)
}

func (i NetAccessPointArray) ToOutput(ctx context.Context) pulumix.Output[[]*NetAccessPoint] {
	return pulumix.Output[[]*NetAccessPoint]{
		OutputState: i.ToNetAccessPointArrayOutputWithContext(ctx).OutputState,
	}
}

// NetAccessPointMapInput is an input type that accepts NetAccessPointMap and NetAccessPointMapOutput values.
// You can construct a concrete instance of `NetAccessPointMapInput` via:
//
//	NetAccessPointMap{ "key": NetAccessPointArgs{...} }
type NetAccessPointMapInput interface {
	pulumi.Input

	ToNetAccessPointMapOutput() NetAccessPointMapOutput
	ToNetAccessPointMapOutputWithContext(context.Context) NetAccessPointMapOutput
}

type NetAccessPointMap map[string]NetAccessPointInput

func (NetAccessPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetAccessPoint)(nil)).Elem()
}

func (i NetAccessPointMap) ToNetAccessPointMapOutput() NetAccessPointMapOutput {
	return i.ToNetAccessPointMapOutputWithContext(context.Background())
}

func (i NetAccessPointMap) ToNetAccessPointMapOutputWithContext(ctx context.Context) NetAccessPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAccessPointMapOutput)
}

func (i NetAccessPointMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NetAccessPoint] {
	return pulumix.Output[map[string]*NetAccessPoint]{
		OutputState: i.ToNetAccessPointMapOutputWithContext(ctx).OutputState,
	}
}

type NetAccessPointOutput struct{ *pulumi.OutputState }

func (NetAccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAccessPoint)(nil)).Elem()
}

func (o NetAccessPointOutput) ToNetAccessPointOutput() NetAccessPointOutput {
	return o
}

func (o NetAccessPointOutput) ToNetAccessPointOutputWithContext(ctx context.Context) NetAccessPointOutput {
	return o
}

func (o NetAccessPointOutput) ToOutput(ctx context.Context) pulumix.Output[*NetAccessPoint] {
	return pulumix.Output[*NetAccessPoint]{
		OutputState: o.OutputState,
	}
}

func (o NetAccessPointOutput) NetAccessPointId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetAccessPoint) pulumi.StringOutput { return v.NetAccessPointId }).(pulumi.StringOutput)
}

func (o NetAccessPointOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetAccessPoint) pulumi.StringOutput { return v.NetId }).(pulumi.StringOutput)
}

func (o NetAccessPointOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetAccessPoint) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o NetAccessPointOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetAccessPoint) pulumi.StringArrayOutput { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

func (o NetAccessPointOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetAccessPoint) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

func (o NetAccessPointOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetAccessPoint) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o NetAccessPointOutput) Tags() NetAccessPointTagArrayOutput {
	return o.ApplyT(func(v *NetAccessPoint) NetAccessPointTagArrayOutput { return v.Tags }).(NetAccessPointTagArrayOutput)
}

type NetAccessPointArrayOutput struct{ *pulumi.OutputState }

func (NetAccessPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetAccessPoint)(nil)).Elem()
}

func (o NetAccessPointArrayOutput) ToNetAccessPointArrayOutput() NetAccessPointArrayOutput {
	return o
}

func (o NetAccessPointArrayOutput) ToNetAccessPointArrayOutputWithContext(ctx context.Context) NetAccessPointArrayOutput {
	return o
}

func (o NetAccessPointArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NetAccessPoint] {
	return pulumix.Output[[]*NetAccessPoint]{
		OutputState: o.OutputState,
	}
}

func (o NetAccessPointArrayOutput) Index(i pulumi.IntInput) NetAccessPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetAccessPoint {
		return vs[0].([]*NetAccessPoint)[vs[1].(int)]
	}).(NetAccessPointOutput)
}

type NetAccessPointMapOutput struct{ *pulumi.OutputState }

func (NetAccessPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetAccessPoint)(nil)).Elem()
}

func (o NetAccessPointMapOutput) ToNetAccessPointMapOutput() NetAccessPointMapOutput {
	return o
}

func (o NetAccessPointMapOutput) ToNetAccessPointMapOutputWithContext(ctx context.Context) NetAccessPointMapOutput {
	return o
}

func (o NetAccessPointMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NetAccessPoint] {
	return pulumix.Output[map[string]*NetAccessPoint]{
		OutputState: o.OutputState,
	}
}

func (o NetAccessPointMapOutput) MapIndex(k pulumi.StringInput) NetAccessPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetAccessPoint {
		return vs[0].(map[string]*NetAccessPoint)[vs[1].(string)]
	}).(NetAccessPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetAccessPointInput)(nil)).Elem(), &NetAccessPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetAccessPointArrayInput)(nil)).Elem(), NetAccessPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetAccessPointMapInput)(nil)).Elem(), NetAccessPointMap{})
	pulumi.RegisterOutputType(NetAccessPointOutput{})
	pulumi.RegisterOutputType(NetAccessPointArrayOutput{})
	pulumi.RegisterOutputType(NetAccessPointMapOutput{})
}
