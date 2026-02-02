/*
 * Copyright 2018-2026 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package upx_test

import (
	"os"
	"testing"

	"github.com/dmikusa/bptest"
	. "github.com/dmikusa/bptest/matchers"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/libpak/v2/log"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/upx/v3/upx"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("detects with upx in the buildplan", func() {
		logger := log.NewPaketoLogger(os.Stdout)
		result := bptest.NewDetectTest().ExecuteT(t, upx.NewDetect(logger))

		Expect(result).To(HavePassed())
		Expect(result).To(HavePlan("upx"))
	})
}
