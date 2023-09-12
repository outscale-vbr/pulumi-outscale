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

type Image struct {
	pulumi.CustomResourceState

	AccountAlias          pulumi.StringOutput                 `pulumi:"accountAlias"`
	AccountId             pulumi.StringOutput                 `pulumi:"accountId"`
	Architecture          pulumi.StringOutput                 `pulumi:"architecture"`
	BlockDeviceMappings   ImageBlockDeviceMappingArrayOutput  `pulumi:"blockDeviceMappings"`
	CreationDate          pulumi.StringOutput                 `pulumi:"creationDate"`
	Description           pulumi.StringOutput                 `pulumi:"description"`
	FileLocation          pulumi.StringOutput                 `pulumi:"fileLocation"`
	ImageId               pulumi.StringOutput                 `pulumi:"imageId"`
	ImageName             pulumi.StringOutput                 `pulumi:"imageName"`
	ImageType             pulumi.StringOutput                 `pulumi:"imageType"`
	IsPublic              pulumi.BoolOutput                   `pulumi:"isPublic"`
	NoReboot              pulumi.BoolOutput                   `pulumi:"noReboot"`
	PermissionsToLaunches ImagePermissionsToLaunchArrayOutput `pulumi:"permissionsToLaunches"`
	ProductCodes          pulumi.StringArrayOutput            `pulumi:"productCodes"`
	RequestId             pulumi.StringOutput                 `pulumi:"requestId"`
	RootDeviceName        pulumi.StringOutput                 `pulumi:"rootDeviceName"`
	RootDeviceType        pulumi.StringOutput                 `pulumi:"rootDeviceType"`
	SourceImageId         pulumi.StringOutput                 `pulumi:"sourceImageId"`
	SourceRegionName      pulumi.StringPtrOutput              `pulumi:"sourceRegionName"`
	State                 pulumi.StringOutput                 `pulumi:"state"`
	StateComments         ImageStateCommentArrayOutput        `pulumi:"stateComments"`
	Tags                  ImageTagArrayOutput                 `pulumi:"tags"`
	VmId                  pulumi.StringOutput                 `pulumi:"vmId"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		args = &ImageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("outscale:index/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("outscale:index/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	AccountAlias          *string                    `pulumi:"accountAlias"`
	AccountId             *string                    `pulumi:"accountId"`
	Architecture          *string                    `pulumi:"architecture"`
	BlockDeviceMappings   []ImageBlockDeviceMapping  `pulumi:"blockDeviceMappings"`
	CreationDate          *string                    `pulumi:"creationDate"`
	Description           *string                    `pulumi:"description"`
	FileLocation          *string                    `pulumi:"fileLocation"`
	ImageId               *string                    `pulumi:"imageId"`
	ImageName             *string                    `pulumi:"imageName"`
	ImageType             *string                    `pulumi:"imageType"`
	IsPublic              *bool                      `pulumi:"isPublic"`
	NoReboot              *bool                      `pulumi:"noReboot"`
	PermissionsToLaunches []ImagePermissionsToLaunch `pulumi:"permissionsToLaunches"`
	ProductCodes          []string                   `pulumi:"productCodes"`
	RequestId             *string                    `pulumi:"requestId"`
	RootDeviceName        *string                    `pulumi:"rootDeviceName"`
	RootDeviceType        *string                    `pulumi:"rootDeviceType"`
	SourceImageId         *string                    `pulumi:"sourceImageId"`
	SourceRegionName      *string                    `pulumi:"sourceRegionName"`
	State                 *string                    `pulumi:"state"`
	StateComments         []ImageStateComment        `pulumi:"stateComments"`
	Tags                  []ImageTag                 `pulumi:"tags"`
	VmId                  *string                    `pulumi:"vmId"`
}

type ImageState struct {
	AccountAlias          pulumi.StringPtrInput
	AccountId             pulumi.StringPtrInput
	Architecture          pulumi.StringPtrInput
	BlockDeviceMappings   ImageBlockDeviceMappingArrayInput
	CreationDate          pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	FileLocation          pulumi.StringPtrInput
	ImageId               pulumi.StringPtrInput
	ImageName             pulumi.StringPtrInput
	ImageType             pulumi.StringPtrInput
	IsPublic              pulumi.BoolPtrInput
	NoReboot              pulumi.BoolPtrInput
	PermissionsToLaunches ImagePermissionsToLaunchArrayInput
	ProductCodes          pulumi.StringArrayInput
	RequestId             pulumi.StringPtrInput
	RootDeviceName        pulumi.StringPtrInput
	RootDeviceType        pulumi.StringPtrInput
	SourceImageId         pulumi.StringPtrInput
	SourceRegionName      pulumi.StringPtrInput
	State                 pulumi.StringPtrInput
	StateComments         ImageStateCommentArrayInput
	Tags                  ImageTagArrayInput
	VmId                  pulumi.StringPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	Architecture        *string                   `pulumi:"architecture"`
	BlockDeviceMappings []ImageBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	Description         *string                   `pulumi:"description"`
	FileLocation        *string                   `pulumi:"fileLocation"`
	ImageName           *string                   `pulumi:"imageName"`
	NoReboot            *bool                     `pulumi:"noReboot"`
	RootDeviceName      *string                   `pulumi:"rootDeviceName"`
	SourceImageId       *string                   `pulumi:"sourceImageId"`
	SourceRegionName    *string                   `pulumi:"sourceRegionName"`
	Tags                []ImageTag                `pulumi:"tags"`
	VmId                *string                   `pulumi:"vmId"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	Architecture        pulumi.StringPtrInput
	BlockDeviceMappings ImageBlockDeviceMappingArrayInput
	Description         pulumi.StringPtrInput
	FileLocation        pulumi.StringPtrInput
	ImageName           pulumi.StringPtrInput
	NoReboot            pulumi.BoolPtrInput
	RootDeviceName      pulumi.StringPtrInput
	SourceImageId       pulumi.StringPtrInput
	SourceRegionName    pulumi.StringPtrInput
	Tags                ImageTagArrayInput
	VmId                pulumi.StringPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

func (i *Image) ToOutput(ctx context.Context) pulumix.Output[*Image] {
	return pulumix.Output[*Image]{
		OutputState: i.ToImageOutputWithContext(ctx).OutputState,
	}
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//	ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

func (i ImageArray) ToOutput(ctx context.Context) pulumix.Output[[]*Image] {
	return pulumix.Output[[]*Image]{
		OutputState: i.ToImageArrayOutputWithContext(ctx).OutputState,
	}
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//	ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

func (i ImageMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Image] {
	return pulumix.Output[map[string]*Image]{
		OutputState: i.ToImageMapOutputWithContext(ctx).OutputState,
	}
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func (o ImageOutput) ToOutput(ctx context.Context) pulumix.Output[*Image] {
	return pulumix.Output[*Image]{
		OutputState: o.OutputState,
	}
}

func (o ImageOutput) AccountAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.AccountAlias }).(pulumi.StringOutput)
}

func (o ImageOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o ImageOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Architecture }).(pulumi.StringOutput)
}

