package chiffmpeg

import (
	"testing"
	"time"
)

var ff = &FfmpegTools{
	FFMpeg:         "/opt/homebrew/bin/ffmpeg",
	FFProbe:        "/opt/homebrew/bin/ffprobe",
	CommandTimeout: time.Duration(30 * time.Second),
}

func TestFfmpegTools_Thumbnail(t *testing.T) {

	type args struct {
		src       string
		duration  float64
		dst       string
		overwrite bool
	}

	tests := []struct {
		name string

		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				src:       "samples/sample.mp4",
				duration:  3.0,
				dst:       "samples/sample.jpg",
				overwrite: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ff.Thumbnail(tt.args.src, tt.args.dst, tt.args.duration, tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("Thumbnail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestFfmpegTools_Duration(t *testing.T) {
	td, _ := time.ParseDuration("11.52s")

	type args struct {
		src string
	}
	tests := []struct {
		name string

		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				src: "samples/sample.mp4",
			},
			want:    td,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ff.Duration(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Duration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFfmpegTools_Transcoding(t *testing.T) {

	type args struct {
		src       string
		dst       string
		overwrite bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				src:       "samples/sample.mov",
				dst:       "samples/sample_out.mp4",
				overwrite: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ff.Transcoding(tt.args.src, tt.args.dst, tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("Transcoding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
