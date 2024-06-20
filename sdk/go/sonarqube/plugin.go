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

// ## # Plugin
//
// Provides a Sonarqube Plugin resource. This can be used to create and manage Sonarqube Plugins.
//
// ## Example: create a project
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
//			_, err := sonarqube.NewPlugin(ctx, "main", &sonarqube.PluginArgs{
//				Key: pulumi.String("cloudformation"),
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
// # Projects can be imported using their plugin key
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Notes
//
// After installing or uninstalling a plugin, the sonarqube server needs to be restarted.
type Plugin struct {
	pulumi.CustomResourceState

	Key pulumi.StringOutput `pulumi:"key"`
}

// NewPlugin registers a new resource with the given unique name, arguments, and options.
func NewPlugin(ctx *pulumi.Context,
	name string, args *PluginArgs, opts ...pulumi.ResourceOption) (*Plugin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Plugin
	err := ctx.RegisterResource("sonarqube:index/plugin:Plugin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlugin gets an existing Plugin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlugin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluginState, opts ...pulumi.ResourceOption) (*Plugin, error) {
	var resource Plugin
	err := ctx.ReadResource("sonarqube:index/plugin:Plugin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plugin resources.
type pluginState struct {
	Key *string `pulumi:"key"`
}

type PluginState struct {
	Key pulumi.StringPtrInput
}

func (PluginState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginState)(nil)).Elem()
}

type pluginArgs struct {
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a Plugin resource.
type PluginArgs struct {
	Key pulumi.StringInput
}

func (PluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginArgs)(nil)).Elem()
}

type PluginInput interface {
	pulumi.Input

	ToPluginOutput() PluginOutput
	ToPluginOutputWithContext(ctx context.Context) PluginOutput
}

func (*Plugin) ElementType() reflect.Type {
	return reflect.TypeOf((**Plugin)(nil)).Elem()
}

func (i *Plugin) ToPluginOutput() PluginOutput {
	return i.ToPluginOutputWithContext(context.Background())
}

func (i *Plugin) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginOutput)
}

// PluginArrayInput is an input type that accepts PluginArray and PluginArrayOutput values.
// You can construct a concrete instance of `PluginArrayInput` via:
//
//	PluginArray{ PluginArgs{...} }
type PluginArrayInput interface {
	pulumi.Input

	ToPluginArrayOutput() PluginArrayOutput
	ToPluginArrayOutputWithContext(context.Context) PluginArrayOutput
}

type PluginArray []PluginInput

func (PluginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Plugin)(nil)).Elem()
}

func (i PluginArray) ToPluginArrayOutput() PluginArrayOutput {
	return i.ToPluginArrayOutputWithContext(context.Background())
}

func (i PluginArray) ToPluginArrayOutputWithContext(ctx context.Context) PluginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginArrayOutput)
}

// PluginMapInput is an input type that accepts PluginMap and PluginMapOutput values.
// You can construct a concrete instance of `PluginMapInput` via:
//
//	PluginMap{ "key": PluginArgs{...} }
type PluginMapInput interface {
	pulumi.Input

	ToPluginMapOutput() PluginMapOutput
	ToPluginMapOutputWithContext(context.Context) PluginMapOutput
}

type PluginMap map[string]PluginInput

func (PluginMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Plugin)(nil)).Elem()
}

func (i PluginMap) ToPluginMapOutput() PluginMapOutput {
	return i.ToPluginMapOutputWithContext(context.Background())
}

func (i PluginMap) ToPluginMapOutputWithContext(ctx context.Context) PluginMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginMapOutput)
}

type PluginOutput struct{ *pulumi.OutputState }

func (PluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plugin)(nil)).Elem()
}

func (o PluginOutput) ToPluginOutput() PluginOutput {
	return o
}

func (o PluginOutput) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return o
}

func (o PluginOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Plugin) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type PluginArrayOutput struct{ *pulumi.OutputState }

func (PluginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Plugin)(nil)).Elem()
}

func (o PluginArrayOutput) ToPluginArrayOutput() PluginArrayOutput {
	return o
}

func (o PluginArrayOutput) ToPluginArrayOutputWithContext(ctx context.Context) PluginArrayOutput {
	return o
}

func (o PluginArrayOutput) Index(i pulumi.IntInput) PluginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Plugin {
		return vs[0].([]*Plugin)[vs[1].(int)]
	}).(PluginOutput)
}

type PluginMapOutput struct{ *pulumi.OutputState }

func (PluginMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Plugin)(nil)).Elem()
}

func (o PluginMapOutput) ToPluginMapOutput() PluginMapOutput {
	return o
}

func (o PluginMapOutput) ToPluginMapOutputWithContext(ctx context.Context) PluginMapOutput {
	return o
}

func (o PluginMapOutput) MapIndex(k pulumi.StringInput) PluginOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Plugin {
		return vs[0].(map[string]*Plugin)[vs[1].(string)]
	}).(PluginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PluginInput)(nil)).Elem(), &Plugin{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginArrayInput)(nil)).Elem(), PluginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginMapInput)(nil)).Elem(), PluginMap{})
	pulumi.RegisterOutputType(PluginOutput{})
	pulumi.RegisterOutputType(PluginArrayOutput{})
	pulumi.RegisterOutputType(PluginMapOutput{})
}
