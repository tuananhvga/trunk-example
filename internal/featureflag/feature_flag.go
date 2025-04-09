package featureflag

import "github.com/open-feature/go-sdk/openfeature"

type FeatureFlagProviderType string

// write enum of feature flag providers
const (
	Flipt FeatureFlagProviderType = "flipt"
	Flagsmith FeatureFlagProviderType = "flagsmith"
)

func New(providerType FeatureFlagProviderType) *openfeature.Client {
	switch providerType {
	case Flipt:
		openfeature.SetProviderAndWait(newFlipOpenFeatureProvider())
	case Flagsmith:
		openfeature.SetProviderAndWait(newFlagsmithOpenFeatureProvider())
	}
	client := openfeature.NewClient(string(providerType))
	return client
}
