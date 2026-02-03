//go:build slow

package tests

import (
	"os"
	"testing"

	"github.com/dmikusa/bptest"
	. "github.com/dmikusa/bptest/matchers"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/libpak/v2/log"
	"github.com/paketo-buildpacks/upx/v3/upx"
)

// These are slower tests and are not going to run in CI for the time being
// They can be run locally with `go test -tags=slow ./...`

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
