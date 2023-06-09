// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sesv2iface provides an interface to enable mocking the Amazon Simple Email Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sesv2iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sesv2"
)

// SESV2API provides an interface to enable mocking the
// sesv2.SESV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Simple Email Service.
//	func myFunc(svc sesv2iface.SESV2API) bool {
//	    // Make svc.BatchGetMetricData request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := sesv2.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSESV2Client struct {
//	    sesv2iface.SESV2API
//	}
//	func (m *mockSESV2Client) BatchGetMetricData(input *sesv2.BatchGetMetricDataInput) (*sesv2.BatchGetMetricDataOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSESV2Client{}
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
type SESV2API interface {
	BatchGetMetricData(*sesv2.BatchGetMetricDataInput) (*sesv2.BatchGetMetricDataOutput, error)
	BatchGetMetricDataWithContext(aws.Context, *sesv2.BatchGetMetricDataInput, ...request.Option) (*sesv2.BatchGetMetricDataOutput, error)
	BatchGetMetricDataRequest(*sesv2.BatchGetMetricDataInput) (*request.Request, *sesv2.BatchGetMetricDataOutput)

	CreateConfigurationSet(*sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error)
	CreateConfigurationSetWithContext(aws.Context, *sesv2.CreateConfigurationSetInput, ...request.Option) (*sesv2.CreateConfigurationSetOutput, error)
	CreateConfigurationSetRequest(*sesv2.CreateConfigurationSetInput) (*request.Request, *sesv2.CreateConfigurationSetOutput)

	CreateConfigurationSetEventDestination(*sesv2.CreateConfigurationSetEventDestinationInput) (*sesv2.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationWithContext(aws.Context, *sesv2.CreateConfigurationSetEventDestinationInput, ...request.Option) (*sesv2.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationRequest(*sesv2.CreateConfigurationSetEventDestinationInput) (*request.Request, *sesv2.CreateConfigurationSetEventDestinationOutput)

	CreateContact(*sesv2.CreateContactInput) (*sesv2.CreateContactOutput, error)
	CreateContactWithContext(aws.Context, *sesv2.CreateContactInput, ...request.Option) (*sesv2.CreateContactOutput, error)
	CreateContactRequest(*sesv2.CreateContactInput) (*request.Request, *sesv2.CreateContactOutput)

	CreateContactList(*sesv2.CreateContactListInput) (*sesv2.CreateContactListOutput, error)
	CreateContactListWithContext(aws.Context, *sesv2.CreateContactListInput, ...request.Option) (*sesv2.CreateContactListOutput, error)
	CreateContactListRequest(*sesv2.CreateContactListInput) (*request.Request, *sesv2.CreateContactListOutput)

	CreateCustomVerificationEmailTemplate(*sesv2.CreateCustomVerificationEmailTemplateInput) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error)
	CreateCustomVerificationEmailTemplateWithContext(aws.Context, *sesv2.CreateCustomVerificationEmailTemplateInput, ...request.Option) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error)
	CreateCustomVerificationEmailTemplateRequest(*sesv2.CreateCustomVerificationEmailTemplateInput) (*request.Request, *sesv2.CreateCustomVerificationEmailTemplateOutput)

	CreateDedicatedIpPool(*sesv2.CreateDedicatedIpPoolInput) (*sesv2.CreateDedicatedIpPoolOutput, error)
	CreateDedicatedIpPoolWithContext(aws.Context, *sesv2.CreateDedicatedIpPoolInput, ...request.Option) (*sesv2.CreateDedicatedIpPoolOutput, error)
	CreateDedicatedIpPoolRequest(*sesv2.CreateDedicatedIpPoolInput) (*request.Request, *sesv2.CreateDedicatedIpPoolOutput)

	CreateDeliverabilityTestReport(*sesv2.CreateDeliverabilityTestReportInput) (*sesv2.CreateDeliverabilityTestReportOutput, error)
	CreateDeliverabilityTestReportWithContext(aws.Context, *sesv2.CreateDeliverabilityTestReportInput, ...request.Option) (*sesv2.CreateDeliverabilityTestReportOutput, error)
	CreateDeliverabilityTestReportRequest(*sesv2.CreateDeliverabilityTestReportInput) (*request.Request, *sesv2.CreateDeliverabilityTestReportOutput)

	CreateEmailIdentity(*sesv2.CreateEmailIdentityInput) (*sesv2.CreateEmailIdentityOutput, error)
	CreateEmailIdentityWithContext(aws.Context, *sesv2.CreateEmailIdentityInput, ...request.Option) (*sesv2.CreateEmailIdentityOutput, error)
	CreateEmailIdentityRequest(*sesv2.CreateEmailIdentityInput) (*request.Request, *sesv2.CreateEmailIdentityOutput)

	CreateEmailIdentityPolicy(*sesv2.CreateEmailIdentityPolicyInput) (*sesv2.CreateEmailIdentityPolicyOutput, error)
	CreateEmailIdentityPolicyWithContext(aws.Context, *sesv2.CreateEmailIdentityPolicyInput, ...request.Option) (*sesv2.CreateEmailIdentityPolicyOutput, error)
	CreateEmailIdentityPolicyRequest(*sesv2.CreateEmailIdentityPolicyInput) (*request.Request, *sesv2.CreateEmailIdentityPolicyOutput)

	CreateEmailTemplate(*sesv2.CreateEmailTemplateInput) (*sesv2.CreateEmailTemplateOutput, error)
	CreateEmailTemplateWithContext(aws.Context, *sesv2.CreateEmailTemplateInput, ...request.Option) (*sesv2.CreateEmailTemplateOutput, error)
	CreateEmailTemplateRequest(*sesv2.CreateEmailTemplateInput) (*request.Request, *sesv2.CreateEmailTemplateOutput)

	CreateImportJob(*sesv2.CreateImportJobInput) (*sesv2.CreateImportJobOutput, error)
	CreateImportJobWithContext(aws.Context, *sesv2.CreateImportJobInput, ...request.Option) (*sesv2.CreateImportJobOutput, error)
	CreateImportJobRequest(*sesv2.CreateImportJobInput) (*request.Request, *sesv2.CreateImportJobOutput)

	DeleteConfigurationSet(*sesv2.DeleteConfigurationSetInput) (*sesv2.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetWithContext(aws.Context, *sesv2.DeleteConfigurationSetInput, ...request.Option) (*sesv2.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetRequest(*sesv2.DeleteConfigurationSetInput) (*request.Request, *sesv2.DeleteConfigurationSetOutput)

	DeleteConfigurationSetEventDestination(*sesv2.DeleteConfigurationSetEventDestinationInput) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationWithContext(aws.Context, *sesv2.DeleteConfigurationSetEventDestinationInput, ...request.Option) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationRequest(*sesv2.DeleteConfigurationSetEventDestinationInput) (*request.Request, *sesv2.DeleteConfigurationSetEventDestinationOutput)

	DeleteContact(*sesv2.DeleteContactInput) (*sesv2.DeleteContactOutput, error)
	DeleteContactWithContext(aws.Context, *sesv2.DeleteContactInput, ...request.Option) (*sesv2.DeleteContactOutput, error)
	DeleteContactRequest(*sesv2.DeleteContactInput) (*request.Request, *sesv2.DeleteContactOutput)

	DeleteContactList(*sesv2.DeleteContactListInput) (*sesv2.DeleteContactListOutput, error)
	DeleteContactListWithContext(aws.Context, *sesv2.DeleteContactListInput, ...request.Option) (*sesv2.DeleteContactListOutput, error)
	DeleteContactListRequest(*sesv2.DeleteContactListInput) (*request.Request, *sesv2.DeleteContactListOutput)

	DeleteCustomVerificationEmailTemplate(*sesv2.DeleteCustomVerificationEmailTemplateInput) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteCustomVerificationEmailTemplateWithContext(aws.Context, *sesv2.DeleteCustomVerificationEmailTemplateInput, ...request.Option) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteCustomVerificationEmailTemplateRequest(*sesv2.DeleteCustomVerificationEmailTemplateInput) (*request.Request, *sesv2.DeleteCustomVerificationEmailTemplateOutput)

	DeleteDedicatedIpPool(*sesv2.DeleteDedicatedIpPoolInput) (*sesv2.DeleteDedicatedIpPoolOutput, error)
	DeleteDedicatedIpPoolWithContext(aws.Context, *sesv2.DeleteDedicatedIpPoolInput, ...request.Option) (*sesv2.DeleteDedicatedIpPoolOutput, error)
	DeleteDedicatedIpPoolRequest(*sesv2.DeleteDedicatedIpPoolInput) (*request.Request, *sesv2.DeleteDedicatedIpPoolOutput)

	DeleteEmailIdentity(*sesv2.DeleteEmailIdentityInput) (*sesv2.DeleteEmailIdentityOutput, error)
	DeleteEmailIdentityWithContext(aws.Context, *sesv2.DeleteEmailIdentityInput, ...request.Option) (*sesv2.DeleteEmailIdentityOutput, error)
	DeleteEmailIdentityRequest(*sesv2.DeleteEmailIdentityInput) (*request.Request, *sesv2.DeleteEmailIdentityOutput)

	DeleteEmailIdentityPolicy(*sesv2.DeleteEmailIdentityPolicyInput) (*sesv2.DeleteEmailIdentityPolicyOutput, error)
	DeleteEmailIdentityPolicyWithContext(aws.Context, *sesv2.DeleteEmailIdentityPolicyInput, ...request.Option) (*sesv2.DeleteEmailIdentityPolicyOutput, error)
	DeleteEmailIdentityPolicyRequest(*sesv2.DeleteEmailIdentityPolicyInput) (*request.Request, *sesv2.DeleteEmailIdentityPolicyOutput)

	DeleteEmailTemplate(*sesv2.DeleteEmailTemplateInput) (*sesv2.DeleteEmailTemplateOutput, error)
	DeleteEmailTemplateWithContext(aws.Context, *sesv2.DeleteEmailTemplateInput, ...request.Option) (*sesv2.DeleteEmailTemplateOutput, error)
	DeleteEmailTemplateRequest(*sesv2.DeleteEmailTemplateInput) (*request.Request, *sesv2.DeleteEmailTemplateOutput)

	DeleteSuppressedDestination(*sesv2.DeleteSuppressedDestinationInput) (*sesv2.DeleteSuppressedDestinationOutput, error)
	DeleteSuppressedDestinationWithContext(aws.Context, *sesv2.DeleteSuppressedDestinationInput, ...request.Option) (*sesv2.DeleteSuppressedDestinationOutput, error)
	DeleteSuppressedDestinationRequest(*sesv2.DeleteSuppressedDestinationInput) (*request.Request, *sesv2.DeleteSuppressedDestinationOutput)

	GetAccount(*sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error)
	GetAccountWithContext(aws.Context, *sesv2.GetAccountInput, ...request.Option) (*sesv2.GetAccountOutput, error)
	GetAccountRequest(*sesv2.GetAccountInput) (*request.Request, *sesv2.GetAccountOutput)

	GetBlacklistReports(*sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error)
	GetBlacklistReportsWithContext(aws.Context, *sesv2.GetBlacklistReportsInput, ...request.Option) (*sesv2.GetBlacklistReportsOutput, error)
	GetBlacklistReportsRequest(*sesv2.GetBlacklistReportsInput) (*request.Request, *sesv2.GetBlacklistReportsOutput)

	GetConfigurationSet(*sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error)
	GetConfigurationSetWithContext(aws.Context, *sesv2.GetConfigurationSetInput, ...request.Option) (*sesv2.GetConfigurationSetOutput, error)
	GetConfigurationSetRequest(*sesv2.GetConfigurationSetInput) (*request.Request, *sesv2.GetConfigurationSetOutput)

	GetConfigurationSetEventDestinations(*sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsWithContext(aws.Context, *sesv2.GetConfigurationSetEventDestinationsInput, ...request.Option) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsRequest(*sesv2.GetConfigurationSetEventDestinationsInput) (*request.Request, *sesv2.GetConfigurationSetEventDestinationsOutput)

	GetContact(*sesv2.GetContactInput) (*sesv2.GetContactOutput, error)
	GetContactWithContext(aws.Context, *sesv2.GetContactInput, ...request.Option) (*sesv2.GetContactOutput, error)
	GetContactRequest(*sesv2.GetContactInput) (*request.Request, *sesv2.GetContactOutput)

	GetContactList(*sesv2.GetContactListInput) (*sesv2.GetContactListOutput, error)
	GetContactListWithContext(aws.Context, *sesv2.GetContactListInput, ...request.Option) (*sesv2.GetContactListOutput, error)
	GetContactListRequest(*sesv2.GetContactListInput) (*request.Request, *sesv2.GetContactListOutput)

	GetCustomVerificationEmailTemplate(*sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
	GetCustomVerificationEmailTemplateWithContext(aws.Context, *sesv2.GetCustomVerificationEmailTemplateInput, ...request.Option) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
	GetCustomVerificationEmailTemplateRequest(*sesv2.GetCustomVerificationEmailTemplateInput) (*request.Request, *sesv2.GetCustomVerificationEmailTemplateOutput)

	GetDedicatedIp(*sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error)
	GetDedicatedIpWithContext(aws.Context, *sesv2.GetDedicatedIpInput, ...request.Option) (*sesv2.GetDedicatedIpOutput, error)
	GetDedicatedIpRequest(*sesv2.GetDedicatedIpInput) (*request.Request, *sesv2.GetDedicatedIpOutput)

	GetDedicatedIpPool(*sesv2.GetDedicatedIpPoolInput) (*sesv2.GetDedicatedIpPoolOutput, error)
	GetDedicatedIpPoolWithContext(aws.Context, *sesv2.GetDedicatedIpPoolInput, ...request.Option) (*sesv2.GetDedicatedIpPoolOutput, error)
	GetDedicatedIpPoolRequest(*sesv2.GetDedicatedIpPoolInput) (*request.Request, *sesv2.GetDedicatedIpPoolOutput)

	GetDedicatedIps(*sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error)
	GetDedicatedIpsWithContext(aws.Context, *sesv2.GetDedicatedIpsInput, ...request.Option) (*sesv2.GetDedicatedIpsOutput, error)
	GetDedicatedIpsRequest(*sesv2.GetDedicatedIpsInput) (*request.Request, *sesv2.GetDedicatedIpsOutput)

	GetDedicatedIpsPages(*sesv2.GetDedicatedIpsInput, func(*sesv2.GetDedicatedIpsOutput, bool) bool) error
	GetDedicatedIpsPagesWithContext(aws.Context, *sesv2.GetDedicatedIpsInput, func(*sesv2.GetDedicatedIpsOutput, bool) bool, ...request.Option) error

	GetDeliverabilityDashboardOptions(*sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityDashboardOptionsWithContext(aws.Context, *sesv2.GetDeliverabilityDashboardOptionsInput, ...request.Option) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityDashboardOptionsRequest(*sesv2.GetDeliverabilityDashboardOptionsInput) (*request.Request, *sesv2.GetDeliverabilityDashboardOptionsOutput)

	GetDeliverabilityTestReport(*sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error)
	GetDeliverabilityTestReportWithContext(aws.Context, *sesv2.GetDeliverabilityTestReportInput, ...request.Option) (*sesv2.GetDeliverabilityTestReportOutput, error)
	GetDeliverabilityTestReportRequest(*sesv2.GetDeliverabilityTestReportInput) (*request.Request, *sesv2.GetDeliverabilityTestReportOutput)

	GetDomainDeliverabilityCampaign(*sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainDeliverabilityCampaignWithContext(aws.Context, *sesv2.GetDomainDeliverabilityCampaignInput, ...request.Option) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainDeliverabilityCampaignRequest(*sesv2.GetDomainDeliverabilityCampaignInput) (*request.Request, *sesv2.GetDomainDeliverabilityCampaignOutput)

	GetDomainStatisticsReport(*sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error)
	GetDomainStatisticsReportWithContext(aws.Context, *sesv2.GetDomainStatisticsReportInput, ...request.Option) (*sesv2.GetDomainStatisticsReportOutput, error)
	GetDomainStatisticsReportRequest(*sesv2.GetDomainStatisticsReportInput) (*request.Request, *sesv2.GetDomainStatisticsReportOutput)

	GetEmailIdentity(*sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error)
	GetEmailIdentityWithContext(aws.Context, *sesv2.GetEmailIdentityInput, ...request.Option) (*sesv2.GetEmailIdentityOutput, error)
	GetEmailIdentityRequest(*sesv2.GetEmailIdentityInput) (*request.Request, *sesv2.GetEmailIdentityOutput)

	GetEmailIdentityPolicies(*sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error)
	GetEmailIdentityPoliciesWithContext(aws.Context, *sesv2.GetEmailIdentityPoliciesInput, ...request.Option) (*sesv2.GetEmailIdentityPoliciesOutput, error)
	GetEmailIdentityPoliciesRequest(*sesv2.GetEmailIdentityPoliciesInput) (*request.Request, *sesv2.GetEmailIdentityPoliciesOutput)

	GetEmailTemplate(*sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error)
	GetEmailTemplateWithContext(aws.Context, *sesv2.GetEmailTemplateInput, ...request.Option) (*sesv2.GetEmailTemplateOutput, error)
	GetEmailTemplateRequest(*sesv2.GetEmailTemplateInput) (*request.Request, *sesv2.GetEmailTemplateOutput)

	GetImportJob(*sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error)
	GetImportJobWithContext(aws.Context, *sesv2.GetImportJobInput, ...request.Option) (*sesv2.GetImportJobOutput, error)
	GetImportJobRequest(*sesv2.GetImportJobInput) (*request.Request, *sesv2.GetImportJobOutput)

	GetSuppressedDestination(*sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error)
	GetSuppressedDestinationWithContext(aws.Context, *sesv2.GetSuppressedDestinationInput, ...request.Option) (*sesv2.GetSuppressedDestinationOutput, error)
	GetSuppressedDestinationRequest(*sesv2.GetSuppressedDestinationInput) (*request.Request, *sesv2.GetSuppressedDestinationOutput)

	ListConfigurationSets(*sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error)
	ListConfigurationSetsWithContext(aws.Context, *sesv2.ListConfigurationSetsInput, ...request.Option) (*sesv2.ListConfigurationSetsOutput, error)
	ListConfigurationSetsRequest(*sesv2.ListConfigurationSetsInput) (*request.Request, *sesv2.ListConfigurationSetsOutput)

	ListConfigurationSetsPages(*sesv2.ListConfigurationSetsInput, func(*sesv2.ListConfigurationSetsOutput, bool) bool) error
	ListConfigurationSetsPagesWithContext(aws.Context, *sesv2.ListConfigurationSetsInput, func(*sesv2.ListConfigurationSetsOutput, bool) bool, ...request.Option) error

	ListContactLists(*sesv2.ListContactListsInput) (*sesv2.ListContactListsOutput, error)
	ListContactListsWithContext(aws.Context, *sesv2.ListContactListsInput, ...request.Option) (*sesv2.ListContactListsOutput, error)
	ListContactListsRequest(*sesv2.ListContactListsInput) (*request.Request, *sesv2.ListContactListsOutput)

	ListContactListsPages(*sesv2.ListContactListsInput, func(*sesv2.ListContactListsOutput, bool) bool) error
	ListContactListsPagesWithContext(aws.Context, *sesv2.ListContactListsInput, func(*sesv2.ListContactListsOutput, bool) bool, ...request.Option) error

	ListContacts(*sesv2.ListContactsInput) (*sesv2.ListContactsOutput, error)
	ListContactsWithContext(aws.Context, *sesv2.ListContactsInput, ...request.Option) (*sesv2.ListContactsOutput, error)
	ListContactsRequest(*sesv2.ListContactsInput) (*request.Request, *sesv2.ListContactsOutput)

	ListContactsPages(*sesv2.ListContactsInput, func(*sesv2.ListContactsOutput, bool) bool) error
	ListContactsPagesWithContext(aws.Context, *sesv2.ListContactsInput, func(*sesv2.ListContactsOutput, bool) bool, ...request.Option) error

	ListCustomVerificationEmailTemplates(*sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
	ListCustomVerificationEmailTemplatesWithContext(aws.Context, *sesv2.ListCustomVerificationEmailTemplatesInput, ...request.Option) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
	ListCustomVerificationEmailTemplatesRequest(*sesv2.ListCustomVerificationEmailTemplatesInput) (*request.Request, *sesv2.ListCustomVerificationEmailTemplatesOutput)

	ListCustomVerificationEmailTemplatesPages(*sesv2.ListCustomVerificationEmailTemplatesInput, func(*sesv2.ListCustomVerificationEmailTemplatesOutput, bool) bool) error
	ListCustomVerificationEmailTemplatesPagesWithContext(aws.Context, *sesv2.ListCustomVerificationEmailTemplatesInput, func(*sesv2.ListCustomVerificationEmailTemplatesOutput, bool) bool, ...request.Option) error

	ListDedicatedIpPools(*sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error)
	ListDedicatedIpPoolsWithContext(aws.Context, *sesv2.ListDedicatedIpPoolsInput, ...request.Option) (*sesv2.ListDedicatedIpPoolsOutput, error)
	ListDedicatedIpPoolsRequest(*sesv2.ListDedicatedIpPoolsInput) (*request.Request, *sesv2.ListDedicatedIpPoolsOutput)

	ListDedicatedIpPoolsPages(*sesv2.ListDedicatedIpPoolsInput, func(*sesv2.ListDedicatedIpPoolsOutput, bool) bool) error
	ListDedicatedIpPoolsPagesWithContext(aws.Context, *sesv2.ListDedicatedIpPoolsInput, func(*sesv2.ListDedicatedIpPoolsOutput, bool) bool, ...request.Option) error

	ListDeliverabilityTestReports(*sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error)
	ListDeliverabilityTestReportsWithContext(aws.Context, *sesv2.ListDeliverabilityTestReportsInput, ...request.Option) (*sesv2.ListDeliverabilityTestReportsOutput, error)
	ListDeliverabilityTestReportsRequest(*sesv2.ListDeliverabilityTestReportsInput) (*request.Request, *sesv2.ListDeliverabilityTestReportsOutput)

	ListDeliverabilityTestReportsPages(*sesv2.ListDeliverabilityTestReportsInput, func(*sesv2.ListDeliverabilityTestReportsOutput, bool) bool) error
	ListDeliverabilityTestReportsPagesWithContext(aws.Context, *sesv2.ListDeliverabilityTestReportsInput, func(*sesv2.ListDeliverabilityTestReportsOutput, bool) bool, ...request.Option) error

	ListDomainDeliverabilityCampaigns(*sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
	ListDomainDeliverabilityCampaignsWithContext(aws.Context, *sesv2.ListDomainDeliverabilityCampaignsInput, ...request.Option) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
	ListDomainDeliverabilityCampaignsRequest(*sesv2.ListDomainDeliverabilityCampaignsInput) (*request.Request, *sesv2.ListDomainDeliverabilityCampaignsOutput)

	ListDomainDeliverabilityCampaignsPages(*sesv2.ListDomainDeliverabilityCampaignsInput, func(*sesv2.ListDomainDeliverabilityCampaignsOutput, bool) bool) error
	ListDomainDeliverabilityCampaignsPagesWithContext(aws.Context, *sesv2.ListDomainDeliverabilityCampaignsInput, func(*sesv2.ListDomainDeliverabilityCampaignsOutput, bool) bool, ...request.Option) error

	ListEmailIdentities(*sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error)
	ListEmailIdentitiesWithContext(aws.Context, *sesv2.ListEmailIdentitiesInput, ...request.Option) (*sesv2.ListEmailIdentitiesOutput, error)
	ListEmailIdentitiesRequest(*sesv2.ListEmailIdentitiesInput) (*request.Request, *sesv2.ListEmailIdentitiesOutput)

	ListEmailIdentitiesPages(*sesv2.ListEmailIdentitiesInput, func(*sesv2.ListEmailIdentitiesOutput, bool) bool) error
	ListEmailIdentitiesPagesWithContext(aws.Context, *sesv2.ListEmailIdentitiesInput, func(*sesv2.ListEmailIdentitiesOutput, bool) bool, ...request.Option) error

	ListEmailTemplates(*sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error)
	ListEmailTemplatesWithContext(aws.Context, *sesv2.ListEmailTemplatesInput, ...request.Option) (*sesv2.ListEmailTemplatesOutput, error)
	ListEmailTemplatesRequest(*sesv2.ListEmailTemplatesInput) (*request.Request, *sesv2.ListEmailTemplatesOutput)

	ListEmailTemplatesPages(*sesv2.ListEmailTemplatesInput, func(*sesv2.ListEmailTemplatesOutput, bool) bool) error
	ListEmailTemplatesPagesWithContext(aws.Context, *sesv2.ListEmailTemplatesInput, func(*sesv2.ListEmailTemplatesOutput, bool) bool, ...request.Option) error

	ListImportJobs(*sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error)
	ListImportJobsWithContext(aws.Context, *sesv2.ListImportJobsInput, ...request.Option) (*sesv2.ListImportJobsOutput, error)
	ListImportJobsRequest(*sesv2.ListImportJobsInput) (*request.Request, *sesv2.ListImportJobsOutput)

	ListImportJobsPages(*sesv2.ListImportJobsInput, func(*sesv2.ListImportJobsOutput, bool) bool) error
	ListImportJobsPagesWithContext(aws.Context, *sesv2.ListImportJobsInput, func(*sesv2.ListImportJobsOutput, bool) bool, ...request.Option) error

	ListRecommendations(*sesv2.ListRecommendationsInput) (*sesv2.ListRecommendationsOutput, error)
	ListRecommendationsWithContext(aws.Context, *sesv2.ListRecommendationsInput, ...request.Option) (*sesv2.ListRecommendationsOutput, error)
	ListRecommendationsRequest(*sesv2.ListRecommendationsInput) (*request.Request, *sesv2.ListRecommendationsOutput)

	ListRecommendationsPages(*sesv2.ListRecommendationsInput, func(*sesv2.ListRecommendationsOutput, bool) bool) error
	ListRecommendationsPagesWithContext(aws.Context, *sesv2.ListRecommendationsInput, func(*sesv2.ListRecommendationsOutput, bool) bool, ...request.Option) error

	ListSuppressedDestinations(*sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error)
	ListSuppressedDestinationsWithContext(aws.Context, *sesv2.ListSuppressedDestinationsInput, ...request.Option) (*sesv2.ListSuppressedDestinationsOutput, error)
	ListSuppressedDestinationsRequest(*sesv2.ListSuppressedDestinationsInput) (*request.Request, *sesv2.ListSuppressedDestinationsOutput)

	ListSuppressedDestinationsPages(*sesv2.ListSuppressedDestinationsInput, func(*sesv2.ListSuppressedDestinationsOutput, bool) bool) error
	ListSuppressedDestinationsPagesWithContext(aws.Context, *sesv2.ListSuppressedDestinationsInput, func(*sesv2.ListSuppressedDestinationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *sesv2.ListTagsForResourceInput, ...request.Option) (*sesv2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*sesv2.ListTagsForResourceInput) (*request.Request, *sesv2.ListTagsForResourceOutput)

	PutAccountDedicatedIpWarmupAttributes(*sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error)
	PutAccountDedicatedIpWarmupAttributesWithContext(aws.Context, *sesv2.PutAccountDedicatedIpWarmupAttributesInput, ...request.Option) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error)
	PutAccountDedicatedIpWarmupAttributesRequest(*sesv2.PutAccountDedicatedIpWarmupAttributesInput) (*request.Request, *sesv2.PutAccountDedicatedIpWarmupAttributesOutput)

	PutAccountDetails(*sesv2.PutAccountDetailsInput) (*sesv2.PutAccountDetailsOutput, error)
	PutAccountDetailsWithContext(aws.Context, *sesv2.PutAccountDetailsInput, ...request.Option) (*sesv2.PutAccountDetailsOutput, error)
	PutAccountDetailsRequest(*sesv2.PutAccountDetailsInput) (*request.Request, *sesv2.PutAccountDetailsOutput)

	PutAccountSendingAttributes(*sesv2.PutAccountSendingAttributesInput) (*sesv2.PutAccountSendingAttributesOutput, error)
	PutAccountSendingAttributesWithContext(aws.Context, *sesv2.PutAccountSendingAttributesInput, ...request.Option) (*sesv2.PutAccountSendingAttributesOutput, error)
	PutAccountSendingAttributesRequest(*sesv2.PutAccountSendingAttributesInput) (*request.Request, *sesv2.PutAccountSendingAttributesOutput)

	PutAccountSuppressionAttributes(*sesv2.PutAccountSuppressionAttributesInput) (*sesv2.PutAccountSuppressionAttributesOutput, error)
	PutAccountSuppressionAttributesWithContext(aws.Context, *sesv2.PutAccountSuppressionAttributesInput, ...request.Option) (*sesv2.PutAccountSuppressionAttributesOutput, error)
	PutAccountSuppressionAttributesRequest(*sesv2.PutAccountSuppressionAttributesInput) (*request.Request, *sesv2.PutAccountSuppressionAttributesOutput)

	PutAccountVdmAttributes(*sesv2.PutAccountVdmAttributesInput) (*sesv2.PutAccountVdmAttributesOutput, error)
	PutAccountVdmAttributesWithContext(aws.Context, *sesv2.PutAccountVdmAttributesInput, ...request.Option) (*sesv2.PutAccountVdmAttributesOutput, error)
	PutAccountVdmAttributesRequest(*sesv2.PutAccountVdmAttributesInput) (*request.Request, *sesv2.PutAccountVdmAttributesOutput)

	PutConfigurationSetDeliveryOptions(*sesv2.PutConfigurationSetDeliveryOptionsInput) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error)
	PutConfigurationSetDeliveryOptionsWithContext(aws.Context, *sesv2.PutConfigurationSetDeliveryOptionsInput, ...request.Option) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error)
	PutConfigurationSetDeliveryOptionsRequest(*sesv2.PutConfigurationSetDeliveryOptionsInput) (*request.Request, *sesv2.PutConfigurationSetDeliveryOptionsOutput)

	PutConfigurationSetReputationOptions(*sesv2.PutConfigurationSetReputationOptionsInput) (*sesv2.PutConfigurationSetReputationOptionsOutput, error)
	PutConfigurationSetReputationOptionsWithContext(aws.Context, *sesv2.PutConfigurationSetReputationOptionsInput, ...request.Option) (*sesv2.PutConfigurationSetReputationOptionsOutput, error)
	PutConfigurationSetReputationOptionsRequest(*sesv2.PutConfigurationSetReputationOptionsInput) (*request.Request, *sesv2.PutConfigurationSetReputationOptionsOutput)

	PutConfigurationSetSendingOptions(*sesv2.PutConfigurationSetSendingOptionsInput) (*sesv2.PutConfigurationSetSendingOptionsOutput, error)
	PutConfigurationSetSendingOptionsWithContext(aws.Context, *sesv2.PutConfigurationSetSendingOptionsInput, ...request.Option) (*sesv2.PutConfigurationSetSendingOptionsOutput, error)
	PutConfigurationSetSendingOptionsRequest(*sesv2.PutConfigurationSetSendingOptionsInput) (*request.Request, *sesv2.PutConfigurationSetSendingOptionsOutput)

	PutConfigurationSetSuppressionOptions(*sesv2.PutConfigurationSetSuppressionOptionsInput) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error)
	PutConfigurationSetSuppressionOptionsWithContext(aws.Context, *sesv2.PutConfigurationSetSuppressionOptionsInput, ...request.Option) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error)
	PutConfigurationSetSuppressionOptionsRequest(*sesv2.PutConfigurationSetSuppressionOptionsInput) (*request.Request, *sesv2.PutConfigurationSetSuppressionOptionsOutput)

	PutConfigurationSetTrackingOptions(*sesv2.PutConfigurationSetTrackingOptionsInput) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error)
	PutConfigurationSetTrackingOptionsWithContext(aws.Context, *sesv2.PutConfigurationSetTrackingOptionsInput, ...request.Option) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error)
	PutConfigurationSetTrackingOptionsRequest(*sesv2.PutConfigurationSetTrackingOptionsInput) (*request.Request, *sesv2.PutConfigurationSetTrackingOptionsOutput)

	PutConfigurationSetVdmOptions(*sesv2.PutConfigurationSetVdmOptionsInput) (*sesv2.PutConfigurationSetVdmOptionsOutput, error)
	PutConfigurationSetVdmOptionsWithContext(aws.Context, *sesv2.PutConfigurationSetVdmOptionsInput, ...request.Option) (*sesv2.PutConfigurationSetVdmOptionsOutput, error)
	PutConfigurationSetVdmOptionsRequest(*sesv2.PutConfigurationSetVdmOptionsInput) (*request.Request, *sesv2.PutConfigurationSetVdmOptionsOutput)

	PutDedicatedIpInPool(*sesv2.PutDedicatedIpInPoolInput) (*sesv2.PutDedicatedIpInPoolOutput, error)
	PutDedicatedIpInPoolWithContext(aws.Context, *sesv2.PutDedicatedIpInPoolInput, ...request.Option) (*sesv2.PutDedicatedIpInPoolOutput, error)
	PutDedicatedIpInPoolRequest(*sesv2.PutDedicatedIpInPoolInput) (*request.Request, *sesv2.PutDedicatedIpInPoolOutput)

	PutDedicatedIpWarmupAttributes(*sesv2.PutDedicatedIpWarmupAttributesInput) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error)
	PutDedicatedIpWarmupAttributesWithContext(aws.Context, *sesv2.PutDedicatedIpWarmupAttributesInput, ...request.Option) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error)
	PutDedicatedIpWarmupAttributesRequest(*sesv2.PutDedicatedIpWarmupAttributesInput) (*request.Request, *sesv2.PutDedicatedIpWarmupAttributesOutput)

	PutDeliverabilityDashboardOption(*sesv2.PutDeliverabilityDashboardOptionInput) (*sesv2.PutDeliverabilityDashboardOptionOutput, error)
	PutDeliverabilityDashboardOptionWithContext(aws.Context, *sesv2.PutDeliverabilityDashboardOptionInput, ...request.Option) (*sesv2.PutDeliverabilityDashboardOptionOutput, error)
	PutDeliverabilityDashboardOptionRequest(*sesv2.PutDeliverabilityDashboardOptionInput) (*request.Request, *sesv2.PutDeliverabilityDashboardOptionOutput)

	PutEmailIdentityConfigurationSetAttributes(*sesv2.PutEmailIdentityConfigurationSetAttributesInput) (*sesv2.PutEmailIdentityConfigurationSetAttributesOutput, error)
	PutEmailIdentityConfigurationSetAttributesWithContext(aws.Context, *sesv2.PutEmailIdentityConfigurationSetAttributesInput, ...request.Option) (*sesv2.PutEmailIdentityConfigurationSetAttributesOutput, error)
	PutEmailIdentityConfigurationSetAttributesRequest(*sesv2.PutEmailIdentityConfigurationSetAttributesInput) (*request.Request, *sesv2.PutEmailIdentityConfigurationSetAttributesOutput)

	PutEmailIdentityDkimAttributes(*sesv2.PutEmailIdentityDkimAttributesInput) (*sesv2.PutEmailIdentityDkimAttributesOutput, error)
	PutEmailIdentityDkimAttributesWithContext(aws.Context, *sesv2.PutEmailIdentityDkimAttributesInput, ...request.Option) (*sesv2.PutEmailIdentityDkimAttributesOutput, error)
	PutEmailIdentityDkimAttributesRequest(*sesv2.PutEmailIdentityDkimAttributesInput) (*request.Request, *sesv2.PutEmailIdentityDkimAttributesOutput)

	PutEmailIdentityDkimSigningAttributes(*sesv2.PutEmailIdentityDkimSigningAttributesInput) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error)
	PutEmailIdentityDkimSigningAttributesWithContext(aws.Context, *sesv2.PutEmailIdentityDkimSigningAttributesInput, ...request.Option) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error)
	PutEmailIdentityDkimSigningAttributesRequest(*sesv2.PutEmailIdentityDkimSigningAttributesInput) (*request.Request, *sesv2.PutEmailIdentityDkimSigningAttributesOutput)

	PutEmailIdentityFeedbackAttributes(*sesv2.PutEmailIdentityFeedbackAttributesInput) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error)
	PutEmailIdentityFeedbackAttributesWithContext(aws.Context, *sesv2.PutEmailIdentityFeedbackAttributesInput, ...request.Option) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error)
	PutEmailIdentityFeedbackAttributesRequest(*sesv2.PutEmailIdentityFeedbackAttributesInput) (*request.Request, *sesv2.PutEmailIdentityFeedbackAttributesOutput)

	PutEmailIdentityMailFromAttributes(*sesv2.PutEmailIdentityMailFromAttributesInput) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error)
	PutEmailIdentityMailFromAttributesWithContext(aws.Context, *sesv2.PutEmailIdentityMailFromAttributesInput, ...request.Option) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error)
	PutEmailIdentityMailFromAttributesRequest(*sesv2.PutEmailIdentityMailFromAttributesInput) (*request.Request, *sesv2.PutEmailIdentityMailFromAttributesOutput)

	PutSuppressedDestination(*sesv2.PutSuppressedDestinationInput) (*sesv2.PutSuppressedDestinationOutput, error)
	PutSuppressedDestinationWithContext(aws.Context, *sesv2.PutSuppressedDestinationInput, ...request.Option) (*sesv2.PutSuppressedDestinationOutput, error)
	PutSuppressedDestinationRequest(*sesv2.PutSuppressedDestinationInput) (*request.Request, *sesv2.PutSuppressedDestinationOutput)

	SendBulkEmail(*sesv2.SendBulkEmailInput) (*sesv2.SendBulkEmailOutput, error)
	SendBulkEmailWithContext(aws.Context, *sesv2.SendBulkEmailInput, ...request.Option) (*sesv2.SendBulkEmailOutput, error)
	SendBulkEmailRequest(*sesv2.SendBulkEmailInput) (*request.Request, *sesv2.SendBulkEmailOutput)

	SendCustomVerificationEmail(*sesv2.SendCustomVerificationEmailInput) (*sesv2.SendCustomVerificationEmailOutput, error)
	SendCustomVerificationEmailWithContext(aws.Context, *sesv2.SendCustomVerificationEmailInput, ...request.Option) (*sesv2.SendCustomVerificationEmailOutput, error)
	SendCustomVerificationEmailRequest(*sesv2.SendCustomVerificationEmailInput) (*request.Request, *sesv2.SendCustomVerificationEmailOutput)

	SendEmail(*sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error)
	SendEmailWithContext(aws.Context, *sesv2.SendEmailInput, ...request.Option) (*sesv2.SendEmailOutput, error)
	SendEmailRequest(*sesv2.SendEmailInput) (*request.Request, *sesv2.SendEmailOutput)

	TagResource(*sesv2.TagResourceInput) (*sesv2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *sesv2.TagResourceInput, ...request.Option) (*sesv2.TagResourceOutput, error)
	TagResourceRequest(*sesv2.TagResourceInput) (*request.Request, *sesv2.TagResourceOutput)

	TestRenderEmailTemplate(*sesv2.TestRenderEmailTemplateInput) (*sesv2.TestRenderEmailTemplateOutput, error)
	TestRenderEmailTemplateWithContext(aws.Context, *sesv2.TestRenderEmailTemplateInput, ...request.Option) (*sesv2.TestRenderEmailTemplateOutput, error)
	TestRenderEmailTemplateRequest(*sesv2.TestRenderEmailTemplateInput) (*request.Request, *sesv2.TestRenderEmailTemplateOutput)

	UntagResource(*sesv2.UntagResourceInput) (*sesv2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *sesv2.UntagResourceInput, ...request.Option) (*sesv2.UntagResourceOutput, error)
	UntagResourceRequest(*sesv2.UntagResourceInput) (*request.Request, *sesv2.UntagResourceOutput)

	UpdateConfigurationSetEventDestination(*sesv2.UpdateConfigurationSetEventDestinationInput) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationWithContext(aws.Context, *sesv2.UpdateConfigurationSetEventDestinationInput, ...request.Option) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationRequest(*sesv2.UpdateConfigurationSetEventDestinationInput) (*request.Request, *sesv2.UpdateConfigurationSetEventDestinationOutput)

	UpdateContact(*sesv2.UpdateContactInput) (*sesv2.UpdateContactOutput, error)
	UpdateContactWithContext(aws.Context, *sesv2.UpdateContactInput, ...request.Option) (*sesv2.UpdateContactOutput, error)
	UpdateContactRequest(*sesv2.UpdateContactInput) (*request.Request, *sesv2.UpdateContactOutput)

	UpdateContactList(*sesv2.UpdateContactListInput) (*sesv2.UpdateContactListOutput, error)
	UpdateContactListWithContext(aws.Context, *sesv2.UpdateContactListInput, ...request.Option) (*sesv2.UpdateContactListOutput, error)
	UpdateContactListRequest(*sesv2.UpdateContactListInput) (*request.Request, *sesv2.UpdateContactListOutput)

	UpdateCustomVerificationEmailTemplate(*sesv2.UpdateCustomVerificationEmailTemplateInput) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateCustomVerificationEmailTemplateWithContext(aws.Context, *sesv2.UpdateCustomVerificationEmailTemplateInput, ...request.Option) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateCustomVerificationEmailTemplateRequest(*sesv2.UpdateCustomVerificationEmailTemplateInput) (*request.Request, *sesv2.UpdateCustomVerificationEmailTemplateOutput)

	UpdateEmailIdentityPolicy(*sesv2.UpdateEmailIdentityPolicyInput) (*sesv2.UpdateEmailIdentityPolicyOutput, error)
	UpdateEmailIdentityPolicyWithContext(aws.Context, *sesv2.UpdateEmailIdentityPolicyInput, ...request.Option) (*sesv2.UpdateEmailIdentityPolicyOutput, error)
	UpdateEmailIdentityPolicyRequest(*sesv2.UpdateEmailIdentityPolicyInput) (*request.Request, *sesv2.UpdateEmailIdentityPolicyOutput)

	UpdateEmailTemplate(*sesv2.UpdateEmailTemplateInput) (*sesv2.UpdateEmailTemplateOutput, error)
	UpdateEmailTemplateWithContext(aws.Context, *sesv2.UpdateEmailTemplateInput, ...request.Option) (*sesv2.UpdateEmailTemplateOutput, error)
	UpdateEmailTemplateRequest(*sesv2.UpdateEmailTemplateInput) (*request.Request, *sesv2.UpdateEmailTemplateOutput)
}

var _ SESV2API = (*sesv2.SESV2)(nil)
