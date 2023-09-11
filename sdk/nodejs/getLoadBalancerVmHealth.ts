// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getLoadBalancerVmHealth(args: GetLoadBalancerVmHealthArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerVmHealthResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getLoadBalancerVmHealth:getLoadBalancerVmHealth", {
        "backendVmIds": args.backendVmIds,
        "filters": args.filters,
        "loadBalancerName": args.loadBalancerName,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancerVmHealth.
 */
export interface GetLoadBalancerVmHealthArgs {
    backendVmIds?: string[];
    filters?: inputs.GetLoadBalancerVmHealthFilter[];
    loadBalancerName: string;
}

/**
 * A collection of values returned by getLoadBalancerVmHealth.
 */
export interface GetLoadBalancerVmHealthResult {
    readonly backendVmHealths: outputs.GetLoadBalancerVmHealthBackendVmHealth[];
    readonly backendVmIds?: string[];
    readonly filters?: outputs.GetLoadBalancerVmHealthFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loadBalancerName: string;
    readonly requestId: string;
}
export function getLoadBalancerVmHealthOutput(args: GetLoadBalancerVmHealthOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadBalancerVmHealthResult> {
    return pulumi.output(args).apply((a: any) => getLoadBalancerVmHealth(a, opts))
}

/**
 * A collection of arguments for invoking getLoadBalancerVmHealth.
 */
export interface GetLoadBalancerVmHealthOutputArgs {
    backendVmIds?: pulumi.Input<pulumi.Input<string>[]>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetLoadBalancerVmHealthFilterArgs>[]>;
    loadBalancerName: pulumi.Input<string>;
}