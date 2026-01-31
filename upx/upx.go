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
	"os"
	"path/filepath"

	"github.com/buildpacks/libcnb/v2"
	"github.com/paketo-buildpacks/libpak/v2"
	"github.com/paketo-buildpacks/libpak/v2/crush"
	"github.com/paketo-buildpacks/libpak/v2/log"
)

type Upx struct {
	LayerContributor libpak.DependencyLayerContributor
	Logger           log.Logger
}

func NewUpx(dependency libpak.BuildModuleDependency, cache libpak.DependencyCache, logger log.Logger) Upx {
	contributor := libpak.NewDependencyLayerContributor(dependency, cache, libcnb.LayerTypes{
		Build: true,
		Cache: true,
	}, logger)
	return Upx{LayerContributor: contributor, Logger: logger}
}

func (u *Upx) Contribute(layer *libcnb.Layer) error {
	return u.LayerContributor.Contribute(layer, func(layer *libcnb.Layer, artifact *os.File) error {
		u.Logger.Bodyf("Expanding to %s", layer.Path)
		if err := crush.Extract(artifact, layer.Path, 1); err != nil {
			return fmt.Errorf("unable to expand UPX\n%w", err)
		}

		binDir := filepath.Join(layer.Path, "bin")

		if err := os.MkdirAll(binDir, 0755); err != nil {
			return fmt.Errorf("unable to mkdir\n%w", err)
		}

		if err := os.Symlink(filepath.Join(layer.Path, "upx"), filepath.Join(binDir, "upx")); err != nil {
			return fmt.Errorf("unable to symlink UPX\n%w", err)
		}

		return nil
	})
}

func (u Upx) Name() string {
	return u.LayerContributor.LayerName()
}
