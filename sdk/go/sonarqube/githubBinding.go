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

// ## # GithubBinding
//
// Provides a Sonarqube GitHub binding resource. This can be used to create and manage the binding between a
// GitHub repository and a SonarQube project
//
// ## Example: Create a GitHub binding
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
//			_, err = sonarqube.NewProject(ctx, "main", &sonarqube.ProjectArgs{
//				Project:    pulumi.String("my_project"),
//				Visibility: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sonarqube.NewGithubBinding(ctx, "github-binding", &sonarqube.GithubBindingArgs{
//				AlmSetting: github_alm.Key,
//				Project:    pulumi.String("my_project"),
//				Repository: pulumi.String("myorg/myrepo"),
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
// # Bindings can be imported using their ID
//
// terraform
//
// ```sh
// $ pulumi import sonarqube:index/githubBinding:GithubBinding github-binding project/repository
// ```
type GithubBinding struct {
	pulumi.CustomResourceState

	AlmSetting            pulumi.StringOutput    `pulumi:"almSetting"`
	Monorepo              pulumi.StringPtrOutput `pulumi:"monorepo"`
	Project               pulumi.StringOutput    `pulumi:"project"`
	Repository            pulumi.StringOutput    `pulumi:"repository"`
	SummaryCommentEnabled pulumi.StringPtrOutput `pulumi:"summaryCommentEnabled"`
}

// NewGithubBinding registers a new resource with the given unique name, arguments, and options.
func NewGithubBinding(ctx *pulumi.Context,
	name string, args *GithubBindingArgs, opts ...pulumi.ResourceOption) (*GithubBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlmSetting == nil {
		return nil, errors.New("invalid value for required argument 'AlmSetting'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GithubBinding
	err := ctx.RegisterResource("sonarqube:index/githubBinding:GithubBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGithubBinding gets an existing GithubBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGithubBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GithubBindingState, opts ...pulumi.ResourceOption) (*GithubBinding, error) {
	var resource GithubBinding
	err := ctx.ReadResource("sonarqube:index/githubBinding:GithubBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GithubBinding resources.
type githubBindingState struct {
	AlmSetting            *string `pulumi:"almSetting"`
	Monorepo              *string `pulumi:"monorepo"`
	Project               *string `pulumi:"project"`
	Repository            *string `pulumi:"repository"`
	SummaryCommentEnabled *string `pulumi:"summaryCommentEnabled"`
}

type GithubBindingState struct {
	AlmSetting            pulumi.StringPtrInput
	Monorepo              pulumi.StringPtrInput
	Project               pulumi.StringPtrInput
	Repository            pulumi.StringPtrInput
	SummaryCommentEnabled pulumi.StringPtrInput
}

func (GithubBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*githubBindingState)(nil)).Elem()
}

type githubBindingArgs struct {
	AlmSetting            string  `pulumi:"almSetting"`
	Monorepo              *string `pulumi:"monorepo"`
	Project               string  `pulumi:"project"`
	Repository            string  `pulumi:"repository"`
	SummaryCommentEnabled *string `pulumi:"summaryCommentEnabled"`
}

// The set of arguments for constructing a GithubBinding resource.
type GithubBindingArgs struct {
	AlmSetting            pulumi.StringInput
	Monorepo              pulumi.StringPtrInput
	Project               pulumi.StringInput
	Repository            pulumi.StringInput
	SummaryCommentEnabled pulumi.StringPtrInput
}

func (GithubBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*githubBindingArgs)(nil)).Elem()
}

type GithubBindingInput interface {
	pulumi.Input

	ToGithubBindingOutput() GithubBindingOutput
	ToGithubBindingOutputWithContext(ctx context.Context) GithubBindingOutput
}

func (*GithubBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubBinding)(nil)).Elem()
}

func (i *GithubBinding) ToGithubBindingOutput() GithubBindingOutput {
	return i.ToGithubBindingOutputWithContext(context.Background())
}

func (i *GithubBinding) ToGithubBindingOutputWithContext(ctx context.Context) GithubBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubBindingOutput)
}

