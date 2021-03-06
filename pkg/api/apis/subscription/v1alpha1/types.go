package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	GroupVersion              = "v1alpha1" // version used in the Scheme for subscriptions
	SubscriptionKind          = "Subscription-v1"
	SubscriptionListKind      = "SubscriptionList-v1"
	SubscriptionCRDName       = "subscription-v1s"
	SubscriptionCRDAPIVersion = "app.coreos.com/v1alpha1" // API version w/ CRD support
)

// SubscriptionState tracks when updates are available, installing, or service is up to date
type SubscriptionState string

const (
	SubscriptionStateNone             = ""
	SubscriptionStateUpgradeAvailable = "UpgradeAvailable"
	SubscriptionStateUpgradePending   = "UpgradePending"
	SubscriptionStateAtLatest         = "AtLatestKnown"
)

// SubscriptionSpec defines an Application that can be installed
type SubscriptionSpec struct {
	CatalogSource string `json:"source"`
	Package       string `json:"name"`
	Channel       string `json:"channel,omitempty"`

	StartingCSV string `json:"startingCSV,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   *SubscriptionSpec  `json:"spec"`
	Status SubscriptionStatus `json:"status"`
}

type SubscriptionStatus struct {
	CurrentCSV string                `json:"installedCSV,omitempty"`
	Install    *InstallPlanReference `json:"installplan,omitempty"`

	State       SubscriptionState `json:"state,omitempty"`
	LastUpdated metav1.Time       `json:"lastUpdated"`
}

type InstallPlanReference struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Name       string    `json:"name"`
	UID        types.UID `json:"uuid"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Subscription `json:"items"`
}
