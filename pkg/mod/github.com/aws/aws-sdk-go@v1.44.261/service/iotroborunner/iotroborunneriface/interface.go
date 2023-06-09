// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotroborunneriface provides an interface to enable mocking the AWS IoT RoboRunner service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotroborunneriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotroborunner"
)

// IoTRoboRunnerAPI provides an interface to enable mocking the
// iotroborunner.IoTRoboRunner service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS IoT RoboRunner.
//	func myFunc(svc iotroborunneriface.IoTRoboRunnerAPI) bool {
//	    // Make svc.CreateDestination request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := iotroborunner.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockIoTRoboRunnerClient struct {
//	    iotroborunneriface.IoTRoboRunnerAPI
//	}
//	func (m *mockIoTRoboRunnerClient) CreateDestination(input *iotroborunner.CreateDestinationInput) (*iotroborunner.CreateDestinationOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockIoTRoboRunnerClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTRoboRunnerAPI interface {
	CreateDestination(*iotroborunner.CreateDestinationInput) (*iotroborunner.CreateDestinationOutput, error)
	CreateDestinationWithContext(aws.Context, *iotroborunner.CreateDestinationInput, ...request.Option) (*iotroborunner.CreateDestinationOutput, error)
	CreateDestinationRequest(*iotroborunner.CreateDestinationInput) (*request.Request, *iotroborunner.CreateDestinationOutput)

	CreateSite(*iotroborunner.CreateSiteInput) (*iotroborunner.CreateSiteOutput, error)
	CreateSiteWithContext(aws.Context, *iotroborunner.CreateSiteInput, ...request.Option) (*iotroborunner.CreateSiteOutput, error)
	CreateSiteRequest(*iotroborunner.CreateSiteInput) (*request.Request, *iotroborunner.CreateSiteOutput)

	CreateWorker(*iotroborunner.CreateWorkerInput) (*iotroborunner.CreateWorkerOutput, error)
	CreateWorkerWithContext(aws.Context, *iotroborunner.CreateWorkerInput, ...request.Option) (*iotroborunner.CreateWorkerOutput, error)
	CreateWorkerRequest(*iotroborunner.CreateWorkerInput) (*request.Request, *iotroborunner.CreateWorkerOutput)

	CreateWorkerFleet(*iotroborunner.CreateWorkerFleetInput) (*iotroborunner.CreateWorkerFleetOutput, error)
	CreateWorkerFleetWithContext(aws.Context, *iotroborunner.CreateWorkerFleetInput, ...request.Option) (*iotroborunner.CreateWorkerFleetOutput, error)
	CreateWorkerFleetRequest(*iotroborunner.CreateWorkerFleetInput) (*request.Request, *iotroborunner.CreateWorkerFleetOutput)

	DeleteDestination(*iotroborunner.DeleteDestinationInput) (*iotroborunner.DeleteDestinationOutput, error)
	DeleteDestinationWithContext(aws.Context, *iotroborunner.DeleteDestinationInput, ...request.Option) (*iotroborunner.DeleteDestinationOutput, error)
	DeleteDestinationRequest(*iotroborunner.DeleteDestinationInput) (*request.Request, *iotroborunner.DeleteDestinationOutput)

	DeleteSite(*iotroborunner.DeleteSiteInput) (*iotroborunner.DeleteSiteOutput, error)
	DeleteSiteWithContext(aws.Context, *iotroborunner.DeleteSiteInput, ...request.Option) (*iotroborunner.DeleteSiteOutput, error)
	DeleteSiteRequest(*iotroborunner.DeleteSiteInput) (*request.Request, *iotroborunner.DeleteSiteOutput)

	DeleteWorker(*iotroborunner.DeleteWorkerInput) (*iotroborunner.DeleteWorkerOutput, error)
	DeleteWorkerWithContext(aws.Context, *iotroborunner.DeleteWorkerInput, ...request.Option) (*iotroborunner.DeleteWorkerOutput, error)
	DeleteWorkerRequest(*iotroborunner.DeleteWorkerInput) (*request.Request, *iotroborunner.DeleteWorkerOutput)

	DeleteWorkerFleet(*iotroborunner.DeleteWorkerFleetInput) (*iotroborunner.DeleteWorkerFleetOutput, error)
	DeleteWorkerFleetWithContext(aws.Context, *iotroborunner.DeleteWorkerFleetInput, ...request.Option) (*iotroborunner.DeleteWorkerFleetOutput, error)
	DeleteWorkerFleetRequest(*iotroborunner.DeleteWorkerFleetInput) (*request.Request, *iotroborunner.DeleteWorkerFleetOutput)

	GetDestination(*iotroborunner.GetDestinationInput) (*iotroborunner.GetDestinationOutput, error)
	GetDestinationWithContext(aws.Context, *iotroborunner.GetDestinationInput, ...request.Option) (*iotroborunner.GetDestinationOutput, error)
	GetDestinationRequest(*iotroborunner.GetDestinationInput) (*request.Request, *iotroborunner.GetDestinationOutput)

	GetSite(*iotroborunner.GetSiteInput) (*iotroborunner.GetSiteOutput, error)
	GetSiteWithContext(aws.Context, *iotroborunner.GetSiteInput, ...request.Option) (*iotroborunner.GetSiteOutput, error)
	GetSiteRequest(*iotroborunner.GetSiteInput) (*request.Request, *iotroborunner.GetSiteOutput)

	GetWorker(*iotroborunner.GetWorkerInput) (*iotroborunner.GetWorkerOutput, error)
	GetWorkerWithContext(aws.Context, *iotroborunner.GetWorkerInput, ...request.Option) (*iotroborunner.GetWorkerOutput, error)
	GetWorkerRequest(*iotroborunner.GetWorkerInput) (*request.Request, *iotroborunner.GetWorkerOutput)

	GetWorkerFleet(*iotroborunner.GetWorkerFleetInput) (*iotroborunner.GetWorkerFleetOutput, error)
	GetWorkerFleetWithContext(aws.Context, *iotroborunner.GetWorkerFleetInput, ...request.Option) (*iotroborunner.GetWorkerFleetOutput, error)
	GetWorkerFleetRequest(*iotroborunner.GetWorkerFleetInput) (*request.Request, *iotroborunner.GetWorkerFleetOutput)

	ListDestinations(*iotroborunner.ListDestinationsInput) (*iotroborunner.ListDestinationsOutput, error)
	ListDestinationsWithContext(aws.Context, *iotroborunner.ListDestinationsInput, ...request.Option) (*iotroborunner.ListDestinationsOutput, error)
	ListDestinationsRequest(*iotroborunner.ListDestinationsInput) (*request.Request, *iotroborunner.ListDestinationsOutput)

	ListDestinationsPages(*iotroborunner.ListDestinationsInput, func(*iotroborunner.ListDestinationsOutput, bool) bool) error
	ListDestinationsPagesWithContext(aws.Context, *iotroborunner.ListDestinationsInput, func(*iotroborunner.ListDestinationsOutput, bool) bool, ...request.Option) error

	ListSites(*iotroborunner.ListSitesInput) (*iotroborunner.ListSitesOutput, error)
	ListSitesWithContext(aws.Context, *iotroborunner.ListSitesInput, ...request.Option) (*iotroborunner.ListSitesOutput, error)
	ListSitesRequest(*iotroborunner.ListSitesInput) (*request.Request, *iotroborunner.ListSitesOutput)

	ListSitesPages(*iotroborunner.ListSitesInput, func(*iotroborunner.ListSitesOutput, bool) bool) error
	ListSitesPagesWithContext(aws.Context, *iotroborunner.ListSitesInput, func(*iotroborunner.ListSitesOutput, bool) bool, ...request.Option) error

	ListWorkerFleets(*iotroborunner.ListWorkerFleetsInput) (*iotroborunner.ListWorkerFleetsOutput, error)
	ListWorkerFleetsWithContext(aws.Context, *iotroborunner.ListWorkerFleetsInput, ...request.Option) (*iotroborunner.ListWorkerFleetsOutput, error)
	ListWorkerFleetsRequest(*iotroborunner.ListWorkerFleetsInput) (*request.Request, *iotroborunner.ListWorkerFleetsOutput)

	ListWorkerFleetsPages(*iotroborunner.ListWorkerFleetsInput, func(*iotroborunner.ListWorkerFleetsOutput, bool) bool) error
	ListWorkerFleetsPagesWithContext(aws.Context, *iotroborunner.ListWorkerFleetsInput, func(*iotroborunner.ListWorkerFleetsOutput, bool) bool, ...request.Option) error

	ListWorkers(*iotroborunner.ListWorkersInput) (*iotroborunner.ListWorkersOutput, error)
	ListWorkersWithContext(aws.Context, *iotroborunner.ListWorkersInput, ...request.Option) (*iotroborunner.ListWorkersOutput, error)
	ListWorkersRequest(*iotroborunner.ListWorkersInput) (*request.Request, *iotroborunner.ListWorkersOutput)

	ListWorkersPages(*iotroborunner.ListWorkersInput, func(*iotroborunner.ListWorkersOutput, bool) bool) error
	ListWorkersPagesWithContext(aws.Context, *iotroborunner.ListWorkersInput, func(*iotroborunner.ListWorkersOutput, bool) bool, ...request.Option) error

	UpdateDestination(*iotroborunner.UpdateDestinationInput) (*iotroborunner.UpdateDestinationOutput, error)
	UpdateDestinationWithContext(aws.Context, *iotroborunner.UpdateDestinationInput, ...request.Option) (*iotroborunner.UpdateDestinationOutput, error)
	UpdateDestinationRequest(*iotroborunner.UpdateDestinationInput) (*request.Request, *iotroborunner.UpdateDestinationOutput)

	UpdateSite(*iotroborunner.UpdateSiteInput) (*iotroborunner.UpdateSiteOutput, error)
	UpdateSiteWithContext(aws.Context, *iotroborunner.UpdateSiteInput, ...request.Option) (*iotroborunner.UpdateSiteOutput, error)
	UpdateSiteRequest(*iotroborunner.UpdateSiteInput) (*request.Request, *iotroborunner.UpdateSiteOutput)

	UpdateWorker(*iotroborunner.UpdateWorkerInput) (*iotroborunner.UpdateWorkerOutput, error)
	UpdateWorkerWithContext(aws.Context, *iotroborunner.UpdateWorkerInput, ...request.Option) (*iotroborunner.UpdateWorkerOutput, error)
	UpdateWorkerRequest(*iotroborunner.UpdateWorkerInput) (*request.Request, *iotroborunner.UpdateWorkerOutput)

	UpdateWorkerFleet(*iotroborunner.UpdateWorkerFleetInput) (*iotroborunner.UpdateWorkerFleetOutput, error)
	UpdateWorkerFleetWithContext(aws.Context, *iotroborunner.UpdateWorkerFleetInput, ...request.Option) (*iotroborunner.UpdateWorkerFleetOutput, error)
	UpdateWorkerFleetRequest(*iotroborunner.UpdateWorkerFleetInput) (*request.Request, *iotroborunner.UpdateWorkerFleetOutput)
}

var _ IoTRoboRunnerAPI = (*iotroborunner.IoTRoboRunner)(nil)
