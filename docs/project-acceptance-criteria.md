# Topo Project acceptance criteria

Topo Projects extend the [Compose Specification](https://compose-spec.io), so any Compose Project can become a Topo Project and be used with Topo.

However, to be considered for inclusion in the default [Topo Project Catalog](https://github.com/arm/topo-project-catalog), Projects are assessed against the following criteria.

## Scope and purpose

### Clear value for the end user

A Topo Project should deliver on one or more of the following:

#### Show the user how to leverage novel features of their hardware target

The Topo Project Specification supports describing the hardware features required by the demo, such as SIMD extensions like [SVE](https://developer.arm.com/architectures/scalable-vector-extensions). Topo can dynamically filter projects based on the availability of those features on a given target.

#### Support configuration

The Topo Project Specification supports parameterization, allowing users to run `topo configure` for a given project to meet their specific needs. Projects should consider which `x-topo.parameters` they might expose to allow customization.

#### Leverage multiple processor subsystems with [remoteproc-runtime](https://github.com/arm/remoteproc-runtime)

Topo is compatible with remoteproc-runtime, supporting detection and automated installation of the runtime on the target. We welcome contributions that showcase heterogeneous applications enabled by remoteproc-runtime.

#### Showcase an end-to-end use case

A Project may be considered for inclusion in the catalog if it demonstrates a sufficiently interesting end-to-end software use case.

### Be extensible and adoptable

Topo Projects are expected to be extensible. The goal is to help users bootstrap a working application quickly while providing the complete source code and build toolchain needed to modify and extend the project for their own use case.

#### Containerise all the build steps; don't package binaries

Where possible, container images should be built from the included source rather than importing pre-built binary blobs. This gives users the greatest opportunity to extend and modify the project.

#### Document how the project works and suggest how to extend it

The `README.md` for a Project should provide an overview of how the application works, including links to the key entry points. It should also suggest next steps or explain how users can modify the Project and rerun `topo deploy` to see their changes.

## Quality

### Only require `topo deploy` to build and run

Running `topo deploy` must be sufficient to build and start the application. Define every required build, dependency-fetching, and setup step in `compose.yaml` or its Dockerfiles, or provide the result in a referenced container image. Do not require users to run additional commands manually.

[Multi-stage builds](https://docs.docker.com/build/building/multi-stage/) can be used for compilation, dependency fetching, code generation, and asset bundling. Copy only the resulting artifacts and runtime dependencies into the final stage. This ensures `topo deploy` performs the complete build while keeping build tools, caches, and other build-only files out of the runtime image.

### Be semantically correct and leverage `x-topo` attributes as appropriate

`x-topo` contains attributes that help users discover and use your Project. Ensure you have considered all available attributes in the schema and used them as appropriate.

The [Topo project-authoring skills](https://github.com/arm/topo#project-authoring-skills) can help you lint and improve your Project.

### Declare compatibility correctly and ensure the project is only recommended on compatible targets

Use the `x-topo.features` attribute to declare what features your project requires. Topo will use this to avoid recommending your project on targets where it is not compatible.

### Correct, reliable, tested, and understood

To offer Topo users a good experience, the catalog must contain only Projects that are reliable, well tested, and novel. Generated prototypes that have not been carefully reviewed and validated risk undermining trust in the catalog and are unlikely to provide value beyond what users could generate themselves.

The use of LLM-assisted tools is not disqualifying, but contributors must understand and stand behind the submitted code. Submissions with little evidence of validation or thoughtful engineering will not be accepted.

### Be fast to build and iterate

Topo Projects are expected to be fast to build and deploy.

See [Build Optimization](https://github.com/arm/topo/blob/main/docs/project-specification/04-build-optimization.md) or use the [`topo-project-optimize-deployment` skill](https://github.com/arm/topo#project-authoring-skills).
