// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sonarqube

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/saggafarsyad/pulumi-sonarqube/sdk/go/sonarqube/internal"
)

// The provider type for the sonarqube package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	Host             pulumi.StringOutput    `pulumi:"host"`
	HttpProxy        pulumi.StringPtrOutput `pulumi:"httpProxy"`
	InstalledEdition pulumi.StringPtrOutput `pulumi:"installedEdition"`
	InstalledVersion pulumi.StringPtrOutput `pulumi:"installedVersion"`
	Pass             pulumi.StringPtrOutput `pulumi:"pass"`
	Token            pulumi.StringPtrOutput `pulumi:"token"`
	User             pulumi.StringPtrOutput `pulumi:"user"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Pass != nil {
		args.Pass = pulumi.ToSecret(args.Pass).(pulumi.StringPtrInput)
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"pass",
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:sonarqube", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Allows anonymizing users on destroy. Requires Sonarqube version >= 9.7.
	AnonymizeUserOnDelete *bool   `pulumi:"anonymizeUserOnDelete"`
	Host                  string  `pulumi:"host"`
	HttpProxy             *string `pulumi:"httpProxy"`
	InstalledEdition      *string `pulumi:"installedEdition"`
	InstalledVersion      *string `pulumi:"installedVersion"`
	Pass                  *string `pulumi:"pass"`
	// Allows ignoring insecure certificates when set to true. Defaults to false. Disabling TLS verification is dangerous and
	// should only be done for local testing.
	TlsInsecureSkipVerify *bool   `pulumi:"tlsInsecureSkipVerify"`
	Token                 *string `pulumi:"token"`
	User                  *string `pulumi:"user"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Allows anonymizing users on destroy. Requires Sonarqube version >= 9.7.
	AnonymizeUserOnDelete pulumi.BoolPtrInput
	Host                  pulumi.StringInput
	HttpProxy             pulumi.StringPtrInput
	InstalledEdition      pulumi.StringPtrInput
	InstalledVersion      pulumi.StringPtrInput
	Pass                  pulumi.StringPtrInput
	// Allows ignoring insecure certificates when set to true. Defaults to false. Disabling TLS verification is dangerous and
	// should only be done for local testing.
	TlsInsecureSkipVerify pulumi.BoolPtrInput
	Token                 pulumi.StringPtrInput
	User                  pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

func (o ProviderOutput) HttpProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.HttpProxy }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) InstalledEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.InstalledEdition }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) InstalledVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.InstalledVersion }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Pass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Pass }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}