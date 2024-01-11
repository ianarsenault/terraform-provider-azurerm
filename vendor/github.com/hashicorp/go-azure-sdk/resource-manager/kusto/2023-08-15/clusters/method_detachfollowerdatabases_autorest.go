package clusters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetachFollowerDatabasesOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DetachFollowerDatabases ...
func (c ClustersClient) DetachFollowerDatabases(ctx context.Context, id commonids.KustoClusterId, input FollowerDatabaseDefinition) (result DetachFollowerDatabasesOperationResponse, err error) {
	req, err := c.preparerForDetachFollowerDatabases(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusters.ClustersClient", "DetachFollowerDatabases", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDetachFollowerDatabases(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusters.ClustersClient", "DetachFollowerDatabases", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DetachFollowerDatabasesThenPoll performs DetachFollowerDatabases then polls until it's completed
func (c ClustersClient) DetachFollowerDatabasesThenPoll(ctx context.Context, id commonids.KustoClusterId, input FollowerDatabaseDefinition) error {
	result, err := c.DetachFollowerDatabases(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DetachFollowerDatabases: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DetachFollowerDatabases: %+v", err)
	}

	return nil
}

// preparerForDetachFollowerDatabases prepares the DetachFollowerDatabases request.
func (c ClustersClient) preparerForDetachFollowerDatabases(ctx context.Context, id commonids.KustoClusterId, input FollowerDatabaseDefinition) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detachFollowerDatabases", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDetachFollowerDatabases sends the DetachFollowerDatabases request. The method will close the
// http.Response Body if it receives an error.
func (c ClustersClient) senderForDetachFollowerDatabases(ctx context.Context, req *http.Request) (future DetachFollowerDatabasesOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
