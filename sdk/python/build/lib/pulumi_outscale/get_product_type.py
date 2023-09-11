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
    'GetProductTypeResult',
    'AwaitableGetProductTypeResult',
    'get_product_type',
    'get_product_type_output',
]

@pulumi.output_type
class GetProductTypeResult:
    """
    A collection of values returned by getProductType.
    """
    def __init__(__self__, description=None, filters=None, id=None, product_type_id=None, request_id=None, vendor=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if product_type_id and not isinstance(product_type_id, str):
            raise TypeError("Expected argument 'product_type_id' to be a str")
        pulumi.set(__self__, "product_type_id", product_type_id)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)
        if vendor and not isinstance(vendor, str):
            raise TypeError("Expected argument 'vendor' to be a str")
        pulumi.set(__self__, "vendor", vendor)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetProductTypeFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="productTypeId")
    def product_type_id(self) -> str:
        return pulumi.get(self, "product_type_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def vendor(self) -> str:
        return pulumi.get(self, "vendor")


class AwaitableGetProductTypeResult(GetProductTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductTypeResult(
            description=self.description,
            filters=self.filters,
            id=self.id,
            product_type_id=self.product_type_id,
            request_id=self.request_id,
            vendor=self.vendor)


def get_product_type(filters: Optional[Sequence[pulumi.InputType['GetProductTypeFilterArgs']]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductTypeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getProductType:getProductType', __args__, opts=opts, typ=GetProductTypeResult).value

    return AwaitableGetProductTypeResult(
        description=pulumi.get(__ret__, 'description'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        product_type_id=pulumi.get(__ret__, 'product_type_id'),
        request_id=pulumi.get(__ret__, 'request_id'),
        vendor=pulumi.get(__ret__, 'vendor'))


@_utilities.lift_output_func(get_product_type)
def get_product_type_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetProductTypeFilterArgs']]]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProductTypeResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
