// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package runtime_env

const (
	RuntimeEnvCredentialIdColumn      = "runtime_env_credential_id"
	RuntimeEnvCredentialContentColumn = "content"
	RuntimeEnvIdColumn                = "runtime_env_id"
	RuntimeEnvLabelValueColumn        = "label_value"
	RuntimeEnvLabelKeyColumn          = "label_key"
)

const (
	NameColumn        = "name"
	DescriptionColumn = "description"
	StatusColumn      = "status"
	StatusTimeColumn  = "status_time"
	LabelKeyColumn    = "label_key"
	LabelValueColumn  = "label_value"
)

const (
	NameMinLength       = "1"
	NameMaxLength       = "255"
	LabelKeyMinLength   = "1"
	LabelKeyMaxLength   = "50"
	LabelValueMinLength = "1"
	LabelValueMaxLength = "255"
	LabelKeyFmt         = "^[a-zA-Z]([-_a-zA-Z0-9]*[a-zA-Z0-9])?$"
)

const (
	LabelRuntime = "runtime"
	LabelZone    = "zone"
)

var VmBaseRuntime = []string{"qingcloud"}
var CmBaseRuntime = []string{"kubernetes"}
