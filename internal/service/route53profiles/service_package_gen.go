// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package route53profiles

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53profiles"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceProfiles,
			TypeName: "aws_route53profiles_profiles",
			Name:     "Profiles",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceAssociation,
			TypeName: "aws_route53profiles_association",
			Name:     "Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceProfile,
			TypeName: "aws_route53profiles_profile",
			Name:     "Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceResourceAssociation,
			TypeName: "aws_route53profiles_resource_association",
			Name:     "ResourceAssociation",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53Profiles
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*route53profiles.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*route53profiles.Options){
		route53profiles.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	}

	optFns = append(optFns, servicePackageExtraOptFns(ctx, p, config)...)

	return route53profiles.NewFromConfig(cfg, optFns...), nil
}

func servicePackageExtraOptFns(ctx context.Context, sp conns.ServicePackage, config map[string]any) []func(*route53profiles.Options) {
	if v, ok := sp.(interface {
		extraOptFns(context.Context, map[string]any) []func(*route53profiles.Options)
	}); ok {
		return v.extraOptFns(ctx, config)
	}

	return nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
