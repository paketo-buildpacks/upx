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

package upx

import (
	"fmt"
	"os"

	"github.com/buildpacks/libcnb/v2"
	"github.com/paketo-buildpacks/libpak/v2"
	"github.com/paketo-buildpacks/libpak/v2/log"
)

func Build(context libcnb.BuildContext, result *libcnb.BuildResult) ([]libpak.Contributable, error) {
	logger := log.NewPaketoLogger(os.Stdout)
	logger.Title(context.Buildpack.Info.Name, context.Buildpack.Info.Version, context.Buildpack.Info.Homepage)

	dc, err := libpak.NewDependencyCache(context.Buildpack.Info.ID, context.Buildpack.Info.Version, context.Buildpack.Path, context.Platform.Bindings, logger)
	if err != nil {
		return []libpak.Contributable{}, fmt.Errorf("unable to create dependency cache\n%w", err)
	}
	dc.Logger = logger

	// Look! I can now set these in a layer contributor
	result.Labels = append(result.Labels, libcnb.Label{
		Key:   "foo",
		Value: "bar",
	})
	result.Processes = append(result.Processes, libcnb.Process{
		Command: []string{"bash -c \"sleep 99\""},
		Default: false,
	})
	result.Unmet = append(result.Unmet, libcnb.UnmetPlanEntry{Name: "baz"})

	pr := libpak.PlanEntryResolver{Plan: context.Plan}

	if _, ok, err := pr.Resolve(PlanEntryUpx); err != nil {
		return []libpak.Contributable{}, fmt.Errorf("unable to resolve UPX plan entry\n%w", err)
	} else if ok {
		buildModuleMetadata, err := libpak.NewBuildModuleMetadata(context.Buildpack.Metadata)
		if err != nil {
			return []libpak.Contributable{}, fmt.Errorf("unable to make build module metadata from %+v\n%w", context.Buildpack.Metadata, err)
		}

		dr, err := libpak.NewDependencyResolver(buildModuleMetadata, context.StackID)
		if err != nil {
			return []libpak.Contributable{}, fmt.Errorf("unable to create dependency resolver\n%w", err)
		}

		upxDependency, err := dr.Resolve(PlanEntryUpx, "")
		if err != nil {
			return []libpak.Contributable{}, fmt.Errorf("unable to find dependency\n%w", err)
		}

		upx := NewUpx(upxDependency, dc, logger)

		return []libpak.Contributable{upx}, nil
	}

	return []libpak.Contributable{}, nil
}
