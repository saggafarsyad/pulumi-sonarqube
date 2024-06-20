// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # sonarqubeUserExternalEntity
 *
 * Updates the *external identity* of a *non local* Sonarqube User. This can be used to set the *Identity Provider* which should be used to
 * authenticate a specific user.
 *
 * The Sonarqube API currently does not provide an endpoint to read the *external identity* setting of an user.
 *
 * ## Example: change the external identity to SAML
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const remoteUserUser = new sonarqube.User("remoteUserUser", {
 *     loginName: "terraform-test",
 *     email: "terraform-test@sonarqube.com",
 *     isLocal: false,
 * });
 * const remoteUserUserExternalIdentity = new sonarqube.UserExternalIdentity("remoteUserUserExternalIdentity", {
 *     loginName: remoteUserUser.loginName,
 *     externalIdentity: "terraform-test@sonarqube.com",
 *     externalProvider: "saml",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class UserExternalIdentity extends pulumi.CustomResource {
    /**
     * Get an existing UserExternalIdentity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserExternalIdentityState, opts?: pulumi.CustomResourceOptions): UserExternalIdentity {
        return new UserExternalIdentity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sonarqube:index/userExternalIdentity:UserExternalIdentity';

    /**
     * Returns true if the given object is an instance of UserExternalIdentity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserExternalIdentity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserExternalIdentity.__pulumiType;
    }

    /**
     * The identifier of the User used by the Authentication Provider. Changing this forces a new resource to be 
     * created.
     */
    public readonly externalIdentity!: pulumi.Output<string>;
    /**
     * The key of the Authentication Provider. The Authentication Provider must be activated on Sonarqube. Changing 
     * this forces a new resource to be created.
     */
    public readonly externalProvider!: pulumi.Output<string>;
    /**
     * The login name of the User to update. Changing this forces a new resource to be created.
     */
    public readonly loginName!: pulumi.Output<string>;

    /**
     * Create a UserExternalIdentity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserExternalIdentityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserExternalIdentityArgs | UserExternalIdentityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserExternalIdentityState | undefined;
            resourceInputs["externalIdentity"] = state ? state.externalIdentity : undefined;
            resourceInputs["externalProvider"] = state ? state.externalProvider : undefined;
            resourceInputs["loginName"] = state ? state.loginName : undefined;
        } else {
            const args = argsOrState as UserExternalIdentityArgs | undefined;
            if ((!args || args.externalIdentity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalIdentity'");
            }
            if ((!args || args.externalProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalProvider'");
            }
            if ((!args || args.loginName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loginName'");
            }
            resourceInputs["externalIdentity"] = args ? args.externalIdentity : undefined;
            resourceInputs["externalProvider"] = args ? args.externalProvider : undefined;
            resourceInputs["loginName"] = args ? args.loginName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserExternalIdentity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserExternalIdentity resources.
 */
export interface UserExternalIdentityState {
    /**
     * The identifier of the User used by the Authentication Provider. Changing this forces a new resource to be 
     * created.
     */
    externalIdentity?: pulumi.Input<string>;
    /**
     * The key of the Authentication Provider. The Authentication Provider must be activated on Sonarqube. Changing 
     * this forces a new resource to be created.
     */
    externalProvider?: pulumi.Input<string>;
    /**
     * The login name of the User to update. Changing this forces a new resource to be created.
     */
    loginName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserExternalIdentity resource.
 */
export interface UserExternalIdentityArgs {
    /**
     * The identifier of the User used by the Authentication Provider. Changing this forces a new resource to be 
     * created.
     */
    externalIdentity: pulumi.Input<string>;
    /**
     * The key of the Authentication Provider. The Authentication Provider must be activated on Sonarqube. Changing 
     * this forces a new resource to be created.
     */
    externalProvider: pulumi.Input<string>;
    /**
     * The login name of the User to update. Changing this forces a new resource to be created.
     */
    loginName: pulumi.Input<string>;
}
