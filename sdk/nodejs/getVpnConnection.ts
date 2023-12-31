// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getVpnConnection(args?: GetVpnConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnConnectionResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getVpnConnection:getVpnConnection", {
        "filters": args.filters,
        "staticRoutesOnly": args.staticRoutesOnly,
        "vpnConnectionId": args.vpnConnectionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpnConnection.
 */
export interface GetVpnConnectionArgs {
    filters?: inputs.GetVpnConnectionFilter[];
    staticRoutesOnly?: boolean;
    vpnConnectionId?: string;
}

/**
 * A collection of values returned by getVpnConnection.
 */
export interface GetVpnConnectionResult {
    readonly clientGatewayConfiguration: string;
    readonly clientGatewayId: string;
    readonly connectionType: string;
    readonly filters?: outputs.GetVpnConnectionFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly requestId: string;
    readonly routes: outputs.GetVpnConnectionRoute[];
    readonly state: string;
    readonly staticRoutesOnly?: boolean;
    readonly tags: outputs.GetVpnConnectionTag[];
    readonly vgwTelemetries: outputs.GetVpnConnectionVgwTelemetry[];
    readonly virtualGatewayId: string;
    readonly vpnConnectionId?: string;
}
export function getVpnConnectionOutput(args?: GetVpnConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpnConnectionResult> {
    return pulumi.output(args).apply((a: any) => getVpnConnection(a, opts))
}

/**
 * A collection of arguments for invoking getVpnConnection.
 */
export interface GetVpnConnectionOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetVpnConnectionFilterArgs>[]>;
    staticRoutesOnly?: pulumi.Input<boolean>;
    vpnConnectionId?: pulumi.Input<string>;
}
