package model

import "time"

// Namespace represents an Okteto namespace
type Namespace struct {
	// Date and time when the namespace was created.
	CreationDate time.Time `json:"creationDate,required" format:"date-time"`
	// Date and time when the namespace was last updated.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Name of the namespace.
	Name string `json:"name,required"`
	// Indicates if the namespace is persistent.
	Persistent bool `json:"persistent,required"`
	// Indicates if the namespace is the default namespace for any user in the
	// instance.
	Personal bool `json:"personal,required"`
	// Status of the namespace. Possible values are "Active", "DestroyAllFailed","DestroyingAll","Deleting","Inactive","Sleeping" and "DeleteFailed"
	Status NamespaceStatus `json:"status,required" casing:"lower"`
	// Type of the namespace. Possible values are 'Development' and 'Preview'.
	Type NamespaceType `json:"type,required" casing:"lower"`
	// Unique identifier of the namespace.
	Uuid string `json:"uuid,required"`
}

// Type of the namespace. Possible values are 'Development' and 'Preview'.
type NamespaceType string

const (
	NamespaceTypeDevelopment NamespaceType = "development"
	NamespaceTypePreview     NamespaceType = "preview"
)

// Status of the namespace. Possible values are 'Active', 'Inactive', 'Deleted'.
type NamespaceStatus string

const (
	NamespaceStatusActive           NamespaceStatus = "Active"
	NamespaceStatusDestroyAllFailed NamespaceStatus = "DestroyAllFailed"
	NamespaceStatusDestroyingAll    NamespaceStatus = "DestroyingAll"
	NamespaceStatusDeleting         NamespaceStatus = "Deleting"
	NamespaceStatusInactive         NamespaceStatus = "Inactive"
	NamespaceStatusSleeping         NamespaceStatus = "Sleeping"
	NamespaceStatusDeleteFailed     NamespaceStatus = "DeleteFailed"
)
