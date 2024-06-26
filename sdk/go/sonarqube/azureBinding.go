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

// ## # AzureBinding
//
// Provides a Sonarqube Azure Devops binding resource. This can be used to create and manage the binding between an
// Azure Devops repository and a SonarQube project
//
// ## Example: Create an Azure Devops binding
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
//			az1, err := sonarqube.NewAlmAzure(ctx, "az1", &sonarqube.AlmAzureArgs{
//				Key:                 pulumi.String("az1"),
//				PersonalAccessToken: pulumi.String("my_pat"),
//				Url:                 pulumi.String("https://dev.azure.com/my-org"),
//			})
//			if err != nil {
//				return err
//			}
//			mainProject, err := sonarqube.NewProject(ctx, "mainProject", &sonarqube.ProjectArgs{
//				Project:    pulumi.String("main"),
//				Visibility: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sonarqube.NewAzureBinding(ctx, "mainAzureBinding", &sonarqube.AzureBindingArgs{
//				AlmSetting:     az1.Key,
//				Project:        mainProject.Project,
//				ProjectName:    pulumi.String("my_azure_project"),
//				RepositoryName: pulumi.String("my_repo"),
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
// $ pulumi import sonarqube:index/azureBinding:AzureBinding main project/project_name/repository
// ```
type AzureBinding struct {
	pulumi.CustomResourceState

	// Azure DevOps setting key
	AlmSetting pulumi.StringOutput `pulumi:"almSetting"`
	// Is this project part of a monorepo
	Monorepo pulumi.BoolPtrOutput `pulumi:"monorepo"`
	// SonarQube project key
	Project pulumi.StringOutput `pulumi:"project"`
	// Azure project name
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Azure repository name
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
}

// NewAzureBinding registers a new resource with the given unique name, arguments, and options.
func NewAzureBinding(ctx *pulumi.Context,
	name string, args *AzureBindingArgs, opts ...pulumi.ResourceOption) (*AzureBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlmSetting == nil {
		return nil, errors.New("invalid value for required argument 'AlmSetting'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.RepositoryName == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureBinding
	err := ctx.RegisterResource("sonarqube:index/azureBinding:AzureBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureBinding gets an existing AzureBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureBindingState, opts ...pulumi.ResourceOption) (*AzureBinding, error) {
	var resource AzureBinding
	err := ctx.ReadResource("sonarqube:index/azureBinding:AzureBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureBinding resources.
type azureBindingState struct {
	// Azure DevOps setting key
	AlmSetting *string `pulumi:"almSetting"`
	// Is this project part of a monorepo
	Monorepo *bool `pulumi:"monorepo"`
	// SonarQube project key
	Project *string `pulumi:"project"`
	// Azure project name
	ProjectName *string `pulumi:"projectName"`
	// Azure repository name
	RepositoryName *string `pulumi:"repositoryName"`
}

type AzureBindingState struct {
	// Azure DevOps setting key
	AlmSetting pulumi.StringPtrInput
	// Is this project part of a monorepo
	Monorepo pulumi.BoolPtrInput
	// SonarQube project key
	Project pulumi.StringPtrInput
	// Azure project name
	ProjectName pulumi.StringPtrInput
	// Azure repository name
	RepositoryName pulumi.StringPtrInput
}

func (AzureBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureBindingState)(nil)).Elem()
}

type azureBindingArgs struct {
	// Azure DevOps setting key
	AlmSetting string `pulumi:"almSetting"`
	// Is this project part of a monorepo
	Monorepo *bool `pulumi:"monorepo"`
	// SonarQube project key
	Project string `pulumi:"project"`
	// Azure project name
	ProjectName string `pulumi:"projectName"`
	// Azure repository name
	RepositoryName string `pulumi:"repositoryName"`
}

// The set of arguments for constructing a AzureBinding resource.
type AzureBindingArgs struct {
	// Azure DevOps setting key
	AlmSetting pulumi.StringInput
	// Is this project part of a monorepo
	Monorepo pulumi.BoolPtrInput
	// SonarQube project key
	Project pulumi.StringInput
	// Azure project name
	ProjectName pulumi.StringInput
	// Azure repository name
	RepositoryName pulumi.StringInput
}

func (AzureBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureBindingArgs)(nil)).Elem()
}

type AzureBindingInput interface {
	pulumi.Input

	ToAzureBindingOutput() AzureBindingOutput
	ToAzureBindingOutputWithContext(ctx context.Context) AzureBindingOutput
}

func (*AzureBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBinding)(nil)).Elem()
}

func (i *AzureBinding) ToAzureBindingOutput() AzureBindingOutput {
	return i.ToAzureBindingOutputWithContext(context.Background())
}

