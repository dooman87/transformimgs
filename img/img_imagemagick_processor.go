package img

import (
	"bytes"
	"flag"
	"github.com/golang/glog"
	"log"
	"os/exec"
)

type ImageMagickProcessor struct {
	convertCmd string
}

var imagemagickConvertCmd string

func init() {
	flag.StringVar(&imagemagickConvertCmd, "imConvert", "", "Imagemagick convert command")
}

func CheckImagemagick() {
	if len(imagemagickConvertCmd) == 0 {
		log.Fatal("Command convert should be set by -imConvert flag")
		return
	}

	_, err := exec.LookPath(imagemagickConvertCmd)
	if err != nil {
		log.Fatalf("Imagemagick is not available '%s'", err.Error())
	}
}

//Using convert util from imagemagick package to resize
//image to specific size.
func (p *ImageMagickProcessor) Resize(data []byte, size string) ([]byte, error) {
	var out, cmderr bytes.Buffer
	cmd := exec.Command(imagemagickConvertCmd, "-", "-resize", size, "-")
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stdout = &out
	cmd.Stderr = &cmderr

	glog.Infof("Running resize command, args '%v'", cmd.Args)
	err := cmd.Run()
	if err != nil {
		glog.Errorf("Error executing convert command: %s", err.Error())
		glog.Errorf("ERROR: %s", cmderr.String())
		return nil, err
	}

	return out.Bytes(), nil
}
