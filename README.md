# Cantor

An exploration into CanJS with webpack, babel, go, etc.

## Setup

- Have npm installed and a working go build environment.
- Check out this repo somewhere.
- Run `make` in the root of the repo.
- Browse to: http://127.0.0.1:9000/

## Artifacts

After running `make` a binary called `cantor` and a directory called
`build` should be created in the root of the repo. Run the binary to
serve both the API and the frontend. The `build` directory contains the
entire set of assets - the `static` directory is not needed to deploy
this application.

## Docker

See [Dockerfile](./Dockerfile) for an example of a bare container image
you can deploy. Hint: try running `make docker` if you have docker
setup. See [Makefile](./Makefile) for more details.

