package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Options is application command line options
type Options struct {
	version      bool
	n            uint
	d            uint
	passFileName string
}

// Logger is all loggers
type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

var (
	options Options
	logger  Logger
)

func init() {
	initOptions(&options)
	initLogger(&logger)
}

func initOptions(opt *Options) {
	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, "Usage:")
		fmt.Fprintf(os.Stdout, "\t%s -version\n", APPNAME)
		fmt.Fprintf(os.Stdout, "\t%s [ -n N ] [ -d D ] <pass-file-name> \n", APPNAME)
		fmt.Fprintln(os.Stdout)

		flag.PrintDefaults()
	}

	flag.BoolVar(&opt.version, "version", false, "Print version number")
	flag.UintVar(&opt.n, "n", 6, "number of words in generated password")
	flag.UintVar(&opt.d, "d", 5, "number of dice to roll to select a word")

	flag.Parse()
	opt.passFileName = strings.Join(flag.Args(), " ")
}

func initLogger(logger *Logger) {
	*logger = Logger{
		Info:    log.New(os.Stderr, "INFO: ", 0),
		Warning: log.New(os.Stderr, "WARNING: ", 0),
		Error:   log.New(os.Stderr, "ERROR: ", 0),
	}
}

func main() {
	opt := options
	log := logger

	if opt.version {
		fmt.Println(APPNAME + " version " + VERSION)
		os.Exit(0)
	}

	f, err := os.Open(opt.passFileName)
	if err != nil {
		log.Error.Fatalf("cannot open password file %q: %v", opt.passFileName, err)
	}
	defer f.Close()

	os.Exit(0)
}
