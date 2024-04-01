// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
)

func FindBackendEnvironmentByAppIDAndEnvironmentName(ctx context.Context, conn *amplify.Client, appID, environmentName string) (*types.BackendEnvironment, error) {
	input := &amplify.GetBackendEnvironmentInput{
		AppId:           aws.String(appID),
		EnvironmentName: aws.String(environmentName),
	}

	output, err := conn.GetBackendEnvironment(ctx, input)

	if errs.IsA[*types.NotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.BackendEnvironment == nil {
		return nil, &retry.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return output.BackendEnvironment, nil
}

func FindBranchByAppIDAndBranchName(ctx context.Context, conn *amplify.Client, appID, branchName string) (*types.Branch, error) {
	input := &amplify.GetBranchInput{
		AppId:      aws.String(appID),
		BranchName: aws.String(branchName),
	}

	output, err := conn.GetBranch(ctx, input)

	if errs.IsA[*types.NotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.Branch == nil {
		return nil, &retry.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return output.Branch, nil
}

func FindDomainAssociationByAppIDAndDomainName(ctx context.Context, conn *amplify.Client, appID, domainName string) (*types.DomainAssociation, error) {
	input := &amplify.GetDomainAssociationInput{
		AppId:      aws.String(appID),
		DomainName: aws.String(domainName),
	}

	output, err := conn.GetDomainAssociation(ctx, input)

	if errs.IsA[*types.NotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.DomainAssociation == nil {
		return nil, &retry.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return output.DomainAssociation, nil
}

func FindWebhookByID(ctx context.Context, conn *amplify.Client, id string) (*types.Webhook, error) {
	input := &amplify.GetWebhookInput{
		WebhookId: aws.String(id),
	}

	output, err := conn.GetWebhook(ctx, input)

	if errs.IsA[*types.NotFoundException](err) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.Webhook == nil {
		return nil, &retry.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return output.Webhook, nil
}
