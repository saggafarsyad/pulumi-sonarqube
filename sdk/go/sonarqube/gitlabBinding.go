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

// ## # GitlabBinding
//
// Provides a Sonarqube GitLab binding resource. This can be used to create and manage the binding between a
// GitLab repository and a SonarQube project
//
// ## Example: Create a GitLab binding
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
//			_, err = sonarqube.NewProject(ctx, "main", &sonarqube.ProjectArgs{
//				Project:    pulumi.String("my_project"),
//				Visibility: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sonarqube.NewGitlabBinding(ctx, "gitlab-binding", &sonarqube.GitlabBindingArgs{
//				AlmSetting: gitlab_alm.Key,
//				Project:    pulumi.String("my_project"),
//				Repository: pulumi.String("123"),
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
// $ pulumi import sonarqube:index/gitlabBinding:GitlabBinding gitlab-binding project/repository
// ```
type GitlabBinding struct {
	pulumi.CustomResourceState

	AlmSetting pulumi.StringOutput    `pulumi:"almSetting"`
	Monorepo   pulumi.StringPtrOutput `pulumi:"monorepo"`
	Project    pulumi.StringOutput    `pulumi:"project"`
	Repository pulumi.StringOutput    `pulumi:"repository"`
}

// NewGitlabBinding registers a new resource with the given unique name, arguments, and options.
func NewGitlabBinding(ctx *pulumi.Context,
	name string, args *GitlabBindingArgs, opts ...pulumi.ResourceOption) (*GitlabBinding, error) {
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
	var resource GitlabBinding
	err := ctx.RegisterResource("sonarqube:index/gitlabBinding:GitlabBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGitlabBinding gets an existing GitlabBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGitlabBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitlabBindingState, opts ...pulumi.ResourceOption) (*GitlabBinding, error) {
	var resource GitlabBinding
	err := ctx.ReadResource("sonarqube:index/gitlabBinding:GitlabBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GitlabBinding resources.
type gitlabBindingState struct {
	AlmSetting *string `pulumi:"almSetting"`
	Monorepo   *string `pulumi:"monorepo"`
	Project    *string `pulumi:"project"`
	Repository *string `pulumi:"repository"`
}

type GitlabBindingState struct {
	AlmSetting pulumi.StringPtrInput
	Monorepo   pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	Repository pulumi.StringPtrInput
}

func (GitlabBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitlabBindingState)(nil)).Elem()
}

type gitlabBindingArgs struct {
	AlmSetting string  `pulumi:"almSetting"`
	Monorepo   *string `pulumi:"monorepo"`
	Project    string  `pulumi:"project"`
	Repository string  `pulumi:"repository"`
}

// The set of arguments for constructing a GitlabBinding resource.
type GitlabBindingArgs struct {
	AlmSetting pulumi.StringInput
	Monorepo   pulumi.StringPtrInput
	Project    pulumi.StringInput
	Repository pulumi.StringInput
}

func (GitlabBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitlabBindingArgs)(nil)).Elem()
}

type GitlabBindingInput interface {
	pulumi.Input

	ToGitlabBindingOutput() GitlabBindingOutput
	ToGitlabBindingOutputWithContext(ctx context.Context) GitlabBindingOutput
}

func (*GitlabBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GitlabBinding)(nil)).Elem()
}

func (i *GitlabBinding) ToGitlabBindingOutput() GitlabBindingOutput {
	return i.ToGitlabBindingOutputWithContext(context.Background())
}

func (i *GitlabBinding) ToGitlabBindingOutputWithContext(ctx context.Context) GitlabBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitlabBindingOutput)
}

// GitlabBindingArrayInput is an input type that accepts GitlabBindingArray and GitlabBindingArrayOutput values.
// You can construct a concrete instance of `GitlabBindingArrayInput` via:
//
//	GitlabBindingArray{ GitlabBindingArgs{...} }
type GitlabBindingArrayInput interface {
	pulumi.Input

	ToGitlabBindingArrayOutput() GitlabBindingArrayOutput
	ToGitlabBindingArrayOutputWithContext(context.Context) GitlabBindingArrayOutput
}

