// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdbiface provides an interface to enable mocking the NIFCLOUD RDB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdbiface

import (
	"github.com/shztki/nifcloud-sdk-go/nifcloud"
	"github.com/shztki/nifcloud-sdk-go/nifcloud/request"
	"github.com/shztki/nifcloud-sdk-go/service/rdb"
)

// RdbAPI provides an interface to enable mocking the
// rdb.Rdb service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // NIFCLOUD RDB.
//    func myFunc(svc rdbiface.RdbAPI) bool {
//        // Make svc.AddSourceIdentifierToSubscription request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rdb.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRdbClient struct {
//        rdbiface.RdbAPI
//    }
//    func (m *mockRdbClient) AddSourceIdentifierToSubscription(input *rdb.AddSourceIdentifierToSubscriptionInput) (*rdb.AddSourceIdentifierToSubscriptionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRdbClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RdbAPI interface {
	AddSourceIdentifierToSubscription(*rdb.AddSourceIdentifierToSubscriptionInput) (*rdb.AddSourceIdentifierToSubscriptionOutput, error)
	AddSourceIdentifierToSubscriptionWithContext(nifcloud.Context, *rdb.AddSourceIdentifierToSubscriptionInput, ...request.Option) (*rdb.AddSourceIdentifierToSubscriptionOutput, error)
	AddSourceIdentifierToSubscriptionRequest(*rdb.AddSourceIdentifierToSubscriptionInput) (*request.Request, *rdb.AddSourceIdentifierToSubscriptionOutput)

	AuthorizeDBSecurityGroupIngress(*rdb.AuthorizeDBSecurityGroupIngressInput) (*rdb.AuthorizeDBSecurityGroupIngressOutput, error)
	AuthorizeDBSecurityGroupIngressWithContext(nifcloud.Context, *rdb.AuthorizeDBSecurityGroupIngressInput, ...request.Option) (*rdb.AuthorizeDBSecurityGroupIngressOutput, error)
	AuthorizeDBSecurityGroupIngressRequest(*rdb.AuthorizeDBSecurityGroupIngressInput) (*request.Request, *rdb.AuthorizeDBSecurityGroupIngressOutput)

	CopyDBSnapshot(*rdb.CopyDBSnapshotInput) (*rdb.CopyDBSnapshotOutput, error)
	CopyDBSnapshotWithContext(nifcloud.Context, *rdb.CopyDBSnapshotInput, ...request.Option) (*rdb.CopyDBSnapshotOutput, error)
	CopyDBSnapshotRequest(*rdb.CopyDBSnapshotInput) (*request.Request, *rdb.CopyDBSnapshotOutput)

	CreateDBInstance(*rdb.CreateDBInstanceInput) (*rdb.CreateDBInstanceOutput, error)
	CreateDBInstanceWithContext(nifcloud.Context, *rdb.CreateDBInstanceInput, ...request.Option) (*rdb.CreateDBInstanceOutput, error)
	CreateDBInstanceRequest(*rdb.CreateDBInstanceInput) (*request.Request, *rdb.CreateDBInstanceOutput)

	CreateDBInstanceReadReplica(*rdb.CreateDBInstanceReadReplicaInput) (*rdb.CreateDBInstanceReadReplicaOutput, error)
	CreateDBInstanceReadReplicaWithContext(nifcloud.Context, *rdb.CreateDBInstanceReadReplicaInput, ...request.Option) (*rdb.CreateDBInstanceReadReplicaOutput, error)
	CreateDBInstanceReadReplicaRequest(*rdb.CreateDBInstanceReadReplicaInput) (*request.Request, *rdb.CreateDBInstanceReadReplicaOutput)

	CreateDBParameterGroup(*rdb.CreateDBParameterGroupInput) (*rdb.CreateDBParameterGroupOutput, error)
	CreateDBParameterGroupWithContext(nifcloud.Context, *rdb.CreateDBParameterGroupInput, ...request.Option) (*rdb.CreateDBParameterGroupOutput, error)
	CreateDBParameterGroupRequest(*rdb.CreateDBParameterGroupInput) (*request.Request, *rdb.CreateDBParameterGroupOutput)

	CreateDBSecurityGroup(*rdb.CreateDBSecurityGroupInput) (*rdb.CreateDBSecurityGroupOutput, error)
	CreateDBSecurityGroupWithContext(nifcloud.Context, *rdb.CreateDBSecurityGroupInput, ...request.Option) (*rdb.CreateDBSecurityGroupOutput, error)
	CreateDBSecurityGroupRequest(*rdb.CreateDBSecurityGroupInput) (*request.Request, *rdb.CreateDBSecurityGroupOutput)

	CreateDBSnapshot(*rdb.CreateDBSnapshotInput) (*rdb.CreateDBSnapshotOutput, error)
	CreateDBSnapshotWithContext(nifcloud.Context, *rdb.CreateDBSnapshotInput, ...request.Option) (*rdb.CreateDBSnapshotOutput, error)
	CreateDBSnapshotRequest(*rdb.CreateDBSnapshotInput) (*request.Request, *rdb.CreateDBSnapshotOutput)

	CreateEventSubscription(*rdb.CreateEventSubscriptionInput) (*rdb.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionWithContext(nifcloud.Context, *rdb.CreateEventSubscriptionInput, ...request.Option) (*rdb.CreateEventSubscriptionOutput, error)
	CreateEventSubscriptionRequest(*rdb.CreateEventSubscriptionInput) (*request.Request, *rdb.CreateEventSubscriptionOutput)

	DeleteDBInstance(*rdb.DeleteDBInstanceInput) (*rdb.DeleteDBInstanceOutput, error)
	DeleteDBInstanceWithContext(nifcloud.Context, *rdb.DeleteDBInstanceInput, ...request.Option) (*rdb.DeleteDBInstanceOutput, error)
	DeleteDBInstanceRequest(*rdb.DeleteDBInstanceInput) (*request.Request, *rdb.DeleteDBInstanceOutput)

	DeleteDBParameterGroup(*rdb.DeleteDBParameterGroupInput) (*rdb.DeleteDBParameterGroupOutput, error)
	DeleteDBParameterGroupWithContext(nifcloud.Context, *rdb.DeleteDBParameterGroupInput, ...request.Option) (*rdb.DeleteDBParameterGroupOutput, error)
	DeleteDBParameterGroupRequest(*rdb.DeleteDBParameterGroupInput) (*request.Request, *rdb.DeleteDBParameterGroupOutput)

	DeleteDBSecurityGroup(*rdb.DeleteDBSecurityGroupInput) (*rdb.DeleteDBSecurityGroupOutput, error)
	DeleteDBSecurityGroupWithContext(nifcloud.Context, *rdb.DeleteDBSecurityGroupInput, ...request.Option) (*rdb.DeleteDBSecurityGroupOutput, error)
	DeleteDBSecurityGroupRequest(*rdb.DeleteDBSecurityGroupInput) (*request.Request, *rdb.DeleteDBSecurityGroupOutput)

	DeleteDBSnapshot(*rdb.DeleteDBSnapshotInput) (*rdb.DeleteDBSnapshotOutput, error)
	DeleteDBSnapshotWithContext(nifcloud.Context, *rdb.DeleteDBSnapshotInput, ...request.Option) (*rdb.DeleteDBSnapshotOutput, error)
	DeleteDBSnapshotRequest(*rdb.DeleteDBSnapshotInput) (*request.Request, *rdb.DeleteDBSnapshotOutput)

	DeleteEventSubscription(*rdb.DeleteEventSubscriptionInput) (*rdb.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionWithContext(nifcloud.Context, *rdb.DeleteEventSubscriptionInput, ...request.Option) (*rdb.DeleteEventSubscriptionOutput, error)
	DeleteEventSubscriptionRequest(*rdb.DeleteEventSubscriptionInput) (*request.Request, *rdb.DeleteEventSubscriptionOutput)

	DescribeCertificates(*rdb.DescribeCertificatesInput) (*rdb.DescribeCertificatesOutput, error)
	DescribeCertificatesWithContext(nifcloud.Context, *rdb.DescribeCertificatesInput, ...request.Option) (*rdb.DescribeCertificatesOutput, error)
	DescribeCertificatesRequest(*rdb.DescribeCertificatesInput) (*request.Request, *rdb.DescribeCertificatesOutput)

	DescribeDBEngineVersions(*rdb.DescribeDBEngineVersionsInput) (*rdb.DescribeDBEngineVersionsOutput, error)
	DescribeDBEngineVersionsWithContext(nifcloud.Context, *rdb.DescribeDBEngineVersionsInput, ...request.Option) (*rdb.DescribeDBEngineVersionsOutput, error)
	DescribeDBEngineVersionsRequest(*rdb.DescribeDBEngineVersionsInput) (*request.Request, *rdb.DescribeDBEngineVersionsOutput)

	DescribeDBInstances(*rdb.DescribeDBInstancesInput) (*rdb.DescribeDBInstancesOutput, error)
	DescribeDBInstancesWithContext(nifcloud.Context, *rdb.DescribeDBInstancesInput, ...request.Option) (*rdb.DescribeDBInstancesOutput, error)
	DescribeDBInstancesRequest(*rdb.DescribeDBInstancesInput) (*request.Request, *rdb.DescribeDBInstancesOutput)

	DescribeDBLogFiles(*rdb.DescribeDBLogFilesInput) (*rdb.DescribeDBLogFilesOutput, error)
	DescribeDBLogFilesWithContext(nifcloud.Context, *rdb.DescribeDBLogFilesInput, ...request.Option) (*rdb.DescribeDBLogFilesOutput, error)
	DescribeDBLogFilesRequest(*rdb.DescribeDBLogFilesInput) (*request.Request, *rdb.DescribeDBLogFilesOutput)

	DescribeDBParameterGroups(*rdb.DescribeDBParameterGroupsInput) (*rdb.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameterGroupsWithContext(nifcloud.Context, *rdb.DescribeDBParameterGroupsInput, ...request.Option) (*rdb.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameterGroupsRequest(*rdb.DescribeDBParameterGroupsInput) (*request.Request, *rdb.DescribeDBParameterGroupsOutput)

	DescribeDBParameters(*rdb.DescribeDBParametersInput) (*rdb.DescribeDBParametersOutput, error)
	DescribeDBParametersWithContext(nifcloud.Context, *rdb.DescribeDBParametersInput, ...request.Option) (*rdb.DescribeDBParametersOutput, error)
	DescribeDBParametersRequest(*rdb.DescribeDBParametersInput) (*request.Request, *rdb.DescribeDBParametersOutput)

	DescribeDBSecurityGroups(*rdb.DescribeDBSecurityGroupsInput) (*rdb.DescribeDBSecurityGroupsOutput, error)
	DescribeDBSecurityGroupsWithContext(nifcloud.Context, *rdb.DescribeDBSecurityGroupsInput, ...request.Option) (*rdb.DescribeDBSecurityGroupsOutput, error)
	DescribeDBSecurityGroupsRequest(*rdb.DescribeDBSecurityGroupsInput) (*request.Request, *rdb.DescribeDBSecurityGroupsOutput)

	DescribeDBSnapshots(*rdb.DescribeDBSnapshotsInput) (*rdb.DescribeDBSnapshotsOutput, error)
	DescribeDBSnapshotsWithContext(nifcloud.Context, *rdb.DescribeDBSnapshotsInput, ...request.Option) (*rdb.DescribeDBSnapshotsOutput, error)
	DescribeDBSnapshotsRequest(*rdb.DescribeDBSnapshotsInput) (*request.Request, *rdb.DescribeDBSnapshotsOutput)

	DescribeEngineDefaultParameters(*rdb.DescribeEngineDefaultParametersInput) (*rdb.DescribeEngineDefaultParametersOutput, error)
	DescribeEngineDefaultParametersWithContext(nifcloud.Context, *rdb.DescribeEngineDefaultParametersInput, ...request.Option) (*rdb.DescribeEngineDefaultParametersOutput, error)
	DescribeEngineDefaultParametersRequest(*rdb.DescribeEngineDefaultParametersInput) (*request.Request, *rdb.DescribeEngineDefaultParametersOutput)

	DescribeEventCategories(*rdb.DescribeEventCategoriesInput) (*rdb.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesWithContext(nifcloud.Context, *rdb.DescribeEventCategoriesInput, ...request.Option) (*rdb.DescribeEventCategoriesOutput, error)
	DescribeEventCategoriesRequest(*rdb.DescribeEventCategoriesInput) (*request.Request, *rdb.DescribeEventCategoriesOutput)

	DescribeEventSubscriptions(*rdb.DescribeEventSubscriptionsInput) (*rdb.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsWithContext(nifcloud.Context, *rdb.DescribeEventSubscriptionsInput, ...request.Option) (*rdb.DescribeEventSubscriptionsOutput, error)
	DescribeEventSubscriptionsRequest(*rdb.DescribeEventSubscriptionsInput) (*request.Request, *rdb.DescribeEventSubscriptionsOutput)

	DescribeEvents(*rdb.DescribeEventsInput) (*rdb.DescribeEventsOutput, error)
	DescribeEventsWithContext(nifcloud.Context, *rdb.DescribeEventsInput, ...request.Option) (*rdb.DescribeEventsOutput, error)
	DescribeEventsRequest(*rdb.DescribeEventsInput) (*request.Request, *rdb.DescribeEventsOutput)

	DescribeOrderableDBInstanceOptions(*rdb.DescribeOrderableDBInstanceOptionsInput) (*rdb.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribeOrderableDBInstanceOptionsWithContext(nifcloud.Context, *rdb.DescribeOrderableDBInstanceOptionsInput, ...request.Option) (*rdb.DescribeOrderableDBInstanceOptionsOutput, error)
	DescribeOrderableDBInstanceOptionsRequest(*rdb.DescribeOrderableDBInstanceOptionsInput) (*request.Request, *rdb.DescribeOrderableDBInstanceOptionsOutput)

	DownloadDBLogFilePortion(*rdb.DownloadDBLogFilePortionInput) (*rdb.DownloadDBLogFilePortionOutput, error)
	DownloadDBLogFilePortionWithContext(nifcloud.Context, *rdb.DownloadDBLogFilePortionInput, ...request.Option) (*rdb.DownloadDBLogFilePortionOutput, error)
	DownloadDBLogFilePortionRequest(*rdb.DownloadDBLogFilePortionInput) (*request.Request, *rdb.DownloadDBLogFilePortionOutput)

	ModifyDBInstance(*rdb.ModifyDBInstanceInput) (*rdb.ModifyDBInstanceOutput, error)
	ModifyDBInstanceWithContext(nifcloud.Context, *rdb.ModifyDBInstanceInput, ...request.Option) (*rdb.ModifyDBInstanceOutput, error)
	ModifyDBInstanceRequest(*rdb.ModifyDBInstanceInput) (*request.Request, *rdb.ModifyDBInstanceOutput)

	ModifyDBParameterGroup(*rdb.ModifyDBParameterGroupInput) (*rdb.ModifyDBParameterGroupOutput, error)
	ModifyDBParameterGroupWithContext(nifcloud.Context, *rdb.ModifyDBParameterGroupInput, ...request.Option) (*rdb.ModifyDBParameterGroupOutput, error)
	ModifyDBParameterGroupRequest(*rdb.ModifyDBParameterGroupInput) (*request.Request, *rdb.ModifyDBParameterGroupOutput)

	ModifyEventSubscription(*rdb.ModifyEventSubscriptionInput) (*rdb.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionWithContext(nifcloud.Context, *rdb.ModifyEventSubscriptionInput, ...request.Option) (*rdb.ModifyEventSubscriptionOutput, error)
	ModifyEventSubscriptionRequest(*rdb.ModifyEventSubscriptionInput) (*request.Request, *rdb.ModifyEventSubscriptionOutput)

	NiftyFailoverDBInstance(*rdb.NiftyFailoverDBInstanceInput) (*rdb.NiftyFailoverDBInstanceOutput, error)
	NiftyFailoverDBInstanceWithContext(nifcloud.Context, *rdb.NiftyFailoverDBInstanceInput, ...request.Option) (*rdb.NiftyFailoverDBInstanceOutput, error)
	NiftyFailoverDBInstanceRequest(*rdb.NiftyFailoverDBInstanceInput) (*request.Request, *rdb.NiftyFailoverDBInstanceOutput)

	NiftyGetMetricStatistics(*rdb.NiftyGetMetricStatisticsInput) (*rdb.NiftyGetMetricStatisticsOutput, error)
	NiftyGetMetricStatisticsWithContext(nifcloud.Context, *rdb.NiftyGetMetricStatisticsInput, ...request.Option) (*rdb.NiftyGetMetricStatisticsOutput, error)
	NiftyGetMetricStatisticsRequest(*rdb.NiftyGetMetricStatisticsInput) (*request.Request, *rdb.NiftyGetMetricStatisticsOutput)

	RebootDBInstance(*rdb.RebootDBInstanceInput) (*rdb.RebootDBInstanceOutput, error)
	RebootDBInstanceWithContext(nifcloud.Context, *rdb.RebootDBInstanceInput, ...request.Option) (*rdb.RebootDBInstanceOutput, error)
	RebootDBInstanceRequest(*rdb.RebootDBInstanceInput) (*request.Request, *rdb.RebootDBInstanceOutput)

	RemoveSourceIdentifierFromSubscription(*rdb.RemoveSourceIdentifierFromSubscriptionInput) (*rdb.RemoveSourceIdentifierFromSubscriptionOutput, error)
	RemoveSourceIdentifierFromSubscriptionWithContext(nifcloud.Context, *rdb.RemoveSourceIdentifierFromSubscriptionInput, ...request.Option) (*rdb.RemoveSourceIdentifierFromSubscriptionOutput, error)
	RemoveSourceIdentifierFromSubscriptionRequest(*rdb.RemoveSourceIdentifierFromSubscriptionInput) (*request.Request, *rdb.RemoveSourceIdentifierFromSubscriptionOutput)

	ResetDBParameterGroup(*rdb.ResetDBParameterGroupInput) (*rdb.ResetDBParameterGroupOutput, error)
	ResetDBParameterGroupWithContext(nifcloud.Context, *rdb.ResetDBParameterGroupInput, ...request.Option) (*rdb.ResetDBParameterGroupOutput, error)
	ResetDBParameterGroupRequest(*rdb.ResetDBParameterGroupInput) (*request.Request, *rdb.ResetDBParameterGroupOutput)

	ResetExternalMaster(*rdb.ResetExternalMasterInput) (*rdb.ResetExternalMasterOutput, error)
	ResetExternalMasterWithContext(nifcloud.Context, *rdb.ResetExternalMasterInput, ...request.Option) (*rdb.ResetExternalMasterOutput, error)
	ResetExternalMasterRequest(*rdb.ResetExternalMasterInput) (*request.Request, *rdb.ResetExternalMasterOutput)

	RestoreDBInstanceFromDBSnapshot(*rdb.RestoreDBInstanceFromDBSnapshotInput) (*rdb.RestoreDBInstanceFromDBSnapshotOutput, error)
	RestoreDBInstanceFromDBSnapshotWithContext(nifcloud.Context, *rdb.RestoreDBInstanceFromDBSnapshotInput, ...request.Option) (*rdb.RestoreDBInstanceFromDBSnapshotOutput, error)
	RestoreDBInstanceFromDBSnapshotRequest(*rdb.RestoreDBInstanceFromDBSnapshotInput) (*request.Request, *rdb.RestoreDBInstanceFromDBSnapshotOutput)

	RestoreDBInstanceToPointInTime(*rdb.RestoreDBInstanceToPointInTimeInput) (*rdb.RestoreDBInstanceToPointInTimeOutput, error)
	RestoreDBInstanceToPointInTimeWithContext(nifcloud.Context, *rdb.RestoreDBInstanceToPointInTimeInput, ...request.Option) (*rdb.RestoreDBInstanceToPointInTimeOutput, error)
	RestoreDBInstanceToPointInTimeRequest(*rdb.RestoreDBInstanceToPointInTimeInput) (*request.Request, *rdb.RestoreDBInstanceToPointInTimeOutput)

	RevokeDBSecurityGroupIngress(*rdb.RevokeDBSecurityGroupIngressInput) (*rdb.RevokeDBSecurityGroupIngressOutput, error)
	RevokeDBSecurityGroupIngressWithContext(nifcloud.Context, *rdb.RevokeDBSecurityGroupIngressInput, ...request.Option) (*rdb.RevokeDBSecurityGroupIngressOutput, error)
	RevokeDBSecurityGroupIngressRequest(*rdb.RevokeDBSecurityGroupIngressInput) (*request.Request, *rdb.RevokeDBSecurityGroupIngressOutput)

	SetExternalMaster(*rdb.SetExternalMasterInput) (*rdb.SetExternalMasterOutput, error)
	SetExternalMasterWithContext(nifcloud.Context, *rdb.SetExternalMasterInput, ...request.Option) (*rdb.SetExternalMasterOutput, error)
	SetExternalMasterRequest(*rdb.SetExternalMasterInput) (*request.Request, *rdb.SetExternalMasterOutput)

	StartReplication(*rdb.StartReplicationInput) (*rdb.StartReplicationOutput, error)
	StartReplicationWithContext(nifcloud.Context, *rdb.StartReplicationInput, ...request.Option) (*rdb.StartReplicationOutput, error)
	StartReplicationRequest(*rdb.StartReplicationInput) (*request.Request, *rdb.StartReplicationOutput)

	StopReplication(*rdb.StopReplicationInput) (*rdb.StopReplicationOutput, error)
	StopReplicationWithContext(nifcloud.Context, *rdb.StopReplicationInput, ...request.Option) (*rdb.StopReplicationOutput, error)
	StopReplicationRequest(*rdb.StopReplicationInput) (*request.Request, *rdb.StopReplicationOutput)
}

var _ RdbAPI = (*rdb.Rdb)(nil)
