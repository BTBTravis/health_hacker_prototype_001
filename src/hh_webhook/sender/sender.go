package sender

import (
    "log"
    "net/http"
    "net/url"
)

func SendMessage(channelId string, message string) {
    params := url.Values{}
    params.Add("text", message)
    params.Add("channel", channelId)
    params.Add("token", "xoxb-752798852689-756983711187-QPM8GGpBrrUV14NzWEIzR7om")
    slackURL := "https://slack.com/api/chat.postMessage?" + params.Encode()

    log.Println(slackURL)
    resp, err := http.Get(slackURL)

    if err != nil {
        log.Println(err)
    }
    log.Println(resp)
}
