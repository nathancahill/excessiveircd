// Copyright (c) 2014 Michael Johnson. All rights reserved.
//
// Use of this source code is governed by the BSD license that can be found in
// the LICENSE file.

package config_test

import (
	"testing"

	"github.com/nightexcessive/excessiveircd/config"
)

func TestSetGet(t *testing.T) {
	const (
		key   = "key"
		value = "value"
	)
	err := config.Set(key, value)
	if err != nil {
		t.Fatalf("Error setting key: %s", err)
	}

	var outStr string
	err = config.Get(key, &outStr)
	if err != nil {
		t.Fatalf("Error getting key: %s", err)
	}
	if outStr != value {
		t.Errorf("Value mismatch: %q != %q", value, outStr)
	}
}
