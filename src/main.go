package main

import (
    "os/exec"
    "bytes"
    "fmt"
    "net/http"
    "os"
    "time"
    "math/rand"
    "strconv"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    paragraph := rand.Intn(4)
    emoji := rand.Intn(5)

    out, err := exec.Command(
        "docker",
        "run",
        "--rm",
        "-i",
        "greymd/ojichat:latest",
        "-p",
        strconv.Itoa(paragraph),
        "-e",
        strconv.Itoa(emoji)).Output()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    channel := os.Getenv("SLACK_CHANNEL")
    name := os.Getenv("SLACK_BOT_NAME")

    jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + string(out) + `"}`

    req, err := http.NewRequest(
        "POST",
        os.Getenv("SLACK_WEBHOOK_URL"),
        bytes.NewBuffer([]byte(jsonStr)),
    )

    if err != nil {
        fmt.Print(err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        fmt.Print(err)
    }

    fmt.Print(resp)

    defer resp.Body.Close()
}
