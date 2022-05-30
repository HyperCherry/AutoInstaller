package main

import (
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Download(url string, target string) {
	urls := strings.Split(url, "/")
	filename := target + "\\" + urls[len(urls)-1]
	client := http.DefaultClient
	client.Timeout = 60 * 10 * time.Second
	reps, err := client.Get(url)
	if err != nil {
		log.Panic(err.Error())
	}
	if reps.StatusCode == http.StatusOK {
		file, err := os.Create(filename)
		if err != nil {
			log.Panic(err.Error())
		}
		defer file.Close()
		length := reps.Header.Get("Content-Length")
		size, _ := strconv.ParseInt(length, 10, 64)
		body := reps.Body
		bar := pb.Full.Start64(size)
		bar.SetWidth(120)
		bar.SetRefreshRate(10 * time.Millisecond)
		defer bar.Finish()
		barReader := bar.NewProxyReader(body)
		writer := io.Writer(file)
		io.Copy(writer, barReader)
	}
}
