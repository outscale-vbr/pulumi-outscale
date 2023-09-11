// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Vm extends pulumi.CustomResource {
    /**
     * Get an existing Vm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VmState, opts?: pulumi.CustomResourceOptions): Vm {
        return new Vm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/vm:Vm';

    /**
     * Returns true if the given object is an instance of Vm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vm.__pulumiType;
    }

    public /*out*/ readonly adminPassword!: pulumi.Output<string>;
    public /*out*/ readonly architecture!: pulumi.Output<string>;
    public readonly blockDeviceMappings!: pulumi.Output<outputs.VmBlockDeviceMapping[] | undefined>;
    public readonly blockDeviceMappingsCreateds!: pulumi.Output<outputs.VmBlockDeviceMappingsCreated[]>;
    public readonly bsuOptimized!: pulumi.Output<boolean>;
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    public readonly deletionProtection!: pulumi.Output<boolean>;
    public readonly getAdminPassword!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly hypervisor!: pulumi.Output<string>;
    public readonly imageId!: pulumi.Output<string>;
    public readonly isSourceDestChecked!: pulumi.Output<boolean>;
    public readonly keypairName!: pulumi.Output<string>;
    public /*out*/ readonly launchNumber!: pulumi.Output<number>;
    public readonly nestedVirtualization!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly netId!: pulumi.Output<string>;
    public readonly nics!: pulumi.Output<outputs.VmNic[]>;
    public /*out*/ readonly osFamily!: pulumi.Output<string>;
    public readonly performance!: pulumi.Output<string>;
    public readonly placementSubregionName!: pulumi.Output<string>;
    public readonly placementTenancy!: pulumi.Output<string>;
    public /*out*/ readonly privateDnsName!: pulumi.Output<string>;
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    public readonly privateIps!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly productCodes!: pulumi.Output<string[]>;
    public /*out*/ readonly publicDnsName!: pulumi.Output<string>;
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly reservationId!: pulumi.Output<string>;
    public /*out*/ readonly rootDeviceName!: pulumi.Output<string>;
    public /*out*/ readonly rootDeviceType!: pulumi.Output<string>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly securityGroupNames!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly securityGroups!: pulumi.Output<outputs.VmSecurityGroup[]>;
    public readonly state!: pulumi.Output<string | undefined>;
    public /*out*/ readonly stateReason!: pulumi.Output<string>;
    public readonly subnetId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.VmTag[] | undefined>;
    public readonly userData!: pulumi.Output<string | undefined>;
    public readonly vmId!: pulumi.Output<string>;
    public readonly vmInitiatedShutdownBehavior!: pulumi.Output<string>;
    public readonly vmType!: pulumi.Output<string>;

    /**
     * Create a Vm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VmArgs | VmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VmState | undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["architecture"] = state ? state.architecture : undefined;
            resourceInputs["blockDeviceMappings"] = state ? state.blockDeviceMappings : undefined;
            resourceInputs["blockDeviceMappingsCreateds"] = state ? state.blockDeviceMappingsCreateds : undefined;
            resourceInputs["bsuOptimized"] = state ? state.bsuOptimized : undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["getAdminPassword"] = state ? state.getAdminPassword : undefined;
            resourceInputs["hypervisor"] = state ? state.hypervisor : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["isSourceDestChecked"] = state ? state.isSourceDestChecked : undefined;
            resourceInputs["keypairName"] = state ? state.keypairName : undefined;
            resourceInputs["launchNumber"] = state ? state.launchNumber : undefined;
            resourceInputs["nestedVirtualization"] = state ? state.nestedVirtualization : undefined;
            resourceInputs["netId"] = state ? state.netId : undefined;
            resourceInputs["nics"] = state ? state.nics : undefined;
            resourceInputs["osFamily"] = state ? state.osFamily : undefined;
            resourceInputs["performance"] = state ? state.performance : undefined;
            resourceInputs["placementSubregionName"] = state ? state.placementSubregionName : undefined;
            resourceInputs["placementTenancy"] = state ? state.placementTenancy : undefined;
            resourceInputs["privateDnsName"] = state ? state.privateDnsName : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["productCodes"] = state ? state.productCodes : undefined;
            resourceInputs["publicDnsName"] = state ? state.publicDnsName : undefined;
            resourceInputs["publicIp"] = state ? state.publicIp : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["reservationId"] = state ? state.reservationId : undefined;
            resourceInputs["rootDeviceName"] = state ? state.rootDeviceName : undefined;
            resourceInputs["rootDeviceType"] = state ? state.rootDeviceType : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["securityGroupNames"] = state ? state.securityGroupNames : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["stateReason"] = state ? state.stateReason : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["vmId"] = state ? state.vmId : undefined;
            resourceInputs["vmInitiatedShutdownBehavior"] = state ? state.vmInitiatedShutdownBehavior : undefined;
            resourceInputs["vmType"] = state ? state.vmType : undefined;
        } else {
            const args = argsOrState as VmArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            resourceInputs["blockDeviceMappings"] = args ? args.blockDeviceMappings : undefined;
            resourceInputs["blockDeviceMappingsCreateds"] = args ? args.blockDeviceMappingsCreateds : undefined;
            resourceInputs["bsuOptimized"] = args ? args.bsuOptimized : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["getAdminPassword"] = args ? args.getAdminPassword : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["isSourceDestChecked"] = args ? args.isSourceDestChecked : undefined;
            resourceInputs["keypairName"] = args ? args.keypairName : undefined;
            resourceInputs["nestedVirtualization"] = args ? args.nestedVirtualization : undefined;
            resourceInputs["nics"] = args ? args.nics : undefined;
            resourceInputs["performance"] = args ? args.performance : undefined;
            resourceInputs["placementSubregionName"] = args ? args.placementSubregionName : undefined;
            resourceInputs["placementTenancy"] = args ? args.placementTenancy : undefined;
            resourceInputs["privateIps"] = args ? args.privateIps : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["securityGroupNames"] = args ? args.securityGroupNames : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["vmId"] = args ? args.vmId : undefined;
            resourceInputs["vmInitiatedShutdownBehavior"] = args ? args.vmInitiatedShutdownBehavior : undefined;
            resourceInputs["vmType"] = args ? args.vmType : undefined;
            resourceInputs["adminPassword"] = undefined /*out*/;
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["clientToken"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["hypervisor"] = undefined /*out*/;
            resourceInputs["launchNumber"] = undefined /*out*/;
            resourceInputs["netId"] = undefined /*out*/;
            resourceInputs["osFamily"] = undefined /*out*/;
            resourceInputs["privateDnsName"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["productCodes"] = undefined /*out*/;
            resourceInputs["publicDnsName"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["reservationId"] = undefined /*out*/;
            resourceInputs["rootDeviceName"] = undefined /*out*/;
            resourceInputs["rootDeviceType"] = undefined /*out*/;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["stateReason"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vm resources.
 */
export interface VmState {
    adminPassword?: pulumi.Input<string>;
    architecture?: pulumi.Input<string>;
    blockDeviceMappings?: pulumi.Input<pulumi.Input<inputs.VmBlockDeviceMapping>[]>;
    blockDeviceMappingsCreateds?: pulumi.Input<pulumi.Input<inputs.VmBlockDeviceMappingsCreated>[]>;
    bsuOptimized?: pulumi.Input<boolean>;
    clientToken?: pulumi.Input<string>;
    creationDate?: pulumi.Input<string>;
    deletionProtection?: pulumi.Input<boolean>;
    getAdminPassword?: pulumi.Input<boolean>;
    hypervisor?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    isSourceDestChecked?: pulumi.Input<boolean>;
    keypairName?: pulumi.Input<string>;
    launchNumber?: pulumi.Input<number>;
    nestedVirtualization?: pulumi.Input<boolean>;
    netId?: pulumi.Input<string>;
    nics?: pulumi.Input<pulumi.Input<inputs.VmNic>[]>;
    osFamily?: pulumi.Input<string>;
    performance?: pulumi.Input<string>;
    placementSubregionName?: pulumi.Input<string>;
    placementTenancy?: pulumi.Input<string>;
    privateDnsName?: pulumi.Input<string>;
    privateIp?: pulumi.Input<string>;
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    productCodes?: pulumi.Input<pulumi.Input<string>[]>;
    publicDnsName?: pulumi.Input<string>;
    publicIp?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    reservationId?: pulumi.Input<string>;
    rootDeviceName?: pulumi.Input<string>;
    rootDeviceType?: pulumi.Input<string>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    securityGroups?: pulumi.Input<pulumi.Input<inputs.VmSecurityGroup>[]>;
    state?: pulumi.Input<string>;
    stateReason?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.VmTag>[]>;
    userData?: pulumi.Input<string>;
    vmId?: pulumi.Input<string>;
    vmInitiatedShutdownBehavior?: pulumi.Input<string>;
    vmType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vm resource.
 */
export interface VmArgs {
    blockDeviceMappings?: pulumi.Input<pulumi.Input<inputs.VmBlockDeviceMapping>[]>;
    blockDeviceMappingsCreateds?: pulumi.Input<pulumi.Input<inputs.VmBlockDeviceMappingsCreated>[]>;
    bsuOptimized?: pulumi.Input<boolean>;
    deletionProtection?: pulumi.Input<boolean>;
    getAdminPassword?: pulumi.Input<boolean>;
    imageId: pulumi.Input<string>;
    isSourceDestChecked?: pulumi.Input<boolean>;
    keypairName?: pulumi.Input<string>;
    nestedVirtualization?: pulumi.Input<boolean>;
    nics?: pulumi.Input<pulumi.Input<inputs.VmNic>[]>;
    performance?: pulumi.Input<string>;
    placementSubregionName?: pulumi.Input<string>;
    placementTenancy?: pulumi.Input<string>;
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    state?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.VmTag>[]>;
    userData?: pulumi.Input<string>;
    vmId?: pulumi.Input<string>;
    vmInitiatedShutdownBehavior?: pulumi.Input<string>;
    vmType?: pulumi.Input<string>;
}