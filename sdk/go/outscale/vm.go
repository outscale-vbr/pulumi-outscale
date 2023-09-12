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

type Vm struct {
	pulumi.CustomResourceState

	AdminPassword               pulumi.StringOutput                     `pulumi:"adminPassword"`
	Architecture                pulumi.StringOutput                     `pulumi:"architecture"`
	BlockDeviceMappings         VmBlockDeviceMappingArrayOutput         `pulumi:"blockDeviceMappings"`
	BlockDeviceMappingsCreateds VmBlockDeviceMappingsCreatedArrayOutput `pulumi:"blockDeviceMappingsCreateds"`
	BsuOptimized                pulumi.BoolOutput                       `pulumi:"bsuOptimized"`
	ClientToken                 pulumi.StringOutput                     `pulumi:"clientToken"`
	CreationDate                pulumi.StringOutput                     `pulumi:"creationDate"`
	DeletionProtection          pulumi.BoolOutput                       `pulumi:"deletionProtection"`
	GetAdminPassword            pulumi.BoolPtrOutput                    `pulumi:"getAdminPassword"`
	Hypervisor                  pulumi.StringOutput                     `pulumi:"hypervisor"`
	ImageId                     pulumi.StringOutput                     `pulumi:"imageId"`
	IsSourceDestChecked         pulumi.BoolOutput                       `pulumi:"isSourceDestChecked"`
	KeypairName                 pulumi.StringOutput                     `pulumi:"keypairName"`
	LaunchNumber                pulumi.IntOutput                        `pulumi:"launchNumber"`
	NestedVirtualization        pulumi.BoolPtrOutput                    `pulumi:"nestedVirtualization"`
	NetId                       pulumi.StringOutput                     `pulumi:"netId"`
	Nics                        VmNicArrayOutput                        `pulumi:"nics"`
	OsFamily                    pulumi.StringOutput                     `pulumi:"osFamily"`
	Performance                 pulumi.StringOutput                     `pulumi:"performance"`
	PlacementSubregionName      pulumi.StringOutput                     `pulumi:"placementSubregionName"`
	PlacementTenancy            pulumi.StringOutput                     `pulumi:"placementTenancy"`
	PrivateDnsName              pulumi.StringOutput                     `pulumi:"privateDnsName"`
	PrivateIp                   pulumi.StringOutput                     `pulumi:"privateIp"`
	PrivateIps                  pulumi.StringArrayOutput                `pulumi:"privateIps"`
	ProductCodes                pulumi.StringArrayOutput                `pulumi:"productCodes"`
	PublicDnsName               pulumi.StringOutput                     `pulumi:"publicDnsName"`
	PublicIp                    pulumi.StringOutput                     `pulumi:"publicIp"`
	RequestId                   pulumi.StringOutput                     `pulumi:"requestId"`
	ReservationId               pulumi.StringOutput                     `pulumi:"reservationId"`
	RootDeviceName              pulumi.StringOutput                     `pulumi:"rootDeviceName"`
	RootDeviceType              pulumi.StringOutput                     `pulumi:"rootDeviceType"`
	SecurityGroupIds            pulumi.StringArrayOutput                `pulumi:"securityGroupIds"`
	SecurityGroupNames          pulumi.StringArrayOutput                `pulumi:"securityGroupNames"`
	SecurityGroups              VmSecurityGroupArrayOutput              `pulumi:"securityGroups"`
	State                       pulumi.StringPtrOutput                  `pulumi:"state"`
	StateReason                 pulumi.StringOutput                     `pulumi:"stateReason"`
	SubnetId                    pulumi.StringOutput                     `pulumi:"subnetId"`
	Tags                        VmTagArrayOutput                        `pulumi:"tags"`
	UserData                    pulumi.StringPtrOutput                  `pulumi:"userData"`
	VmId                        pulumi.StringOutput                     `pulumi:"vmId"`
	VmInitiatedShutdownBehavior pulumi.StringOutput                     `pulumi:"vmInitiatedShutdownBehavior"`
	VmType                      pulumi.StringOutput                     `pulumi:"vmType"`
}

