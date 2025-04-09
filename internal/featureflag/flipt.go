package featureflag

import (
	flipt "github.com/open-feature/go-sdk-contrib/providers/flipt/pkg/provider"
)

func newFlipOpenFeatureProvider() *flipt.Provider {
	provider := flipt.NewProvider(
		flipt.WithAddress("localhost:9000"),
		flipt.ForNamespace("default"),
	)
	return provider
}
