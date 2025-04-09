package featureflag

import (
	flagsmithClient "github.com/Flagsmith/flagsmith-go-client/v3"
	flagsmith "github.com/open-feature/go-sdk-contrib/providers/flagsmith/pkg"
	of "github.com/open-feature/go-sdk/openfeature"
)

func newFlagsmithOpenFeatureProvider() of.FeatureProvider {
	apiKey := "ser.EAhMXAZ2ddQYtK9BwPhyeU"
	client := flagsmithClient.NewClient(apiKey,
		flagsmithClient.WithBaseURL("http://localhost:8001/api/v1/"),
	)
	provider := flagsmith.NewProvider(client, flagsmith.WithUsingBooleanConfigValue())
	return provider
}
