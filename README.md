# GoLogFile

gologfile — file adapter for [golog](https://github.com/memUsins/golog)

## Installation

```shell
$ go get github.com/memUsins/gologfile
```

## Usage

```go
// simple
logger := golog.NewLogger(gologfile.NewDefaultFileAdapterWithLevel(golog.ErrorLevel))

// with default config
logger := golog.NewLogger(gologfile.NewDefaultFileAdapter())

// with custom config
logger := golog.NewLogger(gologfile.NewFileAdapter(&gologfile.FileConfig{
	Enable: true, 
	Level: golog.ErrorLevel,
	LJLogger: &lumberjack.Logger{
		Filename:   "logs/error.log", 
		MaxSize:    10, 
		MaxBackups: 3, 
		MaxAge:     28, 
		LocalTime:  false,
		Compress:   false,
	},
}))
```