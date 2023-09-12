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

type Nic struct {
	pulumi.CustomResourceState

	AccountId           pulumi.StringOutput         `pulumi:"accountId"`
	Description         pulumi.StringOutput         `pulumi:"description"`
	IsSourceDestChecked pulumi.BoolOutput           `pulumi:"isSourceDestChecked"`
	LinkNic             NicLinkNicOutput            `pulumi:"linkNic"`
	LinkPublicIp        NicLinkPublicIpOutput       `pulumi:"linkPublicIp"`
	MacAddress          pulumi.StringOutput         `pulumi:"macAddress"`
	NetId               pulumi.StringOutput         `pulumi:"netId"`
	NicId               pulumi.StringOutput         `pulumi:"nicId"`
	PrivateDnsName      pulumi.StringOutput         `pulumi:"privateDnsName"`
	PrivateIp           pulumi.StringOutput         `pulumi:"privateIp"`
	PrivateIps          NicPrivateIpTypeArrayOutput `pulumi:"privateIps"`
	RequestId           pulumi.StringOutput         `pulumi:"requestId"`
	RequesterManaged    pulumi.BoolOutput           `pulumi:"requesterManaged"`
	SecurityGroupIds    pulumi.StringArrayOutput    `pulumi:"securityGroupIds"`
	SecurityGroups      NicSecurityGroupArrayOutput `pulumi:"securityGroups"`
	State               pulumi.StringOutput         `pulumi:"state"`
	SubnetId            pulumi.StringOutput         `pulumi:"subnetId"`
	SubregionName       pulumi.StringOutput         `pulumi:"subregionName"`
	Tags                NicTagArrayOutput           `pulumi:"tags"`
}

// NewNic registers a new resource with the given unique name, arguments, and options.
func NewNic(ctx *pulumi.Context,
	name string, args *NicArgs, opts ...pulumi.ResourceOption) (*Nic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nic
	err := ctx.RegisterResource("outscale:index/nic:Nic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNic gets an existing Nic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NicState, opts ...pulumi.ResourceOption) (*Nic, error) {
	var resource Nic
	err := ctx.ReadResource("outscale:index/nic:Nic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nic resources.
type nicState struct {
	AccountId           *string            `pulumi:"accountId"`
	Description         *string            `pulumi:"description"`
	IsSourceDestChecked *bool              `pulumi:"isSourceDestChecked"`
	LinkNic             *NicLinkNic        `pulumi:"linkNic"`
	LinkPublicIp        *NicLinkPublicIp   `pulumi:"linkPublicIp"`
	MacAddress          *string            `pulumi:"macAddress"`
	NetId               *string            `pulumi:"netId"`
	NicId               *string            `pulumi:"nicId"`
	PrivateDnsName      *string            `pulumi:"privateDnsName"`
	PrivateIp           *string            `pulumi:"privateIp"`
	PrivateIps          []NicPrivateIpType `pulumi:"privateIps"`
	RequestId           *string            `pulumi:"requestId"`
	RequesterManaged    *bool              `pulumi:"requesterManaged"`
	SecurityGroupIds    []string           `pulumi:"securityGroupIds"`
	SecurityGroups      []NicSecurityGroup `pulumi:"securityGroups"`
	State               *string            `pulumi:"state"`
	SubnetId            *string            `pulumi:"subnetId"`
	SubregionName       *string            `pulumi:"subregionName"`
	Tags                []NicTag           `pulumi:"tags"`
}

type NicState struct {
	AccountId           pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	IsSourceDestChecked pulumi.BoolPtrInput
	LinkNic             NicLinkNicPtrInput
	LinkPublicIp        NicLinkPublicIpPtrInput
	MacAddress          pulumi.StringPtrInput
	NetId               pulumi.StringPtrInput
	NicId               pulumi.StringPtrInput
	PrivateDnsName      pulumi.StringPtrInput
	PrivateIp           pulumi.StringPtrInput
	PrivateIps          NicPrivateIpTypeArrayInput
	RequestId           pulumi.StringPtrInput
	RequesterManaged    pulumi.BoolPtrInput
	SecurityGroupIds    pulumi.StringArrayInput
	SecurityGroups      NicSecurityGroupArrayInput
	State               pulumi.StringPtrInput
	SubnetId            pulumi.StringPtrInput
	SubregionName       pulumi.StringPtrInput
	Tags                NicTagArrayInput
}

func (NicState) ElementType() reflect.Type {
	return reflect.TypeOf((*nicState)(nil)).Elem()
}

type nicArgs struct {
	Description      *string            `pulumi:"description"`
	PrivateIp        *string            `pulumi:"privateIp"`
	PrivateIps       []NicPrivateIpType `pulumi:"privateIps"`
	SecurityGroupIds []string           `pulumi:"securityGroupIds"`
	SubnetId         string             `pulumi:"subnetId"`
	Tags             []NicTag           `pulumi:"tags"`
}

// The set of arguments for constructing a Nic resource.
type NicArgs struct {
	Description      pulumi.StringPtrInput
	PrivateIp        pulumi.StringPtrInput
	PrivateIps       NicPrivateIpTypeArrayInput
	SecurityGroupIds pulumi.StringArrayInput
	SubnetId         pulumi.StringInput
	Tags             NicTagArrayInput
}

func (NicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nicArgs)(nil)).Elem()
}

type NicInput interface {
	pulumi.Input

	ToNicOutput() NicOutput
	ToNicOutputWithContext(ctx context.Context) NicOutput
}

func (*Nic) ElementType() reflect.Type {
	return reflect.TypeOf((**Nic)(nil)).Elem()
}

func (i *Nic) ToNicOutput() NicOutput {
	return i.ToNicOutputWithContext(context.Background())
}

func (i *Nic) ToNicOutputWithContext(ctx context.Context) NicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicOutput)
}

