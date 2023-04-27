Toaster-Controller
==================

This project aims to demonstrate how a simple custom controller can be implemented.

# Getting Started

This project contains a (devcontainer)[https://containers.dev] definition. To get started, use either the VS Code
Dev Containers extension, or the devcontainer CLI which is installable using npm with the `@devcontainers/cli` package.

The container includes the Operator Framework SDK and its dependencies; Go, Kubectl and Docker CLI.

*It is likely that the docker socket is not writeable after the container is built. To fix this, run a shell outside
the devcontainer as follows:*

```shell
docker exec -ti <name of container> 

# Instructions

(Start the lessons)[docs/step1/index.md], or look at the branch **TBD** for a finished project.