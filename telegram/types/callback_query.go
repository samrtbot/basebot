package types

import (
	"encoding/json"
)

// 查询回调
type CallbackQuery struct {
	ID              string   `json:"id"`                          // 查询唯一ID
	From            *Chat    `json:"from"`                        // 发送者
	Message         *Message `json:"message"`                     // 消息
	ChatInstance    string   `json:"chat_instance,omitempty"`     // 全局标识符
	Data            string   `json:"data,omitempty"`              // 查询数据
	InlineMessageID *string  `json:"inline_message_id,omitempty"` // 内联消息ID
}

// 转换为JSON
func (callbackQuery *CallbackQuery) ToJSON() ([]byte, error) {
	return json.Marshal(callbackQuery)
}

// 从JSON反序列化
func (callbackQuery *CallbackQuery) FromJSON(jsb []byte) error {
	return json.Unmarshal(jsb, callbackQuery)
}
