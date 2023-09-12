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
    'GetVmResult',
    'AwaitableGetVmResult',
    'get_vm',
    'get_vm_output',
]

@pulumi.output_type
class GetVmResult:
    """
    A collection of values returned by getVm.
    """
    def __init__(__self__, architecture=None, block_device_mappings_createds=None, bsu_optimized=None, client_token=None, creation_date=None, deletion_protection=None, filters=None, hypervisor=None, id=None, image_id=None, is_source_dest_checked=None, keypair_name=None, launch_number=None, nested_virtualization=None, net_id=None, nics=None, os_family=None, performance=None, placement_subregion_name=None, placement_tenancy=None, private_dns_name=None, private_ip=None, private_ips=None, product_codes=None, public_dns_name=None, public_ip=None, request_id=None, reservation_id=None, root_device_name=None, root_device_type=None, security_group_ids=None, security_group_names=None, security_groups=None, state=None, state_reason=None, subnet_id=None, tags=None, user_data=None, vm_id=None, vm_initiated_shutdown_behavior=None, vm_type=None):
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if block_device_mappings_createds and not isinstance(block_device_mappings_createds, list):
            raise TypeError("Expected argument 'block_device_mappings_createds' to be a list")
        pulumi.set(__self__, "block_device_mappings_createds", block_device_mappings_createds)
        if bsu_optimized and not isinstance(bsu_optimized, bool):
            raise TypeError("Expected argument 'bsu_optimized' to be a bool")
        pulumi.set(__self__, "bsu_optimized", bsu_optimized)
        if client_token and not isinstance(client_token, str):
            raise TypeError("Expected argument 'client_token' to be a str")
        pulumi.set(__self__, "client_token", client_token)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if hypervisor and not isinstance(hypervisor, str):
            raise TypeError("Expected argument 'hypervisor' to be a str")
        pulumi.set(__self__, "hypervisor", hypervisor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if is_source_dest_checked and not isinstance(is_source_dest_checked, bool):
            raise TypeError("Expected argument 'is_source_dest_checked' to be a bool")
        pulumi.set(__self__, "is_source_dest_checked", is_source_dest_checked)
        if keypair_name and not isinstance(keypair_name, str):
            raise TypeError("Expected argument 'keypair_name' to be a str")
        pulumi.set(__self__, "keypair_name", keypair_name)
        if launch_number and not isinstance(launch_number, int):
            raise TypeError("Expected argument 'launch_number' to be a int")
        pulumi.set(__self__, "launch_number", launch_number)
        if nested_virtualization and not isinstance(nested_virtualization, bool):
            raise TypeError("Expected argument 'nested_virtualization' to be a bool")
        pulumi.set(__self__, "nested_virtualization", nested_virtualization)
        if net_id and not isinstance(net_id, str):
            raise TypeError("Expected argument 'net_id' to be a str")
        pulumi.set(__self__, "net_id", net_id)
        if nics and not isinstance(nics, list):
            raise TypeError("Expected argument 'nics' to be a list")
        pulumi.set(__self__, "nics", nics)
        if os_family and not isinstance(os_family, str):
            raise TypeError("Expected argument 'os_family' to be a str")
        pulumi.set(__self__, "os_family", os_family)
        if performance and not isinstance(performance, str):
            raise TypeError("Expected argument 'performance' to be a str")
        pulumi.set(__self__, "performance", performance)
        if placement_subregion_name and not isinstance(placement_subregion_name, str):
            raise TypeError("Expected argument 'placement_subregion_name' to be a str")
        pulumi.set(__self__, "placement_subregion_name", placement_subregion_name)
        if placement_tenancy and not isinstance(placement_tenancy, str):
            raise TypeError("Expected argument 'placement_tenancy' to be a str")
        pulumi.set(__self__, "placement_tenancy", placement_tenancy)
        if private_dns_name and not isinstance(private_dns_name, str):
            raise TypeError("Expected argument 'private_dns_name' to be a str")
        pulumi.set(__self__, "private_dns_name", private_dns_name)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        pulumi.set(__self__, "private_ips", private_ips)
        if product_codes and not isinstance(product_codes, list):
            raise TypeError("Expected argument 'product_codes' to be a list")
        pulumi.set(__self__, "product_codes", product_codes)
        if public_dns_name and not isinstance(public_dns_name, str):
            raise TypeError("Expected argument 'public_dns_name' to be a str")
        pulumi.set(__self__, "public_dns_name", public_dns_name)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)
        if reservation_id and not isinstance(reservation_id, str):
            raise TypeError("Expected argument 'reservation_id' to be a str")
        pulumi.set(__self__, "reservation_id", reservation_id)
        if root_device_name and not isinstance(root_device_name, str):
            raise TypeError("Expected argument 'root_device_name' to be a str")
        pulumi.set(__self__, "root_device_name", root_device_name)
        if root_device_type and not isinstance(root_device_type, str):
            raise TypeError("Expected argument 'root_device_type' to be a str")
        pulumi.set(__self__, "root_device_type", root_device_type)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if security_group_names and not isinstance(security_group_names, list):
            raise TypeError("Expected argument 'security_group_names' to be a list")
        pulumi.set(__self__, "security_group_names", security_group_names)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_reason and not isinstance(state_reason, str):
            raise TypeError("Expected argument 'state_reason' to be a str")
        pulumi.set(__self__, "state_reason", state_reason)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)
        if vm_id and not isinstance(vm_id, str):
            raise TypeError("Expected argument 'vm_id' to be a str")
        pulumi.set(__self__, "vm_id", vm_id)
        if vm_initiated_shutdown_behavior and not isinstance(vm_initiated_shutdown_behavior, str):
            raise TypeError("Expected argument 'vm_initiated_shutdown_behavior' to be a str")
        pulumi.set(__self__, "vm_initiated_shutdown_behavior", vm_initiated_shutdown_behavior)
        if vm_type and not isinstance(vm_type, str):
            raise TypeError("Expected argument 'vm_type' to be a str")
        pulumi.set(__self__, "vm_type", vm_type)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="blockDeviceMappingsCreateds")
    def block_device_mappings_createds(self) -> Sequence['outputs.GetVmBlockDeviceMappingsCreatedResult']:
        return pulumi.get(self, "block_device_mappings_createds")

    @property
    @pulumi.getter(name="bsuOptimized")
    def bsu_optimized(self) -> bool:
        return pulumi.get(self, "bsu_optimized")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> str:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> bool:
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetVmFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def hypervisor(self) -> str:
        return pulumi.get(self, "hypervisor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="isSourceDestChecked")
    def is_source_dest_checked(self) -> bool:
        return pulumi.get(self, "is_source_dest_checked")

    @property
    @pulumi.getter(name="keypairName")
    def keypair_name(self) -> str:
        return pulumi.get(self, "keypair_name")

    @property
    @pulumi.getter(name="launchNumber")
    def launch_number(self) -> int:
        return pulumi.get(self, "launch_number")

    @property
    @pulumi.getter(name="nestedVirtualization")
    def nested_virtualization(self) -> bool:
        return pulumi.get(self, "nested_virtualization")

    @property
    @pulumi.getter(name="netId")
    def net_id(self) -> str:
        return pulumi.get(self, "net_id")

    @property
    @pulumi.getter
    def nics(self) -> Sequence['outputs.GetVmNicResult']:
        return pulumi.get(self, "nics")

    @property
    @pulumi.getter(name="osFamily")
    def os_family(self) -> str:
        return pulumi.get(self, "os_family")

    @property
    @pulumi.getter
    def performance(self) -> str:
        return pulumi.get(self, "performance")

    @property
    @pulumi.getter(name="placementSubregionName")
    def placement_subregion_name(self) -> str:
        return pulumi.get(self, "placement_subregion_name")

    @property
    @pulumi.getter(name="placementTenancy")
    def placement_tenancy(self) -> str:
        return pulumi.get(self, "placement_tenancy")

    @property
    @pulumi.getter(name="privateDnsName")
    def private_dns_name(self) -> str:
        return pulumi.get(self, "private_dns_name")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter(name="productCodes")
    def product_codes(self) -> Sequence[str]:
        return pulumi.get(self, "product_codes")

    @property
    @pulumi.getter(name="publicDnsName")
    def public_dns_name(self) -> str:
        return pulumi.get(self, "public_dns_name")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="reservationId")
    def reservation_id(self) -> str:
        return pulumi.get(self, "reservation_id")

    @property
    @pulumi.getter(name="rootDeviceName")
    def root_device_name(self) -> str:
        return pulumi.get(self, "root_device_name")

    @property
    @pulumi.getter(name="rootDeviceType")
    def root_device_type(self) -> str:
        return pulumi.get(self, "root_device_type")

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
    def security_groups(self) -> Sequence['outputs.GetVmSecurityGroupResult']:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateReason")
    def state_reason(self) -> str:
        return pulumi.get(self, "state_reason")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetVmTagResult']:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> str:
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> str:
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter(name="vmInitiatedShutdownBehavior")
    def vm_initiated_shutdown_behavior(self) -> str:
        return pulumi.get(self, "vm_initiated_shutdown_behavior")

    @property
    @pulumi.getter(name="vmType")
    def vm_type(self) -> str:
        return pulumi.get(self, "vm_type")


