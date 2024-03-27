package core

import (
	"os"
	"path"
)

type App struct {
	RootDir  string
	Password string
	Port     string
	UseTLS   bool
	Cameras  []Camera
}

func (app *App) GetPath(relative_path ...string) string {
	return path.Join(app.RootDir, path.Join(relative_path...))
}

func (app *App) ReadFile(relative_path string) (data []byte, err error) {
	return os.ReadFile(app.GetPath(relative_path))
}

func (app *App) LoadConfig() error {
	app.RootDir = "_data"
	// data, err := app.ReadFile("config.toml")
	// if err != nil {
	// 	fmt.Println("[ERR] Failed to open config.toml", err)
	// 	return err
	// }
	// err = toml.Unmarshal(data, app)
	// if err != nil {
	// 	fmt.Println("[ERR] Failed to parse config.toml", err)
	// 	return err
	// }
	// fmt.Println("Loaded config.toml")
	return nil
}
