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

// ## # QualityprofileProjectAssociation
//
// Provides a Sonarqube Quality Profile Project association resource. This can be used to associate a Quality Profile to a Project
//
// ## Example: create a quality profile project association
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
//			mainQualityprofile, err := sonarqube.NewQualityprofile(ctx, "mainQualityprofile", &sonarqube.QualityprofileArgs{
//				Language: pulumi.String("js"),
//			})
//			if err != nil {
//				return err
//			}
//			mainProject, err := sonarqube.NewProject(ctx, "mainProject", &sonarqube.ProjectArgs{
//				Project:    pulumi.String("my_project"),
//				Visibility: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sonarqube.NewQualityprofileProjectAssociation(ctx, "mainQualityprofileProjectAssociation", &sonarqube.QualityprofileProjectAssociationArgs{
//				QualityProfile: mainQualityprofile.Name,
//				Project:        mainProject.Project,
//				Language:       pulumi.String("js"),
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
type QualityprofileProjectAssociation struct {
	pulumi.CustomResourceState

	// Quality profile language
	Language pulumi.StringOutput `pulumi:"language"`
	// Project key
	Project pulumi.StringOutput `pulumi:"project"`
	// Quality profile name
	QualityProfile pulumi.StringOutput `pulumi:"qualityProfile"`
}

// NewQualityprofileProjectAssociation registers a new resource with the given unique name, arguments, and options.
func NewQualityprofileProjectAssociation(ctx *pulumi.Context,
	name string, args *QualityprofileProjectAssociationArgs, opts ...pulumi.ResourceOption) (*QualityprofileProjectAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Language == nil {
		return nil, errors.New("invalid value for required argument 'Language'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.QualityProfile == nil {
		return nil, errors.New("invalid value for required argument 'QualityProfile'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QualityprofileProjectAssociation
	err := ctx.RegisterResource("sonarqube:index/qualityprofileProjectAssociation:QualityprofileProjectAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQualityprofileProjectAssociation gets an existing QualityprofileProjectAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQualityprofileProjectAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QualityprofileProjectAssociationState, opts ...pulumi.ResourceOption) (*QualityprofileProjectAssociation, error) {
	var resource QualityprofileProjectAssociation
	err := ctx.ReadResource("sonarqube:index/qualityprofileProjectAssociation:QualityprofileProjectAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QualityprofileProjectAssociation resources.
type qualityprofileProjectAssociationState struct {
	// Quality profile language
	Language *string `pulumi:"language"`
	// Project key
	Project *string `pulumi:"project"`
	// Quality profile name
	QualityProfile *string `pulumi:"qualityProfile"`
}

type QualityprofileProjectAssociationState struct {
	// Quality profile language
	Language pulumi.StringPtrInput
	// Project key
	Project pulumi.StringPtrInput
	// Quality profile name
	QualityProfile pulumi.StringPtrInput
}

func (QualityprofileProjectAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*qualityprofileProjectAssociationState)(nil)).Elem()
}

type qualityprofileProjectAssociationArgs struct {
	// Quality profile language
	Language string `pulumi:"language"`
	// Project key
	Project string `pulumi:"project"`
	// Quality profile name
	QualityProfile string `pulumi:"qualityProfile"`
}

// The set of arguments for constructing a QualityprofileProjectAssociation resource.
type QualityprofileProjectAssociationArgs struct {
	// Quality profile language
	Language pulumi.StringInput
	// Project key
	Project pulumi.StringInput
	// Quality profile name
	QualityProfile pulumi.StringInput
}

func (QualityprofileProjectAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qualityprofileProjectAssociationArgs)(nil)).Elem()
}

type QualityprofileProjectAssociationInput interface {
	pulumi.Input

	ToQualityprofileProjectAssociationOutput() QualityprofileProjectAssociationOutput
	ToQualityprofileProjectAssociationOutputWithContext(ctx context.Context) QualityprofileProjectAssociationOutput
}

func (*QualityprofileProjectAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**QualityprofileProjectAssociation)(nil)).Elem()
}

func (i *QualityprofileProjectAssociation) ToQualityprofileProjectAssociationOutput() QualityprofileProjectAssociationOutput {
	return i.ToQualityprofileProjectAssociationOutputWithContext(context.Background())
}

func (i *QualityprofileProjectAssociation) ToQualityprofileProjectAssociationOutputWithContext(ctx context.Context) QualityprofileProjectAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualityprofileProjectAssociationOutput)
}

// QualityprofileProjectAssociationArrayInput is an input type that accepts QualityprofileProjectAssociationArray and QualityprofileProjectAssociationArrayOutput values.
// You can construct a concrete instance of `QualityprofileProjectAssociationArrayInput` via:
//
//	QualityprofileProjectAssociationArray{ QualityprofileProjectAssociationArgs{...} }
type QualityprofileProjectAssociationArrayInput interface {
	pulumi.Input

	ToQualityprofileProjectAssociationArrayOutput() QualityprofileProjectAssociationArrayOutput
	ToQualityprofileProjectAssociationArrayOutputWithContext(context.Context) QualityprofileProjectAssociationArrayOutput
}

type QualityprofileProjectAssociationArray []QualityprofileProjectAssociationInput

func (QualityprofileProjectAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QualityprofileProjectAssociation)(nil)).Elem()
}

