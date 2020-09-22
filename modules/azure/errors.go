package azure

import "fmt"

// SubscriptionIDNotFound is an error that occurs when the Azure Subscription ID could not be found or was not provided
type SubscriptionIDNotFound struct{}

func (err SubscriptionIDNotFound) Error() string {
	return fmt.Sprintf("Could not find an Azure Subscription ID in expected environment variable %s and one was not provided for this test.", AzureSubscriptionID)
}

// ResourceGroupNameNotFound is an error that occurs when the target Azure Resource Group name could not be found or was not provided
type ResourceGroupNameNotFound struct{}

func (err ResourceGroupNameNotFound) Error() string {
	return fmt.Sprintf("Could not find an Azure Resource Group name in expected environment variable %s and one was not provided for this test.", AzureResGroupName)
}

// FailedToParseError is returned when an object cannot be parsed
type FailedToParseError struct {
	objectType string
	objectID   string
}

func (err FailedToParseError) Error() string {
	return fmt.Sprintf("Failed to parse %s with ID %s", err.objectType, err.objectID)
}

// NewFailedToParseError creates a new not found error when an expected object is not found in the search space
func NewFailedToParseError(objectType string, objectID string) FailedToParseError {
	return FailedToParseError{objectType, objectID}
}
