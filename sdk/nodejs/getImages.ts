// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getImages(args?: GetImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetImagesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getImages:getImages", {
        "accountIds": args.accountIds,
        "filters": args.filters,
        "imageIds": args.imageIds,
        "permissions": args.permissions,
    }, opts);
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesArgs {
    accountIds?: string[];
    filters?: inputs.GetImagesFilter[];
    imageIds?: string[];
    permissions?: string[];
}

/**
 * A collection of values returned by getImages.
 */
export interface GetImagesResult {
    readonly accountIds?: string[];
    readonly filters?: outputs.GetImagesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageIds?: string[];
    readonly images: outputs.GetImagesImage[];
    readonly permissions?: string[];
    readonly requestId: string;
}
export function getImagesOutput(args?: GetImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImagesResult> {
    return pulumi.output(args).apply((a: any) => getImages(a, opts))
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesOutputArgs {
    accountIds?: pulumi.Input<pulumi.Input<string>[]>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetImagesFilterArgs>[]>;
    imageIds?: pulumi.Input<pulumi.Input<string>[]>;
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
}
