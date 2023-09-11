// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getImageExportTask(args?: GetImageExportTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetImageExportTaskResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getImageExportTask:getImageExportTask", {
        "dryRun": args.dryRun,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageExportTask.
 */
export interface GetImageExportTaskArgs {
    dryRun?: boolean;
    filters?: inputs.GetImageExportTaskFilter[];
}

/**
 * A collection of values returned by getImageExportTask.
 */
export interface GetImageExportTaskResult {
    readonly comment: string;
    readonly dryRun: boolean;
    readonly filters?: outputs.GetImageExportTaskFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId: string;
    readonly osuExports: outputs.GetImageExportTaskOsuExport[];
    readonly progress: number;
    readonly requestId: string;
    readonly state: string;
    readonly tags: outputs.GetImageExportTaskTag[];
    readonly taskId: string;
}
export function getImageExportTaskOutput(args?: GetImageExportTaskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageExportTaskResult> {
    return pulumi.output(args).apply((a: any) => getImageExportTask(a, opts))
}

/**
 * A collection of arguments for invoking getImageExportTask.
 */
export interface GetImageExportTaskOutputArgs {
    dryRun?: pulumi.Input<boolean>;
    filters?: pulumi.Input<pulumi.Input<inputs.GetImageExportTaskFilterArgs>[]>;
}