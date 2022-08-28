package create

import (
	"fmt"
	"os/exec"
)

func CreateAPI(name string) error {
	if err := CreateFile(&FileConfig{
		Name: "openapi.yml.tmpl",
		Tmpl: "tools/create/templates/openapi.yml.tmpl",
		Dir:  "api",
		File: fmt.Sprintf("%s.yml", name),
		Data: map[string]string{
			"name": name,
		},
	}); err != nil {
		return err
	}
	return nil
}

func CreateNewServer(name string) error {
	mainFileCfg := &FileConfig{
		Name: "main.go.tmpl",
		Tmpl: "tools/create/templates/main.go.tmpl",
		Dir:  fmt.Sprintf("cmd/%s", name),
		File: "main.go",
		Data: map[string]string{
			"name": name,
		},
	}
	genFileCfg := &FileConfig{
		Name: "generate.go.tmpl",
		Tmpl: "tools/create/templates/generate.go.tmpl",
		Dir:  fmt.Sprintf("pkg/%s", name),
		File: "generate.go",
		Data: map[string]string{
			"name": name,
		},
	}
	serverFileCfg := &FileConfig{
		Name: "server.go.tmpl",
		Tmpl: "tools/create/templates/server.go.tmpl",
		Dir:  fmt.Sprintf("pkg/%s", name),
		File: "server.go",
		Data: map[string]string{
			"name": name,
		},
	}
	typesFileCfg := &FileConfig{
		Name: "types.go.tmpl",
		Tmpl: "tools/create/templates/types.go.tmpl",
		Dir:  fmt.Sprintf("pkg/%s", name),
		File: "types.go",
		Data: map[string]string{
			"name": name,
		},
	}
	dockerFileCfg := &FileConfig{
		Name: "dockerfile.tmpl",
		Tmpl: "tools/create/templates/dockerfile.tmpl",
		Dir:  fmt.Sprintf("cmd/%s", name),
		File: "Dockerfile",
		Data: map[string]string{
			"name": name,
		},
	}
	configFileCfg := &FileConfig{
		Name: "config.yml.tmpl",
		Tmpl: "tools/create/templates/config.yml.tmpl",
		Dir:  "configs",
		File: fmt.Sprintf("%s.yml", name),
		Data: map[string]string{
			"name": name,
		},
	}
	goDocFileCfg := &FileConfig{
		Name: "doc.go.tmpl",
		Tmpl: "tools/create/templates/doc.go.tmpl",
		Dir:  fmt.Sprintf("pkg/%s", name),
		File: "doc.go",
		Data: map[string]string{
			"name": name,
		},
	}

	if err := CreateFiles([]*FileConfig{genFileCfg, mainFileCfg, serverFileCfg, typesFileCfg, dockerFileCfg, configFileCfg, goDocFileCfg}); err != nil {
		return err
	}

	cmd := exec.Command("go", "generate")
	cmd.Dir = fmt.Sprintf("pkg/%s", name)
	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
