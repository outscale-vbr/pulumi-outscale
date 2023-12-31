// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getNetPeering(args?: GetNetPeeringArgs, opts?: pulumi.InvokeOptions): Promise<GetNetPeeringResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getNetPeering:getNetPeering", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetPeering.
 */
export interface GetNetPeeringArgs {
    filters?: inputs.GetNetPeeringFilter[];
}

/**
 * A collection of values returned by getNetPeering.
 */
export interface GetNetPeeringResult {
    readonly accepterNet: outputs.GetNetPeeringAccepterNet;
    readonly filters?: outputs.GetNetPeeringFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly netPeeringId: string;
    readonly requestId: string;
    readonly sourceNet: outputs.GetNetPeeringSourceNet;
    readonly state: outputs.GetNetPeeringState;
    readonly tags: outputs.GetNetPeeringTag[];
}
export function getNetPeeringOutput(args?: GetNetPeeringOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetPeeringResult> {
    return pulumi.output(args).apply((a: any) => getNetPeering(a, opts))
}

/**
 * A collection of arguments for invoking getNetPeering.
 */
export interface GetNetPeeringOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetNetPeeringFilterArgs>[]>;
}
