package gormt

import (
	"time"
)

// Prompt [...]
type Prompt struct {
}

// TableName get sql table name.获取数据库表名
func (m *Prompt) TableName() string {
	return "prompt"
}

type IPrompt interface {
}

// ToMap  struct to map 结构体转成map
func (m *Prompt) ToMap() map[string]any {
	return map[string]any{}
}

// ToMapWithoutModel  struct to map 结构体转成map, 不带gorm.Model
func (m *Prompt) ToMapWithoutModel() map[string]any {
	return map[string]any{}
}

// PromptColumns get sql column name.获取数据库列名
var PromptColumns = struct {
}{}
