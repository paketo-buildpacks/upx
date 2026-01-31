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

package upx

import (
	"fmt"

	"github.com/buildpacks/libcnb/v2"
	"github.com/paketo-buildpacks/libpak/v2"
	"github.com/paketo-buildpacks/libpak/v2/log"
)

func NewBuild(logger log.Logger) libcnb.BuildFunc {
	return libpak.ContributableBuildFunc(func(context libcnb.BuildContext, result *libcnb.BuildResult) ([]libpak.Contributable, error) {
		logger.Title(context.Buildpack.Info.Name, context.Buildpack.Info.Version, context.Buildpack.Info.Homepage)

		md, err := libpak.NewBuildModuleMetadata(context.Buildpack.Metadata)
		if err != nil {
			return nil, fmt.Errorf("unable to create build module metadata\n%w", err)
		}

		dc, err := libpak.NewDependencyCache(context.Buildpack.Info.ID, context.Buildpack.Info.Version, context.Buildpack.Path, context.Platform.Bindings, logger)
		if err != nil {
			return nil, fmt.Errorf("unable to create dependency cache\n%w", err)
		}

		pr := libpak.PlanEntryResolver{Plan: context.Plan}

		var contributables []libpak.Contributable

		if _, ok, err := pr.Resolve(PlanEntryUpx); err != nil {
			return nil, fmt.Errorf("unable to resolve UPX plan entry\n%w", err)
		} else if ok {
			dr, err := libpak.NewDependencyResolver(md, context.StackID)
			if err != nil {
				return nil, fmt.Errorf("unable to create dependency resolver\n%w", err)
			}

			upxDependency, err := dr.Resolve(PlanEntryUpx, "")
			if err != nil {
				return nil, fmt.Errorf("unable to find dependency\n%w", err)
			}

			u := NewUpx(upxDependency, dc, logger)
			contributables = append(contributables, &u)
		}

		return contributables, nil
	})
}
