package docker

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/telepresenceio/telepresence/v2/pkg/proc"
)

// BuildImage builds an image from source. Stdout is silenced during those operations. The
// image ID is returned.
func BuildImage(ctx context.Context, context string, buildArgs []string) (string, error) {
	var dockerBuildCmd string

	if strings.HasPrefix(context, "gosh://") {
		dockerBuildCmd = "gosh"
	} else {
		dockerBuildCmd = "docker"
	}

	args := append([]string{"build", "--quiet"}, buildArgs...)
	cmd := proc.StdCommand(ctx, dockerBuildCmd, append(args, context)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// PullImage checks if the given image exists locally by doing docker image inspect. A docker pull is
// performed if no local image is found. Stdout is silenced during those operations.
func PullImage(ctx context.Context, image string) error {
	cmd := proc.StdCommand(ctx, "docker", "image", "inspect", image)
	cmd.Stderr = io.Discard
	cmd.Stdout = io.Discard
	if cmd.Run() == nil {
		// Image exists in the local cache, so don't bother pulling it.
		return nil
	}
	cmd = proc.StdCommand(ctx, "docker", "pull", image)
	// Docker run will put the pull logs in stderr, but docker pull will put them in stdout.
	// We discard them here, so they don't spam the user. They'll get errors through stderr if it comes to it.
	cmd.Stdout = io.Discard
	return cmd.Run()
}
