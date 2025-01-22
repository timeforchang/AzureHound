package client

import (
	"context"
	"fmt"

	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
)

// ListAzureUserAssignedManagedIdentities https://learn.microsoft.com/en-us/rest/api/managedidentity/user-assigned-identities/list-by-subscription
func (s *azureClient) ListAzureUserAssignedManagedIdentities(ctx context.Context, subscriptionId string) <-chan AzureResult[azure.UserAssignedIdentityModel] {
	var (
		out    = make(chan AzureResult[azure.UserAssignedIdentityModel])
		path   = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.ManagedIdentity/userAssignedIdentities", subscriptionId)
		params = query.RMParams{ApiVersion: "2023-01-31"}
	)

	go getAzureObjectList[azure.UserAssignedIdentityModel](s.resourceManager, ctx, path, params, out)

	return out
}
