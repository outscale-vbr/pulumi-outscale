// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getNatService(args?: GetNatServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetNatServiceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getNatService:getNatService", {
        "filters": args.filters,
        "natServiceId": args.natServiceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatService.
 */
export interface GetNatServiceArgs {
    filters?: inputs.GetNatServiceFilter[];
    natServiceId?: string;
}

/**
 * A collection of values returned by getNatService.
 */
export interface GetNatServiceResult {
    readonly filters?: outputs.GetNatServiceFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly natServiceId?: string;
    readonly netId: string;
    readonly publicIps: outputs.GetNatServicePublicIp[];
    readonly requestId: string;
    readonly state: string;
    readonly subnetId: string;
    readonly tags: outputs.GetNatServiceTag[];
}
export function getNatServiceOutput(args?: GetNatServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNatServiceResult> {
    return pulumi.output(args).apply((a: any) => getNatService(a, opts))
}

/**
 * A collection of arguments for invoking getNatService.
 */
export interface GetNatServiceOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetNatServiceFilterArgs>[]>;
    natServiceId?: pulumi.Input<string>;
}
