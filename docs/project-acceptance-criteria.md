# Topo Project acceptance criteria

Topo Projects extend the [Compose Specification](https://compose-spec.io), so any Compose Project can become a Topo Project and be used with Topo.

However, to be considered for inclusion in the default [Topo Project Catalog](https://github.com/arm/topo-project-catalog), Projects are assessed against the following criteria.

## Follow general best practices

The Project specification defines a set of best practices:

[Topo Project Authoring Best Practices](https://github.com/arm/topo/blob/main/docs/project-specification/05-authoring-best-practices.md)

Projects in the catalog are expected to adhere to those practices.

## Clear value for the end user

A Topo Project should deliver on one or more of the following:

#### Show the user how to leverage novel features of their hardware target

The Topo Project Specification supports describing the hardware features required by the demo, such as SIMD extensions like [SVE](https://developer.arm.com/architectures/scalable-vector-extensions). Topo can dynamically filter projects based on the availability of those features on a given target.

#### Support configuration

The Topo Project Specification supports parameterization, allowing users to run `topo configure` for a given project to meet their specific needs. Projects should consider which `x-topo.parameters` they might expose to allow customization.

#### Leverage multiple processor subsystems with [remoteproc-runtime](https://github.com/arm/remoteproc-runtime)

Topo is compatible with remoteproc-runtime, supporting detection and automated installation of the runtime on the target. We welcome contributions that showcase heterogeneous applications enabled by remoteproc-runtime.

#### Showcase an end-to-end use case

A Project may be considered for inclusion in the catalog if it demonstrates a sufficiently interesting end-to-end software use case.

## Be extensible and adoptable

Topo Projects are expected to be extensible. The goal is to help users bootstrap a working application quickly while providing the complete source code and build toolchain needed to modify and extend the project for their own use case.

### Document how the project works and suggest how to extend it

The `README.md` for a Project should provide an overview of how the application works, including links to the key entry points. It should also suggest next steps or explain how users can modify the Project and rerun `topo deploy` to see their changes.

### Containerize everything necessary to modify the application

Ensure your project contains all the source and build steps necessary to allow the end user to extend or modify the core features. This means avoiding pre-built binaries or fetching packages which the user cannot modify, but does not extend to libraries or other images which support the application but which you would not reasonably expect the user to want to change.

## Correct, reliable, tested, and understood

To offer Topo users a good experience, the catalog must contain only Projects that are reliable, well tested, and novel. Generated prototypes that have not been carefully reviewed and validated risk undermining trust in the catalog and are unlikely to provide value beyond what users could generate themselves.

The use of LLM-assisted tools is not disqualifying, but contributors must understand and stand behind the submitted code. Submissions with little evidence of validation or thoughtful engineering will not be accepted.