// NewVm registers a new resource with the given unique name, arguments, and options.
func NewVm(ctx *pulumi.Context,
	name string, args *VmArgs, opts ...pulumi.ResourceOption) (*Vm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vm
	err := ctx.RegisterResource("outscale:index/vm:Vm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVm gets an existing Vm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmState, opts ...pulumi.ResourceOption) (*Vm, error) {
	var resource Vm
	err := ctx.ReadResource("outscale:index/vm:Vm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vm resources.
type vmState struct {
	AdminPassword               *string                        `pulumi:"adminPassword"`
	Architecture                *string                        `pulumi:"architecture"`
	BlockDeviceMappings         []VmBlockDeviceMapping         `pulumi:"blockDeviceMappings"`
	BlockDeviceMappingsCreateds []VmBlockDeviceMappingsCreated `pulumi:"blockDeviceMappingsCreateds"`
	BsuOptimized                *bool                          `pulumi:"bsuOptimized"`
	ClientToken                 *string                        `pulumi:"clientToken"`
	CreationDate                *string                        `pulumi:"creationDate"`
	DeletionProtection          *bool                          `pulumi:"deletionProtection"`
	GetAdminPassword            *bool                          `pulumi:"getAdminPassword"`
	Hypervisor                  *string                        `pulumi:"hypervisor"`
	ImageId                     *string                        `pulumi:"imageId"`
	IsSourceDestChecked         *bool                          `pulumi:"isSourceDestChecked"`
	KeypairName                 *string                        `pulumi:"keypairName"`
	LaunchNumber                *int                           `pulumi:"launchNumber"`
	NestedVirtualization        *bool                          `pulumi:"nestedVirtualization"`
	NetId                       *string                        `pulumi:"netId"`
	Nics                        []VmNic                        `pulumi:"nics"`
	OsFamily                    *string                        `pulumi:"osFamily"`
	Performance                 *string                        `pulumi:"performance"`
	PlacementSubregionName      *string                        `pulumi:"placementSubregionName"`
	PlacementTenancy            *string                        `pulumi:"placementTenancy"`
	PrivateDnsName              *string                        `pulumi:"privateDnsName"`
	PrivateIp                   *string                        `pulumi:"privateIp"`
	PrivateIps                  []string                       `pulumi:"privateIps"`
	ProductCodes                []string                       `pulumi:"productCodes"`
	PublicDnsName               *string                        `pulumi:"publicDnsName"`
	PublicIp                    *string                        `pulumi:"publicIp"`
	RequestId                   *string                        `pulumi:"requestId"`
	ReservationId               *string                        `pulumi:"reservationId"`
	RootDeviceName              *string                        `pulumi:"rootDeviceName"`
	RootDeviceType              *string                        `pulumi:"rootDeviceType"`
	SecurityGroupIds            []string                       `pulumi:"securityGroupIds"`
	SecurityGroupNames          []string                       `pulumi:"securityGroupNames"`
	SecurityGroups              []VmSecurityGroup              `pulumi:"securityGroups"`
	State                       *string                        `pulumi:"state"`
	StateReason                 *string                        `pulumi:"stateReason"`
	SubnetId                    *string                        `pulumi:"subnetId"`
	Tags                        []VmTag                        `pulumi:"tags"`
	UserData                    *string                        `pulumi:"userData"`
	VmId                        *string                        `pulumi:"vmId"`
	VmInitiatedShutdownBehavior *string                        `pulumi:"vmInitiatedShutdownBehavior"`
	VmType                      *string                        `pulumi:"vmType"`
}

type VmState struct {
	AdminPassword               pulumi.StringPtrInput
	Architecture                pulumi.StringPtrInput
	BlockDeviceMappings         VmBlockDeviceMappingArrayInput
	BlockDeviceMappingsCreateds VmBlockDeviceMappingsCreatedArrayInput
	BsuOptimized                pulumi.BoolPtrInput
	ClientToken                 pulumi.StringPtrInput
	CreationDate                pulumi.StringPtrInput
	DeletionProtection          pulumi.BoolPtrInput
	GetAdminPassword            pulumi.BoolPtrInput
	Hypervisor                  pulumi.StringPtrInput
	ImageId                     pulumi.StringPtrInput
	IsSourceDestChecked         pulumi.BoolPtrInput
	KeypairName                 pulumi.StringPtrInput
	LaunchNumber                pulumi.IntPtrInput
	NestedVirtualization        pulumi.BoolPtrInput
	NetId                       pulumi.StringPtrInput
	Nics                        VmNicArrayInput
	OsFamily                    pulumi.StringPtrInput
	Performance                 pulumi.StringPtrInput
	PlacementSubregionName      pulumi.StringPtrInput
	PlacementTenancy            pulumi.StringPtrInput
	PrivateDnsName              pulumi.StringPtrInput
	PrivateIp                   pulumi.StringPtrInput
	PrivateIps                  pulumi.StringArrayInput
	ProductCodes                pulumi.StringArrayInput
	PublicDnsName               pulumi.StringPtrInput
	PublicIp                    pulumi.StringPtrInput
	RequestId                   pulumi.StringPtrInput
	ReservationId               pulumi.StringPtrInput
	RootDeviceName              pulumi.StringPtrInput
	RootDeviceType              pulumi.StringPtrInput
	SecurityGroupIds            pulumi.StringArrayInput
	SecurityGroupNames          pulumi.StringArrayInput
	SecurityGroups              VmSecurityGroupArrayInput
	State                       pulumi.StringPtrInput
	StateReason                 pulumi.StringPtrInput
	SubnetId                    pulumi.StringPtrInput
	Tags                        VmTagArrayInput
	UserData                    pulumi.StringPtrInput
	VmId                        pulumi.StringPtrInput
	VmInitiatedShutdownBehavior pulumi.StringPtrInput
	VmType                      pulumi.StringPtrInput
}

func (VmState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmState)(nil)).Elem()
}

type vmArgs struct {
	BlockDeviceMappings         []VmBlockDeviceMapping         `pulumi:"blockDeviceMappings"`
	BlockDeviceMappingsCreateds []VmBlockDeviceMappingsCreated `pulumi:"blockDeviceMappingsCreateds"`
	BsuOptimized                *bool                          `pulumi:"bsuOptimized"`
	DeletionProtection          *bool                          `pulumi:"deletionProtection"`
	GetAdminPassword            *bool                          `pulumi:"getAdminPassword"`
	ImageId                     string                         `pulumi:"imageId"`
	IsSourceDestChecked         *bool                          `pulumi:"isSourceDestChecked"`
	KeypairName                 *string                        `pulumi:"keypairName"`
	NestedVirtualization        *bool                          `pulumi:"nestedVirtualization"`
	Nics                        []VmNic                        `pulumi:"nics"`
	Performance                 *string                        `pulumi:"performance"`
	PlacementSubregionName      *string                        `pulumi:"placementSubregionName"`
	PlacementTenancy            *string                        `pulumi:"placementTenancy"`
	PrivateIps                  []string                       `pulumi:"privateIps"`
	SecurityGroupIds            []string                       `pulumi:"securityGroupIds"`
	SecurityGroupNames          []string                       `pulumi:"securityGroupNames"`
	State                       *string                        `pulumi:"state"`
	SubnetId                    *string                        `pulumi:"subnetId"`
	Tags                        []VmTag                        `pulumi:"tags"`
	UserData                    *string                        `pulumi:"userData"`
	VmId                        *string                        `pulumi:"vmId"`
	VmInitiatedShutdownBehavior *string                        `pulumi:"vmInitiatedShutdownBehavior"`
	VmType                      *string                        `pulumi:"vmType"`
}

// The set of arguments for constructing a Vm resource.
type VmArgs struct {
	BlockDeviceMappings         VmBlockDeviceMappingArrayInput
	BlockDeviceMappingsCreateds VmBlockDeviceMappingsCreatedArrayInput
	BsuOptimized                pulumi.BoolPtrInput
	DeletionProtection          pulumi.BoolPtrInput
	GetAdminPassword            pulumi.BoolPtrInput
	ImageId                     pulumi.StringInput
	IsSourceDestChecked         pulumi.BoolPtrInput
	KeypairName                 pulumi.StringPtrInput
	NestedVirtualization        pulumi.BoolPtrInput
	Nics                        VmNicArrayInput
	Performance                 pulumi.StringPtrInput
	PlacementSubregionName      pulumi.StringPtrInput
	PlacementTenancy            pulumi.StringPtrInput
	PrivateIps                  pulumi.StringArrayInput
	SecurityGroupIds            pulumi.StringArrayInput
	SecurityGroupNames          pulumi.StringArrayInput
	State                       pulumi.StringPtrInput
	SubnetId                    pulumi.StringPtrInput
	Tags                        VmTagArrayInput
	UserData                    pulumi.StringPtrInput
	VmId                        pulumi.StringPtrInput
	VmInitiatedShutdownBehavior pulumi.StringPtrInput
	VmType                      pulumi.StringPtrInput
}

func (VmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmArgs)(nil)).Elem()
}

