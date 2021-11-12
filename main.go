package main

import (
	_ "embed"
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

//go:embed eff_large_wordlist.txt
var wordlist string

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
	flag.UintVar(&opt.n, "n", 4, "number of words in generated password")

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
	if options.version {
		version()
		os.Exit(0)
	}

	printPassword()
	os.Exit(0)
}

func version() {
	fmt.Println(APPNAME + " version " + VERSION)
}

func printPassword() {
	if options.passFileName == "" {
		if err := printPasswordFromString(wordlist, options.n); err != nil {
			logger.Error.Fatalf("%v: internal word list: number of words: %d: %v", APPNAME, options.n, err)
		}
	} else {
		if err := printPasswordFromFile(options.passFileName, options.n); err != nil {
			logger.Error.Fatalf("%v: word list file: %q: number of words: %d: %v", APPNAME, options.passFileName, options.n, err)
		}
	}
}

func printPasswordFromString(words string, n uint) error {
	pw, err := internal.GeneratePassword(strings.NewReader(words), n)
	if err != nil {
		return err
	}

	fmt.Println(pw)
	return nil
}

func printPasswordFromFile(file string, n uint) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	pw, err := internal.GeneratePassword(f, n)
	if err != nil {
		return err
	}

	fmt.Println(pw)
	return nil
}
