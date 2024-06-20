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

// ## # AlmAzure
//
// Provides a Sonarqube Azure Devops Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops
// Platform Integration for Azure Devops.
//
// ## Example: Create an Azure Devops Alm Integration
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
//			_, err := sonarqube.NewAlmAzure(ctx, "az1", &sonarqube.AlmAzureArgs{
//				Key:                 pulumi.String("az1"),
//				PersonalAccessToken: pulumi.String("my_pat"),
//				Url:                 pulumi.String("https://dev.azure.com/my-org"),
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
//
// ## Import
//
// # Resource can be imported providing their Alm Instance Key and Azure DevOps Personal Access Token
//
// terraform
//
// ```sh
// $ pulumi import sonarqube:index/almAzure:AlmAzure az1 key/personal_access_token
// ```
type AlmAzure struct {
	pulumi.CustomResourceState

	// Unique key of the Azure Devops instance setting
	Key pulumi.StringOutput `pulumi:"key"`
	// Azure Devops personal access token
	PersonalAccessToken pulumi.StringOutput `pulumi:"personalAccessToken"`
	// Azure API URL
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewAlmAzure registers a new resource with the given unique name, arguments, and options.
func NewAlmAzure(ctx *pulumi.Context,
	name string, args *AlmAzureArgs, opts ...pulumi.ResourceOption) (*AlmAzure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.PersonalAccessToken == nil {
		return nil, errors.New("invalid value for required argument 'PersonalAccessToken'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.PersonalAccessToken != nil {
		args.PersonalAccessToken = pulumi.ToSecret(args.PersonalAccessToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"personalAccessToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AlmAzure
	err := ctx.RegisterResource("sonarqube:index/almAzure:AlmAzure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlmAzure gets an existing AlmAzure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlmAzure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlmAzureState, opts ...pulumi.ResourceOption) (*AlmAzure, error) {
	var resource AlmAzure
	err := ctx.ReadResource("sonarqube:index/almAzure:AlmAzure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlmAzure resources.
type almAzureState struct {
	// Unique key of the Azure Devops instance setting
	Key *string `pulumi:"key"`
	// Azure Devops personal access token
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
	// Azure API URL
	Url *string `pulumi:"url"`
}

type AlmAzureState struct {
	// Unique key of the Azure Devops instance setting
	Key pulumi.StringPtrInput
	// Azure Devops personal access token
	PersonalAccessToken pulumi.StringPtrInput
	// Azure API URL
	Url pulumi.StringPtrInput
}

func (AlmAzureState) ElementType() reflect.Type {
	return reflect.TypeOf((*almAzureState)(nil)).Elem()
}

type almAzureArgs struct {
	// Unique key of the Azure Devops instance setting
	Key string `pulumi:"key"`
	// Azure Devops personal access token
	PersonalAccessToken string `pulumi:"personalAccessToken"`
	// Azure API URL
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a AlmAzure resource.
type AlmAzureArgs struct {
	// Unique key of the Azure Devops instance setting
	Key pulumi.StringInput
	// Azure Devops personal access token
	PersonalAccessToken pulumi.StringInput
	// Azure API URL
	Url pulumi.StringInput
}

func (AlmAzureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*almAzureArgs)(nil)).Elem()
}

type AlmAzureInput interface {
	pulumi.Input

	ToAlmAzureOutput() AlmAzureOutput
	ToAlmAzureOutputWithContext(ctx context.Context) AlmAzureOutput
}

func (*AlmAzure) ElementType() reflect.Type {
	return reflect.TypeOf((**AlmAzure)(nil)).Elem()
}

func (i *AlmAzure) ToAlmAzureOutput() AlmAzureOutput {
	return i.ToAlmAzureOutputWithContext(context.Background())
}

func (i *AlmAzure) ToAlmAzureOutputWithContext(ctx context.Context) AlmAzureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmAzureOutput)
}

// AlmAzureArrayInput is an input type that accepts AlmAzureArray and AlmAzureArrayOutput values.
// You can construct a concrete instance of `AlmAzureArrayInput` via:
//
//	AlmAzureArray{ AlmAzureArgs{...} }
type AlmAzureArrayInput interface {
	pulumi.Input

	ToAlmAzureArrayOutput() AlmAzureArrayOutput
	ToAlmAzureArrayOutputWithContext(context.Context) AlmAzureArrayOutput
}

type AlmAzureArray []AlmAzureInput

func (AlmAzureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlmAzure)(nil)).Elem()
}

func (i AlmAzureArray) ToAlmAzureArrayOutput() AlmAzureArrayOutput {
	return i.ToAlmAzureArrayOutputWithContext(context.Background())
}

func (i AlmAzureArray) ToAlmAzureArrayOutputWithContext(ctx context.Context) AlmAzureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmAzureArrayOutput)
}