func (i *Nic) ToOutput(ctx context.Context) pulumix.Output[*Nic] {
	return pulumix.Output[*Nic]{
		OutputState: i.ToNicOutputWithContext(ctx).OutputState,
	}
}

// NicArrayInput is an input type that accepts NicArray and NicArrayOutput values.
// You can construct a concrete instance of `NicArrayInput` via:
//
//	NicArray{ NicArgs{...} }
type NicArrayInput interface {
	pulumi.Input

	ToNicArrayOutput() NicArrayOutput
	ToNicArrayOutputWithContext(context.Context) NicArrayOutput
}

type NicArray []NicInput

func (NicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nic)(nil)).Elem()
}

func (i NicArray) ToNicArrayOutput() NicArrayOutput {
	return i.ToNicArrayOutputWithContext(context.Background())
}

func (i NicArray) ToNicArrayOutputWithContext(ctx context.Context) NicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicArrayOutput)
}

func (i NicArray) ToOutput(ctx context.Context) pulumix.Output[[]*Nic] {
	return pulumix.Output[[]*Nic]{
		OutputState: i.ToNicArrayOutputWithContext(ctx).OutputState,
	}
}

// NicMapInput is an input type that accepts NicMap and NicMapOutput values.
// You can construct a concrete instance of `NicMapInput` via:
//
//	NicMap{ "key": NicArgs{...} }
type NicMapInput interface {
	pulumi.Input

	ToNicMapOutput() NicMapOutput
	ToNicMapOutputWithContext(context.Context) NicMapOutput
}

type NicMap map[string]NicInput

func (NicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nic)(nil)).Elem()
}

func (i NicMap) ToNicMapOutput() NicMapOutput {
	return i.ToNicMapOutputWithContext(context.Background())
}

func (i NicMap) ToNicMapOutputWithContext(ctx context.Context) NicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicMapOutput)
}

func (i NicMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Nic] {
	return pulumix.Output[map[string]*Nic]{
		OutputState: i.ToNicMapOutputWithContext(ctx).OutputState,
	}
}

type NicOutput struct{ *pulumi.OutputState }

func (NicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nic)(nil)).Elem()
}

func (o NicOutput) ToNicOutput() NicOutput {
	return o
}

