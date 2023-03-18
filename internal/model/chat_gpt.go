package model

type ChatGPTMessage struct {
	Content        string
	ConversationId string // 会话ID
	ParentId       string // 上一句的ID
}

type ChatGPTResp struct {
	Id             string `json:"id"`
	ResponseId     string `json:"response_id"`
	ConversationId string `json:"conversation_id"`
	Content        string `json:"content"`
	Error          string `json:"error"`
}
