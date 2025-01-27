// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceApplication,
			TypeName: "aws_servicecatalogappregistry_application",
			Name:     "Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newDataSourceAttributeGroup,
			TypeName: "aws_servicecatalogappregistry_attribute_group",
			Name:     "Attribute Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newDataSourceAttributeGroupAssociations,
			TypeName: "aws_servicecatalogappregistry_attribute_group_associations",
			Name:     "Attribute Group Associations",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceApplication,
			TypeName: "aws_servicecatalogappregistry_application",
			Name:     "Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceAttributeGroup,
			TypeName: "aws_servicecatalogappregistry_attribute_group",
			Name:     "Attribute Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceAttributeGroupAssociation,
			TypeName: "aws_servicecatalogappregistry_attribute_group_association",
			Name:     "Attribute Group Association",
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
	return names.ServiceCatalogAppRegistry
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*servicecatalogappregistry.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*servicecatalogappregistry.Options){
		servicecatalogappregistry.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	}

	optFns = append(optFns, servicePackageExtraOptFns(ctx, p, config)...)

	return servicecatalogappregistry.NewFromConfig(cfg, optFns...), nil
}

func servicePackageExtraOptFns(ctx context.Context, sp conns.ServicePackage, config map[string]any) []func(*servicecatalogappregistry.Options) {
	if v, ok := sp.(interface {
		extraOptFns(context.Context, map[string]any) []func(*servicecatalogappregistry.Options)
	}); ok {
		return v.extraOptFns(ctx, config)
	}

	return nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
