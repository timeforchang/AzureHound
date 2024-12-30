// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bloodhoundad/azurehound/v2/client (interfaces: AzureClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/bloodhoundad/azurehound/v2/client"
	"github.com/bloodhoundad/azurehound/v2/client/query"
	"github.com/bloodhoundad/azurehound/v2/models/azure"
	"go.uber.org/mock/gomock"
)

// MockAzureClient is a mock of AzureClient interface.
type MockAzureClient struct {
	ctrl     *gomock.Controller
	recorder *MockAzureClientMockRecorder
}

// MockAzureClientMockRecorder is the mock recorder for MockAzureClient.
type MockAzureClientMockRecorder struct {
	mock *MockAzureClient
}

// NewMockAzureClient creates a new mock instance.
func NewMockAzureClient(ctrl *gomock.Controller) *MockAzureClient {
	mock := &MockAzureClient{ctrl: ctrl}
	mock.recorder = &MockAzureClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzureClient) EXPECT() *MockAzureClientMockRecorder {
	return m.recorder
}

// CloseIdleConnections mocks base method.
func (m *MockAzureClient) CloseIdleConnections() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseIdleConnections")
}

// CloseIdleConnections indicates an expected call of CloseIdleConnections.
func (mr *MockAzureClientMockRecorder) CloseIdleConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseIdleConnections", reflect.TypeOf((*MockAzureClient)(nil).CloseIdleConnections))
}

// GetAzureADOrganization mocks base method.
func (m *MockAzureClient) GetAzureADOrganization(arg0 context.Context, arg1 []string) (*azure.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADOrganization", arg0, arg1)
	ret0, _ := ret[0].(*azure.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADOrganization indicates an expected call of GetAzureADOrganization.
func (mr *MockAzureClientMockRecorder) GetAzureADOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADOrganization", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADOrganization), arg0, arg1)
}

// GetAzureADTenants mocks base method.
func (m *MockAzureClient) GetAzureADTenants(arg0 context.Context, arg1 bool) (azure.TenantList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADTenants", arg0, arg1)
	ret0, _ := ret[0].(azure.TenantList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADTenants indicates an expected call of GetAzureADTenants.
func (mr *MockAzureClientMockRecorder) GetAzureADTenants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADTenants", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADTenants), arg0, arg1)
}

// ListAzureADAppOwners mocks base method.
func (m *MockAzureClient) ListAzureADAppOwners(arg0 context.Context, arg1 string, arg2 query.GraphParams) <-chan client.AzureResult[json.RawMessage] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADAppOwners", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[json.RawMessage])
	return ret0
}

// ListAzureADAppOwners indicates an expected call of ListAzureADAppOwners.
func (mr *MockAzureClientMockRecorder) ListAzureADAppOwners(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADAppOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADAppOwners), arg0, arg1, arg2)
}

// ListAzureADAppRoleAssignments mocks base method.
func (m *MockAzureClient) ListAzureADAppRoleAssignments(arg0 context.Context, arg1 string, arg2 query.GraphParams) <-chan client.AzureResult[azure.AppRoleAssignment] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADAppRoleAssignments", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.AppRoleAssignment])
	return ret0
}

// ListAzureADAppRoleAssignments indicates an expected call of ListAzureADAppRoleAssignments.
func (mr *MockAzureClientMockRecorder) ListAzureADAppRoleAssignments(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADAppRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADAppRoleAssignments), arg0, arg1, arg2)
}

// ListAzureADApps mocks base method.
func (m *MockAzureClient) ListAzureADApps(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.Application] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADApps", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.Application])
	return ret0
}

// ListAzureADApps indicates an expected call of ListAzureADApps.
func (mr *MockAzureClientMockRecorder) ListAzureADApps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADApps), arg0, arg1)
}

// ListAzureADGroupMembers mocks base method.
func (m *MockAzureClient) ListAzureADGroupMembers(arg0 context.Context, arg1 string, arg2 query.GraphParams) <-chan client.AzureResult[json.RawMessage] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADGroupMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[json.RawMessage])
	return ret0
}

