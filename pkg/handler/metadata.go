package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/quantinium03/conduit/utils"
)

var supportedFormats = struct {
	VideoContainer []string
	AudioContainer []string
	AudioCodec     []string
	VideoCodec     []string
}{
	VideoContainer: []string{"mp4", "mkv", "webm"},
	AudioContainer: []string{"mp3", "flac", "aac"},
	AudioCodec:     []string{"aac", "mp3", "flac"},
	VideoCodec:     []string{"h264", "h265", "vp9"},
}

type initRes struct {
	Error     string  `json:"error,omitempty"`
	Supported bool    `json:"supported,omitempty"`
	Type      string  `json:"type,omitempty"`
	BufLen    float64 `json:"bufLen,omitempty"`
}

type output struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

type Stream struct {
	CodecType string `json:"codec_type"`
	CodecName string `json:"codec_name"`
	Duration  string `json:"duration"`
}

type Format struct {
	FormatName string `json:"format_name"`
	Duration   string `json:"duration"`
}

func GetMetaData(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")
	cmd := exec.Command("ffprobe", "-v", "error", "-show_format", "-show_streams", "-of", "json", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		utils.ResponseWithErr(w, 500, fmt.Sprintf("ffprobe execution failed: %v", err))
		return
	}

	var ffprobeOut output
	err = json.Unmarshal(out, &ffprobeOut)
	if err != nil {
		utils.ResponseWithErr(w, 500, fmt.Sprintf("failed to parse ffprobe output: %v", err))
		return
	}

	format := strings.Split(ffprobeOut.Format.FormatName, ",")[0]
	var audioStream *Stream
	var videoStream *Stream

	for _, stream := range ffprobeOut.Streams {
		if stream.CodecType == "audio" {
			audioStream = &stream
		} else if stream.CodecType == "video" {
			videoStream = &stream
		}
	}

	var duration float64
	if videoStream != nil {
		duration, _ = parseDuration(videoStream.Duration)
	} else if audioStream != nil {
		duration, _ = parseDuration(audioStream.Duration)
	} else {
		duration, _ = parseDuration(ffprobeOut.Format.Duration)
	}

	isVideo := videoStream != nil && duration > 0.5
	if !isVideo && audioStream == nil {
		utils.ResponseWithErr(w, 404, fmt.Sprintf("Neither video nor audio stream found"))
	}

	supported := false
	if isVideo && contains(supportedFormats.VideoContainer, format) || (!isVideo && contains(supportedFormats.AudioCodec, format)) {
		if (audioStream == nil || contains(supportedFormats.AudioCodec, audioStream.CodecName)) && (videoStream == nil || contains(supportedFormats.VideoCodec, videoStream.CodecName)) {
			supported = true
		}
	}

	utils.ResponseWithJSON(w, 200, &initRes{
		Type:      getStreamType(isVideo),
		Supported: supported,
		BufLen:    10.0,
	})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getStreamType(isVideo bool) string {
	if isVideo {
		return "video"
	}
	return "audio"
}

func parseDuration(durationStr string) (float64, error) {
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid duration: %v", err)
	}
	return duration, nil
}
