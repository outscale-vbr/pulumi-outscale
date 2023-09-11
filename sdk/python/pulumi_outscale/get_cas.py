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
    'GetCasResult',
    'AwaitableGetCasResult',
    'get_cas',
    'get_cas_output',
]

@pulumi.output_type
class GetCasResult:
    """
    A collection of values returned by getCas.
    """
    def __init__(__self__, cas=None, filters=None, id=None, request_id=None):
        if cas and not isinstance(cas, list):
            raise TypeError("Expected argument 'cas' to be a list")
        pulumi.set(__self__, "cas", cas)
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
    @pulumi.getter
    def cas(self) -> Sequence['outputs.GetCasCaResult']:
        return pulumi.get(self, "cas")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetCasFilterResult']]:
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


class AwaitableGetCasResult(GetCasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCasResult(
            cas=self.cas,
            filters=self.filters,
            id=self.id,
            request_id=self.request_id)


def get_cas(filters: Optional[Sequence[pulumi.InputType['GetCasFilterArgs']]] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCasResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getCas:getCas', __args__, opts=opts, typ=GetCasResult).value

    return AwaitableGetCasResult(
        cas=pulumi.get(__ret__, 'cas'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        request_id=pulumi.get(__ret__, 'request_id'))


@_utilities.lift_output_func(get_cas)
def get_cas_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetCasFilterArgs']]]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCasResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...