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

package s3

var backendConfigTmpl = `
  backend "s3" {
    bucket = "{{ .Bucket }}"
    key    = "{{ .Key }}"
    region = "{{ .Region }}"
    {{- if .AWSCredsPath }}
    shared_credentials_file = "{{ .AWSCredsPath }}"
    {{- end }}
    {{- if .DynamoDBTable }}
    dynamodb_table = "{{ .DynamoDBTable }}"
    {{- end }}
  }
`
