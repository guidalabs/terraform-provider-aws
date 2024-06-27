// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"
	"fmt"
	"net"
	"net/url"

	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
)

var _ endpoints_sdkv1.Resolver = resolverSDKv1{}

type resolverSDKv1 struct {
	ctx context.Context
}

func newEndpointResolverSDKv1(ctx context.Context) resolverSDKv1 {
	return resolverSDKv1{
		ctx: ctx,
	}
}

func (r resolverSDKv1) EndpointFor(service, region string, opts ...func(*endpoints_sdkv1.Options)) (endpoint endpoints_sdkv1.ResolvedEndpoint, err error) {
	ctx := r.ctx

	var opt endpoints_sdkv1.Options
	opt.Set(opts...)

	useFIPS := opt.UseFIPSEndpoint == endpoints_sdkv1.FIPSEndpointStateEnabled

	defaultResolver := endpoints_sdkv1.DefaultResolver()

	if useFIPS {
		ctx = tflog.SetField(ctx, "tf_aws.use_fips", useFIPS)

		endpoint, err = defaultResolver.EndpointFor(service, region, opts...)
		if err != nil {
			return endpoint, err
		}

		tflog.Debug(ctx, "endpoint resolved", map[string]any{
			"tf_aws.endpoint": endpoint.URL,
		})

		var endpointURL *url.URL
		endpointURL, err = url.Parse(endpoint.URL)
		if err != nil {
			return endpoint, err
		}

		hostname := endpointURL.Hostname()
		_, err = net.LookupHost(hostname)
		if err != nil {
			if dnsErr, ok := errs.As[*net.DNSError](err); ok && dnsErr.IsNotFound {
				tflog.Debug(ctx, "default endpoint host not found, disabling FIPS", map[string]any{
					"tf_aws.hostname": hostname,
				})
				opts = append(opts, func(o *endpoints_sdkv1.Options) {
					o.UseFIPSEndpoint = endpoints_sdkv1.FIPSEndpointStateDisabled
				})
			} else {
				err = fmt.Errorf("looking up accessanalyzer endpoint %q: %s", hostname, err)
				return
			}
		} else {
			return endpoint, err
		}
	}

	return defaultResolver.EndpointFor(service, region, opts...)
}
