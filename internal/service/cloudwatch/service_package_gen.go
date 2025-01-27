// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCompositeAlarm,
			TypeName: "aws_cloudwatch_composite_alarm",
			Name:     "Composite Alarm",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDashboard,
			TypeName: "aws_cloudwatch_dashboard",
			Name:     "Dashboard",
		},
		{
			Factory:  resourceMetricAlarm,
			TypeName: "aws_cloudwatch_metric_alarm",
			Name:     "Metric Alarm",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceMetricStream,
			TypeName: "aws_cloudwatch_metric_stream",
			Name:     "Metric Stream",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudWatch
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudwatch.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*cloudwatch.Options){
		cloudwatch.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	}

	optFns = append(optFns, servicePackageExtraOptFns(ctx, p, config)...)

	return cloudwatch.NewFromConfig(cfg, optFns...), nil
}

func servicePackageExtraOptFns(ctx context.Context, sp conns.ServicePackage, config map[string]any) []func(*cloudwatch.Options) {
	if v, ok := sp.(interface {
		extraOptFns(context.Context, map[string]any) []func(*cloudwatch.Options)
	}); ok {
		return v.extraOptFns(ctx, config)
	}

	return nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
