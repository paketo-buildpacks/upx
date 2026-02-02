//go:build integration

package integration

import (
	"os"
	"testing"

	"github.com/dmikusa/bptest"
	. "github.com/dmikusa/bptest/matchers"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/libpak/v2/log"
	"github.com/paketo-buildpacks/upx/v3/upx"
)

func TestBuildApplication(t *testing.T) {
	logger := log.NewPaketoLogger(os.Stdout)

	result := bptest.NewBuildTestFromBuildpack("../buildpack.toml").
		WithAppFileString("test.py", "print('Hello world!')").
		WithPlanEntry("upx", map[string]any{}).
		ExecuteT(t, upx.NewBuild(logger))

	g := NewWithT(t)

	g.Expect(result).To(HaveSucceeded())
	g.Expect(result).To(HaveExactlyLayers("upx"))
	g.Expect(result.Layer("upx")).To(HaveFileWithPerms("bin/upx", 0755))
}
