package chiffmpeg

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type FfmpegTools struct {
	FFMpeg         string        `json:"ffmpeg"`
	FFProbe        string        `json:"ffprobe"`
	CommandTimeout time.Duration `json:"command_timeout"`
}

func (ff *FfmpegTools) Transcoding(src string, dst string, overwrite bool) (output []byte, err error) {
	args := []string{"-i", src, "-c:v", "libx264", "-strict", "-2", dst}
	if overwrite {
		args = append([]string{"-y"}, args...)
	}
	ctx, _ := context.WithTimeout(context.Background(), ff.CommandTimeout)
	cmd := exec.CommandContext(ctx, ff.FFMpeg, args...)
	output, err = cmd.Output()
	if ctx.Err() != nil {
		return output, ctx.Err()
	} else if err != nil {
		return output, err
	}
	return output, err
}

func (ff *FfmpegTools) Duration(src string) (time.Duration, error) {
	c := fmt.Sprintf(`%s -i "%s" -show_format -v quiet | sed -n 's/duration=//p'`, ff.FFProbe, src)
	out, err := exec.Command("bash", "-c", c).Output()
	if err != nil {
		return time.Duration(0), err
	}
	o := strings.TrimSpace(string(out))
	f64, err := strconv.ParseFloat(o, 64)
	fp := f64 * math.Pow(1000.0, 3.0)
	td := time.Duration(int64(math.Round(fp)))

	if err != nil {
		return td, err
	}

	return td, nil
}

func (ff *FfmpegTools) Thumbnail(src string, dst string, duration float64, overwrite bool) error {
	args := []string{"-i", src, "-ss", fmt.Sprintf("%f", duration), "-vframes", "1", dst}
	if overwrite {
		args = append([]string{"-y"}, args...)
	}
	cmd := exec.Command(ff.FFMpeg, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
