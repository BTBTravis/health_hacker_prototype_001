package sender

import (
    "os"
    "log"
    "net/http"
    "net/url"
)

func SendMessage(channelId string, message string) {
    log.Printf("send message: ", channelId, message)
    params := url.Values{}
    params.Add("text", message)
    params.Add("channel", channelId)
    key, key_exists := os.LookupEnv("SLACK_API_KEY")

    if key_exists {
        params.Add("token", key)
    } else {
        log.Println("Missing SLACK_API_KEY")
    }

    slackURL := "https://slack.com/api/chat.postMessage?" + params.Encode()

    log.Println(slackURL)
    resp, err := http.Get(slackURL)

    if err != nil {
        log.Println(err)
    }
    log.Println(resp)
}
