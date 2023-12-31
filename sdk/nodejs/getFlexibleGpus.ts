// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFlexibleGpus(args?: GetFlexibleGpusArgs, opts?: pulumi.InvokeOptions): Promise<GetFlexibleGpusResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getFlexibleGpus:getFlexibleGpus", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlexibleGpus.
 */
export interface GetFlexibleGpusArgs {
    filters?: inputs.GetFlexibleGpusFilter[];
}

/**
 * A collection of values returned by getFlexibleGpus.
 */
export interface GetFlexibleGpusResult {
    readonly filters?: outputs.GetFlexibleGpusFilter[];
    readonly flexibleGpuses: outputs.GetFlexibleGpusFlexibleGpus[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly requestId: string;
}
export function getFlexibleGpusOutput(args?: GetFlexibleGpusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlexibleGpusResult> {
    return pulumi.output(args).apply((a: any) => getFlexibleGpus(a, opts))
}

/**
 * A collection of arguments for invoking getFlexibleGpus.
 */
export interface GetFlexibleGpusOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetFlexibleGpusFilterArgs>[]>;
}
