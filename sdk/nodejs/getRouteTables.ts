// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRouteTables(args?: GetRouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTablesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getRouteTables:getRouteTables", {
        "filters": args.filters,
        "routeTableIds": args.routeTableIds,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesArgs {
    filters?: inputs.GetRouteTablesFilter[];
    routeTableIds?: string[];
}

/**
 * A collection of values returned by getRouteTables.
 */
export interface GetRouteTablesResult {
    readonly filters?: outputs.GetRouteTablesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly requestId: string;
    readonly routeTableIds?: string[];
    readonly routeTables: outputs.GetRouteTablesRouteTable[];
}
export function getRouteTablesOutput(args?: GetRouteTablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteTablesResult> {
    return pulumi.output(args).apply((a: any) => getRouteTables(a, opts))
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetRouteTablesFilterArgs>[]>;
    routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
}
