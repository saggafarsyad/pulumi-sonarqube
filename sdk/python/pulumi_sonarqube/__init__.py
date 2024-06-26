# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alm_azure import *
from .alm_github import *
from .alm_gitlab import *
from .azure_binding import *
from .get_group import *
from .get_portfolio import *
from .get_project import *
from .get_qualitygate import *
from .get_qualityprofile import *
from .get_rule import *
from .get_user import *
from .github_binding import *
from .gitlab_binding import *
from .group import *
from .group_member import *
from .new_code_periods import *
from .permission_template import *
from .permissions import *
from .plugin import *
from .portfolio import *
from .project import *
from .project_main_branch import *
from .provider import *
from .qualitygate import *
from .qualitygate_project_association import *
from .qualitygate_usergroup_association import *
from .qualityprofile import *
from .qualityprofile_activate_rule import *
from .qualityprofile_project_association import *
from .rule import *
from .setting import *
from .user import *
from .user_external_identity import *
from .user_token import *
from .webhook import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_sonarqube.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_sonarqube.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "sonarqube",
  "mod": "index/almAzure",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/almAzure:AlmAzure": "AlmAzure"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/almGithub",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/almGithub:AlmGithub": "AlmGithub"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/almGitlab",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/almGitlab:AlmGitlab": "AlmGitlab"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/azureBinding",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/azureBinding:AzureBinding": "AzureBinding"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/githubBinding",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/githubBinding:GithubBinding": "GithubBinding"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/gitlabBinding",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/gitlabBinding:GitlabBinding": "GitlabBinding"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/group",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/group:Group": "Group"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/groupMember",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/groupMember:GroupMember": "GroupMember"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/newCodePeriods",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/newCodePeriods:NewCodePeriods": "NewCodePeriods"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/permissionTemplate",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/permissionTemplate:PermissionTemplate": "PermissionTemplate"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/permissions",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/permissions:Permissions": "Permissions"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/plugin",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/plugin:Plugin": "Plugin"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/portfolio",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/portfolio:Portfolio": "Portfolio"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/project",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/project:Project": "Project"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/projectMainBranch",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/projectMainBranch:ProjectMainBranch": "ProjectMainBranch"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/qualitygate",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/qualitygate:Qualitygate": "Qualitygate"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/qualitygateProjectAssociation",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation": "QualitygateProjectAssociation"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/qualitygateUsergroupAssociation",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/qualitygateUsergroupAssociation:QualitygateUsergroupAssociation": "QualitygateUsergroupAssociation"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/qualityprofile",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/qualityprofile:Qualityprofile": "Qualityprofile"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/qualityprofileActivateRule",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/qualityprofileActivateRule:QualityprofileActivateRule": "QualityprofileActivateRule"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/qualityprofileProjectAssociation",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/qualityprofileProjectAssociation:QualityprofileProjectAssociation": "QualityprofileProjectAssociation"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/rule",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/rule:Rule": "Rule"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/setting",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/setting:Setting": "Setting"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/user",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/user:User": "User"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/userExternalIdentity",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/userExternalIdentity:UserExternalIdentity": "UserExternalIdentity"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/userToken",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/userToken:UserToken": "UserToken"
  }
 },
 {
  "pkg": "sonarqube",
  "mod": "index/webhook",
  "fqn": "pulumi_sonarqube",
  "classes": {
   "sonarqube:index/webhook:Webhook": "Webhook"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "sonarqube",
  "token": "pulumi:providers:sonarqube",
  "fqn": "pulumi_sonarqube",
  "class": "Provider"
 }
]
"""
)
