// Copyright 2020 The Lokomotive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package openebs deals with configuring Velero openebs plugin.
package openebs

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"

	"github.com/kinvolk/lokomotive/internal"
)

const indentation = 6

// Configuration contains OpenEBS specific parameters.
type Configuration struct {
	Credentials            string                  `hcl:"credentials"`
	BackupStorageLocation  *BackupStorageLocation  `hcl:"backup_storage_location,block"`
	VolumeSnapshotLocation *VolumeSnapshotLocation `hcl:"volume_snapshot_location,block"`
	CredentialsIndented    string
}

// BackupStorageLocation configures the backup storage location for OpenEBS plugin.
type BackupStorageLocation struct {
	Bucket   string `hcl:"bucket"`
	Region   string `hcl:"region"`
	Provider string `hcl:"provider,optional"`
	Name     string `hcl:"name,optional"`
}

// VolumeSnapshotLocation configures the volume snapshot location for OpenEBS plugin.
type VolumeSnapshotLocation struct {
	Bucket           string `hcl:"bucket"`
	Region           string `hcl:"region"`
	Provider         string `hcl:"provider,optional"`
	Name             string `hcl:"name,optional"`
	Prefix           string `hcl:"prefix,optional"`
	Local            bool   `hcl:"local,optional"`
	OpenEBSNamespace string `hcl:"openebs_namespace,optional"`
	S3URL            string `hcl:"s3_url,optional"`
}

// NewConfiguration returns the new default Configuration.
func NewConfiguration() *Configuration {
	return &Configuration{
		BackupStorageLocation: &BackupStorageLocation{
			Provider: "aws",
			Name:     "default",
		},
		VolumeSnapshotLocation: &VolumeSnapshotLocation{
			Provider:         "aws",
			Name:             "default",
			Prefix:           "cstor",
			OpenEBSNamespace: "openebs",
		},
	}
}

// ChartValuesTemplate returns the chart values template.
func (c *Configuration) ChartValuesTemplate() string {
	return chartValuesTmpl
}

// IndentCredentials indents the provided credentials.
func (c *Configuration) IndentCredentials() {
	c.CredentialsIndented = internal.Indent(c.Credentials, indentation)
}

// Validate validates OpenEBS specific parts in the configuration.
func (c *Configuration) Validate() hcl.Diagnostics {
	var diagnostics hcl.Diagnostics
	if c.Credentials == "" {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "'credentials' cannot be empty",
			Detail:   "No credentials found.",
		})
	}

	if c.BackupStorageLocation.Bucket == "" {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "'openebs.backup_storage_location.bucket' cannot be empty",
			Detail:   "Make sure to the set the field to valid non-empty value",
		})
	}

	if c.BackupStorageLocation.Region == "" {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "'openebs.backup_storage_location.region' cannot be empty",
			Detail:   "Make sure to the set the field to valid non-empty value",
		})
	}

	if c.VolumeSnapshotLocation.Bucket == "" {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "'openebs.backup_storage_location.bucket' cannot be empty",
			Detail:   "Make sure to the set the field to valid non-empty value",
		})
	}

	if c.VolumeSnapshotLocation.Region == "" {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "'openebs.backup_storage_location.region' cannot be empty",
			Detail:   "Make sure to the set the field to valid non-empty value",
		})
	}

	if !isSupportedProvider(c.BackupStorageLocation.Provider) {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary: fmt.Sprintf("openebs.backup_storage_location.provider must be one of: '%s'",
				openEBSSupportedProviders()),
			Detail: "Make sure to set provider to one of supported values",
		})
	}

	if !isSupportedProvider(c.VolumeSnapshotLocation.Provider) {
		diagnostics = append(diagnostics, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary: fmt.Sprintf("openebs.volume_snapshot_location.provider must be one of: '%s'",
				openEBSSupportedProviders()),
			Detail: "Make sure to set provider to one of supported values",
		})
	}

	return diagnostics
}

// isSupportedProvider checks if the provider is supported or not.
func isSupportedProvider(provider string) bool {
	for _, p := range openEBSSupportedProviders() {
		if provider == p {
			return true
		}
	}

	return false
}

// openEBSSupportedProviders returns the list of supported providers.
func openEBSSupportedProviders() []string {
	return []string{"aws", "gcp"}
}
