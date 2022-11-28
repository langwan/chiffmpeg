package chiffmpeg

import (
	"bytes"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type FfmpegTools struct {
	Ffmpeg  string `json:"ffmpeg"`
	Ffprobe string `json:"ffprobe"`
}

func (ff *FfmpegTools) Transcoding(src string, dst string) error {
	args := []string{"-i", src, "-c:v", "libx264", "-strict", "-2", dst}
	cmd := exec.Command(ff.Ffmpeg, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (ff *FfmpegTools) Duration(src string) (time.Duration, error) {
	c := fmt.Sprintf(`%s -i "%s" -show_format -v quiet | sed -n 's/duration=//p'`, ff.Ffprobe, src)
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

func (ff *FfmpegTools) Thumbnail(src string, st float64, dst string) error {
	args := []string{"-i", src, "-ss", fmt.Sprintf("%f", st), "-vframes", "1", dst}
	cmd := exec.Command(ff.Ffmpeg, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