// GithubBindingArrayInput is an input type that accepts GithubBindingArray and GithubBindingArrayOutput values.
// You can construct a concrete instance of `GithubBindingArrayInput` via:
//
//	GithubBindingArray{ GithubBindingArgs{...} }
type GithubBindingArrayInput interface {
	pulumi.Input

	ToGithubBindingArrayOutput() GithubBindingArrayOutput
	ToGithubBindingArrayOutputWithContext(context.Context) GithubBindingArrayOutput
}

type GithubBindingArray []GithubBindingInput

func (GithubBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GithubBinding)(nil)).Elem()
}

func (i GithubBindingArray) ToGithubBindingArrayOutput() GithubBindingArrayOutput {
	return i.ToGithubBindingArrayOutputWithContext(context.Background())
}

func (i GithubBindingArray) ToGithubBindingArrayOutputWithContext(ctx context.Context) GithubBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubBindingArrayOutput)
}

// GithubBindingMapInput is an input type that accepts GithubBindingMap and GithubBindingMapOutput values.
// You can construct a concrete instance of `GithubBindingMapInput` via:
//
//	GithubBindingMap{ "key": GithubBindingArgs{...} }
type GithubBindingMapInput interface {
	pulumi.Input

	ToGithubBindingMapOutput() GithubBindingMapOutput
	ToGithubBindingMapOutputWithContext(context.Context) GithubBindingMapOutput
}

type GithubBindingMap map[string]GithubBindingInput

func (GithubBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GithubBinding)(nil)).Elem()
}

func (i GithubBindingMap) ToGithubBindingMapOutput() GithubBindingMapOutput {
	return i.ToGithubBindingMapOutputWithContext(context.Background())
}

func (i GithubBindingMap) ToGithubBindingMapOutputWithContext(ctx context.Context) GithubBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GithubBindingMapOutput)
}

type GithubBindingOutput struct{ *pulumi.OutputState }

func (GithubBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GithubBinding)(nil)).Elem()
}

func (o GithubBindingOutput) ToGithubBindingOutput() GithubBindingOutput {
	return o
}

func (o GithubBindingOutput) ToGithubBindingOutputWithContext(ctx context.Context) GithubBindingOutput {
	return o
}

func (o GithubBindingOutput) AlmSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *GithubBinding) pulumi.StringOutput { return v.AlmSetting }).(pulumi.StringOutput)
}

func (o GithubBindingOutput) Monorepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubBinding) pulumi.StringPtrOutput { return v.Monorepo }).(pulumi.StringPtrOutput)
}

func (o GithubBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GithubBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o GithubBindingOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *GithubBinding) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

func (o GithubBindingOutput) SummaryCommentEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GithubBinding) pulumi.StringPtrOutput { return v.SummaryCommentEnabled }).(pulumi.StringPtrOutput)
}

type GithubBindingArrayOutput struct{ *pulumi.OutputState }

func (GithubBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GithubBinding)(nil)).Elem()
}

func (o GithubBindingArrayOutput) ToGithubBindingArrayOutput() GithubBindingArrayOutput {
	return o
}

func (o GithubBindingArrayOutput) ToGithubBindingArrayOutputWithContext(ctx context.Context) GithubBindingArrayOutput {
	return o
}

func (o GithubBindingArrayOutput) Index(i pulumi.IntInput) GithubBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GithubBinding {
		return vs[0].([]*GithubBinding)[vs[1].(int)]
	}).(GithubBindingOutput)
}

type GithubBindingMapOutput struct{ *pulumi.OutputState }

func (GithubBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GithubBinding)(nil)).Elem()
}

func (o GithubBindingMapOutput) ToGithubBindingMapOutput() GithubBindingMapOutput {
	return o
}

func (o GithubBindingMapOutput) ToGithubBindingMapOutputWithContext(ctx context.Context) GithubBindingMapOutput {
	return o
}

func (o GithubBindingMapOutput) MapIndex(k pulumi.StringInput) GithubBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GithubBinding {
		return vs[0].(map[string]*GithubBinding)[vs[1].(string)]
	}).(GithubBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GithubBindingInput)(nil)).Elem(), &GithubBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*GithubBindingArrayInput)(nil)).Elem(), GithubBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GithubBindingMapInput)(nil)).Elem(), GithubBindingMap{})
	pulumi.RegisterOutputType(GithubBindingOutput{})
	pulumi.RegisterOutputType(GithubBindingArrayOutput{})
	pulumi.RegisterOutputType(GithubBindingMapOutput{})
}