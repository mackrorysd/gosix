# GOSIX

My goal here is to build a very light UNIX-like / POSIX-like environment,
similar to [BusyBox](https://www.busybox.net/), but in [Go](https://go.dev/).

It may grow into more - I've always wanted to build my own OS - but for now the
goal is to sharpen my Go skills after a long hiatus and become better acquainted
with the details of the POSIX spec.

Right now this is a minimal proof of concept demonstrating using hard links to
have a single binary implement multiple commands, enough shell functionality for
the Dockerfile to run, and basic I/O to be tested.

## Usage

Build and test everything:

```bash
make
```

Run an interactive shell:

```bash
make run
```