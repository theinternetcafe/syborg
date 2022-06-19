package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

const (
	ERROR_TYPE    = -1
	VIDEO_TYPE    = 0
	PLAYLIST_TYPE = 1
)

type (
	videoResponse struct {
		Formats []struct {
			Url string `json:"url"`
		} `json:"formats"`
		Title string `json:"title"`
	}

	VideoResult struct {
		Media string
		Title string
	}

	PlaylistVideo struct {
		Id string `json:"id"`
	}

	Youtube struct {
		Conf *Config
	}
)

func (youtube Youtube) getType(input string) int {
	if strings.Contains(input, "upload_date") {
		return VIDEO_TYPE
	}
	if strings.Contains(input, "_type") {
		return PLAYLIST_TYPE
	}
	return ERROR_TYPE
}

func (youtube Youtube) Get(input string) (int, *string, error) {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ERROR_TYPE, nil, err
	}
	str := out.String()
	return youtube.getType(str), &str, nil
}

func (youtube Youtube) Video(input string) (*VideoResult, error) {
	var resp videoResponse
	err := json.Unmarshal([]byte(input), &resp)
	if err != nil {
		return nil, err
	}
	return &VideoResult{resp.Formats[0].Url, resp.Title}, nil
}

func (youtube Youtube) Playlist(input string) (*[]PlaylistVideo, error) {
	lines := strings.Split(input, "\n")
	videos := make([]PlaylistVideo, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var video PlaylistVideo
		fmt.Println("line,", line)
		err := json.Unmarshal([]byte(line), &video)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return &videos, nil
}
