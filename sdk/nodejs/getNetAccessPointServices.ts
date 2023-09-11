// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getNetAccessPointServices(args?: GetNetAccessPointServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetNetAccessPointServicesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getNetAccessPointServices:getNetAccessPointServices", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetAccessPointServices.
 */
export interface GetNetAccessPointServicesArgs {
    filters?: inputs.GetNetAccessPointServicesFilter[];
}

/**
 * A collection of values returned by getNetAccessPointServices.
 */
export interface GetNetAccessPointServicesResult {
    readonly filters?: outputs.GetNetAccessPointServicesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly requestId: string;
    readonly services: outputs.GetNetAccessPointServicesService[];
}
export function getNetAccessPointServicesOutput(args?: GetNetAccessPointServicesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetAccessPointServicesResult> {
    return pulumi.output(args).apply((a: any) => getNetAccessPointServices(a, opts))
}

/**
 * A collection of arguments for invoking getNetAccessPointServices.
 */
export interface GetNetAccessPointServicesOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetNetAccessPointServicesFilterArgs>[]>;
}