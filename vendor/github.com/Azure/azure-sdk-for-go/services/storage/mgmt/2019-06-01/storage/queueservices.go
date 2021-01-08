package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// QueueServicesClient is the the Azure Storage Management API.
type QueueServicesClient struct {
	BaseClient
}

// NewQueueServicesClient creates an instance of the QueueServicesClient client.
func NewQueueServicesClient(subscriptionID string) QueueServicesClient {
	return NewQueueServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueueServicesClientWithBaseURI creates an instance of the QueueServicesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueueServicesClientWithBaseURI(baseURI string, subscriptionID string) QueueServicesClient {
	return QueueServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetServiceProperties gets the properties of a storage account’s Queue service, including properties for Storage
// Analytics and CORS (Cross-Origin Resource Sharing) rules.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
func (client QueueServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string) (result QueueServiceProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueueServicesClient.GetServiceProperties")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.QueueServicesClient", "GetServiceProperties", err.Error())
	}

	req, err := client.GetServicePropertiesPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "GetServiceProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetServicePropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "GetServiceProperties", resp, "Failure sending request")
		return
	}

	result, err = client.GetServicePropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "GetServiceProperties", resp, "Failure responding to request")
		return
	}

	return
}

// GetServicePropertiesPreparer prepares the GetServiceProperties request.
func (client QueueServicesClient) GetServicePropertiesPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"queueServiceName":  autorest.Encode("path", "default"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetServicePropertiesSender sends the GetServiceProperties request. The method will close the
// http.Response Body if it receives an error.
func (client QueueServicesClient) GetServicePropertiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetServicePropertiesResponder handles the response to the GetServiceProperties request. The method always
// closes the http.Response Body.
func (client QueueServicesClient) GetServicePropertiesResponder(resp *http.Response) (result QueueServiceProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all queue services for the storage account
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
func (client QueueServicesClient) List(ctx context.Context, resourceGroupName string, accountName string) (result ListQueueServices, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueueServicesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.QueueServicesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client QueueServicesClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client QueueServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client QueueServicesClient) ListResponder(resp *http.Response) (result ListQueueServices, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SetServiceProperties sets the properties of a storage account’s Queue service, including properties for Storage
// Analytics and CORS (Cross-Origin Resource Sharing) rules.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// accountName - the name of the storage account within the specified resource group. Storage account names
// must be between 3 and 24 characters in length and use numbers and lower-case letters only.
// parameters - the properties of a storage account’s Queue service, only properties for Storage Analytics and
// CORS (Cross-Origin Resource Sharing) rules can be specified.
func (client QueueServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties) (result QueueServiceProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueueServicesClient.SetServiceProperties")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.QueueServicesClient", "SetServiceProperties", err.Error())
	}

	req, err := client.SetServicePropertiesPreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "SetServiceProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetServicePropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "SetServiceProperties", resp, "Failure sending request")
		return
	}

	result, err = client.SetServicePropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.QueueServicesClient", "SetServiceProperties", resp, "Failure responding to request")
		return
	}

	return
}

// SetServicePropertiesPreparer prepares the SetServiceProperties request.
func (client QueueServicesClient) SetServicePropertiesPreparer(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"queueServiceName":  autorest.Encode("path", "default"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetServicePropertiesSender sends the SetServiceProperties request. The method will close the
// http.Response Body if it receives an error.
func (client QueueServicesClient) SetServicePropertiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SetServicePropertiesResponder handles the response to the SetServiceProperties request. The method always
// closes the http.Response Body.
func (client QueueServicesClient) SetServicePropertiesResponder(resp *http.Response) (result QueueServiceProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
