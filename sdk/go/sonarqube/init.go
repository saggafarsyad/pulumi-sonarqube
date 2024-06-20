// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sonarqube

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/saggafarsyad/pulumi-sonarqube/sdk/go/sonarqube/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "sonarqube:index/almAzure:AlmAzure":
		r = &AlmAzure{}
	case "sonarqube:index/almGithub:AlmGithub":
		r = &AlmGithub{}
	case "sonarqube:index/almGitlab:AlmGitlab":
		r = &AlmGitlab{}
	case "sonarqube:index/azureBinding:AzureBinding":
		r = &AzureBinding{}
	case "sonarqube:index/githubBinding:GithubBinding":
		r = &GithubBinding{}
	case "sonarqube:index/gitlabBinding:GitlabBinding":
		r = &GitlabBinding{}
	case "sonarqube:index/group:Group":
		r = &Group{}
	case "sonarqube:index/groupMember:GroupMember":
		r = &GroupMember{}
	case "sonarqube:index/newCodePeriods:NewCodePeriods":
		r = &NewCodePeriods{}
	case "sonarqube:index/permissionTemplate:PermissionTemplate":
		r = &PermissionTemplate{}
	case "sonarqube:index/permissions:Permissions":
		r = &Permissions{}
	case "sonarqube:index/plugin:Plugin":
		r = &Plugin{}
	case "sonarqube:index/portfolio:Portfolio":
		r = &Portfolio{}
	case "sonarqube:index/project:Project":
		r = &Project{}
	case "sonarqube:index/projectMainBranch:ProjectMainBranch":
		r = &ProjectMainBranch{}
	case "sonarqube:index/qualitygate:Qualitygate":
		r = &Qualitygate{}
	case "sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation":
		r = &QualitygateProjectAssociation{}
	case "sonarqube:index/qualitygateUsergroupAssociation:QualitygateUsergroupAssociation":
		r = &QualitygateUsergroupAssociation{}
	case "sonarqube:index/qualityprofile:Qualityprofile":
		r = &Qualityprofile{}
	case "sonarqube:index/qualityprofileActivateRule:QualityprofileActivateRule":
		r = &QualityprofileActivateRule{}
	case "sonarqube:index/qualityprofileProjectAssociation:QualityprofileProjectAssociation":
		r = &QualityprofileProjectAssociation{}
	case "sonarqube:index/rule:Rule":
		r = &Rule{}
	case "sonarqube:index/setting:Setting":
		r = &Setting{}
	case "sonarqube:index/user:User":
		r = &User{}
	case "sonarqube:index/userExternalIdentity:UserExternalIdentity":
		r = &UserExternalIdentity{}
	case "sonarqube:index/userToken:UserToken":
		r = &UserToken{}
	case "sonarqube:index/webhook:Webhook":
		r = &Webhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:sonarqube" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/almAzure",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/almGithub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/almGitlab",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/azureBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/githubBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/gitlabBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/groupMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/newCodePeriods",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/permissionTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/permissions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/plugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/portfolio",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/projectMainBranch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/qualitygate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/qualitygateProjectAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/qualitygateUsergroupAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/qualityprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/qualityprofileActivateRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/qualityprofileProjectAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/setting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/userExternalIdentity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/userToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sonarqube",
		"index/webhook",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"sonarqube",
		&pkg{version},
	)
}
