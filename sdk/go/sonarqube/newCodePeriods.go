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

// ## # NewCodePeriods
//
// Provides a Sonarqube New Code Periods resource. This can be used to manage Sonarqube New Code Periods.
//
// ## Example: Set the global new code period to a number of days
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
//			_, err := sonarqube.NewNewCodePeriods(ctx, "codePeriod", &sonarqube.NewCodePeriodsArgs{
//				Type:  pulumi.String("NUMBER_OF_DAYS"),
//				Value: pulumi.String("7"),
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
type NewCodePeriods struct {
	pulumi.CustomResourceState

	Branch  pulumi.StringPtrOutput `pulumi:"branch"`
	Project pulumi.StringPtrOutput `pulumi:"project"`
	Type    pulumi.StringOutput    `pulumi:"type"`
	Value   pulumi.StringPtrOutput `pulumi:"value"`
}

// NewNewCodePeriods registers a new resource with the given unique name, arguments, and options.
func NewNewCodePeriods(ctx *pulumi.Context,
	name string, args *NewCodePeriodsArgs, opts ...pulumi.ResourceOption) (*NewCodePeriods, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NewCodePeriods
	err := ctx.RegisterResource("sonarqube:index/newCodePeriods:NewCodePeriods", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNewCodePeriods gets an existing NewCodePeriods resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNewCodePeriods(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NewCodePeriodsState, opts ...pulumi.ResourceOption) (*NewCodePeriods, error) {
	var resource NewCodePeriods
	err := ctx.ReadResource("sonarqube:index/newCodePeriods:NewCodePeriods", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NewCodePeriods resources.
type newCodePeriodsState struct {
	Branch  *string `pulumi:"branch"`
	Project *string `pulumi:"project"`
	Type    *string `pulumi:"type"`
	Value   *string `pulumi:"value"`
}

type NewCodePeriodsState struct {
	Branch  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Type    pulumi.StringPtrInput
	Value   pulumi.StringPtrInput
}

func (NewCodePeriodsState) ElementType() reflect.Type {
	return reflect.TypeOf((*newCodePeriodsState)(nil)).Elem()
}

type newCodePeriodsArgs struct {
	Branch  *string `pulumi:"branch"`
	Project *string `pulumi:"project"`
	Type    string  `pulumi:"type"`
	Value   *string `pulumi:"value"`
}

// The set of arguments for constructing a NewCodePeriods resource.
type NewCodePeriodsArgs struct {
	Branch  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Type    pulumi.StringInput
	Value   pulumi.StringPtrInput
}

func (NewCodePeriodsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*newCodePeriodsArgs)(nil)).Elem()
}

type NewCodePeriodsInput interface {
	pulumi.Input

	ToNewCodePeriodsOutput() NewCodePeriodsOutput
	ToNewCodePeriodsOutputWithContext(ctx context.Context) NewCodePeriodsOutput
}

func (*NewCodePeriods) ElementType() reflect.Type {
	return reflect.TypeOf((**NewCodePeriods)(nil)).Elem()
}

func (i *NewCodePeriods) ToNewCodePeriodsOutput() NewCodePeriodsOutput {
	return i.ToNewCodePeriodsOutputWithContext(context.Background())
}

func (i *NewCodePeriods) ToNewCodePeriodsOutputWithContext(ctx context.Context) NewCodePeriodsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NewCodePeriodsOutput)
}

// NewCodePeriodsArrayInput is an input type that accepts NewCodePeriodsArray and NewCodePeriodsArrayOutput values.
// You can construct a concrete instance of `NewCodePeriodsArrayInput` via:
//
//	NewCodePeriodsArray{ NewCodePeriodsArgs{...} }
type NewCodePeriodsArrayInput interface {
	pulumi.Input

	ToNewCodePeriodsArrayOutput() NewCodePeriodsArrayOutput
	ToNewCodePeriodsArrayOutputWithContext(context.Context) NewCodePeriodsArrayOutput
}

type NewCodePeriodsArray []NewCodePeriodsInput

func (NewCodePeriodsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NewCodePeriods)(nil)).Elem()
}

func (i NewCodePeriodsArray) ToNewCodePeriodsArrayOutput() NewCodePeriodsArrayOutput {
	return i.ToNewCodePeriodsArrayOutputWithContext(context.Background())
}

