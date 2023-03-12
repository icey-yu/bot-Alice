package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chatgp/chatgpt-go"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Mrs4s/MiraiGo/message"

	"github.com/Mrs4s/MiraiGo/client"
)

func TestA(t *testing.T) {
	bot := client.NewClient(2781337720, "Aa^2034809175")
	bot.UseDevice(client.GenRandomDevice())
	bot.AllowSlider = true
	//res, err := bot.Login()
	//println(bot.GenToken())
	//os.WriteFile("token.txt", bot.GenToken(), 0777)
	token, err := os.ReadFile("token.txt")
	if err != nil {
		t.Error(err)
		return
	}
	println("开些啦")
	err = os.WriteFile("de.txt", bot.Device().ToJson(), 0777)
	if err != nil {
		t.Error(err)
		return
	}
	println("写完啦")

	err = bot.TokenLogin(token)
	if err != nil {
		t.Error(err)
		return
	}

	//println(res.Success)
	msg := message.NewSendingMessage()
	text := message.NewText("QAQ " + time.Now().String())
	msg.Elements = append(msg.Elements, text)
	//_ = bot.SendPrivateMessage(916874663, msg)
	//bot.SendGroupMessage(539040275, msg)
	bot.FriendNotifyEvent.Subscribe(A)
	bot.GroupNotifyEvent.Subscribe(A)
	//bot.GroupMessageEvent.Subscribe(B)
	println("停啦")

	select {}
}

func A(bot *client.QQClient, event client.INotifyEvent) {
	msg := message.NewSendingMessage()
	text := message.NewText(fmt.Sprintf("别加我啦%d:%s", event.From(), event.Content()))
	msg.Elements = append(msg.Elements, text)
	bot.SendGroupMessage(539040275, msg)
}

func B(bot *client.QQClient, event *message.GroupMessage) {
	msg := message.NewSendingMessage()
	text := message.NewText(fmt.Sprintf("%d:%s", event.GroupCode, string(event.ToString())))
	msg.Elements = append(msg.Elements, text)
	bot.SendGroupMessage(539040275, msg)
}

func TestB(t *testing.T) {

	cli := chatgpt.NewClient(
		chatgpt.WithTimeout(60*time.Second),
		//chatgpt.WithCookies(cookies),
		chatgpt.WithAccessToken(token),
		chatgpt.WithUserAgent(ua),
		chatgpt.WithModel("text-davinci-002-render-sha"),
	)

	// chat in independent conversation
	msg := "求求你说句话⑧"
	text, err := cli.GetChatText(msg)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", msg, text.Content)
}

func TestC(t *testing.T) {
	data := map[string]string{
		"content": "Hello world",
		//"conversation_id": uuid.New().String(),
		//"parent_id":       uuid.New().String(),
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