type VmInput interface {
	pulumi.Input

	ToVmOutput() VmOutput
	ToVmOutputWithContext(ctx context.Context) VmOutput
}

func (*Vm) ElementType() reflect.Type {
	return reflect.TypeOf((**Vm)(nil)).Elem()
}

func (i *Vm) ToVmOutput() VmOutput {
	return i.ToVmOutputWithContext(context.Background())
}

func (i *Vm) ToVmOutputWithContext(ctx context.Context) VmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmOutput)
}

func (i *Vm) ToOutput(ctx context.Context) pulumix.Output[*Vm] {
	return pulumix.Output[*Vm]{
		OutputState: i.ToVmOutputWithContext(ctx).OutputState,
	}
}

// VmArrayInput is an input type that accepts VmArray and VmArrayOutput values.
// You can construct a concrete instance of `VmArrayInput` via:
//
//	VmArray{ VmArgs{...} }
type VmArrayInput interface {
	pulumi.Input

	ToVmArrayOutput() VmArrayOutput
	ToVmArrayOutputWithContext(context.Context) VmArrayOutput
}

type VmArray []VmInput

func (VmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vm)(nil)).Elem()
}

func (i VmArray) ToVmArrayOutput() VmArrayOutput {
	return i.ToVmArrayOutputWithContext(context.Background())
}

