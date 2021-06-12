# `gcr.io/paketo-buildpacks/compress-executable`

The Paketo Compress Executable Buildpack is a Cloud Native Buildpack that will compress an executable file, saving space in the resulting image.

## Behavior

This buildpack will participate all the following conditions are met

* TBD

The buildpack will do the following:

* If `upx` is selected as the compression mechanism, `upx` will be installed
* If `upx` is not selected `gzexe` will be used instead
* The target executable will be compressed using the selected tool and placed into the resulting layer

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
