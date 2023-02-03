package videos

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var ffmpegPath string

func ensureFfmpeg() error {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		if path, err = exec.LookPath("ffmpeg.exe"); err != nil {
			return fmt.Errorf("no ffmpeg or ffmpeg.exe found")
		}
	}
	ffmpegPath = path
	if err := runArgs(nil, "-version"); err != nil {
		return fmt.Errorf("ffmpeg command failed: %v", err)
	}
	return nil
}

func runArgs(stderr io.Writer, args ...string) error {
	cmd := exec.Command(ffmpegPath, args...)
	if cmd.Err != nil {
		hlog.Error("unable to look up ffmpeg", cmd.Err)
		return cmd.Err
	}
	if stderr != nil {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}
	if err := cmd.Run(); err != nil {
		hlog.Error("ffmpeg command failed", err)
		return err
	}
	return nil
}

func ValidateVideo(path string) error {
	buf := NewBuffer(128)
	if err := runArgs(buf,
		"-v", "error", // Only print errors
		"-i", path, // Input
		"-map", "0:1", "-f", "null", "-", // Pseudo output
	); err != nil {
		return fmt.Errorf("ffmpeg failed: %s", buf.String())
	}
	return nil
}

func GenerateCover(path string, output string) error {
	buf := NewBuffer(128)
	if err := runArgs(buf,
		"-v", "error", // Only print errors
		"-i", path, // Input
		"-vsync", "vfr", // Remove some useless frames

		// Extract one scene-changing keyframe as the cover,
		// as well as scale the output to fit into 256x256
		"-skip_frame", "nokey", "-frames:v", "1",
		"-vf", "select='gt(scene,0.4)',scale=w=256:h=256:force_original_aspect_ratio=decrease",
		"-y", output, // Overwrites
	); err != nil {
		return fmt.Errorf("ffmpeg failed: %s", buf.String())
	}
	if _, err := os.Stat(output); err != nil {
		// Probably the video is too short to provide any cover image.
		// We are to choose the first keyframe.
		buf := NewBuffer(128)
		if err := runArgs(buf,
			"-v", "error", // Only print errors
			"-i", path, // Input
			"-vsync", "vfr", // Remove some useless frames
			"-skip_frame", "nokey", "-frames:v", "1",
			"-vf", "scale=w=256:h=256:force_original_aspect_ratio=decrease",
			"-y", output, // Overwrites
		); err != nil {
			return fmt.Errorf("ffmpeg failed: %s", buf.String())
		}
		if _, err := os.Stat(output); err != nil {
			return fmt.Errorf("unable to produce a cover image")
		}
	}
	return nil
}
