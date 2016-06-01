package utiles

import (
    "bytes"
    "crypto/md5"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)

const (
    umengServer = "http://msg.umeng.com/api/send"
    umengAppKey = "56a97651e0f55aac87002b78"
    umnegMKey   = "9pknjpw2b69tb397mvwaoan3kbiafcfg"
)

type Umeng struct {
    deviceToken string
}

func InitUmeng(token string) *Umeng {
    return &Umeng{deviceToken: token}
}


func (umeng *Umeng) ClientNotify(msg string,title string) ([]byte, error) {
    // aps := map[string]string{"alert": msg}
    body := map[string]string{
        "ticker": "bye_school",
        "title": title,
        "text": msg,
        "after_open": "go_app",
    }
    payload := map[string]interface{}{
        "display_type": "notification",
        "body": body,
    }

    data := map[string]interface{}{
        "type":            "unicast",
        "production_mode": "true",
        "timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
        "device_tokens":   umeng.deviceToken,
        "appkey":          umengAppKey,
        "payload":         payload,
    }

    return umeng.postData(data)
}

func (umeng *Umeng) postData(data map[string]interface{}) ([]byte, error) {
    dataBytes, err := json.Marshal(data)

    if err != nil {
        return nil, err
    }

    toSignStr := fmt.Sprintf("POST%s%s%s", umengServer, string(dataBytes), umnegMKey)

    md5h := md5.New()
    io.WriteString(md5h, toSignStr)
    sign := fmt.Sprintf("%x", md5h.Sum(nil))

    postURL := fmt.Sprintf("%s?sign=%s", umengServer, sign)
    resp, err := http.Post(postURL, "application/json;charset=utf-8", bytes.NewBuffer(dataBytes))
    defer resp.Body.Close()

    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    var dat map[string]interface{}

    if err := json.Unmarshal(body, &dat); err != nil {
        return nil, err
    }

    if dat["ret"] == "FAIL" {
        // Logger.Errorf("Umeng Logger resp, %v", string(body))
        return nil, errors.New("Umeng push error")
    }

    return body, nil
}
