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

__all__ = ['TagArgs', 'Tag']

@pulumi.input_type
class TagArgs:
    def __init__(__self__, *,
                 resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag: Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]] = None):
        """
        The set of arguments for constructing a Tag resource.
        """
        if resource_ids is not None:
            pulumi.set(__self__, "resource_ids", resource_ids)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter(name="resourceIds")
    def resource_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_ids")

    @resource_ids.setter
    def resource_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_ids", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]]:
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]]):
        pulumi.set(self, "tag", value)


@pulumi.input_type
class _TagState:
    def __init__(__self__, *,
                 request_id: Optional[pulumi.Input[str]] = None,
                 resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag: Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering Tag resources.
        """
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if resource_ids is not None:
            pulumi.set(__self__, "resource_ids", resource_ids)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="resourceIds")
    def resource_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_ids")

    @resource_ids.setter
    def resource_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_ids", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]]:
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TagTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Tag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTagArgs']]]]] = None,
                 __props__=None):
        """
        Create a Tag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TagArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Tag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagArgs.__new__(TagArgs)

            __props__.__dict__["resource_ids"] = resource_ids
            __props__.__dict__["tag"] = tag
            __props__.__dict__["request_id"] = None
            __props__.__dict__["tags"] = None
        super(Tag, __self__).__init__(
            'outscale:index/tag:Tag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tag: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTagArgs']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TagTagArgs']]]]] = None) -> 'Tag':
        """
        Get an existing Tag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagState.__new__(_TagState)

        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["resource_ids"] = resource_ids
        __props__.__dict__["tag"] = tag
        __props__.__dict__["tags"] = tags
        return Tag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="resourceIds")
    def resource_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "resource_ids")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[Optional[Sequence['outputs.TagTag']]]:
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.TagTag']]:
        return pulumi.get(self, "tags")

