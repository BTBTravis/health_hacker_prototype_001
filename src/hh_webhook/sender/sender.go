package sender

import (
    "bytes"
    "net/http"
    "log"
    "encoding/json"
)



type SlackUser struct{
    ChannelId string
    UserId string
}

var Travis = SlackUser{
    ChannelId: "DN7L5L8AW",
    UserId: "UMYDK4Q66",
}

var Alex = SlackUser{
    ChannelId: "DNLGGKBGD",
    UserId: "UND2UQZB9",
}


func SendMessage(message string) {

    type Msg struct{
	Text string `json:"text"`
	Channel string `json:"channel"`
    }

    var msg = Msg{
        Text: message,
        Channel: Travis.ChannelId,
    }

    jsonValue, _ := json.Marshal(msg)
    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(jsonValue)


    // TODO move secret to env var
    var bearer = "Bearer " + "xoxb-752798852689-756983711187-MfD6lVvnEY5hVaBRrpjrrWSs"
    client := &http.Client{}
    req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", buf)
    req.Header.Set("Authorization", bearer)
    req.Header.Set("User-Agent", "application/json")
    log.Println(req)
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
    }
    log.Println(resp)

    //res, err := http.Post("https://slack.com/api/chat.postMessage", "application/json", bytes.NewBuffer(bytesRepresentation))
    //res, err := http.Post("https://hooks.slack.com/services/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", "application/json", bytes.NewBuffer(bytesRepresentation))
}
