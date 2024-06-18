package main

import (
	"os"
	"os/exec"

	"todox-with-sqlite/internal"
	"github.com/paganotoni/tailo"
)

func main() {
	tailo.Build(internal.TailoOptions...)

	cmd := exec.Command("go", "build")
	cmd.Args = append(
		cmd.Args,

		`--ldflags`, `-linkmode=external -extldflags="-static"`,
		`-tags`, `osusergo,netgo,musl`,
		`-buildvcs=false`,
		"-o", "bin/app",
		"cmd/app/main.go",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
