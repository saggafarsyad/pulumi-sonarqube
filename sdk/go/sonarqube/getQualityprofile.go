// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sonarqube

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/saggafarsyad/pulumi-sonarqube/sdk/go/sonarqube/internal"
)

// ## # Data Source: Qualityprofile
//
// # Use this data source to get a Sonarqube qualityprofile resource
//
// ## Example Usage
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
//			_, err := sonarqube.LookupQualityprofile(ctx, &sonarqube.LookupQualityprofileArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupQualityprofile(ctx *pulumi.Context, args *LookupQualityprofileArgs, opts ...pulumi.InvokeOption) (*LookupQualityprofileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQualityprofileResult
	err := ctx.Invoke("sonarqube:index/getQualityprofile:getQualityprofile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQualityprofile.
type LookupQualityprofileArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getQualityprofile.
type LookupQualityprofileResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	IsDefault bool   `pulumi:"isDefault"`
	Key       string `pulumi:"key"`
	Language  string `pulumi:"language"`
	Name      string `pulumi:"name"`
}

func LookupQualityprofileOutput(ctx *pulumi.Context, args LookupQualityprofileOutputArgs, opts ...pulumi.InvokeOption) LookupQualityprofileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQualityprofileResult, error) {
			args := v.(LookupQualityprofileArgs)
			r, err := LookupQualityprofile(ctx, &args, opts...)
			var s LookupQualityprofileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQualityprofileResultOutput)
}

// A collection of arguments for invoking getQualityprofile.
type LookupQualityprofileOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupQualityprofileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQualityprofileArgs)(nil)).Elem()
}

// A collection of values returned by getQualityprofile.
type LookupQualityprofileResultOutput struct{ *pulumi.OutputState }

func (LookupQualityprofileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQualityprofileResult)(nil)).Elem()
}

func (o LookupQualityprofileResultOutput) ToLookupQualityprofileResultOutput() LookupQualityprofileResultOutput {
	return o
}

func (o LookupQualityprofileResultOutput) ToLookupQualityprofileResultOutputWithContext(ctx context.Context) LookupQualityprofileResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQualityprofileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQualityprofileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQualityprofileResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQualityprofileResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

func (o LookupQualityprofileResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQualityprofileResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupQualityprofileResultOutput) Language() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQualityprofileResult) string { return v.Language }).(pulumi.StringOutput)
}

func (o LookupQualityprofileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQualityprofileResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQualityprofileResultOutput{})
}
