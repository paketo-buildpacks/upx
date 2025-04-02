# Paketo Buildpack for UPX

## Buildpack ID: `paketo-buildpacks/upx`
## Registry URLs: `docker.io/paketobuildpacks/upx`

The Paketo Buildpack for UPX is a Cloud Native Buildpack that providex UPX a tool that can be used to compress executables.

## Behavior

This buildpack will participate all the following conditions are met

* Another buildpack requires `upx`

The buildpack will do the following:

* Contributes UPX to a layer marked `build` and `cache` with command on `$PATH`

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
