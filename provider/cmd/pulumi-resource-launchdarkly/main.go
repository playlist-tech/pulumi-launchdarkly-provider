//go:generate go run ./generate.go

package main

import (
	"context"
	_ "embed"

	launchdarkly "github.com/lbrlabs/pulumi-launchdarkly/provider"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
)

//go:embed schema-embed.json
var pulumiSchema []byte

func main() {
	pfbridge.MainWithMuxer(context.Background(), "launchdarkly", launchdarkly.Provider(), pulumiSchema)
}
