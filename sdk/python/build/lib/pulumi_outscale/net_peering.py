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

__all__ = ['NetPeeringArgs', 'NetPeering']

@pulumi.input_type
class NetPeeringArgs:
    def __init__(__self__, *,
                 accepter_net_id: pulumi.Input[str],
                 source_net_id: pulumi.Input[str],
                 source_net_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['NetPeeringTagArgs']]]] = None):
        """
        The set of arguments for constructing a NetPeering resource.
        """
        pulumi.set(__self__, "accepter_net_id", accepter_net_id)
        pulumi.set(__self__, "source_net_id", source_net_id)
        if source_net_account_id is not None:
            pulumi.set(__self__, "source_net_account_id", source_net_account_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accepterNetId")
    def accepter_net_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "accepter_net_id")

    @accepter_net_id.setter
    def accepter_net_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "accepter_net_id", value)

    @property
    @pulumi.getter(name="sourceNetId")
    def source_net_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_net_id")

    @source_net_id.setter
    def source_net_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_net_id", value)

    @property
    @pulumi.getter(name="sourceNetAccountId")
    def source_net_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_net_account_id")

    @source_net_account_id.setter
    def source_net_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_net_account_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetPeeringTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetPeeringTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _NetPeeringState:
    def __init__(__self__, *,
                 accepter_net: Optional[pulumi.Input['NetPeeringAccepterNetArgs']] = None,
                 accepter_net_id: Optional[pulumi.Input[str]] = None,
                 net_peering_id: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 source_net: Optional[pulumi.Input['NetPeeringSourceNetArgs']] = None,
                 source_net_account_id: Optional[pulumi.Input[str]] = None,
                 source_net_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['NetPeeringStateArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['NetPeeringTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering NetPeering resources.
        """
        if accepter_net is not None:
            pulumi.set(__self__, "accepter_net", accepter_net)
        if accepter_net_id is not None:
            pulumi.set(__self__, "accepter_net_id", accepter_net_id)
        if net_peering_id is not None:
            pulumi.set(__self__, "net_peering_id", net_peering_id)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if source_net is not None:
            pulumi.set(__self__, "source_net", source_net)
        if source_net_account_id is not None:
            pulumi.set(__self__, "source_net_account_id", source_net_account_id)
        if source_net_id is not None:
            pulumi.set(__self__, "source_net_id", source_net_id)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accepterNet")
    def accepter_net(self) -> Optional[pulumi.Input['NetPeeringAccepterNetArgs']]:
        return pulumi.get(self, "accepter_net")

    @accepter_net.setter
    def accepter_net(self, value: Optional[pulumi.Input['NetPeeringAccepterNetArgs']]):
        pulumi.set(self, "accepter_net", value)

    @property
    @pulumi.getter(name="accepterNetId")
    def accepter_net_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "accepter_net_id")

    @accepter_net_id.setter
    def accepter_net_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accepter_net_id", value)

    @property
    @pulumi.getter(name="netPeeringId")
    def net_peering_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "net_peering_id")

    @net_peering_id.setter
    def net_peering_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_peering_id", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="sourceNet")
    def source_net(self) -> Optional[pulumi.Input['NetPeeringSourceNetArgs']]:
        return pulumi.get(self, "source_net")

    @source_net.setter
    def source_net(self, value: Optional[pulumi.Input['NetPeeringSourceNetArgs']]):
        pulumi.set(self, "source_net", value)

    @property
    @pulumi.getter(name="sourceNetAccountId")
    def source_net_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_net_account_id")

    @source_net_account_id.setter
    def source_net_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_net_account_id", value)

    @property
    @pulumi.getter(name="sourceNetId")
    def source_net_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_net_id")

    @source_net_id.setter
    def source_net_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_net_id", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['NetPeeringStateArgs']]:
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['NetPeeringStateArgs']]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetPeeringTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetPeeringTagArgs']]]]):
        pulumi.set(self, "tags", value)


class NetPeering(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter_net_id: Optional[pulumi.Input[str]] = None,
                 source_net_account_id: Optional[pulumi.Input[str]] = None,
                 source_net_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetPeeringTagArgs']]]]] = None,
                 __props__=None):
        """
        Create a NetPeering resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetPeeringArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NetPeering resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NetPeeringArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetPeeringArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter_net_id: Optional[pulumi.Input[str]] = None,
                 source_net_account_id: Optional[pulumi.Input[str]] = None,
                 source_net_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetPeeringTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetPeeringArgs.__new__(NetPeeringArgs)

            if accepter_net_id is None and not opts.urn:
                raise TypeError("Missing required property 'accepter_net_id'")
            __props__.__dict__["accepter_net_id"] = accepter_net_id
            __props__.__dict__["source_net_account_id"] = source_net_account_id
            if source_net_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_net_id'")
            __props__.__dict__["source_net_id"] = source_net_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["accepter_net"] = None
            __props__.__dict__["net_peering_id"] = None
            __props__.__dict__["request_id"] = None
            __props__.__dict__["source_net"] = None
            __props__.__dict__["state"] = None
        super(NetPeering, __self__).__init__(
            'outscale:index/netPeering:NetPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accepter_net: Optional[pulumi.Input[pulumi.InputType['NetPeeringAccepterNetArgs']]] = None,
            accepter_net_id: Optional[pulumi.Input[str]] = None,
            net_peering_id: Optional[pulumi.Input[str]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            source_net: Optional[pulumi.Input[pulumi.InputType['NetPeeringSourceNetArgs']]] = None,
            source_net_account_id: Optional[pulumi.Input[str]] = None,
            source_net_id: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[pulumi.InputType['NetPeeringStateArgs']]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetPeeringTagArgs']]]]] = None) -> 'NetPeering':
        """
        Get an existing NetPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetPeeringState.__new__(_NetPeeringState)

        __props__.__dict__["accepter_net"] = accepter_net
        __props__.__dict__["accepter_net_id"] = accepter_net_id
        __props__.__dict__["net_peering_id"] = net_peering_id
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["source_net"] = source_net
        __props__.__dict__["source_net_account_id"] = source_net_account_id
        __props__.__dict__["source_net_id"] = source_net_id
        __props__.__dict__["state"] = state
        __props__.__dict__["tags"] = tags
        return NetPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accepterNet")
    def accepter_net(self) -> pulumi.Output['outputs.NetPeeringAccepterNet']:
        return pulumi.get(self, "accepter_net")

    @property
    @pulumi.getter(name="accepterNetId")
    def accepter_net_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "accepter_net_id")

    @property
    @pulumi.getter(name="netPeeringId")
    def net_peering_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "net_peering_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="sourceNet")
    def source_net(self) -> pulumi.Output['outputs.NetPeeringSourceNet']:
        return pulumi.get(self, "source_net")

    @property
    @pulumi.getter(name="sourceNetAccountId")
    def source_net_account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_net_account_id")

    @property
    @pulumi.getter(name="sourceNetId")
    def source_net_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_net_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output['outputs.NetPeeringState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.NetPeeringTag']]]:
        return pulumi.get(self, "tags")
