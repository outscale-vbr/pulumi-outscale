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
    'GetSecurityGroupsResult',
    'AwaitableGetSecurityGroupsResult',
    'get_security_groups',
    'get_security_groups_output',
]

@pulumi.output_type
class GetSecurityGroupsResult:
    """
    A collection of values returned by getSecurityGroups.
    """
    def __init__(__self__, filters=None, id=None, request_id=None, security_group_ids=None, security_group_names=None, security_groups=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if security_group_names and not isinstance(security_group_names, list):
            raise TypeError("Expected argument 'security_group_names' to be a list")
        pulumi.set(__self__, "security_group_names", security_group_names)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetSecurityGroupsFilterResult']]:
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

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="securityGroupNames")
    def security_group_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_group_names")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence['outputs.GetSecurityGroupsSecurityGroupResult']:
        return pulumi.get(self, "security_groups")


class AwaitableGetSecurityGroupsResult(GetSecurityGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupsResult(
            filters=self.filters,
            id=self.id,
            request_id=self.request_id,
            security_group_ids=self.security_group_ids,
            security_group_names=self.security_group_names,
            security_groups=self.security_groups)


def get_security_groups(filters: Optional[Sequence[pulumi.InputType['GetSecurityGroupsFilterArgs']]] = None,
                        security_group_ids: Optional[Sequence[str]] = None,
                        security_group_names: Optional[Sequence[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['securityGroupIds'] = security_group_ids
    __args__['securityGroupNames'] = security_group_names
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getSecurityGroups:getSecurityGroups', __args__, opts=opts, typ=GetSecurityGroupsResult).value

    return AwaitableGetSecurityGroupsResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        request_id=pulumi.get(__ret__, 'request_id'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        security_group_names=pulumi.get(__ret__, 'security_group_names'),
        security_groups=pulumi.get(__ret__, 'security_groups'))


@_utilities.lift_output_func(get_security_groups)
def get_security_groups_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetSecurityGroupsFilterArgs']]]]] = None,
                               security_group_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               security_group_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
