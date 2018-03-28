package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"fec"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read data")
		report, err := fec.ReadReport(data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("parsed report")
			out, err := json.MarshalIndent(report, "", "  ")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("jsonified report")
				os.Stdout.Write(out)
				fmt.Println("")
				fecfmt := []byte{}
				for _, rec := range report.Records {
					fecfmt = append(fecfmt, rec.Record()...)
				}
				os.Stdout.Write(fecfmt)
				fmt.Println("")
				if len(data) != len(fecfmt) {
					fmt.Printf("output length (%d bytes) != input length (%d bytes)\n", len(fecfmt), len(data))
				}
				for i, b := range fecfmt {
					if i > len(data) {
						fmt.Println("trailing data in output", fecfmt[i:])
						break
					} else if b != data[i] {
						fmt.Printf("output differs from input at byte %d (%d != %d)\n", i, b, data[i])
						break
					}
				}
			}
		}
	}
}