class AwaitableGetVmResult(GetVmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVmResult(
            architecture=self.architecture,
            block_device_mappings_createds=self.block_device_mappings_createds,
            bsu_optimized=self.bsu_optimized,
            client_token=self.client_token,
            creation_date=self.creation_date,
            deletion_protection=self.deletion_protection,
            filters=self.filters,
            hypervisor=self.hypervisor,
            id=self.id,
            image_id=self.image_id,
            is_source_dest_checked=self.is_source_dest_checked,
            keypair_name=self.keypair_name,
            launch_number=self.launch_number,
            nested_virtualization=self.nested_virtualization,
            net_id=self.net_id,
            nics=self.nics,
            os_family=self.os_family,
            performance=self.performance,
            placement_subregion_name=self.placement_subregion_name,
            placement_tenancy=self.placement_tenancy,
            private_dns_name=self.private_dns_name,
            private_ip=self.private_ip,
            private_ips=self.private_ips,
            product_codes=self.product_codes,
            public_dns_name=self.public_dns_name,
            public_ip=self.public_ip,
            request_id=self.request_id,
            reservation_id=self.reservation_id,
            root_device_name=self.root_device_name,
            root_device_type=self.root_device_type,
            security_group_ids=self.security_group_ids,
            security_group_names=self.security_group_names,
            security_groups=self.security_groups,
            state=self.state,
            state_reason=self.state_reason,
            subnet_id=self.subnet_id,
            tags=self.tags,
            user_data=self.user_data,
            vm_id=self.vm_id,
            vm_initiated_shutdown_behavior=self.vm_initiated_shutdown_behavior,
            vm_type=self.vm_type)


def get_vm(block_device_mappings_createds: Optional[Sequence[pulumi.InputType['GetVmBlockDeviceMappingsCreatedArgs']]] = None,
           bsu_optimized: Optional[bool] = None,
           deletion_protection: Optional[bool] = None,
           filters: Optional[Sequence[pulumi.InputType['GetVmFilterArgs']]] = None,
           image_id: Optional[str] = None,
           is_source_dest_checked: Optional[bool] = None,
           keypair_name: Optional[str] = None,
           nics: Optional[Sequence[pulumi.InputType['GetVmNicArgs']]] = None,
           placement_subregion_name: Optional[str] = None,
           placement_tenancy: Optional[str] = None,
           private_ips: Optional[Sequence[str]] = None,
           security_group_ids: Optional[Sequence[str]] = None,
           security_group_names: Optional[Sequence[str]] = None,
           subnet_id: Optional[str] = None,
           user_data: Optional[str] = None,
           vm_id: Optional[str] = None,
           vm_initiated_shutdown_behavior: Optional[str] = None,
           vm_type: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVmResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['blockDeviceMappingsCreateds'] = block_device_mappings_createds
    __args__['bsuOptimized'] = bsu_optimized
    __args__['deletionProtection'] = deletion_protection
    __args__['filters'] = filters
    __args__['imageId'] = image_id
    __args__['isSourceDestChecked'] = is_source_dest_checked
    __args__['keypairName'] = keypair_name
    __args__['nics'] = nics
    __args__['placementSubregionName'] = placement_subregion_name
    __args__['placementTenancy'] = placement_tenancy
    __args__['privateIps'] = private_ips
    __args__['securityGroupIds'] = security_group_ids
    __args__['securityGroupNames'] = security_group_names
    __args__['subnetId'] = subnet_id
    __args__['userData'] = user_data
    __args__['vmId'] = vm_id
    __args__['vmInitiatedShutdownBehavior'] = vm_initiated_shutdown_behavior
    __args__['vmType'] = vm_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getVm:getVm', __args__, opts=opts, typ=GetVmResult).value

    return AwaitableGetVmResult(
        architecture=pulumi.get(__ret__, 'architecture'),
        block_device_mappings_createds=pulumi.get(__ret__, 'block_device_mappings_createds'),
        bsu_optimized=pulumi.get(__ret__, 'bsu_optimized'),
        client_token=pulumi.get(__ret__, 'client_token'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        filters=pulumi.get(__ret__, 'filters'),
        hypervisor=pulumi.get(__ret__, 'hypervisor'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        is_source_dest_checked=pulumi.get(__ret__, 'is_source_dest_checked'),
        keypair_name=pulumi.get(__ret__, 'keypair_name'),
        launch_number=pulumi.get(__ret__, 'launch_number'),
        nested_virtualization=pulumi.get(__ret__, 'nested_virtualization'),
        net_id=pulumi.get(__ret__, 'net_id'),
        nics=pulumi.get(__ret__, 'nics'),
        os_family=pulumi.get(__ret__, 'os_family'),
        performance=pulumi.get(__ret__, 'performance'),
        placement_subregion_name=pulumi.get(__ret__, 'placement_subregion_name'),
        placement_tenancy=pulumi.get(__ret__, 'placement_tenancy'),
        private_dns_name=pulumi.get(__ret__, 'private_dns_name'),
        private_ip=pulumi.get(__ret__, 'private_ip'),
        private_ips=pulumi.get(__ret__, 'private_ips'),
        product_codes=pulumi.get(__ret__, 'product_codes'),
        public_dns_name=pulumi.get(__ret__, 'public_dns_name'),
        public_ip=pulumi.get(__ret__, 'public_ip'),
        request_id=pulumi.get(__ret__, 'request_id'),
        reservation_id=pulumi.get(__ret__, 'reservation_id'),
        root_device_name=pulumi.get(__ret__, 'root_device_name'),
        root_device_type=pulumi.get(__ret__, 'root_device_type'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        security_group_names=pulumi.get(__ret__, 'security_group_names'),
        security_groups=pulumi.get(__ret__, 'security_groups'),
        state=pulumi.get(__ret__, 'state'),
        state_reason=pulumi.get(__ret__, 'state_reason'),
        subnet_id=pulumi.get(__ret__, 'subnet_id'),
        tags=pulumi.get(__ret__, 'tags'),
        user_data=pulumi.get(__ret__, 'user_data'),
        vm_id=pulumi.get(__ret__, 'vm_id'),
        vm_initiated_shutdown_behavior=pulumi.get(__ret__, 'vm_initiated_shutdown_behavior'),
        vm_type=pulumi.get(__ret__, 'vm_type'))


@_utilities.lift_output_func(get_vm)
def get_vm_output(block_device_mappings_createds: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetVmBlockDeviceMappingsCreatedArgs']]]]] = None,
                  bsu_optimized: Optional[pulumi.Input[Optional[bool]]] = None,
                  deletion_protection: Optional[pulumi.Input[Optional[bool]]] = None,
                  filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetVmFilterArgs']]]]] = None,
                  image_id: Optional[pulumi.Input[Optional[str]]] = None,
                  is_source_dest_checked: Optional[pulumi.Input[Optional[bool]]] = None,
                  keypair_name: Optional[pulumi.Input[Optional[str]]] = None,
                  nics: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetVmNicArgs']]]]] = None,
                  placement_subregion_name: Optional[pulumi.Input[Optional[str]]] = None,
                  placement_tenancy: Optional[pulumi.Input[Optional[str]]] = None,
                  private_ips: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                  security_group_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                  security_group_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                  subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                  user_data: Optional[pulumi.Input[Optional[str]]] = None,
                  vm_id: Optional[pulumi.Input[Optional[str]]] = None,
                  vm_initiated_shutdown_behavior: Optional[pulumi.Input[Optional[str]]] = None,
                  vm_type: Optional[pulumi.Input[Optional[str]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVmResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
