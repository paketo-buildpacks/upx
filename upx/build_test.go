/*
 * Copyright 2018-2023 the original author or authors.
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
	"testing"

	"github.com/buildpacks/libcnb/v2"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/upx/v4/upx"
	"github.com/sclevine/spec"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx    libcnb.BuildContext
		result libcnb.BuildResult
	)

	it.Before(func() {
		var err error

		ctx.ApplicationPath = t.TempDir()
		Expect(err).NotTo(HaveOccurred())

		ctx.Plan.Entries = append(ctx.Plan.Entries, libcnb.BuildpackPlanEntry{Name: "upx"})
		ctx.Buildpack.Metadata = map[string]interface{}{
			"dependencies": []map[string]interface{}{
				{
					"id":      "upx",
					"version": "3.96",
					"stacks":  []interface{}{"test-stack-id"},
				},
			},
		}
		ctx.StackID = "test-stack-id"

		result = libcnb.NewBuildResult()
	})

	it("contributes UPX", func() {
		ctx.Buildpack.Metadata = map[string]interface{}{
			"dependencies": []map[string]interface{}{
				{
					"id":      "upx",
					"version": "3.96",
					"stacks":  []interface{}{"test-stack-id"},
					"cpes":    []string{"cpe:2.3:a:upx_project:upx:3.96:*:*:*:*:*:*:*"},
					"purl":    "pkg:generic/upx@3.96?arch=amd64",
				},
			},
		}
		ctx.Buildpack.API = "0.7"
		layerContributors, err := upx.Build(ctx, &result)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Labels).To(ContainElement(libcnb.Label{
			Key:   "foo",
			Value: "bar",
		}))
		Expect(result.Processes).To(ContainElement(libcnb.Process{
			Command: []string{"bash -c \"sleep 99\""},
			Default: false,
		}))
		Expect(result.Unmet).To(ContainElement(libcnb.UnmetPlanEntry{Name: "baz"}))
		Expect(layerContributors).To(HaveLen(1))
		Expect(layerContributors[0].Name()).To(Equal("upx"))
	})
}
