// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newLifecyclePolicyDocumentDataSource,
			TypeName: "aws_ecr_lifecycle_policy_document",
			Name:     "Lifecycle Policy Document",
		},
		{
			Factory:  newRepositoriesDataSource,
			TypeName: "aws_ecr_repositories",
			Name:     "Repositories",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newAccountSettingResource,
			TypeName: "aws_ecr_account_setting",
			Name:     "Account Setting",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAuthorizationToken,
			TypeName: "aws_ecr_authorization_token",
			Name:     "Authorization Token",
		},
		{
			Factory:  dataSourceImage,
			TypeName: "aws_ecr_image",
			Name:     "Image",
		},
		{
			Factory:  dataSourcePullThroughCacheRule,
			TypeName: "aws_ecr_pull_through_cache_rule",
			Name:     "Pull Through Cache Rule",
		},
		{
			Factory:  dataSourceRepository,
			TypeName: "aws_ecr_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceRepositoryCreationTemplate,
			TypeName: "aws_ecr_repository_creation_template",
			Name:     "Repository Creation Template",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceLifecyclePolicy,
			TypeName: "aws_ecr_lifecycle_policy",
			Name:     "Lifecycle Policy",
		},
		{
			Factory:  resourcePullThroughCacheRule,
			TypeName: "aws_ecr_pull_through_cache_rule",
			Name:     "Pull Through Cache Rule",
		},
		{
			Factory:  resourceRegistryPolicy,
			TypeName: "aws_ecr_registry_policy",
			Name:     "Registry Policy",
		},
		{
			Factory:  resourceRegistryScanningConfiguration,
			TypeName: "aws_ecr_registry_scanning_configuration",
			Name:     "Registry Scanning Configuration",
		},
		{
			Factory:  resourceReplicationConfiguration,
			TypeName: "aws_ecr_replication_configuration",
			Name:     "Replication Configuration",
		},
		{
			Factory:  resourceRepository,
			TypeName: "aws_ecr_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRepositoryCreationTemplate,
			TypeName: "aws_ecr_repository_creation_template",
			Name:     "Repository Creation Template",
		},
		{
			Factory:  resourceRepositoryPolicy,
			TypeName: "aws_ecr_repository_policy",
			Name:     "Repsitory Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECR
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ecr.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*ecr.Options){
		ecr.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	}

	optFns = append(optFns, servicePackageExtraOptFns(ctx, p, config)...)

	return ecr.NewFromConfig(cfg, optFns...), nil
}

func servicePackageExtraOptFns(ctx context.Context, sp conns.ServicePackage, config map[string]any) []func(*ecr.Options) {
	if v, ok := sp.(interface {
		extraOptFns(context.Context, map[string]any) []func(*ecr.Options)
	}); ok {
		return v.extraOptFns(ctx, config)
	}

	return nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
