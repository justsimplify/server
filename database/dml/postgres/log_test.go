// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package postgres

import (
	"reflect"
	"testing"
)

func TestPostgres_createLogService(t *testing.T) {
	// setup types
	want := &Service{
		List: map[string]string{
			"all":   ListLogs,
			"build": ListBuildLogs,
		},
		Select: map[string]string{
			"step":    SelectStepLog,
			"service": SelectServiceLog,
		},
		Delete: DeleteLog,
	}

	// run test
	got := createLogService()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("createLogService is %v, want %v", got, want)
	}
}
