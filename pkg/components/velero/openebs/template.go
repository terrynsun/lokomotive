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

package openebs

const chartValuesTmpl = `
configuration:
  provider: {{ .OpenEBS.BackupStorageLocation.Provider }}
  backupStorageLocation:
    {{- if .OpenEBS.BackupStorageLocation.Provider}}
    provider: {{ .OpenEBS.BackupStorageLocation.Provider }}
    {{- end }}
    {{- if .OpenEBS.BackupStorageLocation.Name}}
    name: {{ .OpenEBS.BackupStorageLocation.Name }}
    {{- end }}
    bucket: {{ .OpenEBS.BackupStorageLocation.Bucket }}
    config:
      region: {{ .OpenEBS.BackupStorageLocation.Region }}
  volumeSnapshotLocation:
    provider: openebs.io/cstor-blockstore
    {{- if .OpenEBS.VolumeSnapshotLocation.Name }}
    name: {{ .OpenEBS.VolumeSnapshotLocation.Name }}
    {{- end }}
    config:
      {{- if .OpenEBS.VolumeSnapshotLocation.Provider}}
      provider: {{ .OpenEBS.VolumeSnapshotLocation.Provider }}
      {{- end }}
      {{- if .OpenEBS.VolumeSnapshotLocation.Bucket }}
      bucket: {{ .OpenEBS.VolumeSnapshotLocation.Bucket }}
      {{- end }}
      {{- if .OpenEBS.VolumeSnapshotLocation.Prefix }}
      prefix: {{ .OpenEBS.VolumeSnapshotLocation.Prefix }}
      {{- end }}
      {{- if .OpenEBS.VolumeSnapshotLocation.Region }}
      region: {{ .OpenEBS.VolumeSnapshotLocation.Region }}
      {{- end }}
      {{- if .OpenEBS.VolumeSnapshotLocation.OpenEBSNamespace }}
      namespace: {{ .OpenEBS.VolumeSnapshotLocation.OpenEBSNamespace }}
      {{- end }}
      {{- if .OpenEBS.VolumeSnapshotLocation.S3URL }}
      s3_url: {{ .OpenEBS.VolumeSnapshotLocation.S3URL }}
      {{- end }}
      {{- if .OpenEBS.VolumeSnapshotLocation.Local }}
      local: {{ .OpenEBS.VolumeSnapshotLocation.Local }}
      {{- end }}
credentials:
  secretContents:
  {{- if .OpenEBS.Credentials }}
    cloud: |
{{ .OpenEBS.CredentialsIndented }}
  {{- end }}
metrics:
  enabled: {{ .Metrics.Enabled }}
  serviceMonitor:
    enabled: {{ .Metrics.ServiceMonitor }}
    additionalLabels:
      release: prometheus-operator
initContainers:
- image: openebs/velero-plugin:2.0.0
  imagePullPolicy: IfNotPresent
  name: velero-plugin-for-openebs
  resources: {}
  terminationMessagePath: /dev/termination-log
  terminationMessagePolicy: File
  volumeMounts:
  - mountPath: /target
    name: plugins
{{- if eq .OpenEBS.BackupStorageLocation.Provider "aws" }}
- image: velero/velero-plugin-for-aws:v1.1.0
  imagePullPolicy: IfNotPresent
  name: velero-plugin-for-aws
  resources: {}
  terminationMessagePath: /dev/termination-log
  terminationMessagePolicy: File
  volumeMounts:
  - mountPath: /target
    name: plugins
{{- end }}
{{- if eq .OpenEBS.BackupStorageLocation.Provider "gcp" }}
- image: velero/velero-plugin-for-gcp:v1.1.0
  imagePullPolicy: IfNotPresent
  name: velero-plugin-for-gcp
  resources: {}
  terminationMessagePath: /dev/termination-log
  terminationMessagePolicy: File
  volumeMounts:
  - mountPath: /target
    name: plugins
{{- end }}
`
