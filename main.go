package main

import (
	"bufio"
	"fusionlsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/philbutterworth/Code/fusionlsp/log.txt")
	logger.Println("Hey! fusionslp started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Log file not found")
	}
	return log.New(logfile, "[fusionlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
