package create

import (
	"bytes"
	"os"
	"path"
	"text/template"
)

func CreateDir(dir string) error {
	_ = os.MkdirAll(dir, os.ModePerm)
	return nil
	//if _, err := os.Stat(dir); os.IsNotExist(err) {
	//	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
	//		return err
	//	}
	//	return err
	//}
	//return nil
}

func CreateFile(config *FileConfig) error {
	var (
		buf bytes.Buffer
	)

	if err := CreateDir(config.Dir); err != nil {
		return err
	}

	//template.New(config.Tmpl)
	tmpl := template.Must(template.New(config.Name).Funcs(FuncMap).ParseFiles(config.Tmpl))
	//.ParseFiles(config.Tmpl))
	err := tmpl.Execute(&buf, config.Data)
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(config.Dir, config.File))
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func CreateFiles(cfgs []*FileConfig) error {
	for _, cfg := range cfgs {
		if err := CreateFile(cfg); err != nil {
			return err
		}
	}
	return nil
}