func (i VmArray) ToVmArrayOutputWithContext(ctx context.Context) VmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmArrayOutput)
}

func (i VmArray) ToOutput(ctx context.Context) pulumix.Output[[]*Vm] {
	return pulumix.Output[[]*Vm]{
		OutputState: i.ToVmArrayOutputWithContext(ctx).OutputState,
	}
}

// VmMapInput is an input type that accepts VmMap and VmMapOutput values.
// You can construct a concrete instance of `VmMapInput` via:
//
//	VmMap{ "key": VmArgs{...} }
type VmMapInput interface {
	pulumi.Input

	ToVmMapOutput() VmMapOutput
	ToVmMapOutputWithContext(context.Context) VmMapOutput
}

type VmMap map[string]VmInput

func (VmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vm)(nil)).Elem()
}

func (i VmMap) ToVmMapOutput() VmMapOutput {
	return i.ToVmMapOutputWithContext(context.Background())
}

func (i VmMap) ToVmMapOutputWithContext(ctx context.Context) VmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmMapOutput)
}

func (i VmMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Vm] {
	return pulumix.Output[map[string]*Vm]{
		OutputState: i.ToVmMapOutputWithContext(ctx).OutputState,
	}
}

type VmOutput struct{ *pulumi.OutputState }

func (VmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vm)(nil)).Elem()
}

func (o VmOutput) ToVmOutput() VmOutput {
	return o
}

func (o VmOutput) ToVmOutputWithContext(ctx context.Context) VmOutput {
	return o
}

func (o VmOutput) ToOutput(ctx context.Context) pulumix.Output[*Vm] {
	return pulumix.Output[*Vm]{
		OutputState: o.OutputState,
	}
}

func (o VmOutput) AdminPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.AdminPassword }).(pulumi.StringOutput)
}

func (o VmOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.Architecture }).(pulumi.StringOutput)
}

func (o VmOutput) BlockDeviceMappings() VmBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v *Vm) VmBlockDeviceMappingArrayOutput { return v.BlockDeviceMappings }).(VmBlockDeviceMappingArrayOutput)
}

func (o VmOutput) BlockDeviceMappingsCreateds() VmBlockDeviceMappingsCreatedArrayOutput {
	return o.ApplyT(func(v *Vm) VmBlockDeviceMappingsCreatedArrayOutput { return v.BlockDeviceMappingsCreateds }).(VmBlockDeviceMappingsCreatedArrayOutput)
}

func (o VmOutput) BsuOptimized() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolOutput { return v.BsuOptimized }).(pulumi.BoolOutput)
}

func (o VmOutput) ClientToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.ClientToken }).(pulumi.StringOutput)
}

func (o VmOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o VmOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o VmOutput) GetAdminPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolPtrOutput { return v.GetAdminPassword }).(pulumi.BoolPtrOutput)
}

func (o VmOutput) Hypervisor() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.Hypervisor }).(pulumi.StringOutput)
}

func (o VmOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

func (o VmOutput) IsSourceDestChecked() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolOutput { return v.IsSourceDestChecked }).(pulumi.BoolOutput)
}

