// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # sonarqube.UserToken
 *
 * Provides a Sonarqube User token resource. This can be used to manage Sonarqube User tokens.
 *
 * ## Example: create a user, user token and output the token value
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const user = new sonarqube.User("user", {
 *     loginName: "terraform-test",
 *     password: "secret-sauce37!",
 * });
 * const token = new sonarqube.UserToken("token", {loginName: user.loginName});
 * export const userToken = token.token;
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Example: create an expiring global analysis token and output the token value
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const token = new sonarqube.UserToken("token", {
 *     type: "GLOBAL_ANALYSIS_TOKEN",
 *     expirationDate: "2099-01-01",
 * });
 * export const globalAnalysisToken = token.token;
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Example: create a project, project analysis token, and output the token value
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const token = new sonarqube.UserToken("token", {
 *     type: "PROJECT_ANALYSIS_TOKEN",
 *     projectKey: "my-project",
 * });
 * export const projectAnalysisToken = token.token;
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Import is not supported for this resource.
 */
export class UserToken extends pulumi.CustomResource {
    /**
     * Get an existing UserToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserTokenState, opts?: pulumi.CustomResourceOptions): UserToken {
        return new UserToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sonarqube:index/userToken:UserToken';

    /**
     * Returns true if the given object is an instance of UserToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserToken.__pulumiType;
    }

    public readonly expirationDate!: pulumi.Output<string>;
    public readonly loginName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly projectKey!: pulumi.Output<string | undefined>;
    public /*out*/ readonly token!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a UserToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserTokenArgs | UserTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserTokenState | undefined;
            resourceInputs["expirationDate"] = state ? state.expirationDate : undefined;
            resourceInputs["loginName"] = state ? state.loginName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as UserTokenArgs | undefined;
            resourceInputs["expirationDate"] = args ? args.expirationDate : undefined;
            resourceInputs["loginName"] = args ? args.loginName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["token"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(UserToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserToken resources.
 */
export interface UserTokenState {
    expirationDate?: pulumi.Input<string>;
    loginName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    projectKey?: pulumi.Input<string>;
    token?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserToken resource.
 */
export interface UserTokenArgs {
    expirationDate?: pulumi.Input<string>;
    loginName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    projectKey?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
