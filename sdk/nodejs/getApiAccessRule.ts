// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getApiAccessRule(args?: GetApiAccessRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetApiAccessRuleResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getApiAccessRule:getApiAccessRule", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiAccessRule.
 */
export interface GetApiAccessRuleArgs {
    filters?: inputs.GetApiAccessRuleFilter[];
}

/**
 * A collection of values returned by getApiAccessRule.
 */
export interface GetApiAccessRuleResult {
    readonly apiAccessRuleId: string;
    readonly caIds: string[];
    readonly cns: string[];
    readonly description: string;
    readonly filters?: outputs.GetApiAccessRuleFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipRanges: string[];
    readonly requestId: string;
}
export function getApiAccessRuleOutput(args?: GetApiAccessRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiAccessRuleResult> {
    return pulumi.output(args).apply((a: any) => getApiAccessRule(a, opts))
}

/**
 * A collection of arguments for invoking getApiAccessRule.
 */
export interface GetApiAccessRuleOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetApiAccessRuleFilterArgs>[]>;
}
