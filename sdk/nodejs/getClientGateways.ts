// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getClientGateways(args?: GetClientGatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GetClientGatewaysResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getClientGateways:getClientGateways", {
        "clientGatewayIds": args.clientGatewayIds,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientGateways.
 */
export interface GetClientGatewaysArgs {
    clientGatewayIds?: string[];
    filters?: inputs.GetClientGatewaysFilter[];
}

/**
 * A collection of values returned by getClientGateways.
 */
export interface GetClientGatewaysResult {
    readonly clientGatewayIds?: string[];
    readonly clientGateways: outputs.GetClientGatewaysClientGateway[];
    readonly filters?: outputs.GetClientGatewaysFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly requestId: string;
}
export function getClientGatewaysOutput(args?: GetClientGatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientGatewaysResult> {
    return pulumi.output(args).apply((a: any) => getClientGateways(a, opts))
}

/**
 * A collection of arguments for invoking getClientGateways.
 */
export interface GetClientGatewaysOutputArgs {
    clientGatewayIds?: pulumi.Input<pulumi.Input<string>[]>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetClientGatewaysFilterArgs>[]>;
}
