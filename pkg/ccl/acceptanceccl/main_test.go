// Copyright 2016 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package acceptanceccl

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/acceptance"
)

func TestMain(m *testing.M) {
	acceptance.MainTest(m)
}

//go:generate ../../util/leaktest/add-leaktest.sh *_test.go
