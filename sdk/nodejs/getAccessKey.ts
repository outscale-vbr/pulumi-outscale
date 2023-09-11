// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getAccessKey(args?: GetAccessKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessKeyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getAccessKey:getAccessKey", {
        "accessKeyId": args.accessKeyId,
        "filters": args.filters,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessKey.
 */
export interface GetAccessKeyArgs {
    accessKeyId?: string;
    filters?: inputs.GetAccessKeyFilter[];
    state?: string;
}

/**
 * A collection of values returned by getAccessKey.
 */
export interface GetAccessKeyResult {
    readonly accessKeyId?: string;
    readonly creationDate: string;
    readonly expirationDate: string;
    readonly filters?: outputs.GetAccessKeyFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastModificationDate: string;
    readonly requestId: string;
    readonly state?: string;
}
export function getAccessKeyOutput(args?: GetAccessKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessKeyResult> {
    return pulumi.output(args).apply((a: any) => getAccessKey(a, opts))
}

/**
 * A collection of arguments for invoking getAccessKey.
 */
export interface GetAccessKeyOutputArgs {
    accessKeyId?: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetAccessKeyFilterArgs>[]>;
    state?: pulumi.Input<string>;
}