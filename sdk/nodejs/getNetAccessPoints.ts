// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getNetAccessPoints(args?: GetNetAccessPointsArgs, opts?: pulumi.InvokeOptions): Promise<GetNetAccessPointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getNetAccessPoints:getNetAccessPoints", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetAccessPoints.
 */
export interface GetNetAccessPointsArgs {
    filters?: inputs.GetNetAccessPointsFilter[];
}

/**
 * A collection of values returned by getNetAccessPoints.
 */
export interface GetNetAccessPointsResult {
    readonly filters?: outputs.GetNetAccessPointsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly netAccessPoints: outputs.GetNetAccessPointsNetAccessPoint[];
    readonly requestId: string;
}
export function getNetAccessPointsOutput(args?: GetNetAccessPointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetAccessPointsResult> {
    return pulumi.output(args).apply((a: any) => getNetAccessPoints(a, opts))
}

/**
 * A collection of arguments for invoking getNetAccessPoints.
 */
export interface GetNetAccessPointsOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetNetAccessPointsFilterArgs>[]>;
}
