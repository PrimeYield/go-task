package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go-task/internal/response/task"
	"net/http"
)

func main() {
	count := flag.Int("n", 1, "request count")
	url := flag.String("url", "", "request URL")

	flag.Parse()

	for i := 0; i < *count; i++ {
		req := task.CreateTaskRequest{
			Title: fmt.Sprintf("Task %d", i+1),
		}

		reqBody, err := json.Marshal(req)
		if err != nil {
			fmt.Println("marshal error:", err)
			continue
		}

		resp, err := http.Post(
			*url,
			"application/json",
			bytes.NewReader(reqBody),
		)
		if err != nil {
			fmt.Println("request error:", err)
			continue
		}

		fmt.Printf("[%d] %s\n", i+1, resp.Status)
		resp.Body.Close()
	}
}
