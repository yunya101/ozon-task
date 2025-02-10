package config

import (
	"context"
	"log"
	"os"
)

var (
	PostgressAddr string
	Port          int
	errLog        *log.Logger     = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog       *log.Logger     = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Ctx           context.Context = context.Background()
	UsePostgres   bool
)
