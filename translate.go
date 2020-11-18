// 文本内容翻译
// 翻译文本复制自yii2框架翻译文本
// todo 增加关键字替换
package transl

import (
	"transl/messages"
)

// 翻译方法
// T("Delete")
func T(str string, language ...string) string {
	k, ok := messages.Messages[str]
	if !ok {
		return str
	}
	lang := "zh-CN"
	if len(language) > 0 {
		lang = language[0]
	}
	if lang == "en" {
		return str
	}
	_, ok = messages.Text[lang]
	if !ok {
		return str
	}

	if messages.Text[lang][k] == "" {
		return str
	}

	return messages.Text[lang][k]
}
