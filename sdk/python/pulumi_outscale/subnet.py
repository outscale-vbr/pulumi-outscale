# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SubnetArgs', 'Subnet']

@pulumi.input_type
class SubnetArgs:
    def __init__(__self__, *,
                 ip_range: pulumi.Input[str],
                 net_id: pulumi.Input[str],
                 map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
                 subregion_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SubnetTagArgs']]]] = None):
        """
        The set of arguments for constructing a Subnet resource.
        """
        pulumi.set(__self__, "ip_range", ip_range)
        pulumi.set(__self__, "net_id", net_id)
        if map_public_ip_on_launch is not None:
            pulumi.set(__self__, "map_public_ip_on_launch", map_public_ip_on_launch)
        if subregion_name is not None:
            pulumi.set(__self__, "subregion_name", subregion_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip_range")

    @ip_range.setter
    def ip_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_range", value)

    @property
    @pulumi.getter(name="netId")
    def net_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "net_id")

    @net_id.setter
    def net_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "net_id", value)

    @property
    @pulumi.getter(name="mapPublicIpOnLaunch")
    def map_public_ip_on_launch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "map_public_ip_on_launch")

    @map_public_ip_on_launch.setter
    def map_public_ip_on_launch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "map_public_ip_on_launch", value)

    @property
    @pulumi.getter(name="subregionName")
    def subregion_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subregion_name")

    @subregion_name.setter
    def subregion_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subregion_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SubnetTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SubnetTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SubnetState:
    def __init__(__self__, *,
                 available_ips_count: Optional[pulumi.Input[int]] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
                 net_id: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 subregion_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SubnetTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering Subnet resources.
        """
        if available_ips_count is not None:
            pulumi.set(__self__, "available_ips_count", available_ips_count)
        if ip_range is not None:
            pulumi.set(__self__, "ip_range", ip_range)
        if map_public_ip_on_launch is not None:
            pulumi.set(__self__, "map_public_ip_on_launch", map_public_ip_on_launch)
        if net_id is not None:
            pulumi.set(__self__, "net_id", net_id)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if subregion_name is not None:
            pulumi.set(__self__, "subregion_name", subregion_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="availableIpsCount")
    def available_ips_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "available_ips_count")

    @available_ips_count.setter
    def available_ips_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "available_ips_count", value)

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_range")

    @ip_range.setter
    def ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_range", value)

    @property
    @pulumi.getter(name="mapPublicIpOnLaunch")
    def map_public_ip_on_launch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "map_public_ip_on_launch")

    @map_public_ip_on_launch.setter
    def map_public_ip_on_launch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "map_public_ip_on_launch", value)

    @property
    @pulumi.getter(name="netId")
    def net_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "net_id")

    @net_id.setter
    def net_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_id", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="subregionName")
    def subregion_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subregion_name")

    @subregion_name.setter
    def subregion_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subregion_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SubnetTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SubnetTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Subnet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
                 net_id: Optional[pulumi.Input[str]] = None,
                 subregion_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetTagArgs']]]]] = None,
                 __props__=None):
        """
        Create a Subnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Subnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubnetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
                 net_id: Optional[pulumi.Input[str]] = None,
                 subregion_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetArgs.__new__(SubnetArgs)

            if ip_range is None and not opts.urn:
                raise TypeError("Missing required property 'ip_range'")
            __props__.__dict__["ip_range"] = ip_range
            __props__.__dict__["map_public_ip_on_launch"] = map_public_ip_on_launch
            if net_id is None and not opts.urn:
                raise TypeError("Missing required property 'net_id'")
            __props__.__dict__["net_id"] = net_id
            __props__.__dict__["subregion_name"] = subregion_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["available_ips_count"] = None
            __props__.__dict__["request_id"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["subnet_id"] = None
        super(Subnet, __self__).__init__(
            'outscale:index/subnet:Subnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            available_ips_count: Optional[pulumi.Input[int]] = None,
            ip_range: Optional[pulumi.Input[str]] = None,
            map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
            net_id: Optional[pulumi.Input[str]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            subregion_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetTagArgs']]]]] = None) -> 'Subnet':
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubnetState.__new__(_SubnetState)

        __props__.__dict__["available_ips_count"] = available_ips_count
        __props__.__dict__["ip_range"] = ip_range
        __props__.__dict__["map_public_ip_on_launch"] = map_public_ip_on_launch
        __props__.__dict__["net_id"] = net_id
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["state"] = state
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["subregion_name"] = subregion_name
        __props__.__dict__["tags"] = tags
        return Subnet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availableIpsCount")
    def available_ips_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "available_ips_count")

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter(name="mapPublicIpOnLaunch")
    def map_public_ip_on_launch(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "map_public_ip_on_launch")

    @property
    @pulumi.getter(name="netId")
    def net_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "net_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="subregionName")
    def subregion_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subregion_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.SubnetTag']]]:
        return pulumi.get(self, "tags")

