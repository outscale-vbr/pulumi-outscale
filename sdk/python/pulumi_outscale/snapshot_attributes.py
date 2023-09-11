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

__all__ = ['SnapshotAttributesArgs', 'SnapshotAttributes']

@pulumi.input_type
class SnapshotAttributesArgs:
    def __init__(__self__, *,
                 snapshot_id: pulumi.Input[str],
                 permissions_to_create_volume_additions: Optional[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']] = None,
                 permissions_to_create_volume_removals: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]] = None):
        """
        The set of arguments for constructing a SnapshotAttributes resource.
        """
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if permissions_to_create_volume_additions is not None:
            pulumi.set(__self__, "permissions_to_create_volume_additions", permissions_to_create_volume_additions)
        if permissions_to_create_volume_removals is not None:
            pulumi.set(__self__, "permissions_to_create_volume_removals", permissions_to_create_volume_removals)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="permissionsToCreateVolumeAdditions")
    def permissions_to_create_volume_additions(self) -> Optional[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]:
        return pulumi.get(self, "permissions_to_create_volume_additions")

    @permissions_to_create_volume_additions.setter
    def permissions_to_create_volume_additions(self, value: Optional[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]):
        pulumi.set(self, "permissions_to_create_volume_additions", value)

    @property
    @pulumi.getter(name="permissionsToCreateVolumeRemovals")
    def permissions_to_create_volume_removals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]:
        return pulumi.get(self, "permissions_to_create_volume_removals")

    @permissions_to_create_volume_removals.setter
    def permissions_to_create_volume_removals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]):
        pulumi.set(self, "permissions_to_create_volume_removals", value)


@pulumi.input_type
class _SnapshotAttributesState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 permissions_to_create_volume_additions: Optional[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']] = None,
                 permissions_to_create_volume_removals: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnapshotAttributes resources.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if permissions_to_create_volume_additions is not None:
            pulumi.set(__self__, "permissions_to_create_volume_additions", permissions_to_create_volume_additions)
        if permissions_to_create_volume_removals is not None:
            pulumi.set(__self__, "permissions_to_create_volume_removals", permissions_to_create_volume_removals)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="permissionsToCreateVolumeAdditions")
    def permissions_to_create_volume_additions(self) -> Optional[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]:
        return pulumi.get(self, "permissions_to_create_volume_additions")

    @permissions_to_create_volume_additions.setter
    def permissions_to_create_volume_additions(self, value: Optional[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]):
        pulumi.set(self, "permissions_to_create_volume_additions", value)

    @property
    @pulumi.getter(name="permissionsToCreateVolumeRemovals")
    def permissions_to_create_volume_removals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]:
        return pulumi.get(self, "permissions_to_create_volume_removals")

    @permissions_to_create_volume_removals.setter
    def permissions_to_create_volume_removals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]):
        pulumi.set(self, "permissions_to_create_volume_removals", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)


class SnapshotAttributes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions_to_create_volume_additions: Optional[pulumi.Input[pulumi.InputType['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]] = None,
                 permissions_to_create_volume_removals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SnapshotAttributes resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotAttributesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SnapshotAttributes resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SnapshotAttributesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotAttributesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permissions_to_create_volume_additions: Optional[pulumi.Input[pulumi.InputType['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]] = None,
                 permissions_to_create_volume_removals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotAttributesArgs.__new__(SnapshotAttributesArgs)

            __props__.__dict__["permissions_to_create_volume_additions"] = permissions_to_create_volume_additions
            __props__.__dict__["permissions_to_create_volume_removals"] = permissions_to_create_volume_removals
            if snapshot_id is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_id'")
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["account_id"] = None
            __props__.__dict__["request_id"] = None
        super(SnapshotAttributes, __self__).__init__(
            'outscale:index/snapshotAttributes:SnapshotAttributes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            permissions_to_create_volume_additions: Optional[pulumi.Input[pulumi.InputType['SnapshotAttributesPermissionsToCreateVolumeAdditionsArgs']]] = None,
            permissions_to_create_volume_removals: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotAttributesPermissionsToCreateVolumeRemovalArgs']]]]] = None,
            request_id: Optional[pulumi.Input[str]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None) -> 'SnapshotAttributes':
        """
        Get an existing SnapshotAttributes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotAttributesState.__new__(_SnapshotAttributesState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["permissions_to_create_volume_additions"] = permissions_to_create_volume_additions
        __props__.__dict__["permissions_to_create_volume_removals"] = permissions_to_create_volume_removals
        __props__.__dict__["request_id"] = request_id
        __props__.__dict__["snapshot_id"] = snapshot_id
        return SnapshotAttributes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="permissionsToCreateVolumeAdditions")
    def permissions_to_create_volume_additions(self) -> pulumi.Output[Optional['outputs.SnapshotAttributesPermissionsToCreateVolumeAdditions']]:
        return pulumi.get(self, "permissions_to_create_volume_additions")

    @property
    @pulumi.getter(name="permissionsToCreateVolumeRemovals")
    def permissions_to_create_volume_removals(self) -> pulumi.Output[Optional[Sequence['outputs.SnapshotAttributesPermissionsToCreateVolumeRemoval']]]:
        return pulumi.get(self, "permissions_to_create_volume_removals")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "snapshot_id")
