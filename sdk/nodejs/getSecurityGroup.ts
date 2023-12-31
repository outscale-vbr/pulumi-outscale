// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSecurityGroup(args?: GetSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("outscale:index/getSecurityGroup:getSecurityGroup", {
        "filters": args.filters,
        "securityGroupId": args.securityGroupId,
        "securityGroupName": args.securityGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupArgs {
    filters?: inputs.GetSecurityGroupFilter[];
    securityGroupId?: string;
    securityGroupName?: string;
}

/**
 * A collection of values returned by getSecurityGroup.
 */
export interface GetSecurityGroupResult {
    readonly accountId: string;
    readonly description: string;
    readonly filters?: outputs.GetSecurityGroupFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly inboundRules: outputs.GetSecurityGroupInboundRule[];
    readonly netId: string;
    readonly outboundRules: outputs.GetSecurityGroupOutboundRule[];
    readonly requestId: string;
    readonly securityGroupId: string;
    readonly securityGroupName: string;
    readonly tags: outputs.GetSecurityGroupTag[];
}
export function getSecurityGroupOutput(args?: GetSecurityGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupResult> {
    return pulumi.output(args).apply((a: any) => getSecurityGroup(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.GetSecurityGroupFilterArgs>[]>;
    securityGroupId?: pulumi.Input<string>;
    securityGroupName?: pulumi.Input<string>;
}