// ListAzureADGroupMembers indicates an expected call of ListAzureADGroupMembers.
func (mr *MockAzureClientMockRecorder) ListAzureADGroupMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADGroupMembers", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADGroupMembers), arg0, arg1, arg2)
}

// ListAzureADGroupOwners mocks base method.
func (m *MockAzureClient) ListAzureADGroupOwners(arg0 context.Context, arg1 string, arg2 query.GraphParams) <-chan client.AzureResult[json.RawMessage] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADGroupOwners", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[json.RawMessage])
	return ret0
}

// ListAzureADGroupOwners indicates an expected call of ListAzureADGroupOwners.
func (mr *MockAzureClientMockRecorder) ListAzureADGroupOwners(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADGroupOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADGroupOwners), arg0, arg1, arg2)
}

// ListAzureADGroups mocks base method.
func (m *MockAzureClient) ListAzureADGroups(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.Group] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADGroups", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.Group])
	return ret0
}

// ListAzureADGroups indicates an expected call of ListAzureADGroups.
func (mr *MockAzureClientMockRecorder) ListAzureADGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADGroups", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADGroups), arg0, arg1)
}

// ListAzureADRoleAssignments mocks base method.
func (m *MockAzureClient) ListAzureADRoleAssignments(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.UnifiedRoleAssignment] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADRoleAssignments", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.UnifiedRoleAssignment])
	return ret0
}

// ListAzureADRoleAssignments indicates an expected call of ListAzureADRoleAssignments.
func (mr *MockAzureClientMockRecorder) ListAzureADRoleAssignments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADRoleAssignments), arg0, arg1)
}

// ListAzureADRoles mocks base method.
func (m *MockAzureClient) ListAzureADRoles(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.Role] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADRoles", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.Role])
	return ret0
}

// ListAzureADRoles indicates an expected call of ListAzureADRoles.
func (mr *MockAzureClientMockRecorder) ListAzureADRoles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADRoles", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADRoles), arg0, arg1)
}

// ListAzureADServicePrincipalOwners mocks base method.
func (m *MockAzureClient) ListAzureADServicePrincipalOwners(arg0 context.Context, arg1 string, arg2 query.GraphParams) <-chan client.AzureResult[json.RawMessage] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADServicePrincipalOwners", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[json.RawMessage])
	return ret0
}

// ListAzureADServicePrincipalOwners indicates an expected call of ListAzureADServicePrincipalOwners.
func (mr *MockAzureClientMockRecorder) ListAzureADServicePrincipalOwners(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADServicePrincipalOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADServicePrincipalOwners), arg0, arg1, arg2)
}

// ListAzureADServicePrincipals mocks base method.
func (m *MockAzureClient) ListAzureADServicePrincipals(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.ServicePrincipal] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADServicePrincipals", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.ServicePrincipal])
	return ret0
}

// ListAzureADServicePrincipals indicates an expected call of ListAzureADServicePrincipals.
func (mr *MockAzureClientMockRecorder) ListAzureADServicePrincipals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADServicePrincipals", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADServicePrincipals), arg0, arg1)
}

// ListAzureADTenants mocks base method.
func (m *MockAzureClient) ListAzureADTenants(arg0 context.Context, arg1 bool) <-chan client.AzureResult[azure.Tenant] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADTenants", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.Tenant])
	return ret0
}

// ListAzureADTenants indicates an expected call of ListAzureADTenants.
func (mr *MockAzureClientMockRecorder) ListAzureADTenants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADTenants", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADTenants), arg0, arg1)
}

// ListAzureADUsers mocks base method.
func (m *MockAzureClient) ListAzureADUsers(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.User] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADUsers", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.User])
	return ret0
}

// ListAzureADUsers indicates an expected call of ListAzureADUsers.
func (mr *MockAzureClientMockRecorder) ListAzureADUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADUsers", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADUsers), arg0, arg1)
}

