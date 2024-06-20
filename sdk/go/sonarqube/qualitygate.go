// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sonarqube

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/saggafarsyad/pulumi-sonarqube/sdk/go/sonarqube/internal"
)

// ## # Qualitygate
//
// Provides a Sonarqube Quality Gate resource. This can be used to create and manage Sonarqube Quality Gates and their Conditions.
//
// ## Example: create a quality gate
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
//			_, err := sonarqube.NewQualitygate(ctx, "main", &sonarqube.QualitygateArgs{
//				Conditions: sonarqube.QualitygateConditionArray{
//					&sonarqube.QualitygateConditionArgs{
//						Metric:    pulumi.String("new_coverage"),
//						Op:        pulumi.String("LT"),
//						Threshold: pulumi.String("50"),
//					},
//					&sonarqube.QualitygateConditionArgs{
//						Metric:    pulumi.String("vulnerabilities"),
//						Op:        pulumi.String("GT"),
//						Threshold: pulumi.String("10"),
//					},
//				},
//				IsDefault: pulumi.Bool(true),
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
// # Quality Gates can be imported using its name
//
// terraform
//
// ```sh
// $ pulumi import sonarqube:index/qualitygate:Qualitygate main my-cool-gate
// ```
type Qualitygate struct {
	pulumi.CustomResourceState

	// A list of conditions that the gate uses
	Conditions QualitygateConditionArrayOutput `pulumi:"conditions"`
	CopyFrom   pulumi.StringPtrOutput          `pulumi:"copyFrom"`
	// When set to true this Quality Gate is set as default
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	Name      pulumi.StringOutput  `pulumi:"name"`
}

// NewQualitygate registers a new resource with the given unique name, arguments, and options.
func NewQualitygate(ctx *pulumi.Context,
	name string, args *QualitygateArgs, opts ...pulumi.ResourceOption) (*Qualitygate, error) {
	if args == nil {
		args = &QualitygateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Qualitygate
	err := ctx.RegisterResource("sonarqube:index/qualitygate:Qualitygate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQualitygate gets an existing Qualitygate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQualitygate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QualitygateState, opts ...pulumi.ResourceOption) (*Qualitygate, error) {
	var resource Qualitygate
	err := ctx.ReadResource("sonarqube:index/qualitygate:Qualitygate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Qualitygate resources.
type qualitygateState struct {
	// A list of conditions that the gate uses
	Conditions []QualitygateCondition `pulumi:"conditions"`
	CopyFrom   *string                `pulumi:"copyFrom"`
	// When set to true this Quality Gate is set as default
	IsDefault *bool   `pulumi:"isDefault"`
	Name      *string `pulumi:"name"`
}

type QualitygateState struct {
	// A list of conditions that the gate uses
	Conditions QualitygateConditionArrayInput
	CopyFrom   pulumi.StringPtrInput
	// When set to true this Quality Gate is set as default
	IsDefault pulumi.BoolPtrInput
	Name      pulumi.StringPtrInput
}

func (QualitygateState) ElementType() reflect.Type {
	return reflect.TypeOf((*qualitygateState)(nil)).Elem()
}

type qualitygateArgs struct {
	// A list of conditions that the gate uses
	Conditions []QualitygateCondition `pulumi:"conditions"`
	CopyFrom   *string                `pulumi:"copyFrom"`
	// When set to true this Quality Gate is set as default
	IsDefault *bool   `pulumi:"isDefault"`
	Name      *string `pulumi:"name"`
}

// The set of arguments for constructing a Qualitygate resource.
type QualitygateArgs struct {
	// A list of conditions that the gate uses
	Conditions QualitygateConditionArrayInput
	CopyFrom   pulumi.StringPtrInput
	// When set to true this Quality Gate is set as default
	IsDefault pulumi.BoolPtrInput
	Name      pulumi.StringPtrInput
}

func (QualitygateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qualitygateArgs)(nil)).Elem()
}

type QualitygateInput interface {
	pulumi.Input

	ToQualitygateOutput() QualitygateOutput
	ToQualitygateOutputWithContext(ctx context.Context) QualitygateOutput
}

func (*Qualitygate) ElementType() reflect.Type {
	return reflect.TypeOf((**Qualitygate)(nil)).Elem()
}

func (i *Qualitygate) ToQualitygateOutput() QualitygateOutput {
	return i.ToQualitygateOutputWithContext(context.Background())
}

func (i *Qualitygate) ToQualitygateOutputWithContext(ctx context.Context) QualitygateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualitygateOutput)
}