func (i QualityprofileProjectAssociationArray) ToQualityprofileProjectAssociationArrayOutput() QualityprofileProjectAssociationArrayOutput {
	return i.ToQualityprofileProjectAssociationArrayOutputWithContext(context.Background())
}

func (i QualityprofileProjectAssociationArray) ToQualityprofileProjectAssociationArrayOutputWithContext(ctx context.Context) QualityprofileProjectAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualityprofileProjectAssociationArrayOutput)
}

// QualityprofileProjectAssociationMapInput is an input type that accepts QualityprofileProjectAssociationMap and QualityprofileProjectAssociationMapOutput values.
// You can construct a concrete instance of `QualityprofileProjectAssociationMapInput` via:
//
//	QualityprofileProjectAssociationMap{ "key": QualityprofileProjectAssociationArgs{...} }
type QualityprofileProjectAssociationMapInput interface {
	pulumi.Input

	ToQualityprofileProjectAssociationMapOutput() QualityprofileProjectAssociationMapOutput
	ToQualityprofileProjectAssociationMapOutputWithContext(context.Context) QualityprofileProjectAssociationMapOutput
}

type QualityprofileProjectAssociationMap map[string]QualityprofileProjectAssociationInput

func (QualityprofileProjectAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QualityprofileProjectAssociation)(nil)).Elem()
}

func (i QualityprofileProjectAssociationMap) ToQualityprofileProjectAssociationMapOutput() QualityprofileProjectAssociationMapOutput {
	return i.ToQualityprofileProjectAssociationMapOutputWithContext(context.Background())
}

func (i QualityprofileProjectAssociationMap) ToQualityprofileProjectAssociationMapOutputWithContext(ctx context.Context) QualityprofileProjectAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualityprofileProjectAssociationMapOutput)
}

type QualityprofileProjectAssociationOutput struct{ *pulumi.OutputState }

func (QualityprofileProjectAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QualityprofileProjectAssociation)(nil)).Elem()
}

func (o QualityprofileProjectAssociationOutput) ToQualityprofileProjectAssociationOutput() QualityprofileProjectAssociationOutput {
	return o
}

func (o QualityprofileProjectAssociationOutput) ToQualityprofileProjectAssociationOutputWithContext(ctx context.Context) QualityprofileProjectAssociationOutput {
	return o
}

// Quality profile language
func (o QualityprofileProjectAssociationOutput) Language() pulumi.StringOutput {
	return o.ApplyT(func(v *QualityprofileProjectAssociation) pulumi.StringOutput { return v.Language }).(pulumi.StringOutput)
}

// Project key
func (o QualityprofileProjectAssociationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *QualityprofileProjectAssociation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Quality profile name
func (o QualityprofileProjectAssociationOutput) QualityProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *QualityprofileProjectAssociation) pulumi.StringOutput { return v.QualityProfile }).(pulumi.StringOutput)
}

type QualityprofileProjectAssociationArrayOutput struct{ *pulumi.OutputState }

func (QualityprofileProjectAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QualityprofileProjectAssociation)(nil)).Elem()
}

func (o QualityprofileProjectAssociationArrayOutput) ToQualityprofileProjectAssociationArrayOutput() QualityprofileProjectAssociationArrayOutput {
	return o
}

func (o QualityprofileProjectAssociationArrayOutput) ToQualityprofileProjectAssociationArrayOutputWithContext(ctx context.Context) QualityprofileProjectAssociationArrayOutput {
	return o
}

func (o QualityprofileProjectAssociationArrayOutput) Index(i pulumi.IntInput) QualityprofileProjectAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QualityprofileProjectAssociation {
		return vs[0].([]*QualityprofileProjectAssociation)[vs[1].(int)]
	}).(QualityprofileProjectAssociationOutput)
}

type QualityprofileProjectAssociationMapOutput struct{ *pulumi.OutputState }

func (QualityprofileProjectAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QualityprofileProjectAssociation)(nil)).Elem()
}

func (o QualityprofileProjectAssociationMapOutput) ToQualityprofileProjectAssociationMapOutput() QualityprofileProjectAssociationMapOutput {
	return o
}

func (o QualityprofileProjectAssociationMapOutput) ToQualityprofileProjectAssociationMapOutputWithContext(ctx context.Context) QualityprofileProjectAssociationMapOutput {
	return o
}

func (o QualityprofileProjectAssociationMapOutput) MapIndex(k pulumi.StringInput) QualityprofileProjectAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QualityprofileProjectAssociation {
		return vs[0].(map[string]*QualityprofileProjectAssociation)[vs[1].(string)]
	}).(QualityprofileProjectAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QualityprofileProjectAssociationInput)(nil)).Elem(), &QualityprofileProjectAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*QualityprofileProjectAssociationArrayInput)(nil)).Elem(), QualityprofileProjectAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QualityprofileProjectAssociationMapInput)(nil)).Elem(), QualityprofileProjectAssociationMap{})
	pulumi.RegisterOutputType(QualityprofileProjectAssociationOutput{})
	pulumi.RegisterOutputType(QualityprofileProjectAssociationArrayOutput{})
	pulumi.RegisterOutputType(QualityprofileProjectAssociationMapOutput{})
}