// ListAzureAutomationAccounts mocks base method.
func (m *MockAzureClient) ListAzureAutomationAccounts(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.AutomationAccount] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureAutomationAccounts", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.AutomationAccount])
	return ret0
}

// ListAzureAutomationAccounts indicates an expected call of ListAzureAutomationAccounts.
func (mr *MockAzureClientMockRecorder) ListAzureAutomationAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureAutomationAccounts", reflect.TypeOf((*MockAzureClient)(nil).ListAzureAutomationAccounts), arg0, arg1)
}

// ListAzureContainerRegistries mocks base method.
func (m *MockAzureClient) ListAzureContainerRegistries(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.ContainerRegistry] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureContainerRegistries", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.ContainerRegistry])
	return ret0
}

// ListAzureContainerRegistries indicates an expected call of ListAzureContainerRegistries.
func (mr *MockAzureClientMockRecorder) ListAzureContainerRegistries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureContainerRegistries", reflect.TypeOf((*MockAzureClient)(nil).ListAzureContainerRegistries), arg0, arg1)
}

// ListAzureDeviceRegisteredOwners mocks base method.
func (m *MockAzureClient) ListAzureDeviceRegisteredOwners(arg0 context.Context, arg1 string, arg2 query.GraphParams) <-chan client.AzureResult[json.RawMessage] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureDeviceRegisteredOwners", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[json.RawMessage])
	return ret0
}

// ListAzureDeviceRegisteredOwners indicates an expected call of ListAzureDeviceRegisteredOwners.
func (mr *MockAzureClientMockRecorder) ListAzureDeviceRegisteredOwners(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureDeviceRegisteredOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureDeviceRegisteredOwners), arg0, arg1, arg2)
}

// ListAzureDevices mocks base method.
func (m *MockAzureClient) ListAzureDevices(arg0 context.Context, arg1 query.GraphParams) <-chan client.AzureResult[azure.Device] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureDevices", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.Device])
	return ret0
}

// ListAzureDevices indicates an expected call of ListAzureDevices.
func (mr *MockAzureClientMockRecorder) ListAzureDevices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureDevices", reflect.TypeOf((*MockAzureClient)(nil).ListAzureDevices), arg0, arg1)
}

// ListAzureFunctionApps mocks base method.
func (m *MockAzureClient) ListAzureFunctionApps(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.FunctionApp] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureFunctionApps", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.FunctionApp])
	return ret0
}

// ListAzureFunctionApps indicates an expected call of ListAzureFunctionApps.
func (mr *MockAzureClientMockRecorder) ListAzureFunctionApps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureFunctionApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureFunctionApps), arg0, arg1)
}

// ListAzureKeyVaults mocks base method.
func (m *MockAzureClient) ListAzureKeyVaults(arg0 context.Context, arg1 string, arg2 query.RMParams) <-chan client.AzureResult[azure.KeyVault] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureKeyVaults", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.KeyVault])
	return ret0
}

// ListAzureKeyVaults indicates an expected call of ListAzureKeyVaults.
func (mr *MockAzureClientMockRecorder) ListAzureKeyVaults(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureKeyVaults", reflect.TypeOf((*MockAzureClient)(nil).ListAzureKeyVaults), arg0, arg1, arg2)
}

// ListAzureLogicApps mocks base method.
func (m *MockAzureClient) ListAzureLogicApps(arg0 context.Context, arg1, arg2 string, arg3 int32) <-chan client.AzureResult[azure.LogicApp] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureLogicApps", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.LogicApp])
	return ret0
}

// ListAzureLogicApps indicates an expected call of ListAzureLogicApps.
func (mr *MockAzureClientMockRecorder) ListAzureLogicApps(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureLogicApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureLogicApps), arg0, arg1, arg2, arg3)
}

// ListAzureManagedClusters mocks base method.
func (m *MockAzureClient) ListAzureManagedClusters(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.ManagedCluster] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureManagedClusters", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.ManagedCluster])
	return ret0
}

