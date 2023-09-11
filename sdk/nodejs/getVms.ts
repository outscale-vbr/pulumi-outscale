// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getVms(args?: GetVmsArgs, opts?: pulumi.InvokeOptions): Promise<GetVmsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getVms:getVms", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getVms.
 */
export interface GetVmsArgs {
    filters?: inputs.GetVmsFilter[];
}

/**
 * A collection of values returned by getVms.
 */
export interface GetVmsResult {
    readonly filters?: outputs.GetVmsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly requestId: string;
    readonly vms: outputs.GetVmsVm[];
}
export function getVmsOutput(args?: GetVmsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVmsResult> {
    return pulumi.output(args).apply((a: any) => getVms(a, opts))
}

/**
 * A collection of arguments for invoking getVms.
 */
export interface GetVmsOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetVmsFilterArgs>[]>;
}
