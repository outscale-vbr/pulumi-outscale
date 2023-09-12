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

type RouteTable struct {
	pulumi.CustomResourceState

	LinkRouteTables                 RouteTableLinkRouteTableArrayOutput                 `pulumi:"linkRouteTables"`
	NetId                           pulumi.StringOutput                                 `pulumi:"netId"`
	RequestId                       pulumi.StringOutput                                 `pulumi:"requestId"`
	RoutePropagatingVirtualGateways RouteTableRoutePropagatingVirtualGatewayArrayOutput `pulumi:"routePropagatingVirtualGateways"`
	RouteTableId                    pulumi.StringOutput                                 `pulumi:"routeTableId"`
	Routes                          RouteTableRouteArrayOutput                          `pulumi:"routes"`
	Tags                            RouteTableTagArrayOutput                            `pulumi:"tags"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetId == nil {
		return nil, errors.New("invalid value for required argument 'NetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteTable
	err := ctx.RegisterResource("outscale:index/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("outscale:index/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	LinkRouteTables                 []RouteTableLinkRouteTable                 `pulumi:"linkRouteTables"`
	NetId                           *string                                    `pulumi:"netId"`
	RequestId                       *string                                    `pulumi:"requestId"`
	RoutePropagatingVirtualGateways []RouteTableRoutePropagatingVirtualGateway `pulumi:"routePropagatingVirtualGateways"`
	RouteTableId                    *string                                    `pulumi:"routeTableId"`
	Routes                          []RouteTableRoute                          `pulumi:"routes"`
	Tags                            []RouteTableTag                            `pulumi:"tags"`
}

type RouteTableState struct {
	LinkRouteTables                 RouteTableLinkRouteTableArrayInput
	NetId                           pulumi.StringPtrInput
	RequestId                       pulumi.StringPtrInput
	RoutePropagatingVirtualGateways RouteTableRoutePropagatingVirtualGatewayArrayInput
	RouteTableId                    pulumi.StringPtrInput
	Routes                          RouteTableRouteArrayInput
	Tags                            RouteTableTagArrayInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	NetId string          `pulumi:"netId"`
	Tags  []RouteTableTag `pulumi:"tags"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	NetId pulumi.StringInput
	Tags  RouteTableTagArrayInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

func (i *RouteTable) ToOutput(ctx context.Context) pulumix.Output[*RouteTable] {
	return pulumix.Output[*RouteTable]{
		OutputState: i.ToRouteTableOutputWithContext(ctx).OutputState,
	}
}

// RouteTableArrayInput is an input type that accepts RouteTableArray and RouteTableArrayOutput values.
// You can construct a concrete instance of `RouteTableArrayInput` via:
//
//	RouteTableArray{ RouteTableArgs{...} }
type RouteTableArrayInput interface {
	pulumi.Input

	ToRouteTableArrayOutput() RouteTableArrayOutput
	ToRouteTableArrayOutputWithContext(context.Context) RouteTableArrayOutput
}

type RouteTableArray []RouteTableInput

func (RouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (i RouteTableArray) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return i.ToRouteTableArrayOutputWithContext(context.Background())
}

func (i RouteTableArray) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableArrayOutput)
}

func (i RouteTableArray) ToOutput(ctx context.Context) pulumix.Output[[]*RouteTable] {
	return pulumix.Output[[]*RouteTable]{
		OutputState: i.ToRouteTableArrayOutputWithContext(ctx).OutputState,
	}
}

// RouteTableMapInput is an input type that accepts RouteTableMap and RouteTableMapOutput values.
// You can construct a concrete instance of `RouteTableMapInput` via:
//
//	RouteTableMap{ "key": RouteTableArgs{...} }
type RouteTableMapInput interface {
	pulumi.Input

	ToRouteTableMapOutput() RouteTableMapOutput
	ToRouteTableMapOutputWithContext(context.Context) RouteTableMapOutput
}

type RouteTableMap map[string]RouteTableInput

func (RouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (i RouteTableMap) ToRouteTableMapOutput() RouteTableMapOutput {
	return i.ToRouteTableMapOutputWithContext(context.Background())
}

func (i RouteTableMap) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableMapOutput)
}

func (i RouteTableMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouteTable] {
	return pulumix.Output[map[string]*RouteTable]{
		OutputState: i.ToRouteTableMapOutputWithContext(ctx).OutputState,
	}
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil)).Elem()
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToOutput(ctx context.Context) pulumix.Output[*RouteTable] {
	return pulumix.Output[*RouteTable]{
		OutputState: o.OutputState,
	}
}

func (o RouteTableOutput) LinkRouteTables() RouteTableLinkRouteTableArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteTableLinkRouteTableArrayOutput { return v.LinkRouteTables }).(RouteTableLinkRouteTableArrayOutput)
}

func (o RouteTableOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.NetId }).(pulumi.StringOutput)
}

func (o RouteTableOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o RouteTableOutput) RoutePropagatingVirtualGateways() RouteTableRoutePropagatingVirtualGatewayArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteTableRoutePropagatingVirtualGatewayArrayOutput {
		return v.RoutePropagatingVirtualGateways
	}).(RouteTableRoutePropagatingVirtualGatewayArrayOutput)
}

func (o RouteTableOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteTable) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

func (o RouteTableOutput) Routes() RouteTableRouteArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteTableRouteArrayOutput { return v.Routes }).(RouteTableRouteArrayOutput)
}

func (o RouteTableOutput) Tags() RouteTableTagArrayOutput {
	return o.ApplyT(func(v *RouteTable) RouteTableTagArrayOutput { return v.Tags }).(RouteTableTagArrayOutput)
}

type RouteTableArrayOutput struct{ *pulumi.OutputState }

func (RouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteTable)(nil)).Elem()
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RouteTable] {
	return pulumix.Output[[]*RouteTable]{
		OutputState: o.OutputState,
	}
}

func (o RouteTableArrayOutput) Index(i pulumi.IntInput) RouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].([]*RouteTable)[vs[1].(int)]
	}).(RouteTableOutput)
}

type RouteTableMapOutput struct{ *pulumi.OutputState }

func (RouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteTable)(nil)).Elem()
}

func (o RouteTableMapOutput) ToRouteTableMapOutput() RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouteTable] {
	return pulumix.Output[map[string]*RouteTable]{
		OutputState: o.OutputState,
	}
}

func (o RouteTableMapOutput) MapIndex(k pulumi.StringInput) RouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteTable {
		return vs[0].(map[string]*RouteTable)[vs[1].(string)]
	}).(RouteTableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableInput)(nil)).Elem(), &RouteTable{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableArrayInput)(nil)).Elem(), RouteTableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTableMapInput)(nil)).Elem(), RouteTableMap{})
	pulumi.RegisterOutputType(RouteTableOutput{})
	pulumi.RegisterOutputType(RouteTableArrayOutput{})
	pulumi.RegisterOutputType(RouteTableMapOutput{})
}
