package script

import (
	"app/pkg/monkey"
	"app/pkg/monkey/object"
	"embed"
	"path/filepath"
)

//go:embed embed/*
var content embed.FS

func Import(filename string, eval monkey.Evaluator) object.Object {
	file, err := content.Open(filepath.Clean("embed/" + filename + ".mky"))
	if err != nil {
		return object.FromError(err)
	}
	defer file.Close()
	return eval.Eval(file)
}
