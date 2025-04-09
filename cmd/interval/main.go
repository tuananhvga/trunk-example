package main

import (
	"context"
	"fmt"
	"time"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/tuananhvga/trunk-example/internal/featureflag"
)

func main() {
	client := featureflag.New(featureflag.Flipt)
	for {
		val, err := client.BooleanValue(context.Background(), "boolean-flag", false, openfeature.NewEvaluationContext("user123", map[string]any{}))
		if err != nil {
			panic(err)
		}
		fmt.Println("booleanFlag:", val)
		time.Sleep(1 * time.Second)
	}
}
