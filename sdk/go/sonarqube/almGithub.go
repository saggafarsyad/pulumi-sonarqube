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

// ## # AlmGithub
//
// Provides a Sonarqube GitHub Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops
// Platform Integration for GitHub.
//
// ## Example: Create a GitHub Alm Integration
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/saggafarsyad/pulumi-sonarqube/sdk/go/sonarqube"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sonarqube.NewAlmGithub(ctx, "github-alm", &sonarqube.AlmGithubArgs{
//				AppId:         pulumi.String("12345"),
//				ClientId:      pulumi.String("56789"),
//				ClientSecret:  pulumi.String("secret"),
//				Key:           pulumi.String("myalm"),
//				PrivateKey:    pulumi.String("myprivate_key"),
//				Url:           pulumi.String("https://api.github.com"),
//				WebhookSecret: pulumi.String("mysecret"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type AlmGithub struct {
	pulumi.CustomResourceState

	AppId         pulumi.StringOutput    `pulumi:"appId"`
	ClientId      pulumi.StringOutput    `pulumi:"clientId"`
	ClientSecret  pulumi.StringOutput    `pulumi:"clientSecret"`
	Key           pulumi.StringOutput    `pulumi:"key"`
	PrivateKey    pulumi.StringOutput    `pulumi:"privateKey"`
	Url           pulumi.StringOutput    `pulumi:"url"`
	WebhookSecret pulumi.StringPtrOutput `pulumi:"webhookSecret"`
}

// NewAlmGithub registers a new resource with the given unique name, arguments, and options.
func NewAlmGithub(ctx *pulumi.Context,
	name string, args *AlmGithubArgs, opts ...pulumi.ResourceOption) (*AlmGithub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AlmGithub
	err := ctx.RegisterResource("sonarqube:index/almGithub:AlmGithub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlmGithub gets an existing AlmGithub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlmGithub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlmGithubState, opts ...pulumi.ResourceOption) (*AlmGithub, error) {
	var resource AlmGithub
	err := ctx.ReadResource("sonarqube:index/almGithub:AlmGithub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlmGithub resources.
type almGithubState struct {
	AppId         *string `pulumi:"appId"`
	ClientId      *string `pulumi:"clientId"`
	ClientSecret  *string `pulumi:"clientSecret"`
	Key           *string `pulumi:"key"`
	PrivateKey    *string `pulumi:"privateKey"`
	Url           *string `pulumi:"url"`
	WebhookSecret *string `pulumi:"webhookSecret"`
}

type AlmGithubState struct {
	AppId         pulumi.StringPtrInput
	ClientId      pulumi.StringPtrInput
	ClientSecret  pulumi.StringPtrInput
	Key           pulumi.StringPtrInput
	PrivateKey    pulumi.StringPtrInput
	Url           pulumi.StringPtrInput
	WebhookSecret pulumi.StringPtrInput
}

func (AlmGithubState) ElementType() reflect.Type {
	return reflect.TypeOf((*almGithubState)(nil)).Elem()
}

type almGithubArgs struct {
	AppId         string  `pulumi:"appId"`
	ClientId      string  `pulumi:"clientId"`
	ClientSecret  string  `pulumi:"clientSecret"`
	Key           string  `pulumi:"key"`
	PrivateKey    string  `pulumi:"privateKey"`
	Url           string  `pulumi:"url"`
	WebhookSecret *string `pulumi:"webhookSecret"`
}

// The set of arguments for constructing a AlmGithub resource.
type AlmGithubArgs struct {
	AppId         pulumi.StringInput
	ClientId      pulumi.StringInput
	ClientSecret  pulumi.StringInput
	Key           pulumi.StringInput
	PrivateKey    pulumi.StringInput
	Url           pulumi.StringInput
	WebhookSecret pulumi.StringPtrInput
}

func (AlmGithubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*almGithubArgs)(nil)).Elem()
}

type AlmGithubInput interface {
	pulumi.Input

	ToAlmGithubOutput() AlmGithubOutput
	ToAlmGithubOutputWithContext(ctx context.Context) AlmGithubOutput
}

func (*AlmGithub) ElementType() reflect.Type {
	return reflect.TypeOf((**AlmGithub)(nil)).Elem()
}

func (i *AlmGithub) ToAlmGithubOutput() AlmGithubOutput {
	return i.ToAlmGithubOutputWithContext(context.Background())
}

func (i *AlmGithub) ToAlmGithubOutputWithContext(ctx context.Context) AlmGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmGithubOutput)
}

