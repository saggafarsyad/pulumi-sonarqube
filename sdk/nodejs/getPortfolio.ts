// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Data Source: sonarqube.Portfolio
 *
 * Use this data source to get a Sonarqube portfolio resource
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const portfolio = sonarqube.getPortfolio({
 *     key: "portfolio-key",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPortfolio(args: GetPortfolioArgs, opts?: pulumi.InvokeOptions): Promise<GetPortfolioResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("sonarqube:index/getPortfolio:getPortfolio", {
        "key": args.key,
    }, opts);
}

/**
 * A collection of arguments for invoking getPortfolio.
 */
export interface GetPortfolioArgs {
    key: string;
}

/**
 * A collection of values returned by getPortfolio.
 */
export interface GetPortfolioResult {
    readonly branch: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    readonly name: string;
    readonly qualifier: string;
    readonly regexp: string;
    readonly selectionMode: string;
    readonly tags: string[];
    readonly visibility: string;
}
/**
 * ## # Data Source: sonarqube.Portfolio
 *
 * Use this data source to get a Sonarqube portfolio resource
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const portfolio = sonarqube.getPortfolio({
 *     key: "portfolio-key",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPortfolioOutput(args: GetPortfolioOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortfolioResult> {
    return pulumi.output(args).apply((a: any) => getPortfolio(a, opts))
}

/**
 * A collection of arguments for invoking getPortfolio.
 */
export interface GetPortfolioOutputArgs {
    key: pulumi.Input<string>;
}
