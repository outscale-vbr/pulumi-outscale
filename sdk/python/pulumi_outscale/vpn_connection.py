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

__all__ = ['VpnConnectionArgs', 'VpnConnection']

@pulumi.input_type
class VpnConnectionArgs:
    def __init__(__self__, *,
                 client_gateway_id: pulumi.Input[str],
                 connection_type: pulumi.Input[str],
                 virtual_gateway_id: pulumi.Input[str],
                 static_routes_only: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionTagArgs']]]] = None):
        """
        The set of arguments for constructing a VpnConnection resource.
        """
        pulumi.set(__self__, "client_gateway_id", client_gateway_id)
        pulumi.set(__self__, "connection_type", connection_type)
        pulumi.set(__self__, "virtual_gateway_id", virtual_gateway_id)
        if static_routes_only is not None:
            pulumi.set(__self__, "static_routes_only", static_routes_only)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clientGatewayId")
    def client_gateway_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_gateway_id")

    @client_gateway_id.setter
    def client_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_gateway_id", value)

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "connection_type")

    @connection_type.setter
    def connection_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_type", value)

    @property
    @pulumi.getter(name="virtualGatewayId")
    def virtual_gateway_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "virtual_gateway_id")

    @virtual_gateway_id.setter
    def virtual_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_gateway_id", value)

    @property
    @pulumi.getter(name="staticRoutesOnly")
    def static_routes_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "static_routes_only")

    @static_routes_only.setter
    def static_routes_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "static_routes_only", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VpnConnectionState:
    def __init__(__self__, *,
                 client_gateway_configuration: Optional[pulumi.Input[str]] = None,
                 client_gateway_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionRouteArgs']]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 static_routes_only: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionTagArgs']]]] = None,
                 vgw_telemetries: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionVgwTelemetryArgs']]]] = None,
                 virtual_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpn_connection_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpnConnection resources.
        """
        if client_gateway_configuration is not None:
            pulumi.set(__self__, "client_gateway_configuration", client_gateway_configuration)
        if client_gateway_id is not None:
            pulumi.set(__self__, "client_gateway_id", client_gateway_id)
        if connection_type is not None:
            pulumi.set(__self__, "connection_type", connection_type)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if routes is not None:
            pulumi.set(__self__, "routes", routes)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if static_routes_only is not None:
            pulumi.set(__self__, "static_routes_only", static_routes_only)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vgw_telemetries is not None:
            pulumi.set(__self__, "vgw_telemetries", vgw_telemetries)
        if virtual_gateway_id is not None:
            pulumi.set(__self__, "virtual_gateway_id", virtual_gateway_id)
        if vpn_connection_id is not None:
            pulumi.set(__self__, "vpn_connection_id", vpn_connection_id)

    @property
    @pulumi.getter(name="clientGatewayConfiguration")
    def client_gateway_configuration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_gateway_configuration")

    @client_gateway_configuration.setter
    def client_gateway_configuration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_gateway_configuration", value)

    @property
    @pulumi.getter(name="clientGatewayId")
    def client_gateway_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_gateway_id")

    @client_gateway_id.setter
    def client_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_gateway_id", value)

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connection_type")

    @connection_type.setter
    def connection_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_type", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionRouteArgs']]]]:
        return pulumi.get(self, "routes")

    @routes.setter
    def routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionRouteArgs']]]]):
        pulumi.set(self, "routes", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="staticRoutesOnly")
    def static_routes_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "static_routes_only")

    @static_routes_only.setter
    def static_routes_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "static_routes_only", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vgwTelemetries")
    def vgw_telemetries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionVgwTelemetryArgs']]]]:
        return pulumi.get(self, "vgw_telemetries")

    @vgw_telemetries.setter
    def vgw_telemetries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnConnectionVgwTelemetryArgs']]]]):
        pulumi.set(self, "vgw_telemetries", value)

    @property
    @pulumi.getter(name="virtualGatewayId")
    def virtual_gateway_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "virtual_gateway_id")

    @virtual_gateway_id.setter
    def virtual_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_gateway_id", value)

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpn_connection_id")

    @vpn_connection_id.setter
    def vpn_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_connection_id", value)


class VpnConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_gateway_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 static_routes_only: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnConnectionTagArgs']]]]] = None,
                 virtual_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VpnConnection resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpnConnection resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpnConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_gateway_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 static_routes_only: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnConnectionTagArgs']]]]] = None,
                 virtual_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpnConnectionArgs.__new__(VpnConnectionArgs)

            if client_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_gateway_id'")
            __props__.__dict__["client_gateway_id"] = client_gateway_id
            if connection_type is None and not opts.urn:
                raise TypeError("Missing required property 'connection_type'")
            __props__.__dict__["connection_type"] = connection_type
            __props__.__dict__["static_routes_only"] = static_routes_only
            __props__.__dict__["tags"] = tags
            if virtual_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_gateway_id'")
            __props__.__dict__["virtual_gateway_id"] = virtual_gateway_id
            __props__.__dict__["client_gateway_configuration"] = None
            __props__.__dict__["request_id"] = None
            __props__.__dict__["routes"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["vgw_telemetries"] = None
            __props__.__dict__["vpn_connection_id"] = None
        super(VpnConnection, __self__).__init__(
            'outscale:index/vpnConnection:VpnConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_gateway_configuration: Optional[pulumi.Input[str]] = None,
            client_gateway_id: Optional[pulumi.Input[str]] = None,
            connection_type: Optional[pulumi.Input[str]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnConnectionRouteArgs']]]]] = None,
            state: Optional[pulumi.Input[str]] = None,
            static_routes_only: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnConnectionTagArgs']]]]] = None,
            vgw_telemetries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnConnectionVgwTelemetryArgs']]]]] = None,
            virtual_gateway_id: Optional[pulumi.Input[str]] = None,
            vpn_connection_id: Optional[pulumi.Input[str]] = None) -> 'VpnConnection':
        """
        Get an existing VpnConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnConnectionState.__new__(_VpnConnectionState)

        __props__.__dict__["client_gateway_configuration"] = client_gateway_configuration
        __props__.__dict__["client_gateway_id"] = client_gateway_id
        __props__.__dict__["connection_type"] = connection_type
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["routes"] = routes
        __props__.__dict__["state"] = state
        __props__.__dict__["static_routes_only"] = static_routes_only
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vgw_telemetries"] = vgw_telemetries
        __props__.__dict__["virtual_gateway_id"] = virtual_gateway_id
        __props__.__dict__["vpn_connection_id"] = vpn_connection_id
        return VpnConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientGatewayConfiguration")
    def client_gateway_configuration(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_gateway_configuration")

    @property
    @pulumi.getter(name="clientGatewayId")
    def client_gateway_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_gateway_id")

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "connection_type")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Output[Sequence['outputs.VpnConnectionRoute']]:
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="staticRoutesOnly")
    def static_routes_only(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "static_routes_only")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.VpnConnectionTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vgwTelemetries")
    def vgw_telemetries(self) -> pulumi.Output[Sequence['outputs.VpnConnectionVgwTelemetry']]:
        return pulumi.get(self, "vgw_telemetries")

    @property
    @pulumi.getter(name="virtualGatewayId")
    def virtual_gateway_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "virtual_gateway_id")

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpn_connection_id")