type GitlabBindingArray []GitlabBindingInput

func (GitlabBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitlabBinding)(nil)).Elem()
}

func (i GitlabBindingArray) ToGitlabBindingArrayOutput() GitlabBindingArrayOutput {
	return i.ToGitlabBindingArrayOutputWithContext(context.Background())
}

func (i GitlabBindingArray) ToGitlabBindingArrayOutputWithContext(ctx context.Context) GitlabBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitlabBindingArrayOutput)
}

// GitlabBindingMapInput is an input type that accepts GitlabBindingMap and GitlabBindingMapOutput values.
// You can construct a concrete instance of `GitlabBindingMapInput` via:
//
//	GitlabBindingMap{ "key": GitlabBindingArgs{...} }
type GitlabBindingMapInput interface {
	pulumi.Input

	ToGitlabBindingMapOutput() GitlabBindingMapOutput
	ToGitlabBindingMapOutputWithContext(context.Context) GitlabBindingMapOutput
}

type GitlabBindingMap map[string]GitlabBindingInput

func (GitlabBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitlabBinding)(nil)).Elem()
}

func (i GitlabBindingMap) ToGitlabBindingMapOutput() GitlabBindingMapOutput {
	return i.ToGitlabBindingMapOutputWithContext(context.Background())
}

func (i GitlabBindingMap) ToGitlabBindingMapOutputWithContext(ctx context.Context) GitlabBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitlabBindingMapOutput)
}

type GitlabBindingOutput struct{ *pulumi.OutputState }

func (GitlabBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitlabBinding)(nil)).Elem()
}

func (o GitlabBindingOutput) ToGitlabBindingOutput() GitlabBindingOutput {
	return o
}

func (o GitlabBindingOutput) ToGitlabBindingOutputWithContext(ctx context.Context) GitlabBindingOutput {
	return o
}

func (o GitlabBindingOutput) AlmSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *GitlabBinding) pulumi.StringOutput { return v.AlmSetting }).(pulumi.StringOutput)
}

func (o GitlabBindingOutput) Monorepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitlabBinding) pulumi.StringPtrOutput { return v.Monorepo }).(pulumi.StringPtrOutput)
}

func (o GitlabBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GitlabBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o GitlabBindingOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *GitlabBinding) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

type GitlabBindingArrayOutput struct{ *pulumi.OutputState }

func (GitlabBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GitlabBinding)(nil)).Elem()
}

func (o GitlabBindingArrayOutput) ToGitlabBindingArrayOutput() GitlabBindingArrayOutput {
	return o
}

func (o GitlabBindingArrayOutput) ToGitlabBindingArrayOutputWithContext(ctx context.Context) GitlabBindingArrayOutput {
	return o
}

func (o GitlabBindingArrayOutput) Index(i pulumi.IntInput) GitlabBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GitlabBinding {
		return vs[0].([]*GitlabBinding)[vs[1].(int)]
	}).(GitlabBindingOutput)
}

type GitlabBindingMapOutput struct{ *pulumi.OutputState }

func (GitlabBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GitlabBinding)(nil)).Elem()
}

func (o GitlabBindingMapOutput) ToGitlabBindingMapOutput() GitlabBindingMapOutput {
	return o
}

func (o GitlabBindingMapOutput) ToGitlabBindingMapOutputWithContext(ctx context.Context) GitlabBindingMapOutput {
	return o
}

func (o GitlabBindingMapOutput) MapIndex(k pulumi.StringInput) GitlabBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GitlabBinding {
		return vs[0].(map[string]*GitlabBinding)[vs[1].(string)]
	}).(GitlabBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GitlabBindingInput)(nil)).Elem(), &GitlabBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitlabBindingArrayInput)(nil)).Elem(), GitlabBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitlabBindingMapInput)(nil)).Elem(), GitlabBindingMap{})
	pulumi.RegisterOutputType(GitlabBindingOutput{})
	pulumi.RegisterOutputType(GitlabBindingArrayOutput{})
	pulumi.RegisterOutputType(GitlabBindingMapOutput{})
}