// AlmGithubArrayInput is an input type that accepts AlmGithubArray and AlmGithubArrayOutput values.
// You can construct a concrete instance of `AlmGithubArrayInput` via:
//
//	AlmGithubArray{ AlmGithubArgs{...} }
type AlmGithubArrayInput interface {
	pulumi.Input

	ToAlmGithubArrayOutput() AlmGithubArrayOutput
	ToAlmGithubArrayOutputWithContext(context.Context) AlmGithubArrayOutput
}

type AlmGithubArray []AlmGithubInput

func (AlmGithubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlmGithub)(nil)).Elem()
}

func (i AlmGithubArray) ToAlmGithubArrayOutput() AlmGithubArrayOutput {
	return i.ToAlmGithubArrayOutputWithContext(context.Background())
}

func (i AlmGithubArray) ToAlmGithubArrayOutputWithContext(ctx context.Context) AlmGithubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmGithubArrayOutput)
}

// AlmGithubMapInput is an input type that accepts AlmGithubMap and AlmGithubMapOutput values.
// You can construct a concrete instance of `AlmGithubMapInput` via:
//
//	AlmGithubMap{ "key": AlmGithubArgs{...} }
type AlmGithubMapInput interface {
	pulumi.Input

	ToAlmGithubMapOutput() AlmGithubMapOutput
	ToAlmGithubMapOutputWithContext(context.Context) AlmGithubMapOutput
}

type AlmGithubMap map[string]AlmGithubInput

func (AlmGithubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlmGithub)(nil)).Elem()
}

func (i AlmGithubMap) ToAlmGithubMapOutput() AlmGithubMapOutput {
	return i.ToAlmGithubMapOutputWithContext(context.Background())
}

func (i AlmGithubMap) ToAlmGithubMapOutputWithContext(ctx context.Context) AlmGithubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmGithubMapOutput)
}

type AlmGithubOutput struct{ *pulumi.OutputState }

func (AlmGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlmGithub)(nil)).Elem()
}

func (o AlmGithubOutput) ToAlmGithubOutput() AlmGithubOutput {
	return o
}

func (o AlmGithubOutput) ToAlmGithubOutputWithContext(ctx context.Context) AlmGithubOutput {
	return o
}

func (o AlmGithubOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o AlmGithubOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o AlmGithubOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o AlmGithubOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o AlmGithubOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o AlmGithubOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o AlmGithubOutput) WebhookSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlmGithub) pulumi.StringPtrOutput { return v.WebhookSecret }).(pulumi.StringPtrOutput)
}

type AlmGithubArrayOutput struct{ *pulumi.OutputState }

func (AlmGithubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlmGithub)(nil)).Elem()
}

func (o AlmGithubArrayOutput) ToAlmGithubArrayOutput() AlmGithubArrayOutput {
	return o
}

func (o AlmGithubArrayOutput) ToAlmGithubArrayOutputWithContext(ctx context.Context) AlmGithubArrayOutput {
	return o
}

func (o AlmGithubArrayOutput) Index(i pulumi.IntInput) AlmGithubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlmGithub {
		return vs[0].([]*AlmGithub)[vs[1].(int)]
	}).(AlmGithubOutput)
}

type AlmGithubMapOutput struct{ *pulumi.OutputState }

func (AlmGithubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlmGithub)(nil)).Elem()
}

func (o AlmGithubMapOutput) ToAlmGithubMapOutput() AlmGithubMapOutput {
	return o
}

func (o AlmGithubMapOutput) ToAlmGithubMapOutputWithContext(ctx context.Context) AlmGithubMapOutput {
	return o
}

func (o AlmGithubMapOutput) MapIndex(k pulumi.StringInput) AlmGithubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlmGithub {
		return vs[0].(map[string]*AlmGithub)[vs[1].(string)]
	}).(AlmGithubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlmGithubInput)(nil)).Elem(), &AlmGithub{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlmGithubArrayInput)(nil)).Elem(), AlmGithubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlmGithubMapInput)(nil)).Elem(), AlmGithubMap{})
	pulumi.RegisterOutputType(AlmGithubOutput{})
	pulumi.RegisterOutputType(AlmGithubArrayOutput{})
	pulumi.RegisterOutputType(AlmGithubMapOutput{})
}
