// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LoadBalancerAttributes extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerAttributes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerAttributesState, opts?: pulumi.CustomResourceOptions): LoadBalancerAttributes {
        return new LoadBalancerAttributes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/loadBalancerAttributes:LoadBalancerAttributes';

    /**
     * Returns true if the given object is an instance of LoadBalancerAttributes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerAttributes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerAttributes.__pulumiType;
    }

    public readonly accessLog!: pulumi.Output<outputs.LoadBalancerAttributesAccessLog>;
    public /*out*/ readonly applicationStickyCookiePolicies!: pulumi.Output<outputs.LoadBalancerAttributesApplicationStickyCookiePolicy[]>;
    public /*out*/ readonly backendVmIds!: pulumi.Output<string[]>;
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    public readonly healthCheck!: pulumi.Output<outputs.LoadBalancerAttributesHealthCheck>;
    public /*out*/ readonly listeners!: pulumi.Output<outputs.LoadBalancerAttributesListener[]>;
    public readonly loadBalancerName!: pulumi.Output<string>;
    public readonly loadBalancerPort!: pulumi.Output<number | undefined>;
    public /*out*/ readonly loadBalancerStickyCookiePolicies!: pulumi.Output<outputs.LoadBalancerAttributesLoadBalancerStickyCookiePolicy[]>;
    public /*out*/ readonly loadBalancerType!: pulumi.Output<string>;
    public readonly policyNames!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly securityGroups!: pulumi.Output<string[]>;
    public readonly serverCertificateId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly sourceSecurityGroup!: pulumi.Output<outputs.LoadBalancerAttributesSourceSecurityGroup>;
    public /*out*/ readonly subnets!: pulumi.Output<string[]>;
    public /*out*/ readonly subregionNames!: pulumi.Output<string[]>;
    public readonly tags!: pulumi.Output<outputs.LoadBalancerAttributesTag[]>;

    /**
     * Create a LoadBalancerAttributes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerAttributesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerAttributesArgs | LoadBalancerAttributesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerAttributesState | undefined;
            resourceInputs["accessLog"] = state ? state.accessLog : undefined;
            resourceInputs["applicationStickyCookiePolicies"] = state ? state.applicationStickyCookiePolicies : undefined;
            resourceInputs["backendVmIds"] = state ? state.backendVmIds : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["healthCheck"] = state ? state.healthCheck : undefined;
            resourceInputs["listeners"] = state ? state.listeners : undefined;
            resourceInputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            resourceInputs["loadBalancerPort"] = state ? state.loadBalancerPort : undefined;
            resourceInputs["loadBalancerStickyCookiePolicies"] = state ? state.loadBalancerStickyCookiePolicies : undefined;
            resourceInputs["loadBalancerType"] = state ? state.loadBalancerType : undefined;
            resourceInputs["policyNames"] = state ? state.policyNames : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["serverCertificateId"] = state ? state.serverCertificateId : undefined;
            resourceInputs["sourceSecurityGroup"] = state ? state.sourceSecurityGroup : undefined;
            resourceInputs["subnets"] = state ? state.subnets : undefined;
            resourceInputs["subregionNames"] = state ? state.subregionNames : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as LoadBalancerAttributesArgs | undefined;
            if ((!args || args.loadBalancerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            resourceInputs["accessLog"] = args ? args.accessLog : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            resourceInputs["loadBalancerPort"] = args ? args.loadBalancerPort : undefined;
            resourceInputs["policyNames"] = args ? args.policyNames : undefined;
            resourceInputs["serverCertificateId"] = args ? args.serverCertificateId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["applicationStickyCookiePolicies"] = undefined /*out*/;
            resourceInputs["backendVmIds"] = undefined /*out*/;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["listeners"] = undefined /*out*/;
            resourceInputs["loadBalancerStickyCookiePolicies"] = undefined /*out*/;
            resourceInputs["loadBalancerType"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["sourceSecurityGroup"] = undefined /*out*/;
            resourceInputs["subnets"] = undefined /*out*/;
            resourceInputs["subregionNames"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancerAttributes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerAttributes resources.
 */
export interface LoadBalancerAttributesState {
    accessLog?: pulumi.Input<inputs.LoadBalancerAttributesAccessLog>;
    applicationStickyCookiePolicies?: pulumi.Input<pulumi.Input<inputs.LoadBalancerAttributesApplicationStickyCookiePolicy>[]>;
    backendVmIds?: pulumi.Input<pulumi.Input<string>[]>;
    dnsName?: pulumi.Input<string>;
    healthCheck?: pulumi.Input<inputs.LoadBalancerAttributesHealthCheck>;
    listeners?: pulumi.Input<pulumi.Input<inputs.LoadBalancerAttributesListener>[]>;
    loadBalancerName?: pulumi.Input<string>;
    loadBalancerPort?: pulumi.Input<number>;
    loadBalancerStickyCookiePolicies?: pulumi.Input<pulumi.Input<inputs.LoadBalancerAttributesLoadBalancerStickyCookiePolicy>[]>;
    loadBalancerType?: pulumi.Input<string>;
    policyNames?: pulumi.Input<pulumi.Input<string>[]>;
    requestId?: pulumi.Input<string>;
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    serverCertificateId?: pulumi.Input<string>;
    sourceSecurityGroup?: pulumi.Input<inputs.LoadBalancerAttributesSourceSecurityGroup>;
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    subregionNames?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.LoadBalancerAttributesTag>[]>;
}

/**
 * The set of arguments for constructing a LoadBalancerAttributes resource.
 */
export interface LoadBalancerAttributesArgs {
    accessLog?: pulumi.Input<inputs.LoadBalancerAttributesAccessLog>;
    healthCheck?: pulumi.Input<inputs.LoadBalancerAttributesHealthCheck>;
    loadBalancerName: pulumi.Input<string>;
    loadBalancerPort?: pulumi.Input<number>;
    policyNames?: pulumi.Input<pulumi.Input<string>[]>;
    serverCertificateId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.LoadBalancerAttributesTag>[]>;
}
