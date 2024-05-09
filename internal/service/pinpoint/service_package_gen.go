// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package pinpoint

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	pinpoint_sdkv1 "github.com/aws/aws-sdk-go/service/pinpoint"
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
			Factory:  ResourceADMChannel,
			TypeName: "aws_pinpoint_adm_channel",
		},
		{
			Factory:  ResourceAPNSChannel,
			TypeName: "aws_pinpoint_apns_channel",
		},
		{
			Factory:  ResourceAPNSSandboxChannel,
			TypeName: "aws_pinpoint_apns_sandbox_channel",
		},
		{
			Factory:  ResourceAPNSVoIPChannel,
			TypeName: "aws_pinpoint_apns_voip_channel",
		},
		{
			Factory:  ResourceAPNSVoIPSandboxChannel,
			TypeName: "aws_pinpoint_apns_voip_sandbox_channel",
		},
		{
			Factory:  ResourceApp,
			TypeName: "aws_pinpoint_app",
			Name:     "App",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceBaiduChannel,
			TypeName: "aws_pinpoint_baidu_channel",
		},
		{
			Factory:  ResourceEmailChannel,
			TypeName: "aws_pinpoint_email_channel",
		},
		{
			Factory:  ResourceEventStream,
			TypeName: "aws_pinpoint_event_stream",
		},
		{
			Factory:  ResourceGCMChannel,
			TypeName: "aws_pinpoint_gcm_channel",
		},
		{
			Factory:  ResourceSMSChannel,
			TypeName: "aws_pinpoint_sms_channel",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Pinpoint
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*pinpoint_sdkv1.Pinpoint, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return pinpoint_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
