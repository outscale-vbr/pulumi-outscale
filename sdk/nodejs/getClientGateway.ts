// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getClientGateway(args?: GetClientGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetClientGatewayResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getClientGateway:getClientGateway", {
        "clientGatewayId": args.clientGatewayId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientGateway.
 */
export interface GetClientGatewayArgs {
    clientGatewayId?: string;
    filters?: inputs.GetClientGatewayFilter[];
}

/**
 * A collection of values returned by getClientGateway.
 */
export interface GetClientGatewayResult {
    readonly bgpAsn: number;
    readonly clientGatewayId?: string;
    readonly connectionType: string;
    readonly filters?: outputs.GetClientGatewayFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly publicIp: string;
    readonly requestId: string;
    readonly state: string;
    readonly tags: outputs.GetClientGatewayTag[];
}
export function getClientGatewayOutput(args?: GetClientGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientGatewayResult> {
    return pulumi.output(args).apply((a: any) => getClientGateway(a, opts))
}

/**
 * A collection of arguments for invoking getClientGateway.
 */
export interface GetClientGatewayOutputArgs {
    clientGatewayId?: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetClientGatewayFilterArgs>[]>;
}