func (o VmOutput) KeypairName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.KeypairName }).(pulumi.StringOutput)
}

func (o VmOutput) LaunchNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Vm) pulumi.IntOutput { return v.LaunchNumber }).(pulumi.IntOutput)
}

func (o VmOutput) NestedVirtualization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.BoolPtrOutput { return v.NestedVirtualization }).(pulumi.BoolPtrOutput)
}

func (o VmOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.NetId }).(pulumi.StringOutput)
}

func (o VmOutput) Nics() VmNicArrayOutput {
	return o.ApplyT(func(v *Vm) VmNicArrayOutput { return v.Nics }).(VmNicArrayOutput)
}

func (o VmOutput) OsFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.OsFamily }).(pulumi.StringOutput)
}

func (o VmOutput) Performance() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.Performance }).(pulumi.StringOutput)
}

func (o VmOutput) PlacementSubregionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.PlacementSubregionName }).(pulumi.StringOutput)
}

func (o VmOutput) PlacementTenancy() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.PlacementTenancy }).(pulumi.StringOutput)
}

func (o VmOutput) PrivateDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.PrivateDnsName }).(pulumi.StringOutput)
}

func (o VmOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o VmOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

func (o VmOutput) ProductCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.ProductCodes }).(pulumi.StringArrayOutput)
}

func (o VmOutput) PublicDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.PublicDnsName }).(pulumi.StringOutput)
}

func (o VmOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

func (o VmOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o VmOutput) ReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.ReservationId }).(pulumi.StringOutput)
}

func (o VmOutput) RootDeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.RootDeviceName }).(pulumi.StringOutput)
}

func (o VmOutput) RootDeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.RootDeviceType }).(pulumi.StringOutput)
}

func (o VmOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o VmOutput) SecurityGroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringArrayOutput { return v.SecurityGroupNames }).(pulumi.StringArrayOutput)
}

func (o VmOutput) SecurityGroups() VmSecurityGroupArrayOutput {
	return o.ApplyT(func(v *Vm) VmSecurityGroupArrayOutput { return v.SecurityGroups }).(VmSecurityGroupArrayOutput)
}

func (o VmOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o VmOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

func (o VmOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VmOutput) Tags() VmTagArrayOutput {
	return o.ApplyT(func(v *Vm) VmTagArrayOutput { return v.Tags }).(VmTagArrayOutput)
}

func (o VmOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

func (o VmOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func (o VmOutput) VmInitiatedShutdownBehavior() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.VmInitiatedShutdownBehavior }).(pulumi.StringOutput)
}

func (o VmOutput) VmType() pulumi.StringOutput {
	return o.ApplyT(func(v *Vm) pulumi.StringOutput { return v.VmType }).(pulumi.StringOutput)
}

type VmArrayOutput struct{ *pulumi.OutputState }

func (VmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vm)(nil)).Elem()
}

func (o VmArrayOutput) ToVmArrayOutput() VmArrayOutput {
	return o
}

func (o VmArrayOutput) ToVmArrayOutputWithContext(ctx context.Context) VmArrayOutput {
	return o
}

func (o VmArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Vm] {
	return pulumix.Output[[]*Vm]{
		OutputState: o.OutputState,
	}
}

func (o VmArrayOutput) Index(i pulumi.IntInput) VmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vm {
		return vs[0].([]*Vm)[vs[1].(int)]
	}).(VmOutput)
}

type VmMapOutput struct{ *pulumi.OutputState }

func (VmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vm)(nil)).Elem()
}

func (o VmMapOutput) ToVmMapOutput() VmMapOutput {
	return o
}

func (o VmMapOutput) ToVmMapOutputWithContext(ctx context.Context) VmMapOutput {
	return o
}

func (o VmMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Vm] {
	return pulumix.Output[map[string]*Vm]{
		OutputState: o.OutputState,
	}
}

func (o VmMapOutput) MapIndex(k pulumi.StringInput) VmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vm {
		return vs[0].(map[string]*Vm)[vs[1].(string)]
	}).(VmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmInput)(nil)).Elem(), &Vm{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmArrayInput)(nil)).Elem(), VmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmMapInput)(nil)).Elem(), VmMap{})
	pulumi.RegisterOutputType(VmOutput{})
	pulumi.RegisterOutputType(VmArrayOutput{})
	pulumi.RegisterOutputType(VmMapOutput{})
}