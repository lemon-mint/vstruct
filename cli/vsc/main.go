package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	_ "embed"

	"github.com/lemon-mint/vstruct/utils"
)

var Version string

var VersionInfo = "vsc" + Version + " " + runtime.GOOS + "/" + runtime.GOARCH

//go:embed LICENSE.txt
var license string

func PrintUsage() {
	fmt.Printf("\nvstruct compiler\n")
	fmt.Printf("Usage:\n\n\t%s [options] <lang> <package name> <input file>\n", os.Args[0])

	fmt.Printf("\n\nOptions:\n")
	fmt.Printf("\t-o <output>\t\tOutput file name (default: <inputfile>.ext where ext is the language extension)\n")
	fmt.Printf("\t-s\t\t\tPrints the generated code to stdout\n")
	fmt.Printf("\t-v\t\t\tPrint version and exit\n")
	fmt.Printf("\t-h\t\t\tPrint help and exit\n")
	fmt.Printf("\t-l\t\t\tPrint license and exit\n")

	fmt.Printf("\n\nLanguages:\n")
	fmt.Printf("\tgo\t\t\tGo (https://go.dev/)\n")
	fmt.Printf("\tpython\t\t\tPython (https://www.python.org/)\n")
	fmt.Printf("\trust\t\t\tRust (https://www.rust-lang.org/)\n")
	fmt.Printf("\tdart\t\t\tDart (https://dart.dev/)\n")

	fmt.Printf("\n")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	options := make(map[string]string)
	var lang string
	var pkgname string
	var inputfile string

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-") {
			switch args[i] {
			case "-v":
				fmt.Printf("vsc version: %s\n", VersionInfo)
				os.Exit(0)
			case "-h":
				PrintUsage()
			case "-l":
				fmt.Printf("%s\n", license)
				os.Exit(0)
			case "-o":
				options[args[i]] = args[i+1]
				i++
			case "-s":
				options[args[i]] = "true"
			default:
				fmt.Printf("Error: Unknown option: %s\n", args[i])
				PrintUsage()
			}
		} else {
			if args[i] == "" {
				continue
			}

			if lang == "" {
				lang = args[i]
				if lang != "go" && lang != "python" && lang != "rust" && lang != "dart" {
					fmt.Printf("Error: Unknown language: %s\n", lang)
					PrintUsage()
					os.Exit(1)
				}
			} else if pkgname == "" {
				pkgname = args[i]
			} else if inputfile == "" {
				inputfile = args[i]
				if _, err := os.Stat(inputfile); os.IsNotExist(err) {
					fmt.Printf("Error: Input file not found: %s\n", inputfile)
					os.Exit(1)
				}
			} else {
				fmt.Printf("Error: too many arguments\n")
				PrintUsage()
				os.Exit(1)
			}
		}
	}

	if lang == "" || pkgname == "" || inputfile == "" {
		fmt.Printf("Error: too few arguments\n")
		PrintUsage()
	}

	var langExt map[string]string = map[string]string{
		"go":     ".go",
		"python": ".py",
		"rust":   ".rs",
		"dart":   ".dart",
	}

	var outputFileName string = inputfile + langExt[lang]

	var outputFile *os.File = nil
	if _, ok := options["-s"]; ok {
		outputFile = os.Stdout
	} else if _, ok := options["-o"]; ok {
		outputFileName = options["-o"]
	}

	if outputFile == nil {
		f, err := os.Create(outputFileName)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()
		outputFile = f
	}

	inputFile, err := os.Open(inputfile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	inputData, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	out := utils.BuildVstructCLI([]string{lang, pkgname, string(inputData)}, outputFile.Name())
	if out.Err != "" {
		fmt.Printf("Error: %s\n", out.Err)
		os.Exit(1)
	}

	_, err = outputFile.Write([]byte(out.Code))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
