// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LoadBalancerVms extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerVms resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerVmsState, opts?: pulumi.CustomResourceOptions): LoadBalancerVms {
        return new LoadBalancerVms(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/loadBalancerVms:LoadBalancerVms';

    /**
     * Returns true if the given object is an instance of LoadBalancerVms.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerVms {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerVms.__pulumiType;
    }

    public readonly backendVmIds!: pulumi.Output<string[]>;
    public readonly loadBalancerName!: pulumi.Output<string>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancerVms resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerVmsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerVmsArgs | LoadBalancerVmsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerVmsState | undefined;
            resourceInputs["backendVmIds"] = state ? state.backendVmIds : undefined;
            resourceInputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
        } else {
            const args = argsOrState as LoadBalancerVmsArgs | undefined;
            if ((!args || args.backendVmIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendVmIds'");
            }
            if ((!args || args.loadBalancerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            resourceInputs["backendVmIds"] = args ? args.backendVmIds : undefined;
            resourceInputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            resourceInputs["requestId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancerVms.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerVms resources.
 */
export interface LoadBalancerVmsState {
    backendVmIds?: pulumi.Input<pulumi.Input<string>[]>;
    loadBalancerName?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancerVms resource.
 */
export interface LoadBalancerVmsArgs {
    backendVmIds: pulumi.Input<pulumi.Input<string>[]>;
    loadBalancerName: pulumi.Input<string>;
}
