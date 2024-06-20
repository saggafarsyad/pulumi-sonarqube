// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/saggafarsyad/pulumi-sonarqube/sdk/go/sonarqube/internal"
)

var _ = internal.GetEnvOrDefault

// Allows anonymizing users on destroy. Requires Sonarqube version >= 9.7.
func GetAnonymizeUserOnDelete(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "sonarqube:anonymizeUserOnDelete")
}
func GetHost(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:host")
}
func GetHttpProxy(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:httpProxy")
}
func GetInstalledEdition(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:installedEdition")
}
func GetInstalledVersion(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:installedVersion")
}
func GetPass(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:pass")
}

// Allows ignoring insecure certificates when set to true. Defaults to false. Disabling TLS verification is dangerous and
// should only be done for local testing.
func GetTlsInsecureSkipVerify(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "sonarqube:tlsInsecureSkipVerify")
}
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:token")
}
func GetUser(ctx *pulumi.Context) string {
	return config.Get(ctx, "sonarqube:user")
}
