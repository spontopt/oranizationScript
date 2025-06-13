package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComunicationModels struct {
	WelcomeMsg       string `bson:"welcomeMsg"`
	PasswordResetMsg string `bson:"passwordResetMsg"`
	NotificationMsg  string `bson:"notificationMsg"`
}

// Organization represents an organization's details.
type Organization struct {
	OrganizationID primitive.ObjectID `bson:"_id,omitempty"`
	SubscriptionID string             `bson:"subscriptionId"`
	Email          string             `bson:"email"`
	Language       string             `bson:"language"`
	Currency       string             `bson:"currency"`
	NumberFormat   string             `bson:"number_format"`
	DateFormat     string             `bson:"date_format"`
	EmailModel     ComunicationModels `bson:"emailModel"`
	WhatsappModel  ComunicationModels `bson:"WhatsappModel"`
	SMSModel       ComunicationModels `bson:"SMSModel"`
}

// UserOrganization represents a user's association with an organization.
type UserOrganization struct {
	UserOrganizationID primitive.ObjectID `bson:"_id,omitempty"`
	OrganizationID     string             `bson:"organizationId"`
	SubscriptionID     string             `bson:"subscriptionId"`
	RoleID             string             `bson:"roleId"`
	UserID             int                `bson:"userId"`
	Name               string             `bson:"name"`
	Email              string             `bson:"email"`
	Role               string             `bson:"role"`
	Status             bool               `bson:"status"`
	Department         string             `bson:"department,omitempty"`
}

// Operation represents an operation with its active status.
type Operation struct {
	Operation string `bson:"operation,omitempty" json:"operation,omitempty"` // Renamed to avoid collision with struct type if methods are added
	Active    bool   `bson:"active" json:"active"`
}

// Feature represents a feature and its associated operations.
type Feature struct {
	Feature    string      `bson:"feature,omitempty" json:"feature,omitempty"` // Renamed to avoid collision with struct type if methods are added
	Operations []Operation `bson:"operations" json:"operations"`
}

// Permission represents a set of features and a general permission string.
type Permission struct {
	Features   []Feature `bson:"features,omitempty" json:"features,omitempty"`
	Permission string    `bson:"permission,omitempty" json:"permission,omitempty"`
}

// Profile represents a user profile.
type Profile struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	LegacyID       string             `bson:"id,omitempty" json:"id,omitempty"` // 'id' from TypeScript might be a string ID, keep for consistency
	Profile        string             `bson:"profile" json:"profile"`           // Required field in TypeScript
	Permissions    Permission         `bson:"permissions" json:"permissions"`   // Required field in TypeScript
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	Custom         bool               `bson:"custom,omitempty" json:"custom,omitempty"`
	DefaultProfile bool               `bson:"defaultProfile,omitempty" json:"defaultProfile,omitempty"`
	OrganizationID string             `bson:"organizationId,omitempty" json:"organizationId,omitempty"`
	SubscriptionID string             `bson:"subscriptionId,omitempty" json:"subscriptionId,omitempty"`
}

type IUser struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	IDUser         int                `bson:"IDUser" json:"IDUser"`
	Email          string             `bson:"email" json:"email"`
	Name           string             `bson:"name" json:"name"`
	IDTenant       string             `bson:"idTenant,omitempty" json:"idTenant,omitempty"`
	SubscriptionID string             `bson:"subscriptionId,omitempty" json:"subscriptionId,omitempty"`
	Profile        *Profile           `bson:"profile,omitempty" json:"profile,omitempty"`
	Language       string             `bson:"language,omitempty" json:"language,omitempty"`
	Currency       string             `bson:"currency,omitempty" json:"currency,omitempty"`
	DateFormat     string             `bson:"date_format,omitempty" json:"date_format,omitempty"`
	NumberFormat   string             `bson:"number_format,omitempty" json:"number_format,omitempty"`
	Active         bool               `bson:"active,omitempty" json:"active,omitempty"`
}
