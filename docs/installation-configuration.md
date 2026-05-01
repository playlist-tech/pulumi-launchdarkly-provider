---
title: Launch Darkly Installation & Configuration
meta_desc: Information on how to install the Launch Darkly provider.
layout: installation
---

## Installation

The Pulumi Launch Darkly provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@playlist-tech/pulumi-launchdarkly`](https://www.npmjs.com/package/@playlist-tech/pulumi-launchdarkly)
* Python: [`playlist_pulumi_launchdarkly`](https://pypi.org/project/playlist-pulumi-launchdarkly/)
* Go: [`github.com/playlist-tech/pulumi-launchdarkly/sdk/go/launchdarkly`](https://pkg.go.dev/github.com/playlist-tech/pulumi-launchdarkly/sdk)
* .NET: [`Playlist.Pulumi.LaunchDarkly`](https://www.nuget.org/packages/Playlist.Pulumi.LaunchDarkly)

### Provider Binary

The Launch Darkly provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource launchdarkly <version>
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi Launch Darkly provider, you need to have Launch Darkly credentials. Launch Darkly maintains documentation on how to create API keys [here](https://docs.launchdarkly.com/home/account-security/api-access-tokens)

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Grafana:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export LAUNCHDARKLY_ACCESS_TOKEN=<LAUNCHDARKLY_ACCESS_TOKEN>
$ export LAUNCHDARKLY_OAUTH_TOKEN=<LAUNCHDARKLY_OAUTH_TOKEN>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export LAUNCHDARKLY_ACCESS_TOKEN=<LAUNCHDARKLY_ACCESS_TOKEN>
$ export LAUNCHDARKLY_OAUTH_TOKEN=<LAUNCHDARKLY_OAUTH_TOKEN>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:LAUNCHDARKLY_ACCESS_TOKEN = "<LAUNCHDARKLY_ACCESS_TOKEN>"
> $env:LAUNCHDARKLY_OAUTH_TOKEN = "<LAUNCHDARKLY_OAUTH_TOKEN>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set launchdarkly:<option>` or pass options to the [constructor of `new launchdarkly.Provider`]({{< relref "/registry/packages/launchdarkly/api-docs/provider" >}}).
