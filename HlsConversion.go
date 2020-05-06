package streaming

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func init() {
	// Check if the FFMPEG is installed on machine
	err := isFfmpegAvailable()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the output folder exist if dont create one
	_, err = os.Stat("video/")
	if os.IsNotExist(err) {
		err = os.Mkdir("video/", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ConvertToHls(videoPath, id string, interval, quality int) ([]byte, error) {

	//check if the output folder exist, if does then remove it then create a new one
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

	//check resolution from inputed quality
	reso, err := getReso(quality)
	if err != nil {
		return nil, err
	}

	if interval == 0 {
		return nil, fmt.Errorf("Interval Cant Be Zero")
	}

	//execut FFMPEG commads to convert video to hls format
	output, err := exec.Command("ffmpeg",
		"-i", videoPath,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-s", reso,
		"-start_number", "0",
		"-hls_time", strconv.Itoa(interval),
		"-hls_list_size", "0",
		"-f", "hls",
		"video/"+id+"/index.m3u8",
	).Output()
	if err != nil {
		return output, err
	}

	// //remove temp video file
	// err = os.Remove(videoPath)
	// if err != nil {
	// 	return nil, err
	// }

	return output, nil
}

func getReso(quality int) (string, error) {
	switch quality {
	case 240:
		return "352x240", nil
	case 360:
		return "480x360", nil
	case 480:
		return "858x480", nil
	case 720:
		return "1280x720", nil
	case 1080:
		return "1920x1080", nil
	case 1440:
		return "2560x1440", nil
	case 2160:
		return "3840x2160", nil
	default:
		return "", fmt.Errorf("Resolution Not Supported")
	}
}

func isFfmpegAvailable() error {
	_, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		// Check if the user is on unix-like machine
		if !strings.Contains(err.Error(), "executable file not found") {
			return fmt.Errorf("FFMPEG is not installed on your machine. Please install the FFMPEG first.")
		}

		// Check if the user is on windows machine
		if !strings.Contains(err.Error(), "is not recognized as an internal") {
			return fmt.Errorf("FFMPEG is not installed on your machine. Please install the FFMPEG first.")
		}

		return err
	}

	return nil
}
