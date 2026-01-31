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

	"github.com/buildpacks/libcnb/v2"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/libpak/v2/log"
	"github.com/paketo-buildpacks/upx/v3/upx"
	"github.com/sclevine/spec"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx    libcnb.BuildContext
		logger log.Logger
	)

	it.Before(func() {
		var err error

		ctx.ApplicationPath, err = os.MkdirTemp("", "build")
		Expect(err).NotTo(HaveOccurred())

		logger = log.NewPaketoLogger(os.Stdout)

		ctx.Buildpack.Metadata = map[string]interface{}{
			"dependencies": []map[string]interface{}{
				{
					"id":      "upx",
					"version": "3.96",
					"stacks":  []interface{}{"test-stack-id", "*"},
					"cpes":    []interface{}{"cpe:2.3:a:upx_project:upx:3.96:*:*:*:*:*:*:*"},
					"purl":    "pkg:generic/upx@3.96",
				},
			},
		}
		ctx.StackID = "test-stack-id"
		t.Setenv("BP_ARCH", "amd64")
	})

	it.After(func() {
		Expect(os.RemoveAll(ctx.ApplicationPath)).To(Succeed())
	})

	it("does not contribute UPX when not in plan", func() {
		// No plan entry for upx
		result, err := upx.NewBuild(logger)(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(BeEmpty())
	})
}
