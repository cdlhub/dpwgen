package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cdlhub/dpwgen/internal"
)

// Options is application command line options
type Options struct {
	version      bool
	n            uint
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
		fmt.Fprintf(os.Stdout, "\t%s [ -n N ] <pass-file-name> \n", APPNAME)
		fmt.Fprintln(os.Stdout)

		flag.PrintDefaults()
	}

	flag.BoolVar(&opt.version, "version", false, "Print version number")
	flag.UintVar(&opt.n, "n", 6, "number of words in generated password")

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
		version()
		os.Exit(0)
	}

	pw, err := dpwgen(opt.passFileName, opt.n)
	if err != nil {
		log.Error.Fatalf("%v: word list file: %q: number of words: %d: %v", APPNAME, opt.passFileName, opt.n, err)
	}
	fmt.Println(pw)

	os.Exit(0)
}

func version() {
	fmt.Println(APPNAME + " version " + VERSION)
}

func dpwgen(fileName string, n uint) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return internal.GeneratePassword(f, n)
}
