// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newAccountSuppressionAttributesResource,
			TypeName: "aws_sesv2_account_suppression_attributes",
			Name:     "Account Suppression Attributes",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
			Name:     "Configuration Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
			Name:     "Dedicated IP Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceEmailIdentity,
			TypeName: "aws_sesv2_email_identity",
			Name:     "Email Identity",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceEmailIdentityMailFromAttributes,
			TypeName: "aws_sesv2_email_identity_mail_from_attributes",
			Name:     "Email Identity Mail From Attributes",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccountVDMAttributes,
			TypeName: "aws_sesv2_account_vdm_attributes",
			Name:     "Account VDM Attributes",
		},
		{
			Factory:  resourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
			Name:     "Configuration Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConfigurationSetEventDestination,
			TypeName: "aws_sesv2_configuration_set_event_destination",
			Name:     "Configuration Set Event Destination",
		},
		{
			Factory:  resourceContactList,
			TypeName: "aws_sesv2_contact_list",
			Name:     "Contact List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDedicatedIPAssignment,
			TypeName: "aws_sesv2_dedicated_ip_assignment",
			Name:     "Dedicated IP Assignment",
		},
		{
			Factory:  resourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
			Name:     "Dedicated IP Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceEmailIdentity,
			TypeName: "aws_sesv2_email_identity",
			Name:     "Email Identity",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceEmailIdentityFeedbackAttributes,
			TypeName: "aws_sesv2_email_identity_feedback_attributes",
			Name:     "Email Identity Feedback Attributes",
		},
		{
			Factory:  resourceEmailIdentityMailFromAttributes,
			TypeName: "aws_sesv2_email_identity_mail_from_attributes",
			Name:     "Email Identity Mail From Attributes",
		},
		{
			Factory:  resourceEmailIdentityPolicy,
			TypeName: "aws_sesv2_email_identity_policy",
			Name:     "Email Identity Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SESV2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sesv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*sesv2.Options){
		sesv2.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return sesv2.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*sesv2.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*sesv2.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *sesv2.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*sesv2.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