func (i *AzureBinding) ToAzureBindingOutputWithContext(ctx context.Context) AzureBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBindingOutput)
}

// AzureBindingArrayInput is an input type that accepts AzureBindingArray and AzureBindingArrayOutput values.
// You can construct a concrete instance of `AzureBindingArrayInput` via:
//
//	AzureBindingArray{ AzureBindingArgs{...} }
type AzureBindingArrayInput interface {
	pulumi.Input

	ToAzureBindingArrayOutput() AzureBindingArrayOutput
	ToAzureBindingArrayOutputWithContext(context.Context) AzureBindingArrayOutput
}

type AzureBindingArray []AzureBindingInput

func (AzureBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureBinding)(nil)).Elem()
}

func (i AzureBindingArray) ToAzureBindingArrayOutput() AzureBindingArrayOutput {
	return i.ToAzureBindingArrayOutputWithContext(context.Background())
}

func (i AzureBindingArray) ToAzureBindingArrayOutputWithContext(ctx context.Context) AzureBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBindingArrayOutput)
}

// AzureBindingMapInput is an input type that accepts AzureBindingMap and AzureBindingMapOutput values.
// You can construct a concrete instance of `AzureBindingMapInput` via:
//
//	AzureBindingMap{ "key": AzureBindingArgs{...} }
type AzureBindingMapInput interface {
	pulumi.Input

	ToAzureBindingMapOutput() AzureBindingMapOutput
	ToAzureBindingMapOutputWithContext(context.Context) AzureBindingMapOutput
}

type AzureBindingMap map[string]AzureBindingInput

func (AzureBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureBinding)(nil)).Elem()
}

func (i AzureBindingMap) ToAzureBindingMapOutput() AzureBindingMapOutput {
	return i.ToAzureBindingMapOutputWithContext(context.Background())
}

func (i AzureBindingMap) ToAzureBindingMapOutputWithContext(ctx context.Context) AzureBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBindingMapOutput)
}

type AzureBindingOutput struct{ *pulumi.OutputState }

func (AzureBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBinding)(nil)).Elem()
}

func (o AzureBindingOutput) ToAzureBindingOutput() AzureBindingOutput {
	return o
}

func (o AzureBindingOutput) ToAzureBindingOutputWithContext(ctx context.Context) AzureBindingOutput {
	return o
}

// Azure DevOps setting key
func (o AzureBindingOutput) AlmSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureBinding) pulumi.StringOutput { return v.AlmSetting }).(pulumi.StringOutput)
}

// Is this project part of a monorepo
func (o AzureBindingOutput) Monorepo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBinding) pulumi.BoolPtrOutput { return v.Monorepo }).(pulumi.BoolPtrOutput)
}

// SonarQube project key
func (o AzureBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Azure project name
func (o AzureBindingOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureBinding) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// Azure repository name
func (o AzureBindingOutput) RepositoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureBinding) pulumi.StringOutput { return v.RepositoryName }).(pulumi.StringOutput)
}

type AzureBindingArrayOutput struct{ *pulumi.OutputState }

func (AzureBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureBinding)(nil)).Elem()
}

func (o AzureBindingArrayOutput) ToAzureBindingArrayOutput() AzureBindingArrayOutput {
	return o
}

func (o AzureBindingArrayOutput) ToAzureBindingArrayOutputWithContext(ctx context.Context) AzureBindingArrayOutput {
	return o
}

func (o AzureBindingArrayOutput) Index(i pulumi.IntInput) AzureBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureBinding {
		return vs[0].([]*AzureBinding)[vs[1].(int)]
	}).(AzureBindingOutput)
}

type AzureBindingMapOutput struct{ *pulumi.OutputState }

func (AzureBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureBinding)(nil)).Elem()
}

func (o AzureBindingMapOutput) ToAzureBindingMapOutput() AzureBindingMapOutput {
	return o
}

func (o AzureBindingMapOutput) ToAzureBindingMapOutputWithContext(ctx context.Context) AzureBindingMapOutput {
	return o
}

func (o AzureBindingMapOutput) MapIndex(k pulumi.StringInput) AzureBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureBinding {
		return vs[0].(map[string]*AzureBinding)[vs[1].(string)]
	}).(AzureBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureBindingInput)(nil)).Elem(), &AzureBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureBindingArrayInput)(nil)).Elem(), AzureBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureBindingMapInput)(nil)).Elem(), AzureBindingMap{})
	pulumi.RegisterOutputType(AzureBindingOutput{})
	pulumi.RegisterOutputType(AzureBindingArrayOutput{})
	pulumi.RegisterOutputType(AzureBindingMapOutput{})
}
