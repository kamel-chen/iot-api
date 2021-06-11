package main

import (
	"fmt"
	"os"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

func main() {
    report, err := runner.Run(
        "gps.GPSService.CreateGPS",
        "192.168.0.103:8080",
        runner.WithProtoFile("gps.proto", []string{"test/ghz"}),
        runner.WithDataFromFile("test/ghz/data.json"),
        runner.WithInsecure(true),
    )

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    printer := printer.ReportPrinter{
        Out:    os.Stdout,
        Report: report,
    }

    printer.Print("summary")
}
