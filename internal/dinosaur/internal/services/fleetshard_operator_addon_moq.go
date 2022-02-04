// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"github.com/bf2fc6cc711aee1a0c2a/fleet-manager/pkg/api"
	serviceError "github.com/bf2fc6cc711aee1a0c2a/fleet-manager/pkg/errors"
	"sync"
)

// Ensure, that FleetshardOperatorAddonMock does implement FleetshardOperatorAddon.
// If this is not the case, regenerate this file with moq.
var _ FleetshardOperatorAddon = &FleetshardOperatorAddonMock{}

// FleetshardOperatorAddonMock is a mock implementation of FleetshardOperatorAddon.
//
// 	func TestSomethingThatUsesFleetshardOperatorAddon(t *testing.T) {
//
// 		// make and configure a mocked FleetshardOperatorAddon
// 		mockedFleetshardOperatorAddon := &FleetshardOperatorAddonMock{
// 			ProvisionFunc: func(cluster api.Cluster) (bool, *serviceError.ServiceError) {
// 				panic("mock out the Provision method")
// 			},
// 			ReconcileParametersFunc: func(cluster api.Cluster) *serviceError.ServiceError {
// 				panic("mock out the ReconcileParameters method")
// 			},
// 			RemoveServiceAccountFunc: func(cluster api.Cluster) *serviceError.ServiceError {
// 				panic("mock out the RemoveServiceAccount method")
// 			},
// 		}
//
// 		// use mockedFleetshardOperatorAddon in code that requires FleetshardOperatorAddon
// 		// and then make assertions.
//
// 	}
type FleetshardOperatorAddonMock struct {
	// ProvisionFunc mocks the Provision method.
	ProvisionFunc func(cluster api.Cluster) (bool, *serviceError.ServiceError)

	// ReconcileParametersFunc mocks the ReconcileParameters method.
	ReconcileParametersFunc func(cluster api.Cluster) *serviceError.ServiceError

	// RemoveServiceAccountFunc mocks the RemoveServiceAccount method.
	RemoveServiceAccountFunc func(cluster api.Cluster) *serviceError.ServiceError

	// calls tracks calls to the methods.
	calls struct {
		// Provision holds details about calls to the Provision method.
		Provision []struct {
			// Cluster is the cluster argument value.
			Cluster api.Cluster
		}
		// ReconcileParameters holds details about calls to the ReconcileParameters method.
		ReconcileParameters []struct {
			// Cluster is the cluster argument value.
			Cluster api.Cluster
		}
		// RemoveServiceAccount holds details about calls to the RemoveServiceAccount method.
		RemoveServiceAccount []struct {
			// Cluster is the cluster argument value.
			Cluster api.Cluster
		}
	}
	lockProvision            sync.RWMutex
	lockReconcileParameters  sync.RWMutex
	lockRemoveServiceAccount sync.RWMutex
}

// Provision calls ProvisionFunc.
func (mock *FleetshardOperatorAddonMock) Provision(cluster api.Cluster) (bool, *serviceError.ServiceError) {
	if mock.ProvisionFunc == nil {
		panic("FleetshardOperatorAddonMock.ProvisionFunc: method is nil but FleetshardOperatorAddon.Provision was just called")
	}
	callInfo := struct {
		Cluster api.Cluster
	}{
		Cluster: cluster,
	}
	mock.lockProvision.Lock()
	mock.calls.Provision = append(mock.calls.Provision, callInfo)
	mock.lockProvision.Unlock()
	return mock.ProvisionFunc(cluster)
}

// ProvisionCalls gets all the calls that were made to Provision.
// Check the length with:
//     len(mockedFleetshardOperatorAddon.ProvisionCalls())
func (mock *FleetshardOperatorAddonMock) ProvisionCalls() []struct {
	Cluster api.Cluster
} {
	var calls []struct {
		Cluster api.Cluster
	}
	mock.lockProvision.RLock()
	calls = mock.calls.Provision
	mock.lockProvision.RUnlock()
	return calls
}

// ReconcileParameters calls ReconcileParametersFunc.
func (mock *FleetshardOperatorAddonMock) ReconcileParameters(cluster api.Cluster) *serviceError.ServiceError {
	if mock.ReconcileParametersFunc == nil {
		panic("FleetshardOperatorAddonMock.ReconcileParametersFunc: method is nil but FleetshardOperatorAddon.ReconcileParameters was just called")
	}
	callInfo := struct {
		Cluster api.Cluster
	}{
		Cluster: cluster,
	}
	mock.lockReconcileParameters.Lock()
	mock.calls.ReconcileParameters = append(mock.calls.ReconcileParameters, callInfo)
	mock.lockReconcileParameters.Unlock()
	return mock.ReconcileParametersFunc(cluster)
}

// ReconcileParametersCalls gets all the calls that were made to ReconcileParameters.
// Check the length with:
//     len(mockedFleetshardOperatorAddon.ReconcileParametersCalls())
func (mock *FleetshardOperatorAddonMock) ReconcileParametersCalls() []struct {
	Cluster api.Cluster
} {
	var calls []struct {
		Cluster api.Cluster
	}
	mock.lockReconcileParameters.RLock()
	calls = mock.calls.ReconcileParameters
	mock.lockReconcileParameters.RUnlock()
	return calls
}

// RemoveServiceAccount calls RemoveServiceAccountFunc.
func (mock *FleetshardOperatorAddonMock) RemoveServiceAccount(cluster api.Cluster) *serviceError.ServiceError {
	if mock.RemoveServiceAccountFunc == nil {
		panic("FleetshardOperatorAddonMock.RemoveServiceAccountFunc: method is nil but FleetshardOperatorAddon.RemoveServiceAccount was just called")
	}
	callInfo := struct {
		Cluster api.Cluster
	}{
		Cluster: cluster,
	}
	mock.lockRemoveServiceAccount.Lock()
	mock.calls.RemoveServiceAccount = append(mock.calls.RemoveServiceAccount, callInfo)
	mock.lockRemoveServiceAccount.Unlock()
	return mock.RemoveServiceAccountFunc(cluster)
}

// RemoveServiceAccountCalls gets all the calls that were made to RemoveServiceAccount.
// Check the length with:
//     len(mockedFleetshardOperatorAddon.RemoveServiceAccountCalls())
func (mock *FleetshardOperatorAddonMock) RemoveServiceAccountCalls() []struct {
	Cluster api.Cluster
} {
	var calls []struct {
		Cluster api.Cluster
	}
	mock.lockRemoveServiceAccount.RLock()
	calls = mock.calls.RemoveServiceAccount
	mock.lockRemoveServiceAccount.RUnlock()
	return calls
}