// QualitygateArrayInput is an input type that accepts QualitygateArray and QualitygateArrayOutput values.
// You can construct a concrete instance of `QualitygateArrayInput` via:
//
//	QualitygateArray{ QualitygateArgs{...} }
type QualitygateArrayInput interface {
	pulumi.Input

	ToQualitygateArrayOutput() QualitygateArrayOutput
	ToQualitygateArrayOutputWithContext(context.Context) QualitygateArrayOutput
}

type QualitygateArray []QualitygateInput

func (QualitygateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qualitygate)(nil)).Elem()
}

func (i QualitygateArray) ToQualitygateArrayOutput() QualitygateArrayOutput {
	return i.ToQualitygateArrayOutputWithContext(context.Background())
}

func (i QualitygateArray) ToQualitygateArrayOutputWithContext(ctx context.Context) QualitygateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualitygateArrayOutput)
}

// QualitygateMapInput is an input type that accepts QualitygateMap and QualitygateMapOutput values.
// You can construct a concrete instance of `QualitygateMapInput` via:
//
//	QualitygateMap{ "key": QualitygateArgs{...} }
type QualitygateMapInput interface {
	pulumi.Input

	ToQualitygateMapOutput() QualitygateMapOutput
	ToQualitygateMapOutputWithContext(context.Context) QualitygateMapOutput
}

type QualitygateMap map[string]QualitygateInput

func (QualitygateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qualitygate)(nil)).Elem()
}

func (i QualitygateMap) ToQualitygateMapOutput() QualitygateMapOutput {
	return i.ToQualitygateMapOutputWithContext(context.Background())
}

func (i QualitygateMap) ToQualitygateMapOutputWithContext(ctx context.Context) QualitygateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QualitygateMapOutput)
}

type QualitygateOutput struct{ *pulumi.OutputState }

func (QualitygateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Qualitygate)(nil)).Elem()
}

func (o QualitygateOutput) ToQualitygateOutput() QualitygateOutput {
	return o
}

func (o QualitygateOutput) ToQualitygateOutputWithContext(ctx context.Context) QualitygateOutput {
	return o
}

// A list of conditions that the gate uses
func (o QualitygateOutput) Conditions() QualitygateConditionArrayOutput {
	return o.ApplyT(func(v *Qualitygate) QualitygateConditionArrayOutput { return v.Conditions }).(QualitygateConditionArrayOutput)
}

func (o QualitygateOutput) CopyFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Qualitygate) pulumi.StringPtrOutput { return v.CopyFrom }).(pulumi.StringPtrOutput)
}

// When set to true this Quality Gate is set as default
func (o QualitygateOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Qualitygate) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o QualitygateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Qualitygate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type QualitygateArrayOutput struct{ *pulumi.OutputState }

func (QualitygateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qualitygate)(nil)).Elem()
}

func (o QualitygateArrayOutput) ToQualitygateArrayOutput() QualitygateArrayOutput {
	return o
}

func (o QualitygateArrayOutput) ToQualitygateArrayOutputWithContext(ctx context.Context) QualitygateArrayOutput {
	return o
}

func (o QualitygateArrayOutput) Index(i pulumi.IntInput) QualitygateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Qualitygate {
		return vs[0].([]*Qualitygate)[vs[1].(int)]
	}).(QualitygateOutput)
}

type QualitygateMapOutput struct{ *pulumi.OutputState }

func (QualitygateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qualitygate)(nil)).Elem()
}

func (o QualitygateMapOutput) ToQualitygateMapOutput() QualitygateMapOutput {
	return o
}

func (o QualitygateMapOutput) ToQualitygateMapOutputWithContext(ctx context.Context) QualitygateMapOutput {
	return o
}

func (o QualitygateMapOutput) MapIndex(k pulumi.StringInput) QualitygateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Qualitygate {
		return vs[0].(map[string]*Qualitygate)[vs[1].(string)]
	}).(QualitygateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QualitygateInput)(nil)).Elem(), &Qualitygate{})
	pulumi.RegisterInputType(reflect.TypeOf((*QualitygateArrayInput)(nil)).Elem(), QualitygateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QualitygateMapInput)(nil)).Elem(), QualitygateMap{})
	pulumi.RegisterOutputType(QualitygateOutput{})
	pulumi.RegisterOutputType(QualitygateArrayOutput{})
	pulumi.RegisterOutputType(QualitygateMapOutput{})
}
