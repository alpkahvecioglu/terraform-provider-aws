// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package verifiedpermissions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourcePolicyStore,
			TypeName: "aws_verifiedpermissions_policy_store",
			Name:     "Policy Store",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceIdentitySource,
			TypeName: "aws_verifiedpermissions_identity_source",
			Name:     "Identity Source",
		},
		{
			Factory:  newResourcePolicy,
			TypeName: "aws_verifiedpermissions_policy",
			Name:     "Policy",
		},
		{
			Factory:  newResourcePolicyStore,
			TypeName: "aws_verifiedpermissions_policy_store",
			Name:     "Policy Store",
		},
		{
			Factory:  newResourcePolicyTemplate,
			TypeName: "aws_verifiedpermissions_policy_template",
			Name:     "Policy Template",
		},
		{
			Factory:  newResourceSchema,
			TypeName: "aws_verifiedpermissions_schema",
			Name:     "Schema",
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
	return names.VerifiedPermissions
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*verifiedpermissions.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*verifiedpermissions.Options){
		verifiedpermissions.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	}

	optFns = append(optFns, servicePackageExtraOptFns(ctx, p, config)...)

	return verifiedpermissions.NewFromConfig(cfg, optFns...), nil
}

func servicePackageExtraOptFns(ctx context.Context, sp conns.ServicePackage, config map[string]any) []func(*verifiedpermissions.Options) {
	if v, ok := sp.(interface {
		extraOptFns(context.Context, map[string]any) []func(*verifiedpermissions.Options)
	}); ok {
		return v.extraOptFns(ctx, config)
	}

	return nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
