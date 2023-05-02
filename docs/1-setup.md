Project setup
=============

Initialise the operator-sdk in the `breakfast/` directory:

```shell
mkdir breakfast
cd breakfast
operator-sdk init \
  --domain newredo.com \
  --plugins go/v4-alpha \
  --repo github.com/NewRedo/toaster-controller/breakfast
```

This creates a Go module, linked to the GitHub repository `NewRedo/toaster-controller` (which is hopefully this
repository; if not, change to suit). The Operator SDK creates skeleton project structure with many placeholders for you
to add your own functionality.

Look at the `1-setup` branch for the output from this step.

# Next step

[Creating your first API](./2-crds.md)