// AlmAzureMapInput is an input type that accepts AlmAzureMap and AlmAzureMapOutput values.
// You can construct a concrete instance of `AlmAzureMapInput` via:
//
//	AlmAzureMap{ "key": AlmAzureArgs{...} }
type AlmAzureMapInput interface {
	pulumi.Input

	ToAlmAzureMapOutput() AlmAzureMapOutput
	ToAlmAzureMapOutputWithContext(context.Context) AlmAzureMapOutput
}

type AlmAzureMap map[string]AlmAzureInput

func (AlmAzureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlmAzure)(nil)).Elem()
}

func (i AlmAzureMap) ToAlmAzureMapOutput() AlmAzureMapOutput {
	return i.ToAlmAzureMapOutputWithContext(context.Background())
}

func (i AlmAzureMap) ToAlmAzureMapOutputWithContext(ctx context.Context) AlmAzureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmAzureMapOutput)
}

type AlmAzureOutput struct{ *pulumi.OutputState }

func (AlmAzureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlmAzure)(nil)).Elem()
}

func (o AlmAzureOutput) ToAlmAzureOutput() AlmAzureOutput {
	return o
}

func (o AlmAzureOutput) ToAlmAzureOutputWithContext(ctx context.Context) AlmAzureOutput {
	return o
}

// Unique key of the Azure Devops instance setting
func (o AlmAzureOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmAzure) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Azure Devops personal access token
func (o AlmAzureOutput) PersonalAccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmAzure) pulumi.StringOutput { return v.PersonalAccessToken }).(pulumi.StringOutput)
}

// Azure API URL
func (o AlmAzureOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmAzure) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type AlmAzureArrayOutput struct{ *pulumi.OutputState }

func (AlmAzureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlmAzure)(nil)).Elem()
}

func (o AlmAzureArrayOutput) ToAlmAzureArrayOutput() AlmAzureArrayOutput {
	return o
}

func (o AlmAzureArrayOutput) ToAlmAzureArrayOutputWithContext(ctx context.Context) AlmAzureArrayOutput {
	return o
}

func (o AlmAzureArrayOutput) Index(i pulumi.IntInput) AlmAzureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlmAzure {
		return vs[0].([]*AlmAzure)[vs[1].(int)]
	}).(AlmAzureOutput)
}

type AlmAzureMapOutput struct{ *pulumi.OutputState }

func (AlmAzureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlmAzure)(nil)).Elem()
}

func (o AlmAzureMapOutput) ToAlmAzureMapOutput() AlmAzureMapOutput {
	return o
}

func (o AlmAzureMapOutput) ToAlmAzureMapOutputWithContext(ctx context.Context) AlmAzureMapOutput {
	return o
}

func (o AlmAzureMapOutput) MapIndex(k pulumi.StringInput) AlmAzureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlmAzure {
		return vs[0].(map[string]*AlmAzure)[vs[1].(string)]
	}).(AlmAzureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlmAzureInput)(nil)).Elem(), &AlmAzure{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlmAzureArrayInput)(nil)).Elem(), AlmAzureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlmAzureMapInput)(nil)).Elem(), AlmAzureMap{})
	pulumi.RegisterOutputType(AlmAzureOutput{})
	pulumi.RegisterOutputType(AlmAzureArrayOutput{})
	pulumi.RegisterOutputType(AlmAzureMapOutput{})
}
