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

__all__ = [
    'GetDhcpOptionsResult',
    'AwaitableGetDhcpOptionsResult',
    'get_dhcp_options',
    'get_dhcp_options_output',
]

@pulumi.output_type
class GetDhcpOptionsResult:
    """
    A collection of values returned by getDhcpOptions.
    """
    def __init__(__self__, dhcp_options=None, dhcp_options_set_ids=None, filters=None, id=None, request_id=None):
        if dhcp_options and not isinstance(dhcp_options, list):
            raise TypeError("Expected argument 'dhcp_options' to be a list")
        pulumi.set(__self__, "dhcp_options", dhcp_options)
        if dhcp_options_set_ids and not isinstance(dhcp_options_set_ids, list):
            raise TypeError("Expected argument 'dhcp_options_set_ids' to be a list")
        pulumi.set(__self__, "dhcp_options_set_ids", dhcp_options_set_ids)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="dhcpOptions")
    def dhcp_options(self) -> Sequence['outputs.GetDhcpOptionsDhcpOptionResult']:
        return pulumi.get(self, "dhcp_options")

    @property
    @pulumi.getter(name="dhcpOptionsSetIds")
    def dhcp_options_set_ids(self) -> Sequence[str]:
        return pulumi.get(self, "dhcp_options_set_ids")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDhcpOptionsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        return pulumi.get(self, "request_id")


class AwaitableGetDhcpOptionsResult(GetDhcpOptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDhcpOptionsResult(
            dhcp_options=self.dhcp_options,
            dhcp_options_set_ids=self.dhcp_options_set_ids,
            filters=self.filters,
            id=self.id,
            request_id=self.request_id)


def get_dhcp_options(filters: Optional[Sequence[pulumi.InputType['GetDhcpOptionsFilterArgs']]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDhcpOptionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getDhcpOptions:getDhcpOptions', __args__, opts=opts, typ=GetDhcpOptionsResult).value

    return AwaitableGetDhcpOptionsResult(
        dhcp_options=pulumi.get(__ret__, 'dhcp_options'),
        dhcp_options_set_ids=pulumi.get(__ret__, 'dhcp_options_set_ids'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        request_id=pulumi.get(__ret__, 'request_id'))


@_utilities.lift_output_func(get_dhcp_options)
def get_dhcp_options_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDhcpOptionsFilterArgs']]]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDhcpOptionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...