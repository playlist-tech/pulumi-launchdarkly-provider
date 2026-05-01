package main

import (
	launchdarkly "github.com/playlist-tech/pulumi-launchdarkly/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"
)

func main() {
	tfgen.MainWithMuxer("launchdarkly", launchdarkly.Provider())
}
