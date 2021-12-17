package golang

import (
	"embed"
	"io"
	"sync"

	"github.com/lemon-mint/strtpl"
	"github.com/lemon-mint/vstruct/ir"
)

//go:embed codes/* codes/**
var tpls embed.FS

var cache map[string]*strtpl.TPL = make(map[string]*strtpl.TPL)
var cacheMutex sync.Mutex

func getTpl(name string) *strtpl.TPL {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	if tpl, ok := cache[name]; ok {
		return tpl
	}
	f, err := tpls.ReadFile(name)
	if err != nil {
		panic(err)
	}
	t := strtpl.NewTPL(string(f))
	cache[name] = t
	return t
}

func Generate(w io.Writer, i *ir.IR) error {
	panic("not implemented")
	return nil
}
