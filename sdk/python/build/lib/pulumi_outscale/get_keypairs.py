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
    'GetKeypairsResult',
    'AwaitableGetKeypairsResult',
    'get_keypairs',
    'get_keypairs_output',
]

@pulumi.output_type
class GetKeypairsResult:
    """
    A collection of values returned by getKeypairs.
    """
    def __init__(__self__, filters=None, id=None, keypair_names=None, keypairs=None, request_id=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keypair_names and not isinstance(keypair_names, list):
            raise TypeError("Expected argument 'keypair_names' to be a list")
        pulumi.set(__self__, "keypair_names", keypair_names)
        if keypairs and not isinstance(keypairs, list):
            raise TypeError("Expected argument 'keypairs' to be a list")
        pulumi.set(__self__, "keypairs", keypairs)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetKeypairsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keypairNames")
    def keypair_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "keypair_names")

    @property
    @pulumi.getter
    def keypairs(self) -> Sequence['outputs.GetKeypairsKeypairResult']:
        return pulumi.get(self, "keypairs")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        return pulumi.get(self, "request_id")


class AwaitableGetKeypairsResult(GetKeypairsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeypairsResult(
            filters=self.filters,
            id=self.id,
            keypair_names=self.keypair_names,
            keypairs=self.keypairs,
            request_id=self.request_id)


def get_keypairs(filters: Optional[Sequence[pulumi.InputType['GetKeypairsFilterArgs']]] = None,
                 keypair_names: Optional[Sequence[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeypairsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['keypairNames'] = keypair_names
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getKeypairs:getKeypairs', __args__, opts=opts, typ=GetKeypairsResult).value

    return AwaitableGetKeypairsResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        keypair_names=pulumi.get(__ret__, 'keypair_names'),
        keypairs=pulumi.get(__ret__, 'keypairs'),
        request_id=pulumi.get(__ret__, 'request_id'))


@_utilities.lift_output_func(get_keypairs)
def get_keypairs_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetKeypairsFilterArgs']]]]] = None,
                        keypair_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeypairsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...