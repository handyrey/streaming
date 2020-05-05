package streaming

import (
	"os"
	"os/exec"
)

func convertToHls(videoFile, id, resolution, hlsTime string) ([]byte, error) {
	_, err := os.Stat("./video/" + id)

	if !os.IsExist(err) {
		err := os.RemoveAll("./video/" + id)
		if err != nil {
			return nil, err
		}
	}

	err = os.Mkdir("./video/"+id, 0755)
	if err != nil {
		return nil, err
	}

	output, err := exec.Command("ffmpeg",
		"-i", videoFile,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-s", resolution,
		"-start_number", "0",
		"-hls_time", hlsTime,
		"-hls_list_size", "0",
		"-f", "hls",
		"video/"+id+"/index.m3u8",
	).Output()

	err = os.Remove(videoFile)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return output, err
	}

	return output, nil
}
