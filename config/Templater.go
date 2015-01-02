package config

import (
	"html/template"
	"path/filepath"
	"sync"
)

var cachedTemplates = map[string]*template.Template{}
var cachedMutex sync.Mutex

func T(name string) *template.Template {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	if t, ok := cachedTemplates[name]; ok {
		return t
	}

	t := template.New("_base.html")

	t = template.Must(template.ParseFiles("./views/_base.html", filepath.Join("./views/", name)))

	cachedTemplates[name] = t

	return t
}