// ListAzureManagedClusters indicates an expected call of ListAzureManagedClusters.
func (mr *MockAzureClientMockRecorder) ListAzureManagedClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureManagedClusters", reflect.TypeOf((*MockAzureClient)(nil).ListAzureManagedClusters), arg0, arg1)
}

// ListAzureManagementGroupDescendants mocks base method.
func (m *MockAzureClient) ListAzureManagementGroupDescendants(arg0 context.Context, arg1 string, arg2 int32) <-chan client.AzureResult[azure.DescendantInfo] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureManagementGroupDescendants", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.DescendantInfo])
	return ret0
}

// ListAzureManagementGroupDescendants indicates an expected call of ListAzureManagementGroupDescendants.
func (mr *MockAzureClientMockRecorder) ListAzureManagementGroupDescendants(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureManagementGroupDescendants", reflect.TypeOf((*MockAzureClient)(nil).ListAzureManagementGroupDescendants), arg0, arg1, arg2)
}

// ListAzureManagementGroups mocks base method.
func (m *MockAzureClient) ListAzureManagementGroups(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.ManagementGroup] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureManagementGroups", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.ManagementGroup])
	return ret0
}

// ListAzureManagementGroups indicates an expected call of ListAzureManagementGroups.
func (mr *MockAzureClientMockRecorder) ListAzureManagementGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureManagementGroups", reflect.TypeOf((*MockAzureClient)(nil).ListAzureManagementGroups), arg0, arg1)
}

// ListAzureRBACRoleDefinitions mocks base method.
func (m *MockAzureClient) ListAzureRBACRoleDefinitions(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.RBACRoleDefinition] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureRBACRoleDefinitions", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.RBACRoleDefinition])
	return ret0
}

// ListAzureRBACRoleDefinitions indicates an expected call of ListAzureRBACRoleDefinitions.
func (mr *MockAzureClientMockRecorder) ListAzureRBACRoleDefinitions(arg0, arg1 interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureRBACRoleDefinitions", reflect.TypeOf((*MockAzureClient)(nil).ListAzureRBACRoleDefinitions), arg0, arg1)
}

// ListAzureResourceGroups mocks base method.
func (m *MockAzureClient) ListAzureResourceGroups(arg0 context.Context, arg1 string, arg2 query.RMParams) <-chan client.AzureResult[azure.ResourceGroup] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureResourceGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.ResourceGroup])
	return ret0
}

// ListAzureResourceGroups indicates an expected call of ListAzureResourceGroups.
func (mr *MockAzureClientMockRecorder) ListAzureResourceGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureResourceGroups", reflect.TypeOf((*MockAzureClient)(nil).ListAzureResourceGroups), arg0, arg1, arg2)
}

// ListAzureSpringApps mocks base method.
func (m *MockAzureClient) ListAzureSpringApps(arg0 context.Context, arg1 string, arg2 string, arg3 string) <-chan client.AzureResult[azure.SpringApp] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureSpringApps", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.SpringApp])
	return ret0
}

// ListAzureSpringApps indicates an expected call of ListAzureSpringApps.
func (mr *MockAzureClientMockRecorder) ListAzureSpringApps(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureSpringApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureSpringApps), arg0, arg1, arg2, arg3)
}

// ListAzureSpringAppServices mocks base method.
func (m *MockAzureClient) ListAzureSpringAppServices(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.SpringAppService] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureSpringAppServices", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.SpringAppService])
	return ret0
}

// ListAzureSpringAppServices indicates an expected call of ListAzureSpringApps.
func (mr *MockAzureClientMockRecorder) ListAzureSpringAppServices(arg0, arg1 interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureSpringAppServices", reflect.TypeOf((*MockAzureClient)(nil).ListAzureSpringAppServices), arg0, arg1)
}

// ListAzureStorageAccounts mocks base method.
func (m *MockAzureClient) ListAzureStorageAccounts(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.StorageAccount] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureStorageAccounts", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.StorageAccount])
	return ret0
}