func (o NicOutput) ToNicOutputWithContext(ctx context.Context) NicOutput {
	return o
}

func (o NicOutput) ToOutput(ctx context.Context) pulumix.Output[*Nic] {
	return pulumix.Output[*Nic]{
		OutputState: o.OutputState,
	}
}

func (o NicOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o NicOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o NicOutput) IsSourceDestChecked() pulumi.BoolOutput {
	return o.ApplyT(func(v *Nic) pulumi.BoolOutput { return v.IsSourceDestChecked }).(pulumi.BoolOutput)
}

func (o NicOutput) LinkNic() NicLinkNicOutput {
	return o.ApplyT(func(v *Nic) NicLinkNicOutput { return v.LinkNic }).(NicLinkNicOutput)
}

func (o NicOutput) LinkPublicIp() NicLinkPublicIpOutput {
	return o.ApplyT(func(v *Nic) NicLinkPublicIpOutput { return v.LinkPublicIp }).(NicLinkPublicIpOutput)
}

func (o NicOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

func (o NicOutput) NetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.NetId }).(pulumi.StringOutput)
}

func (o NicOutput) NicId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.NicId }).(pulumi.StringOutput)
}

func (o NicOutput) PrivateDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.PrivateDnsName }).(pulumi.StringOutput)
}

func (o NicOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o NicOutput) PrivateIps() NicPrivateIpTypeArrayOutput {
	return o.ApplyT(func(v *Nic) NicPrivateIpTypeArrayOutput { return v.PrivateIps }).(NicPrivateIpTypeArrayOutput)
}

func (o NicOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o NicOutput) RequesterManaged() pulumi.BoolOutput {
	return o.ApplyT(func(v *Nic) pulumi.BoolOutput { return v.RequesterManaged }).(pulumi.BoolOutput)
}

func (o NicOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o NicOutput) SecurityGroups() NicSecurityGroupArrayOutput {
	return o.ApplyT(func(v *Nic) NicSecurityGroupArrayOutput { return v.SecurityGroups }).(NicSecurityGroupArrayOutput)
}

func (o NicOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o NicOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o NicOutput) SubregionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.SubregionName }).(pulumi.StringOutput)
}

func (o NicOutput) Tags() NicTagArrayOutput {
	return o.ApplyT(func(v *Nic) NicTagArrayOutput { return v.Tags }).(NicTagArrayOutput)
}

type NicArrayOutput struct{ *pulumi.OutputState }

func (NicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nic)(nil)).Elem()
}

func (o NicArrayOutput) ToNicArrayOutput() NicArrayOutput {
	return o
}

func (o NicArrayOutput) ToNicArrayOutputWithContext(ctx context.Context) NicArrayOutput {
	return o
}

func (o NicArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Nic] {
	return pulumix.Output[[]*Nic]{
		OutputState: o.OutputState,
	}
}

func (o NicArrayOutput) Index(i pulumi.IntInput) NicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nic {
		return vs[0].([]*Nic)[vs[1].(int)]
	}).(NicOutput)
}

type NicMapOutput struct{ *pulumi.OutputState }

func (NicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nic)(nil)).Elem()
}

func (o NicMapOutput) ToNicMapOutput() NicMapOutput {
	return o
}

func (o NicMapOutput) ToNicMapOutputWithContext(ctx context.Context) NicMapOutput {
	return o
}

func (o NicMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Nic] {
	return pulumix.Output[map[string]*Nic]{
		OutputState: o.OutputState,
	}
}

func (o NicMapOutput) MapIndex(k pulumi.StringInput) NicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nic {
		return vs[0].(map[string]*Nic)[vs[1].(string)]
	}).(NicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NicInput)(nil)).Elem(), &Nic{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicArrayInput)(nil)).Elem(), NicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicMapInput)(nil)).Elem(), NicMap{})
	pulumi.RegisterOutputType(NicOutput{})
	pulumi.RegisterOutputType(NicArrayOutput{})
	pulumi.RegisterOutputType(NicMapOutput{})
}