# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LoadBalancerVmsArgs', 'LoadBalancerVms']

@pulumi.input_type
class LoadBalancerVmsArgs:
    def __init__(__self__, *,
                 backend_vm_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 load_balancer_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a LoadBalancerVms resource.
        """
        pulumi.set(__self__, "backend_vm_ids", backend_vm_ids)
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)

    @property
    @pulumi.getter(name="backendVmIds")
    def backend_vm_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "backend_vm_ids")

    @backend_vm_ids.setter
    def backend_vm_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "backend_vm_ids", value)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "load_balancer_name")

    @load_balancer_name.setter
    def load_balancer_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_name", value)


@pulumi.input_type
class _LoadBalancerVmsState:
    def __init__(__self__, *,
                 backend_vm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LoadBalancerVms resources.
        """
        if backend_vm_ids is not None:
            pulumi.set(__self__, "backend_vm_ids", backend_vm_ids)
        if load_balancer_name is not None:
            pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="backendVmIds")
    def backend_vm_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "backend_vm_ids")

    @backend_vm_ids.setter
    def backend_vm_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "backend_vm_ids", value)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "load_balancer_name")

    @load_balancer_name.setter
    def load_balancer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_name", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class LoadBalancerVms(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_vm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LoadBalancerVms resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadBalancerVmsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LoadBalancerVms resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LoadBalancerVmsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadBalancerVmsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_vm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadBalancerVmsArgs.__new__(LoadBalancerVmsArgs)

            if backend_vm_ids is None and not opts.urn:
                raise TypeError("Missing required property 'backend_vm_ids'")
            __props__.__dict__["backend_vm_ids"] = backend_vm_ids
            if load_balancer_name is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_name'")
            __props__.__dict__["load_balancer_name"] = load_balancer_name
            __props__.__dict__["request_id"] = None
        super(LoadBalancerVms, __self__).__init__(
            'outscale:index/loadBalancerVms:LoadBalancerVms',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_vm_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            load_balancer_name: Optional[pulumi.Input[str]] = None,
            request_id: Optional[pulumi.Input[str]] = None) -> 'LoadBalancerVms':
        """
        Get an existing LoadBalancerVms resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoadBalancerVmsState.__new__(_LoadBalancerVmsState)

        __props__.__dict__["backend_vm_ids"] = backend_vm_ids
        __props__.__dict__["load_balancer_name"] = load_balancer_name
        __props__.__dict__["request_id"] = request_id
        return LoadBalancerVms(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendVmIds")
    def backend_vm_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "backend_vm_ids")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")
