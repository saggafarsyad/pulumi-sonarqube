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

// ## # AlmGitlab
//
// Provides a Sonarqube GitLab Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops
// Platform Integration for GitLab.
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
//			_, err := sonarqube.NewAlmGitlab(ctx, "gitlab-alm", &sonarqube.AlmGitlabArgs{
//				Key:                 pulumi.String("myalm"),
//				PersonalAccessToken: pulumi.String("my_personal_access_token"),
//				Url:                 pulumi.String("https://gitlab.com/api/v4"),
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
type AlmGitlab struct {
	pulumi.CustomResourceState

	Key                 pulumi.StringOutput `pulumi:"key"`
	PersonalAccessToken pulumi.StringOutput `pulumi:"personalAccessToken"`
	Url                 pulumi.StringOutput `pulumi:"url"`
}

// NewAlmGitlab registers a new resource with the given unique name, arguments, and options.
func NewAlmGitlab(ctx *pulumi.Context,
	name string, args *AlmGitlabArgs, opts ...pulumi.ResourceOption) (*AlmGitlab, error) {
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
	var resource AlmGitlab
	err := ctx.RegisterResource("sonarqube:index/almGitlab:AlmGitlab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlmGitlab gets an existing AlmGitlab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlmGitlab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlmGitlabState, opts ...pulumi.ResourceOption) (*AlmGitlab, error) {
	var resource AlmGitlab
	err := ctx.ReadResource("sonarqube:index/almGitlab:AlmGitlab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlmGitlab resources.
type almGitlabState struct {
	Key                 *string `pulumi:"key"`
	PersonalAccessToken *string `pulumi:"personalAccessToken"`
	Url                 *string `pulumi:"url"`
}

type AlmGitlabState struct {
	Key                 pulumi.StringPtrInput
	PersonalAccessToken pulumi.StringPtrInput
	Url                 pulumi.StringPtrInput
}

func (AlmGitlabState) ElementType() reflect.Type {
	return reflect.TypeOf((*almGitlabState)(nil)).Elem()
}

type almGitlabArgs struct {
	Key                 string `pulumi:"key"`
	PersonalAccessToken string `pulumi:"personalAccessToken"`
	Url                 string `pulumi:"url"`
}

// The set of arguments for constructing a AlmGitlab resource.
type AlmGitlabArgs struct {
	Key                 pulumi.StringInput
	PersonalAccessToken pulumi.StringInput
	Url                 pulumi.StringInput
}

func (AlmGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*almGitlabArgs)(nil)).Elem()
}

type AlmGitlabInput interface {
	pulumi.Input

	ToAlmGitlabOutput() AlmGitlabOutput
	ToAlmGitlabOutputWithContext(ctx context.Context) AlmGitlabOutput
}

func (*AlmGitlab) ElementType() reflect.Type {
	return reflect.TypeOf((**AlmGitlab)(nil)).Elem()
}

func (i *AlmGitlab) ToAlmGitlabOutput() AlmGitlabOutput {
	return i.ToAlmGitlabOutputWithContext(context.Background())
}

func (i *AlmGitlab) ToAlmGitlabOutputWithContext(ctx context.Context) AlmGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmGitlabOutput)
}

// AlmGitlabArrayInput is an input type that accepts AlmGitlabArray and AlmGitlabArrayOutput values.
// You can construct a concrete instance of `AlmGitlabArrayInput` via:
//
//	AlmGitlabArray{ AlmGitlabArgs{...} }
type AlmGitlabArrayInput interface {
	pulumi.Input

	ToAlmGitlabArrayOutput() AlmGitlabArrayOutput
	ToAlmGitlabArrayOutputWithContext(context.Context) AlmGitlabArrayOutput
}

type AlmGitlabArray []AlmGitlabInput

func (AlmGitlabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlmGitlab)(nil)).Elem()
}

func (i AlmGitlabArray) ToAlmGitlabArrayOutput() AlmGitlabArrayOutput {
	return i.ToAlmGitlabArrayOutputWithContext(context.Background())
}

func (i AlmGitlabArray) ToAlmGitlabArrayOutputWithContext(ctx context.Context) AlmGitlabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmGitlabArrayOutput)
}

// AlmGitlabMapInput is an input type that accepts AlmGitlabMap and AlmGitlabMapOutput values.
// You can construct a concrete instance of `AlmGitlabMapInput` via:
//
//	AlmGitlabMap{ "key": AlmGitlabArgs{...} }
type AlmGitlabMapInput interface {
	pulumi.Input

	ToAlmGitlabMapOutput() AlmGitlabMapOutput
	ToAlmGitlabMapOutputWithContext(context.Context) AlmGitlabMapOutput
}

type AlmGitlabMap map[string]AlmGitlabInput

func (AlmGitlabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlmGitlab)(nil)).Elem()
}

func (i AlmGitlabMap) ToAlmGitlabMapOutput() AlmGitlabMapOutput {
	return i.ToAlmGitlabMapOutputWithContext(context.Background())
}

func (i AlmGitlabMap) ToAlmGitlabMapOutputWithContext(ctx context.Context) AlmGitlabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlmGitlabMapOutput)
}

type AlmGitlabOutput struct{ *pulumi.OutputState }

func (AlmGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlmGitlab)(nil)).Elem()
}

func (o AlmGitlabOutput) ToAlmGitlabOutput() AlmGitlabOutput {
	return o
}

func (o AlmGitlabOutput) ToAlmGitlabOutputWithContext(ctx context.Context) AlmGitlabOutput {
	return o
}

func (o AlmGitlabOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGitlab) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o AlmGitlabOutput) PersonalAccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGitlab) pulumi.StringOutput { return v.PersonalAccessToken }).(pulumi.StringOutput)
}

func (o AlmGitlabOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *AlmGitlab) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type AlmGitlabArrayOutput struct{ *pulumi.OutputState }

func (AlmGitlabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlmGitlab)(nil)).Elem()
}

func (o AlmGitlabArrayOutput) ToAlmGitlabArrayOutput() AlmGitlabArrayOutput {
	return o
}

func (o AlmGitlabArrayOutput) ToAlmGitlabArrayOutputWithContext(ctx context.Context) AlmGitlabArrayOutput {
	return o
}

func (o AlmGitlabArrayOutput) Index(i pulumi.IntInput) AlmGitlabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlmGitlab {
		return vs[0].([]*AlmGitlab)[vs[1].(int)]
	}).(AlmGitlabOutput)
}

type AlmGitlabMapOutput struct{ *pulumi.OutputState }

func (AlmGitlabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlmGitlab)(nil)).Elem()
}

func (o AlmGitlabMapOutput) ToAlmGitlabMapOutput() AlmGitlabMapOutput {
	return o
}

func (o AlmGitlabMapOutput) ToAlmGitlabMapOutputWithContext(ctx context.Context) AlmGitlabMapOutput {
	return o
}

func (o AlmGitlabMapOutput) MapIndex(k pulumi.StringInput) AlmGitlabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlmGitlab {
		return vs[0].(map[string]*AlmGitlab)[vs[1].(string)]
	}).(AlmGitlabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlmGitlabInput)(nil)).Elem(), &AlmGitlab{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlmGitlabArrayInput)(nil)).Elem(), AlmGitlabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlmGitlabMapInput)(nil)).Elem(), AlmGitlabMap{})
	pulumi.RegisterOutputType(AlmGitlabOutput{})
	pulumi.RegisterOutputType(AlmGitlabArrayOutput{})
	pulumi.RegisterOutputType(AlmGitlabMapOutput{})
}