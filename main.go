package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

var filename string

func init() {
	CPUprofiling()
}

func main() {

	flag.StringVar(&filename, "f", "", "Path to the input file")
	flag.Parse()
	if filename == "" {
		flag.Usage()
		log.Fatalln()
	}
	for k, v := range getResults(filename) {
		fmt.Printf("%-20s\t%d\n", k, v)
	}
	defer MEMprofiling()
}

func getResults(path string) map[string]int {
	var results map[string]int = make(map[string]int)
	fd, err := os.Open(filename)
	defer fd.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if _, ok := results[scanner.Text()]; ok {
			results[removeNonWordChars(scanner.Text())]++
		} else {
			results[removeNonWordChars(scanner.Text())] = 1
		}
	}
	return results
}

func removeNonWordChars(input string) string {
	var result string
	for _, c := range input {
		if lower := strings.ToLower(string(c)); lower == "a" ||
			lower == "b" ||
			lower == "c" ||
			lower == "d" ||
			lower == "e" ||
			lower == "f" ||
			lower == "g" ||
			lower == "h" ||
			lower == "i" ||
			lower == "j" ||
			lower == "k" ||
			lower == "l" ||
			lower == "m" ||
			lower == "n" ||
			lower == "ñ" ||
			lower == "o" ||
			lower == "p" ||
			lower == "q" ||
			lower == "r" ||
			lower == "s" ||
			lower == "t" ||
			lower == "u" ||
			lower == "v" ||
			lower == "w" ||
			lower == "x" ||
			lower == "y" ||
			lower == "z" {
			result = result + string(c)
		}
	}
	return result
}

func CPUprofiling() {
	fd, err := os.Create(".cpu.prof")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer fd.Close()
	pprof.StartCPUProfile(fd)
	defer pprof.StopCPUProfile()
}

func MEMprofiling() {
	fd, err := os.Create(".mem.prof")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer fd.Close()
	pprof.WriteHeapProfile(fd)
}
