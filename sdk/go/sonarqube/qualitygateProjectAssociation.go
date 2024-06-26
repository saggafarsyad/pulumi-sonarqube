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

// ## # QualitygateProjectAssociation
//
// Provides a Sonarqube Quality Gate Project association resource. This can be used to associate a Quality Gate to a Project
//
// ## Example: create a quality gate project association
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
//			mainQualitygate, err := sonarqube.NewQualitygate(ctx, "mainQualitygate", &sonarqube.QualitygateArgs{
//				Conditions: sonarqube.QualitygateConditionArray{
//					&sonarqube.QualitygateConditionArgs{
//						Metric:    pulumi.String("new_coverage"),
//						Op:        pulumi.String("LT"),
//						Threshold: pulumi.String("30"),
//					},
//				},
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
//			_, err = sonarqube.NewQualitygateProjectAssociation(ctx, "mainQualitygateProjectAssociation", &sonarqube.QualitygateProjectAssociationArgs{
//				Gatename:   mainQualitygate.ID(),
//				Projectkey: mainProject.Project,
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
type QualitygateProjectAssociation struct {
	pulumi.CustomResourceState

	Gateid     pulumi.StringPtrOutput `pulumi:"gateid"`
	Gatename   pulumi.StringPtrOutput `pulumi:"gatename"`
	Projectkey pulumi.StringOutput    `pulumi:"projectkey"`
}

// NewQualitygateProjectAssociation registers a new resource with the given unique name, arguments, and options.
func NewQualitygateProjectAssociation(ctx *pulumi.Context,
	name string, args *QualitygateProjectAssociationArgs, opts ...pulumi.ResourceOption) (*QualitygateProjectAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Projectkey == nil {
		return nil, errors.New("invalid value for required argument 'Projectkey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QualitygateProjectAssociation
	err := ctx.RegisterResource("sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQualitygateProjectAssociation gets an existing QualitygateProjectAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQualitygateProjectAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QualitygateProjectAssociationState, opts ...pulumi.ResourceOption) (*QualitygateProjectAssociation, error) {
	var resource QualitygateProjectAssociation
	err := ctx.ReadResource("sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QualitygateProjectAssociation resources.
type qualitygateProjectAssociationState struct {
	Gateid     *string `pulumi:"gateid"`
	Gatename   *string `pulumi:"gatename"`
	Projectkey *string `pulumi:"projectkey"`
}

type QualitygateProjectAssociationState struct {
	Gateid     pulumi.StringPtrInput
	Gatename   pulumi.StringPtrInput
	Projectkey pulumi.StringPtrInput
}

func (QualitygateProjectAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*qualitygateProjectAssociationState)(nil)).Elem()
}

type qualitygateProjectAssociationArgs struct {
	Gateid     *string `pulumi:"gateid"`
	Gatename   *string `pulumi:"gatename"`
	Projectkey string  `pulumi:"projectkey"`
}

// The set of arguments for constructing a QualitygateProjectAssociation resource.
type QualitygateProjectAssociationArgs struct {
	Gateid     pulumi.StringPtrInput
	Gatename   pulumi.StringPtrInput
	Projectkey pulumi.StringInput
}

func (QualitygateProjectAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qualitygateProjectAssociationArgs)(nil)).Elem()
}

type QualitygateProjectAssociationInput interface {
	pulumi.Input

	ToQualitygateProjectAssociationOutput() QualitygateProjectAssociationOutput
	ToQualitygateProjectAssociationOutputWithContext(ctx context.Context) QualitygateProjectAssociationOutput
}

func (*QualitygateProjectAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**QualitygateProjectAssociation)(nil)).Elem()
}

func (i *QualitygateProjectAssociation) ToQualitygateProjectAssociationOutput() QualitygateProjectAssociationOutput {
	return i.ToQualitygateProjectAssociationOutputWithContext(context.Background())
}

func (i *QualitygateProjectAssociation) ToQualitygateProjectAssociationOutputWithContext(ctx context.Context) QualitygateProjectAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualitygateProjectAssociationOutput)
}

// QualitygateProjectAssociationArrayInput is an input type that accepts QualitygateProjectAssociationArray and QualitygateProjectAssociationArrayOutput values.
// You can construct a concrete instance of `QualitygateProjectAssociationArrayInput` via:
//
//	QualitygateProjectAssociationArray{ QualitygateProjectAssociationArgs{...} }
type QualitygateProjectAssociationArrayInput interface {
	pulumi.Input

	ToQualitygateProjectAssociationArrayOutput() QualitygateProjectAssociationArrayOutput
	ToQualitygateProjectAssociationArrayOutputWithContext(context.Context) QualitygateProjectAssociationArrayOutput
}

type QualitygateProjectAssociationArray []QualitygateProjectAssociationInput

func (QualitygateProjectAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QualitygateProjectAssociation)(nil)).Elem()
}

