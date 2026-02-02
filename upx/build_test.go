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
	"path/filepath"
	"testing"

	"github.com/buildpacks/libcnb/v2"
	"github.com/dmikusa/bptest"
	. "github.com/dmikusa/bptest/matchers"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/libpak/v2/log"
	"github.com/paketo-buildpacks/upx/v3/upx"
	"github.com/sclevine/spec"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		logger    log.Logger
		buildTest *bptest.BuildTest
	)

	it.Before(func() {
		var err error

		testDepsDir, err := filepath.Abs("testdata")
		Expect(err).NotTo(HaveOccurred())

		logger = log.NewPaketoLogger(os.Stdout)

		buildTest = bptest.NewBuildTest().
			WithBuildpack(bptest.BuildpackConfig{
				API:     "0.8",
				ID:      "example/my-buildpack",
				Version: "1.0.0"}).
			WithBuildpackDependencyCache(testDepsDir).
			WithDependencyMetadata(bptest.DependencyMetadata{
				ID:      "upx",
				Version: "3.96",
				Stacks:  []string{"test-stack-id", "*"},
				CPEs:    []string{"cpe:2.3:a:upx_project:upx:3.96:*:*:*:*:*:*:*"},
				PURL:    "pkg:generic/upx@3.96",
				URI:     "https://localhost/stub-upx.tar.xz",
				SHA256:  "9645730740af103136b4afff7072bb5c511290907a4fde2c7dd6d89ce8e30eca",
			}).
			WithPlanEntry("upx", map[string]any{})
	})

	it("installs UPX when requested in the build plan", func() {
		result := buildTest.ExecuteT(t, upx.NewBuild(logger))

		Expect(result).To(HaveSucceeded())
		Expect(result).To(HaveExactlyLayers("upx"))
		Expect(result.Layer("upx")).To(HaveFile(filepath.Join("bin", "upx")))
	})

	it("does not install UPX when absent from the build plan", func() {
		result := buildTest.
			WithPlan(libcnb.BuildpackPlan{}).
			ExecuteT(t, upx.NewBuild(logger))

		Expect(result).To(HaveSucceeded())
		Expect(result.Layers().Count()).To(Equal(0))
	})
}
