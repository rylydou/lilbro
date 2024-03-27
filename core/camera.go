package core

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

type Camera struct {
	app    *App
	Id     string
	Path   string
	Name   string
	Buffer *[]byte
}

func (cam *Camera) Capture() []byte {
	args := [...]string{
		"-q",
		"--no-banner",
		"-d", cam.Path,
		"-r", "1920x1080",
		"-S", "2",
		"-F", "10",
		"--scale", "1280x720",
		"--webp", "-1",
		"-",
	}
	out, err := exec.Command("fswebcam", args[:]...).Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cam.Buffer = &out
	return out
}

func (cam *Camera) Archive() error {
	if cam.Buffer == nil {
		fmt.Println("Cannot save buffer - Buffer is nil!")
		return errors.New("nil buffer")
	}

	file_dir_path := path.Join(
		time.Now().Format("2006-01"),
		time.Now().Format("02"),
	)

	file_name := time.Now().Format("15-04") + ".webp"

	var err error
	err = os.MkdirAll(cam.app.GetPath("archive", file_dir_path), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = os.WriteFile(cam.app.GetPath("archive", file_dir_path, file_name), *cam.Buffer, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (app *App) AddCamera() Camera {
	cam := Camera{}
	cam.app = app
	return cam
}