func (o ImageOutput) BlockDeviceMappings() ImageBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v *Image) ImageBlockDeviceMappingArrayOutput { return v.BlockDeviceMappings }).(ImageBlockDeviceMappingArrayOutput)
}

func (o ImageOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ImageOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ImageOutput) FileLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.FileLocation }).(pulumi.StringOutput)
}

func (o ImageOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

func (o ImageOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageType }).(pulumi.StringOutput)
}

func (o ImageOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolOutput { return v.IsPublic }).(pulumi.BoolOutput)
}

func (o ImageOutput) NoReboot() pulumi.BoolOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolOutput { return v.NoReboot }).(pulumi.BoolOutput)
}

func (o ImageOutput) PermissionsToLaunches() ImagePermissionsToLaunchArrayOutput {
	return o.ApplyT(func(v *Image) ImagePermissionsToLaunchArrayOutput { return v.PermissionsToLaunches }).(ImagePermissionsToLaunchArrayOutput)
}

func (o ImageOutput) ProductCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.ProductCodes }).(pulumi.StringArrayOutput)
}

func (o ImageOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.RequestId }).(pulumi.StringOutput)
}

func (o ImageOutput) RootDeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.RootDeviceName }).(pulumi.StringOutput)
}

func (o ImageOutput) RootDeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.RootDeviceType }).(pulumi.StringOutput)
}

func (o ImageOutput) SourceImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.SourceImageId }).(pulumi.StringOutput)
}

func (o ImageOutput) SourceRegionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.SourceRegionName }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ImageOutput) StateComments() ImageStateCommentArrayOutput {
	return o.ApplyT(func(v *Image) ImageStateCommentArrayOutput { return v.StateComments }).(ImageStateCommentArrayOutput)
}

func (o ImageOutput) Tags() ImageTagArrayOutput {
	return o.ApplyT(func(v *Image) ImageTagArrayOutput { return v.Tags }).(ImageTagArrayOutput)
}

func (o ImageOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Image] {
	return pulumix.Output[[]*Image]{
		OutputState: o.OutputState,
	}
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Image] {
	return pulumix.Output[map[string]*Image]{
		OutputState: o.OutputState,
	}
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
