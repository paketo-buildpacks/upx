/*
 * Copyright 2018-2020 the original author or authors.
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

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
	"github.com/paketo-buildpacks/libpak/crush"
)

type Upx struct {
	LayerContributor libpak.DependencyLayerContributor
	Logger           bard.Logger
}

func NewUpx(dependency libpak.BuildpackDependency, cache libpak.DependencyCache) (Upx, libcnb.BOMEntry) {
	contributor, entry := libpak.NewDependencyLayer(dependency, cache, libcnb.LayerTypes{
		Build: true,
		Cache: true,
	})
	return Upx{LayerContributor: contributor}, entry
}

func (u Upx) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	u.LayerContributor.Logger = u.Logger

	return u.LayerContributor.Contribute(layer, func(artifact *os.File) (libcnb.Layer, error) {
		u.Logger.Bodyf("Expanding to %s", layer.Path)
		if err := crush.ExtractTarXz(artifact, layer.Path, 1); err != nil {
			return libcnb.Layer{}, fmt.Errorf("unable to expand UPX\n%w", err)
		}

		return layer, nil
	})
}

func (u Upx) Name() string {
	return u.LayerContributor.LayerName()
}