// ListAzureStorageAccounts indicates an expected call of ListAzureStorageAccounts.
func (mr *MockAzureClientMockRecorder) ListAzureStorageAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureStorageAccounts", reflect.TypeOf((*MockAzureClient)(nil).ListAzureStorageAccounts), arg0, arg1)
}

// ListAzureStorageContainers mocks base method.
func (m *MockAzureClient) ListAzureStorageContainers(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string) <-chan client.AzureResult[azure.StorageContainer] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureStorageContainers", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.StorageContainer])
	return ret0
}

// ListAzureStorageContainers indicates an expected call of ListAzureStorageContainers.
func (mr *MockAzureClientMockRecorder) ListAzureStorageContainers(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureStorageContainers", reflect.TypeOf((*MockAzureClient)(nil).ListAzureStorageContainers), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ListAzureSubscriptions mocks base method.
func (m *MockAzureClient) ListAzureSubscriptions(arg0 context.Context) <-chan client.AzureResult[azure.Subscription] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureSubscriptions", arg0)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.Subscription])
	return ret0
}

// ListAzureSubscriptions indicates an expected call of ListAzureSubscriptions.
func (mr *MockAzureClientMockRecorder) ListAzureSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureSubscriptions", reflect.TypeOf((*MockAzureClient)(nil).ListAzureSubscriptions), arg0)
}

// ListAzureVMScaleSets mocks base method.
func (m *MockAzureClient) ListAzureVMScaleSets(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.VMScaleSet] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureVMScaleSets", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.VMScaleSet])
	return ret0
}

// ListAzureVMScaleSets indicates an expected call of ListAzureVMScaleSets.
func (mr *MockAzureClientMockRecorder) ListAzureVMScaleSets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureVMScaleSets", reflect.TypeOf((*MockAzureClient)(nil).ListAzureVMScaleSets), arg0, arg1)
}

// ListAzureVirtualMachines mocks base method.
func (m *MockAzureClient) ListAzureVirtualMachines(arg0 context.Context, arg1 string, arg2 query.RMParams) <-chan client.AzureResult[azure.VirtualMachine] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureVirtualMachines", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.VirtualMachine])
	return ret0
}

// ListAzureVirtualMachines indicates an expected call of ListAzureVirtualMachines.
func (mr *MockAzureClientMockRecorder) ListAzureVirtualMachines(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureVirtualMachines", reflect.TypeOf((*MockAzureClient)(nil).ListAzureVirtualMachines), arg0, arg1, arg2)
}

// ListAzureWebApps mocks base method.
func (m *MockAzureClient) ListAzureWebApps(arg0 context.Context, arg1 string) <-chan client.AzureResult[azure.WebApp] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureWebApps", arg0, arg1)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.WebApp])
	return ret0
}

// ListAzureWebApps indicates an expected call of ListAzureWebApps.
func (mr *MockAzureClientMockRecorder) ListAzureWebApps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureWebApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureWebApps), arg0, arg1)
}

// ListRoleAssignmentsForResource mocks base method.
func (m *MockAzureClient) ListRoleAssignmentsForResource(arg0 context.Context, arg1, arg2, arg3 string) <-chan client.AzureResult[azure.RoleAssignment] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleAssignmentsForResource", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan client.AzureResult[azure.RoleAssignment])
	return ret0
}

// ListRoleAssignmentsForResource indicates an expected call of ListRoleAssignmentsForResource.
func (mr *MockAzureClientMockRecorder) ListRoleAssignmentsForResource(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleAssignmentsForResource", reflect.TypeOf((*MockAzureClient)(nil).ListRoleAssignmentsForResource), arg0, arg1, arg2, arg3)
}

// TenantInfo mocks base method.
func (m *MockAzureClient) TenantInfo() azure.Tenant {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantInfo")
	ret0, _ := ret[0].(azure.Tenant)
	return ret0
}

// TenantInfo indicates an expected call of TenantInfo.
func (mr *MockAzureClientMockRecorder) TenantInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantInfo", reflect.TypeOf((*MockAzureClient)(nil).TenantInfo))
}