func (i NewCodePeriodsArray) ToNewCodePeriodsArrayOutputWithContext(ctx context.Context) NewCodePeriodsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NewCodePeriodsArrayOutput)
}

// NewCodePeriodsMapInput is an input type that accepts NewCodePeriodsMap and NewCodePeriodsMapOutput values.
// You can construct a concrete instance of `NewCodePeriodsMapInput` via:
//
//	NewCodePeriodsMap{ "key": NewCodePeriodsArgs{...} }
type NewCodePeriodsMapInput interface {
	pulumi.Input

	ToNewCodePeriodsMapOutput() NewCodePeriodsMapOutput
	ToNewCodePeriodsMapOutputWithContext(context.Context) NewCodePeriodsMapOutput
}

type NewCodePeriodsMap map[string]NewCodePeriodsInput

func (NewCodePeriodsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NewCodePeriods)(nil)).Elem()
}

func (i NewCodePeriodsMap) ToNewCodePeriodsMapOutput() NewCodePeriodsMapOutput {
	return i.ToNewCodePeriodsMapOutputWithContext(context.Background())
}

func (i NewCodePeriodsMap) ToNewCodePeriodsMapOutputWithContext(ctx context.Context) NewCodePeriodsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NewCodePeriodsMapOutput)
}

type NewCodePeriodsOutput struct{ *pulumi.OutputState }

func (NewCodePeriodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NewCodePeriods)(nil)).Elem()
}

func (o NewCodePeriodsOutput) ToNewCodePeriodsOutput() NewCodePeriodsOutput {
	return o
}

func (o NewCodePeriodsOutput) ToNewCodePeriodsOutputWithContext(ctx context.Context) NewCodePeriodsOutput {
	return o
}

func (o NewCodePeriodsOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NewCodePeriods) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o NewCodePeriodsOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NewCodePeriods) pulumi.StringPtrOutput { return v.Project }).(pulumi.StringPtrOutput)
}

func (o NewCodePeriodsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NewCodePeriods) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NewCodePeriodsOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NewCodePeriods) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type NewCodePeriodsArrayOutput struct{ *pulumi.OutputState }

func (NewCodePeriodsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NewCodePeriods)(nil)).Elem()
}

func (o NewCodePeriodsArrayOutput) ToNewCodePeriodsArrayOutput() NewCodePeriodsArrayOutput {
	return o
}

func (o NewCodePeriodsArrayOutput) ToNewCodePeriodsArrayOutputWithContext(ctx context.Context) NewCodePeriodsArrayOutput {
	return o
}

func (o NewCodePeriodsArrayOutput) Index(i pulumi.IntInput) NewCodePeriodsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NewCodePeriods {
		return vs[0].([]*NewCodePeriods)[vs[1].(int)]
	}).(NewCodePeriodsOutput)
}

type NewCodePeriodsMapOutput struct{ *pulumi.OutputState }

func (NewCodePeriodsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NewCodePeriods)(nil)).Elem()
}

func (o NewCodePeriodsMapOutput) ToNewCodePeriodsMapOutput() NewCodePeriodsMapOutput {
	return o
}

func (o NewCodePeriodsMapOutput) ToNewCodePeriodsMapOutputWithContext(ctx context.Context) NewCodePeriodsMapOutput {
	return o
}

func (o NewCodePeriodsMapOutput) MapIndex(k pulumi.StringInput) NewCodePeriodsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NewCodePeriods {
		return vs[0].(map[string]*NewCodePeriods)[vs[1].(string)]
	}).(NewCodePeriodsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NewCodePeriodsInput)(nil)).Elem(), &NewCodePeriods{})
	pulumi.RegisterInputType(reflect.TypeOf((*NewCodePeriodsArrayInput)(nil)).Elem(), NewCodePeriodsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NewCodePeriodsMapInput)(nil)).Elem(), NewCodePeriodsMap{})
	pulumi.RegisterOutputType(NewCodePeriodsOutput{})
	pulumi.RegisterOutputType(NewCodePeriodsArrayOutput{})
	pulumi.RegisterOutputType(NewCodePeriodsMapOutput{})
}
