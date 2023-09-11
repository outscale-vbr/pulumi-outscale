# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FlexibleGpuLinkArgs', 'FlexibleGpuLink']

@pulumi.input_type
class FlexibleGpuLinkArgs:
    def __init__(__self__, *,
                 flexible_gpu_id: pulumi.Input[str],
                 vm_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a FlexibleGpuLink resource.
        """
        pulumi.set(__self__, "flexible_gpu_id", flexible_gpu_id)
        pulumi.set(__self__, "vm_id", vm_id)

    @property
    @pulumi.getter(name="flexibleGpuId")
    def flexible_gpu_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "flexible_gpu_id")

    @flexible_gpu_id.setter
    def flexible_gpu_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "flexible_gpu_id", value)

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vm_id")

    @vm_id.setter
    def vm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vm_id", value)


@pulumi.input_type
class _FlexibleGpuLinkState:
    def __init__(__self__, *,
                 flexible_gpu_id: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 vm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FlexibleGpuLink resources.
        """
        if flexible_gpu_id is not None:
            pulumi.set(__self__, "flexible_gpu_id", flexible_gpu_id)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if vm_id is not None:
            pulumi.set(__self__, "vm_id", vm_id)

    @property
    @pulumi.getter(name="flexibleGpuId")
    def flexible_gpu_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "flexible_gpu_id")

    @flexible_gpu_id.setter
    def flexible_gpu_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flexible_gpu_id", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vm_id")

    @vm_id.setter
    def vm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vm_id", value)


class FlexibleGpuLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flexible_gpu_id: Optional[pulumi.Input[str]] = None,
                 vm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FlexibleGpuLink resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlexibleGpuLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FlexibleGpuLink resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FlexibleGpuLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlexibleGpuLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flexible_gpu_id: Optional[pulumi.Input[str]] = None,
                 vm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlexibleGpuLinkArgs.__new__(FlexibleGpuLinkArgs)

            if flexible_gpu_id is None and not opts.urn:
                raise TypeError("Missing required property 'flexible_gpu_id'")
            __props__.__dict__["flexible_gpu_id"] = flexible_gpu_id
            if vm_id is None and not opts.urn:
                raise TypeError("Missing required property 'vm_id'")
            __props__.__dict__["vm_id"] = vm_id
            __props__.__dict__["request_id"] = None
        super(FlexibleGpuLink, __self__).__init__(
            'outscale:index/flexibleGpuLink:FlexibleGpuLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            flexible_gpu_id: Optional[pulumi.Input[str]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            vm_id: Optional[pulumi.Input[str]] = None) -> 'FlexibleGpuLink':
        """
        Get an existing FlexibleGpuLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlexibleGpuLinkState.__new__(_FlexibleGpuLinkState)

        __props__.__dict__["flexible_gpu_id"] = flexible_gpu_id
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["vm_id"] = vm_id
        return FlexibleGpuLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="flexibleGpuId")
    def flexible_gpu_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "flexible_gpu_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vm_id")

