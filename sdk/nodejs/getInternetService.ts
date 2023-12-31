// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getInternetService(args?: GetInternetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetInternetServiceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getInternetService:getInternetService", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getInternetService.
 */
export interface GetInternetServiceArgs {
    filters?: inputs.GetInternetServiceFilter[];
}

/**
 * A collection of values returned by getInternetService.
 */
export interface GetInternetServiceResult {
    readonly filters?: outputs.GetInternetServiceFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internetServiceId: string;
    readonly netId: string;
    readonly requestId: string;
    readonly state: string;
    readonly tags: outputs.GetInternetServiceTag[];
}
export function getInternetServiceOutput(args?: GetInternetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInternetServiceResult> {
    return pulumi.output(args).apply((a: any) => getInternetService(a, opts))
}

/**
 * A collection of arguments for invoking getInternetService.
 */
export interface GetInternetServiceOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetInternetServiceFilterArgs>[]>;
}
