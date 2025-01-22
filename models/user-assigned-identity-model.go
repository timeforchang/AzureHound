package models

import "github.com/bloodhoundad/azurehound/v2/models/azure"

type UserAssignedIdentityModel struct {
	azure.UserAssignedIdentityModel
	SubscriptionId string `json:"subscriptionId"`
	TenantId       string `json:"tenantId"`
}
