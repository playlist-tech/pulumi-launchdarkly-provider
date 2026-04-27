package launchdarkly

import (
	"context"
	_ "embed"
	"fmt"
	"path"
	"unicode"

	"github.com/launchdarkly/terraform-provider-launchdarkly/launchdarkly"
	"github.com/lbrlabs/pulumi-launchdarkly/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

//go:embed cmd/pulumi-resource-launchdarkly/bridge-metadata.json
var metadata []byte

const (
	mainPkg = "launchdarkly"
	mainMod = "index"
)

func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

func launchDarklyMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

func launchDarklyType(mod string, typ string) tokens.Type {
	return tokens.Type(launchDarklyMember(mod, typ))
}

func launchDarklyDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return launchDarklyMember(mod+"/"+fn, res)
}

func launchDarklyResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return launchDarklyType(mod+"/"+fn, res)
}

func Provider() tfbridge.ProviderInfo {
	p := pfbridge.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(launchdarkly.Provider()),
		launchdarkly.NewPluginProvider("0.0.0")(),
	)

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "launchdarkly",
		Version:           version.Version,
		DisplayName:       "Launch Darkly",
		Publisher:         "lbrlabs",
		LogoURL:           "https://raw.githubusercontent.com/lbrlabs/pulumi-launchdarkly/master/assets/logo.png",
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		Description:       "A Pulumi package for creating and managing launch darkly cloud resources.",
		Keywords:          []string{"pulumi", "launchdarkly", "lbrlabs"},
		Homepage:          "https://www.pulumi.com",
		Repository:        "https://github.com/lbrlabs/pulumi-launchdarkly",
		GitHubOrg:         "launchdarkly",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"access_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"LAUNCHDARKLY_ACCESS_TOKEN"},
				},
			},
			"oauth_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"LAUNCHDARKLY_OAUTH_TOKEN"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"launchdarkly_access_token":              {Tok: launchDarklyResource(mainMod, "AccessToken")},
			"launchdarkly_ai_config":                 {Tok: launchDarklyResource(mainMod, "AiConfig")},
			"launchdarkly_ai_config_variation":       {Tok: launchDarklyResource(mainMod, "AiConfigVariation")},
			"launchdarkly_ai_tool":                   {Tok: launchDarklyResource(mainMod, "AiTool")},
			"launchdarkly_audit_log_subscription":    {Tok: launchDarklyResource(mainMod, "AuditLogSubscription")},
			"launchdarkly_custom_role":               {Tok: launchDarklyResource(mainMod, "CustomRole")},
			"launchdarkly_destination":               {Tok: launchDarklyResource(mainMod, "Destination")},
			"launchdarkly_environment":               {Tok: launchDarklyResource(mainMod, "Environment")},
			"launchdarkly_feature_flag":              {Tok: launchDarklyResource(mainMod, "FeatureFlag")},
			"launchdarkly_feature_flag_environment":  {Tok: launchDarklyResource(mainMod, "FeatureFlagEnvironment")},
			"launchdarkly_flag_templates":            {Tok: launchDarklyResource(mainMod, "FlagTemplates")},
			"launchdarkly_flag_trigger":              {Tok: launchDarklyResource(mainMod, "FlagTrigger")},
			"launchdarkly_metric":                    {Tok: launchDarklyResource(mainMod, "Metric")},
			"launchdarkly_model_config":              {Tok: launchDarklyResource(mainMod, "ModelConfig")},
			"launchdarkly_project":                   {Tok: launchDarklyResource(mainMod, "Project")},
			"launchdarkly_relay_proxy_configuration": {Tok: launchDarklyResource(mainMod, "RelayProxyConfiguration")},
			"launchdarkly_segment":                   {Tok: launchDarklyResource(mainMod, "Segment")},
			"launchdarkly_team":                      {Tok: launchDarklyResource(mainMod, "Team")},
			"launchdarkly_team_member":               {Tok: launchDarklyResource(mainMod, "TeamMember")},
			"launchdarkly_team_role_mapping":         {Tok: launchDarklyResource(mainMod, "TeamRoleMapping")},
			"launchdarkly_view":                      {Tok: launchDarklyResource(mainMod, "View")},
			"launchdarkly_view_filter_links":         {Tok: launchDarklyResource(mainMod, "ViewFilterLinks")},
			"launchdarkly_view_links":                {Tok: launchDarklyResource(mainMod, "ViewLinks")},
			"launchdarkly_webhook":                   {Tok: launchDarklyResource(mainMod, "Webhook")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"launchdarkly_ai_config":            {Tok: launchDarklyDataSource(mainMod, "getAiConfig")},
			"launchdarkly_ai_config_variation":  {Tok: launchDarklyDataSource(mainMod, "getAiConfigVariation")},
			"launchdarkly_ai_tool":              {Tok: launchDarklyDataSource(mainMod, "getAiTool")},
			"launchdarkly_audit_log_subscription": {
				Tok: launchDarklyDataSource(mainMod, "getAuditLogSubscription"),
			},
			"launchdarkly_environment":  {Tok: launchDarklyDataSource(mainMod, "getEnvironment")},
			"launchdarkly_feature_flag": {Tok: launchDarklyDataSource(mainMod, "getFeatureFlag")},
			"launchdarkly_feature_flag_environment": {
				Tok: launchDarklyDataSource(mainMod, "getFeatureFlagEnvironment"),
			},
			"launchdarkly_flag_templates":            {Tok: launchDarklyDataSource(mainMod, "getFlagTemplates")},
			"launchdarkly_flag_trigger":              {Tok: launchDarklyDataSource(mainMod, "getFlagTrigger")},
			"launchdarkly_metric":                    {Tok: launchDarklyDataSource(mainMod, "getMetric")},
			"launchdarkly_model_config":              {Tok: launchDarklyDataSource(mainMod, "getModelConfig")},
			"launchdarkly_project":                   {Tok: launchDarklyDataSource(mainMod, "getProject")},
			"launchdarkly_relay_proxy_configuration": {Tok: launchDarklyDataSource(mainMod, "getRelayProxyConfiguration")},
			"launchdarkly_segment":                   {Tok: launchDarklyDataSource(mainMod, "getSegment")},
			"launchdarkly_team":                      {Tok: launchDarklyDataSource(mainMod, "getTeam")},
			"launchdarkly_team_member":               {Tok: launchDarklyDataSource(mainMod, "getTeamMember")},
			"launchdarkly_team_members":              {Tok: launchDarklyDataSource(mainMod, "getTeamMembers")},
			"launchdarkly_view":                      {Tok: launchDarklyDataSource(mainMod, "getView")},
			"launchdarkly_webhook":                   {Tok: launchDarklyDataSource(mainMod, "getWebhook")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
			PackageName: "@lbrlabs/pulumi-launchdarkly",
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PackageName: "lbrlabs_pulumi_launchdarkly",
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Lbrlabs.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
