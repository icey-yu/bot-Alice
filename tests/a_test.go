package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/Mrs4s/MiraiGo/message"

	"github.com/Mrs4s/MiraiGo/client"
)

func TestA(t *testing.T) {
	//bot := client.NewClient(2781337720, "Aa^2034809175")
	bot := client.NewClientEmpty()
	device := client.GenRandomDevice()
	device.Protocol = client.AndroidWatch
	bot.UseDevice(device)
	//bot.Device().Protocol = client.AndroidPad

	err := os.WriteFile("device.txt", bot.Device().ToJson(), 0777)
	if err != nil {
		t.Error(err)
		return
	}

	bot.AllowSlider = true
	res, err := bot.Login()
	if res.Error.String() != "" {
		println(res.Error.String())
		//bot = client.NewClient(2781337720, "Aa^2034809175")
		//bot.UseDevice(client.GenRandomDevice())
		qrCode, err := bot.FetchQRCode()
		println("获取qrcode")
		if err != nil {
			t.Error(err)
			return
		}

		println(qrCode.ImageData)
		err = os.WriteFile("qrCode.png", qrCode.ImageData, 0777)
		if err != nil {
			t.Error(err)
			return
		}

		time.Sleep(time.Second * 20)
	}
	println(bot.GenToken())
	err = os.WriteFile("token.txt", bot.GenToken(), 0777)
	//token, err := os.ReadFile("token.txt")
	if err != nil {
		t.Error(err)
		return
	}
	println("写完啦")

	//err = bot.TokenLogin(token)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

	select {}
}

func A(bot *client.QQClient, event client.INotifyEvent) {
	msg := message.NewSendingMessage()
	text := message.NewText(fmt.Sprintf("别加我啦%d:%s", event.From(), event.Content()))
	msg.Elements = append(msg.Elements, text)
	bot.SendGroupMessage(539040275, msg)
}

func B(bot *client.QQClient, event *message.GroupMessage) {
	println("?")
	info, err := bot.GetGroupInfo(event.GroupCode)
	if err != nil {
		g.Log().Error(gctx.New(), err.Error())
	}
	members, err := bot.GetGroupMembers(info)
	if err != nil {
		g.Log().Error(gctx.New(), err.Error())
	}
	info.Members = members
	//println(info.SelfPermission())
	println(info.FindMember(bot.Uin).CardName)
	msg := message.NewSendingMessage()
	text := message.NewText(fmt.Sprintf("%d:%s", event.GroupCode, string(event.ToString())))
	println(event.ToString())
	msg.Elements = append(msg.Elements, text)
	bot.SendGroupMessage(539040275, msg)
}

func C(bot *client.QQClient, event *message.GroupMessage) {
	println("CCC")
	println("CCC")
	println("CCC")
	println("CCC")
	println("CCC")
	println("CCC")
	println("CCC")
}

func TestC(t *testing.T) {
	data := map[string]string{
		"content":         "你好，我感觉很头疼",
		"conversation_id": "",
		"parent_id":       "",
	}
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	response, err := post("http://127.0.0.1:8848/api/ask", body)
	if err != nil {
		t.Error(err)
		return
	}
	bd, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}
	println(string(bd))

	//err := addUser()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
}

func addUser() error {
	data := map[string]string{
		"admin_key": "Alice",
	}
	b, _ := json.Marshal(data)
	body := bytes.NewReader(b)
	response, err := post("http://127.0.0.1:8080/admin/users/add", body)
	if err != nil {
		return err
	}
	bd, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	println(string(bd))
	println("添加完畢")
	return nil
}

func post(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(url), body)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Authorization", token)
	req.Header.Set("Authorization", "Alice")
	cl := http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
