# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .access_key import *
from .api_access_policy import *
from .api_access_rule import *
from .ca import *
from .client_gateway import *
from .dhcp_option import *
from .flexible_gpu import *
from .flexible_gpu_link import *
from .get_access_key import *
from .get_access_keys import *
from .get_account import *
from .get_accounts import *
from .get_api_access_policy import *
from .get_api_access_rule import *
from .get_api_access_rules import *
from .get_ca import *
from .get_cas import *
from .get_client_gateway import *
from .get_client_gateways import *
from .get_dhcp_option import *
from .get_dhcp_options import *
from .get_flexible_gpu import *
from .get_flexible_gpu_catalog import *
from .get_flexible_gpus import *
from .get_image import *
from .get_image_export_task import *
from .get_image_export_tasks import *
from .get_images import *
from .get_internet_service import *
from .get_internet_services import *
from .get_keypair import *
from .get_keypairs import *
from .get_load_balancer import *
from .get_load_balancer_tags import *
from .get_load_balancer_vm_health import *
from .get_load_balancers import *
from .get_nat_service import *
from .get_nat_services import *
from .get_net import *
from .get_net_access_point import *
from .get_net_access_point_services import *
from .get_net_access_points import *
from .get_net_attributes import *
from .get_net_peering import *
from .get_net_peerings import *
from .get_nets import *
from .get_nic import *
from .get_nics import *
from .get_product_type import *
from .get_product_types import *
from .get_public_ip import *
from .get_public_ips import *
from .get_pulic_catalog import *
from .get_quota import *
from .get_quotas import *
from .get_regions import *
from .get_route_table import *
from .get_route_tables import *
from .get_security_group import *
from .get_security_groups import *
from .get_server_certificate import *
from .get_server_certificates import *
from .get_snapshot import *
from .get_snapshot_export_task import *
from .get_snapshot_export_tasks import *
from .get_snapshots import *
from .get_subnet import *
from .get_subnets import *
from .get_subregions import *
from .get_tag import *
from .get_virtual_gateway import *
from .get_virtual_gateways import *
from .get_vm import *
from .get_vm_state import *
from .get_vm_states import *
from .get_vm_types import *
from .get_vms import *
from .get_volume import *
from .get_volumes import *
from .get_vpn_connection import *
from .get_vpn_connections import *
from .image import *
from .image_export_task import *
from .image_launch_permission import *
from .internet_service import *
from .internet_service_link import *
from .keypair import *
from .load_balancer import *
from .load_balancer_attributes import *
from .load_balancer_policy import *
from .load_balancer_vms import *
from .nat_service import *
from .net import *
from .net_access_point import *
from .net_attributes import *
from .net_peering import *
from .net_peering_acception import *
from .nic import *
from .nic_link import *
from .nic_private_ip import *
from .outbound_rule import *
from .provider import *
from .public_ip import *
from .public_ip_link import *
from .route import *
from .route_table import *
from .route_table_link import *
from .security_group import *
from .security_group_rule import *
from .server_certificate import *
from .snapshot import *
from .snapshot_attributes import *
from .snapshot_export_task import *
from .subnet import *
from .tag import *
from .virtual_gateway import *
from .virtual_gateway_link import *
from .virtual_gateway_route_propagation import *
from .vm import *
from .volume import *
from .volumes_link import *
from .vpn_connection import *
from .vpn_connection_route import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_outscale.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_outscale.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "outscale",
  "mod": "index/accessKey",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/accessKey:AccessKey": "AccessKey"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/apiAccessPolicy",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/apiAccessPolicy:ApiAccessPolicy": "ApiAccessPolicy"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/apiAccessRule",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/apiAccessRule:ApiAccessRule": "ApiAccessRule"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/ca",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/ca:Ca": "Ca"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/clientGateway",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/clientGateway:ClientGateway": "ClientGateway"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/dhcpOption",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/dhcpOption:DhcpOption": "DhcpOption"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/flexibleGpu",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/flexibleGpu:FlexibleGpu": "FlexibleGpu"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/flexibleGpuLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/flexibleGpuLink:FlexibleGpuLink": "FlexibleGpuLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/image",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/image:Image": "Image"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/imageExportTask",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/imageExportTask:ImageExportTask": "ImageExportTask"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/imageLaunchPermission",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/imageLaunchPermission:ImageLaunchPermission": "ImageLaunchPermission"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/internetService",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/internetService:InternetService": "InternetService"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/internetServiceLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/internetServiceLink:InternetServiceLink": "InternetServiceLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/keypair",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/keypair:Keypair": "Keypair"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/loadBalancer",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/loadBalancer:LoadBalancer": "LoadBalancer"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/loadBalancerAttributes",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/loadBalancerAttributes:LoadBalancerAttributes": "LoadBalancerAttributes"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/loadBalancerPolicy",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/loadBalancerPolicy:LoadBalancerPolicy": "LoadBalancerPolicy"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/loadBalancerVms",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/loadBalancerVms:LoadBalancerVms": "LoadBalancerVms"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/natService",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/natService:NatService": "NatService"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/net",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/net:Net": "Net"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/netAccessPoint",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/netAccessPoint:NetAccessPoint": "NetAccessPoint"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/netAttributes",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/netAttributes:NetAttributes": "NetAttributes"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/netPeering",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/netPeering:NetPeering": "NetPeering"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/netPeeringAcception",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/netPeeringAcception:NetPeeringAcception": "NetPeeringAcception"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/nic",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/nic:Nic": "Nic"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/nicLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/nicLink:NicLink": "NicLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/nicPrivateIp",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/nicPrivateIp:NicPrivateIp": "NicPrivateIp"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/outboundRule",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/outboundRule:OutboundRule": "OutboundRule"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/publicIp",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/publicIp:PublicIp": "PublicIp"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/publicIpLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/publicIpLink:PublicIpLink": "PublicIpLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/route",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/route:Route": "Route"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/routeTable",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/routeTable:RouteTable": "RouteTable"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/routeTableLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/routeTableLink:RouteTableLink": "RouteTableLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/securityGroup",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/securityGroup:SecurityGroup": "SecurityGroup"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/securityGroupRule",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/securityGroupRule:SecurityGroupRule": "SecurityGroupRule"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/serverCertificate",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/serverCertificate:ServerCertificate": "ServerCertificate"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/snapshot",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/snapshot:Snapshot": "Snapshot"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/snapshotAttributes",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/snapshotAttributes:SnapshotAttributes": "SnapshotAttributes"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/snapshotExportTask",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/snapshotExportTask:SnapshotExportTask": "SnapshotExportTask"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/subnet",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/subnet:Subnet": "Subnet"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/tag",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/tag:Tag": "Tag"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/virtualGateway",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/virtualGateway:VirtualGateway": "VirtualGateway"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/virtualGatewayLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/virtualGatewayLink:VirtualGatewayLink": "VirtualGatewayLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/virtualGatewayRoutePropagation",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/virtualGatewayRoutePropagation:VirtualGatewayRoutePropagation": "VirtualGatewayRoutePropagation"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/vm",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/vm:Vm": "Vm"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/volume",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/volume:Volume": "Volume"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/volumesLink",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/volumesLink:VolumesLink": "VolumesLink"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/vpnConnection",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/vpnConnection:VpnConnection": "VpnConnection"
  }
 },
 {
  "pkg": "outscale",
  "mod": "index/vpnConnectionRoute",
  "fqn": "pulumi_outscale",
  "classes": {
   "outscale:index/vpnConnectionRoute:VpnConnectionRoute": "VpnConnectionRoute"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "outscale",
  "token": "pulumi:providers:outscale",
  "fqn": "pulumi_outscale",
  "class": "Provider"
 }
]
"""
)
