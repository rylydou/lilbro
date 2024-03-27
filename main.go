package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/robfig/cron"

	"github.com/rylydou/lilbro/core"
)

func main() {
	schedule := cron.New()

	var app = core.App{}
	app.LoadConfig()

	var cam = app.AddCamera()
	cam.Id = "cam1"
	cam.Name = "Primary Camera"
	cam.Path = "/dev/video0"
	schedule.AddFunc("@every 15s", func() {
		cam.Capture()
	})
	schedule.AddFunc("@every 1m", func() {
		cam.Archive()
	})

	fmt.Println("Taking initial capture")
	cam.Capture()
	cam.Archive()

	fmt.Println("Starting CRON")
	schedule.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Path[1:]
		// Serve latest
		if query == "" {
			if cam.Buffer == nil {
				fmt.Fprintf(w, "No picture has been taken yet!")
				return
			}
			w.Header().Set("Content-Type", "image/webp")
			w.Write(*cam.Buffer)
		}

		// Server from archive
		path := app.GetPath("archive/", query) + ".webp"
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(w, "Image not found")
			return
		}
		w.Header().Set("Content-Type", "image/webp")
		w.Write(data)
	})

	fmt.Println("Listening on port 8080", app.Port)
	if app.UseTLS {
		cert_key, _ := app.ReadFile("cert.key")
		cert_pem, _ := app.ReadFile("cert.pem")
		http.ListenAndServeTLS(":8080", string(cert_pem), string(cert_key), nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
