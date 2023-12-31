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
    'GetLoadBalancerResult',
    'AwaitableGetLoadBalancerResult',
    'get_load_balancer',
    'get_load_balancer_output',
]

@pulumi.output_type
class GetLoadBalancerResult:
    """
    A collection of values returned by getLoadBalancer.
    """
    def __init__(__self__, access_log=None, application_sticky_cookie_policies=None, backend_vm_ids=None, dns_name=None, filters=None, health_check=None, id=None, listeners=None, load_balancer_name=None, load_balancer_sticky_cookie_policies=None, load_balancer_type=None, net_id=None, public_ip=None, request_id=None, secured_cookies=None, security_groups=None, source_security_group=None, subnets=None, subregion_names=None, tags=None):
        if access_log and not isinstance(access_log, dict):
            raise TypeError("Expected argument 'access_log' to be a dict")
        pulumi.set(__self__, "access_log", access_log)
        if application_sticky_cookie_policies and not isinstance(application_sticky_cookie_policies, list):
            raise TypeError("Expected argument 'application_sticky_cookie_policies' to be a list")
        pulumi.set(__self__, "application_sticky_cookie_policies", application_sticky_cookie_policies)
        if backend_vm_ids and not isinstance(backend_vm_ids, list):
            raise TypeError("Expected argument 'backend_vm_ids' to be a list")
        pulumi.set(__self__, "backend_vm_ids", backend_vm_ids)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if health_check and not isinstance(health_check, dict):
            raise TypeError("Expected argument 'health_check' to be a dict")
        pulumi.set(__self__, "health_check", health_check)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listeners and not isinstance(listeners, list):
            raise TypeError("Expected argument 'listeners' to be a list")
        pulumi.set(__self__, "listeners", listeners)
        if load_balancer_name and not isinstance(load_balancer_name, str):
            raise TypeError("Expected argument 'load_balancer_name' to be a str")
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        if load_balancer_sticky_cookie_policies and not isinstance(load_balancer_sticky_cookie_policies, list):
            raise TypeError("Expected argument 'load_balancer_sticky_cookie_policies' to be a list")
        pulumi.set(__self__, "load_balancer_sticky_cookie_policies", load_balancer_sticky_cookie_policies)
        if load_balancer_type and not isinstance(load_balancer_type, str):
            raise TypeError("Expected argument 'load_balancer_type' to be a str")
        pulumi.set(__self__, "load_balancer_type", load_balancer_type)
        if net_id and not isinstance(net_id, str):
            raise TypeError("Expected argument 'net_id' to be a str")
        pulumi.set(__self__, "net_id", net_id)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)
        if secured_cookies and not isinstance(secured_cookies, bool):
            raise TypeError("Expected argument 'secured_cookies' to be a bool")
        pulumi.set(__self__, "secured_cookies", secured_cookies)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if source_security_group and not isinstance(source_security_group, dict):
            raise TypeError("Expected argument 'source_security_group' to be a dict")
        pulumi.set(__self__, "source_security_group", source_security_group)
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if subregion_names and not isinstance(subregion_names, list):
            raise TypeError("Expected argument 'subregion_names' to be a list")
        pulumi.set(__self__, "subregion_names", subregion_names)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessLog")
    def access_log(self) -> 'outputs.GetLoadBalancerAccessLogResult':
        return pulumi.get(self, "access_log")

    @property
    @pulumi.getter(name="applicationStickyCookiePolicies")
    def application_sticky_cookie_policies(self) -> Sequence['outputs.GetLoadBalancerApplicationStickyCookiePolicyResult']:
        return pulumi.get(self, "application_sticky_cookie_policies")

    @property
    @pulumi.getter(name="backendVmIds")
    def backend_vm_ids(self) -> Sequence[str]:
        return pulumi.get(self, "backend_vm_ids")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetLoadBalancerFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="healthCheck")
    def health_check(self) -> 'outputs.GetLoadBalancerHealthCheckResult':
        return pulumi.get(self, "health_check")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def listeners(self) -> Sequence['outputs.GetLoadBalancerListenerResult']:
        return pulumi.get(self, "listeners")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> str:
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="loadBalancerStickyCookiePolicies")
    def load_balancer_sticky_cookie_policies(self) -> Sequence['outputs.GetLoadBalancerLoadBalancerStickyCookiePolicyResult']:
        return pulumi.get(self, "load_balancer_sticky_cookie_policies")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> str:
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter(name="netId")
    def net_id(self) -> str:
        return pulumi.get(self, "net_id")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="securedCookies")
    def secured_cookies(self) -> bool:
        return pulumi.get(self, "secured_cookies")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="sourceSecurityGroup")
    def source_security_group(self) -> 'outputs.GetLoadBalancerSourceSecurityGroupResult':
        return pulumi.get(self, "source_security_group")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence[str]:
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter(name="subregionNames")
    def subregion_names(self) -> Sequence[str]:
        return pulumi.get(self, "subregion_names")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetLoadBalancerTagResult']:
        return pulumi.get(self, "tags")


