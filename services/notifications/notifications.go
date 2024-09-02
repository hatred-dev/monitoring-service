package notifications

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "monitoring-service/configuration"
    "monitoring-service/logger"
    "monitoring-service/models/api"
    "net/http"
    "strings"
    "text/template"
    "time"
)

type Message struct {
    Status  string
    Message string
}

var EscapeChars = []byte{'_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!'}

// EscapeMarkdown https://core.telegram.org/bots/api/#markdownv2-style
func EscapeMarkdown(text string) string {
    var builder strings.Builder
    for _, char := range text {
        if bytes.ContainsRune(EscapeChars, char) {
            builder.WriteRune('\\')
        }
        builder.WriteRune(char)
    }
    return builder.String()
}

func CreateMessage(msg Message) string {
    msg.Message = EscapeMarkdown(msg.Message)
    tmpl := `*{{.Status}}*  
{{.Message}}`
    t, _ := template.New("message").Parse(tmpl)
    body := &bytes.Buffer{}
    _ = t.Execute(body, msg)
    return body.String()
}

func SendTelegramNotification(text string) error {
    tgConf := configuration.GetTelegramConfig()
    var resp *http.Response
    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tgConf.BotToken)
    body, err := json.Marshal(api.TelegramMessage{
        ParseMode: "MarkdownV2",
        Text:      text,
        ChatId:    tgConf.ChatId,
    })
    if err != nil {
        return err
    }
    for {
        resp, err = http.Post(url, "application/json", bytes.NewBuffer(body))
        if err != nil {
            return err
        }
        if resp.StatusCode == http.StatusTooManyRequests {
            time.Sleep(time.Second * 5)
        } else {
            break
        }
    }
    var response []byte
    if resp != nil {
        response, err = io.ReadAll(resp.Body)
    }
    if err != nil {
        return err
    }
    logger.Log.Info(string(response))
    return nil
}

func SendNotifications(projectName, serviceName, message string, active bool) {
    err := SendTelegramNotification(message)
    if err != nil {
        logger.Log.Warn(err)
    }
}