func (i QualitygateProjectAssociationArray) ToQualitygateProjectAssociationArrayOutput() QualitygateProjectAssociationArrayOutput {
	return i.ToQualitygateProjectAssociationArrayOutputWithContext(context.Background())
}

func (i QualitygateProjectAssociationArray) ToQualitygateProjectAssociationArrayOutputWithContext(ctx context.Context) QualitygateProjectAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualitygateProjectAssociationArrayOutput)
}

// QualitygateProjectAssociationMapInput is an input type that accepts QualitygateProjectAssociationMap and QualitygateProjectAssociationMapOutput values.
// You can construct a concrete instance of `QualitygateProjectAssociationMapInput` via:
//
//	QualitygateProjectAssociationMap{ "key": QualitygateProjectAssociationArgs{...} }
type QualitygateProjectAssociationMapInput interface {
	pulumi.Input

	ToQualitygateProjectAssociationMapOutput() QualitygateProjectAssociationMapOutput
	ToQualitygateProjectAssociationMapOutputWithContext(context.Context) QualitygateProjectAssociationMapOutput
}

type QualitygateProjectAssociationMap map[string]QualitygateProjectAssociationInput

func (QualitygateProjectAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QualitygateProjectAssociation)(nil)).Elem()
}

func (i QualitygateProjectAssociationMap) ToQualitygateProjectAssociationMapOutput() QualitygateProjectAssociationMapOutput {
	return i.ToQualitygateProjectAssociationMapOutputWithContext(context.Background())
}

func (i QualitygateProjectAssociationMap) ToQualitygateProjectAssociationMapOutputWithContext(ctx context.Context) QualitygateProjectAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualitygateProjectAssociationMapOutput)
}

type QualitygateProjectAssociationOutput struct{ *pulumi.OutputState }

func (QualitygateProjectAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QualitygateProjectAssociation)(nil)).Elem()
}

func (o QualitygateProjectAssociationOutput) ToQualitygateProjectAssociationOutput() QualitygateProjectAssociationOutput {
	return o
}

func (o QualitygateProjectAssociationOutput) ToQualitygateProjectAssociationOutputWithContext(ctx context.Context) QualitygateProjectAssociationOutput {
	return o
}

func (o QualitygateProjectAssociationOutput) Gateid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QualitygateProjectAssociation) pulumi.StringPtrOutput { return v.Gateid }).(pulumi.StringPtrOutput)
}

func (o QualitygateProjectAssociationOutput) Gatename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QualitygateProjectAssociation) pulumi.StringPtrOutput { return v.Gatename }).(pulumi.StringPtrOutput)
}

func (o QualitygateProjectAssociationOutput) Projectkey() pulumi.StringOutput {
	return o.ApplyT(func(v *QualitygateProjectAssociation) pulumi.StringOutput { return v.Projectkey }).(pulumi.StringOutput)
}

type QualitygateProjectAssociationArrayOutput struct{ *pulumi.OutputState }

func (QualitygateProjectAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QualitygateProjectAssociation)(nil)).Elem()
}

func (o QualitygateProjectAssociationArrayOutput) ToQualitygateProjectAssociationArrayOutput() QualitygateProjectAssociationArrayOutput {
	return o
}

func (o QualitygateProjectAssociationArrayOutput) ToQualitygateProjectAssociationArrayOutputWithContext(ctx context.Context) QualitygateProjectAssociationArrayOutput {
	return o
}

func (o QualitygateProjectAssociationArrayOutput) Index(i pulumi.IntInput) QualitygateProjectAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QualitygateProjectAssociation {
		return vs[0].([]*QualitygateProjectAssociation)[vs[1].(int)]
	}).(QualitygateProjectAssociationOutput)
}

type QualitygateProjectAssociationMapOutput struct{ *pulumi.OutputState }

func (QualitygateProjectAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QualitygateProjectAssociation)(nil)).Elem()
}

func (o QualitygateProjectAssociationMapOutput) ToQualitygateProjectAssociationMapOutput() QualitygateProjectAssociationMapOutput {
	return o
}

func (o QualitygateProjectAssociationMapOutput) ToQualitygateProjectAssociationMapOutputWithContext(ctx context.Context) QualitygateProjectAssociationMapOutput {
	return o
}

func (o QualitygateProjectAssociationMapOutput) MapIndex(k pulumi.StringInput) QualitygateProjectAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QualitygateProjectAssociation {
		return vs[0].(map[string]*QualitygateProjectAssociation)[vs[1].(string)]
	}).(QualitygateProjectAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QualitygateProjectAssociationInput)(nil)).Elem(), &QualitygateProjectAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*QualitygateProjectAssociationArrayInput)(nil)).Elem(), QualitygateProjectAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QualitygateProjectAssociationMapInput)(nil)).Elem(), QualitygateProjectAssociationMap{})
	pulumi.RegisterOutputType(QualitygateProjectAssociationOutput{})
	pulumi.RegisterOutputType(QualitygateProjectAssociationArrayOutput{})
	pulumi.RegisterOutputType(QualitygateProjectAssociationMapOutput{})
}
