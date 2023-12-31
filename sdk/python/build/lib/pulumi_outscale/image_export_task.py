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

__all__ = ['ImageExportTaskArgs', 'ImageExportTask']

@pulumi.input_type
class ImageExportTaskArgs:
    def __init__(__self__, *,
                 image_id: pulumi.Input[str],
                 osu_exports: pulumi.Input[Sequence[pulumi.Input['ImageExportTaskOsuExportArgs']]],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskTagArgs']]]] = None):
        """
        The set of arguments for constructing a ImageExportTask resource.
        """
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "osu_exports", osu_exports)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="osuExports")
    def osu_exports(self) -> pulumi.Input[Sequence[pulumi.Input['ImageExportTaskOsuExportArgs']]]:
        return pulumi.get(self, "osu_exports")

    @osu_exports.setter
    def osu_exports(self, value: pulumi.Input[Sequence[pulumi.Input['ImageExportTaskOsuExportArgs']]]):
        pulumi.set(self, "osu_exports", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ImageExportTaskState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 osu_exports: Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskOsuExportArgs']]]] = None,
                 progress: Optional[pulumi.Input[int]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskTagArgs']]]] = None,
                 task_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ImageExportTask resources.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if osu_exports is not None:
            pulumi.set(__self__, "osu_exports", osu_exports)
        if progress is not None:
            pulumi.set(__self__, "progress", progress)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if task_id is not None:
            pulumi.set(__self__, "task_id", task_id)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="osuExports")
    def osu_exports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskOsuExportArgs']]]]:
        return pulumi.get(self, "osu_exports")

    @osu_exports.setter
    def osu_exports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskOsuExportArgs']]]]):
        pulumi.set(self, "osu_exports", value)

    @property
    @pulumi.getter
    def progress(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "progress")

    @progress.setter
    def progress(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "progress", value)

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
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageExportTaskTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "task_id")

    @task_id.setter
    def task_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_id", value)


class ImageExportTask(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 osu_exports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageExportTaskOsuExportArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageExportTaskTagArgs']]]]] = None,
                 __props__=None):
        """
        Create a ImageExportTask resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageExportTaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ImageExportTask resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ImageExportTaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageExportTaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 osu_exports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageExportTaskOsuExportArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageExportTaskTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageExportTaskArgs.__new__(ImageExportTaskArgs)

            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
            if osu_exports is None and not opts.urn:
                raise TypeError("Missing required property 'osu_exports'")
            __props__.__dict__["osu_exports"] = osu_exports
            __props__.__dict__["tags"] = tags
            __props__.__dict__["comment"] = None
            __props__.__dict__["progress"] = None
            __props__.__dict__["request_id"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["task_id"] = None
        super(ImageExportTask, __self__).__init__(
            'outscale:index/imageExportTask:ImageExportTask',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            osu_exports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageExportTaskOsuExportArgs']]]]] = None,
            progress: Optional[pulumi.Input[int]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageExportTaskTagArgs']]]]] = None,
            task_id: Optional[pulumi.Input[str]] = None) -> 'ImageExportTask':
        """
        Get an existing ImageExportTask resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageExportTaskState.__new__(_ImageExportTaskState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["osu_exports"] = osu_exports
        __props__.__dict__["progress"] = progress
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["state"] = state
        __props__.__dict__["tags"] = tags
        __props__.__dict__["task_id"] = task_id
        return ImageExportTask(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="osuExports")
    def osu_exports(self) -> pulumi.Output[Sequence['outputs.ImageExportTaskOsuExport']]:
        return pulumi.get(self, "osu_exports")

    @property
    @pulumi.getter
    def progress(self) -> pulumi.Output[int]:
        return pulumi.get(self, "progress")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ImageExportTaskTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "task_id")