class AwaitableGetLoadBalancerResult(GetLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerResult(
            access_log=self.access_log,
            application_sticky_cookie_policies=self.application_sticky_cookie_policies,
            backend_vm_ids=self.backend_vm_ids,
            dns_name=self.dns_name,
            filters=self.filters,
            health_check=self.health_check,
            id=self.id,
            listeners=self.listeners,
            load_balancer_name=self.load_balancer_name,
            load_balancer_sticky_cookie_policies=self.load_balancer_sticky_cookie_policies,
            load_balancer_type=self.load_balancer_type,
            net_id=self.net_id,
            public_ip=self.public_ip,
            request_id=self.request_id,
            secured_cookies=self.secured_cookies,
            security_groups=self.security_groups,
            source_security_group=self.source_security_group,
            subnets=self.subnets,
            subregion_names=self.subregion_names,
            tags=self.tags)


def get_load_balancer(access_log: Optional[pulumi.InputType['GetLoadBalancerAccessLogArgs']] = None,
                      backend_vm_ids: Optional[Sequence[str]] = None,
                      dns_name: Optional[str] = None,
                      filters: Optional[Sequence[pulumi.InputType['GetLoadBalancerFilterArgs']]] = None,
                      health_check: Optional[pulumi.InputType['GetLoadBalancerHealthCheckArgs']] = None,
                      listeners: Optional[Sequence[pulumi.InputType['GetLoadBalancerListenerArgs']]] = None,
                      load_balancer_name: Optional[str] = None,
                      load_balancer_type: Optional[str] = None,
                      net_id: Optional[str] = None,
                      security_groups: Optional[Sequence[str]] = None,
                      subnets: Optional[Sequence[str]] = None,
                      tags: Optional[Sequence[pulumi.InputType['GetLoadBalancerTagArgs']]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['accessLog'] = access_log
    __args__['backendVmIds'] = backend_vm_ids
    __args__['dnsName'] = dns_name
    __args__['filters'] = filters
    __args__['healthCheck'] = health_check
    __args__['listeners'] = listeners
    __args__['loadBalancerName'] = load_balancer_name
    __args__['loadBalancerType'] = load_balancer_type
    __args__['netId'] = net_id
    __args__['securityGroups'] = security_groups
    __args__['subnets'] = subnets
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('outscale:index/getLoadBalancer:getLoadBalancer', __args__, opts=opts, typ=GetLoadBalancerResult).value

    return AwaitableGetLoadBalancerResult(
        access_log=pulumi.get(__ret__, 'access_log'),
        application_sticky_cookie_policies=pulumi.get(__ret__, 'application_sticky_cookie_policies'),
        backend_vm_ids=pulumi.get(__ret__, 'backend_vm_ids'),
        dns_name=pulumi.get(__ret__, 'dns_name'),
        filters=pulumi.get(__ret__, 'filters'),
        health_check=pulumi.get(__ret__, 'health_check'),
        id=pulumi.get(__ret__, 'id'),
        listeners=pulumi.get(__ret__, 'listeners'),
        load_balancer_name=pulumi.get(__ret__, 'load_balancer_name'),
        load_balancer_sticky_cookie_policies=pulumi.get(__ret__, 'load_balancer_sticky_cookie_policies'),
        load_balancer_type=pulumi.get(__ret__, 'load_balancer_type'),
        net_id=pulumi.get(__ret__, 'net_id'),
        public_ip=pulumi.get(__ret__, 'public_ip'),
        request_id=pulumi.get(__ret__, 'request_id'),
        secured_cookies=pulumi.get(__ret__, 'secured_cookies'),
        security_groups=pulumi.get(__ret__, 'security_groups'),
        source_security_group=pulumi.get(__ret__, 'source_security_group'),
        subnets=pulumi.get(__ret__, 'subnets'),
        subregion_names=pulumi.get(__ret__, 'subregion_names'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_load_balancer)
def get_load_balancer_output(access_log: Optional[pulumi.Input[Optional[pulumi.InputType['GetLoadBalancerAccessLogArgs']]]] = None,
                             backend_vm_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             dns_name: Optional[pulumi.Input[Optional[str]]] = None,
                             filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetLoadBalancerFilterArgs']]]]] = None,
                             health_check: Optional[pulumi.Input[Optional[pulumi.InputType['GetLoadBalancerHealthCheckArgs']]]] = None,
                             listeners: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetLoadBalancerListenerArgs']]]]] = None,
                             load_balancer_name: Optional[pulumi.Input[Optional[str]]] = None,
                             load_balancer_type: Optional[pulumi.Input[Optional[str]]] = None,
                             net_id: Optional[pulumi.Input[Optional[str]]] = None,
                             security_groups: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             subnets: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetLoadBalancerTagArgs']]]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLoadBalancerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
