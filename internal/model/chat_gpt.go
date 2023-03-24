package model

type ChatGPTMessage struct {
	Content        string `json:"content"`
	ConversationId string `json:"conversation_id"` // 会话ID
	ParentId       string `json:"parent_id"`       // 上一句的ID
}

type ChatGPTResp struct {
	Id             string `json:"id"`
	ResponseId     string `json:"response_id"`
	ConversationId string `json:"conversation_id"`
	Content        string `json:"content"`
	Error          string `json:"error"`
}
