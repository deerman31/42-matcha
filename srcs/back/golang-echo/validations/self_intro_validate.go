package validations

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

func validateSelfIntro(fl validator.FieldLevel) bool {
	text := fl.Field().String()

	// 1. 基本的な長さチェック
	if len(text) < 20 || len(text) > 1000 {
		return false
	}
	// 2. アルファベット、数字、基本的な記号のみを許可
	// 全角文字は禁止
	matched, _ := regexp.MatchString("^[a-zA-Z0-9\\s.,!?'\"()\\[\\]\\-_]+$", text)
	if !matched {
		return false
	}
	// 3. 禁止パターンのチェック
	forbiddenPatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)\b(http|https)://`),          // URL
		regexp.MustCompile(`\d{10,}`),                        // 連続する数字（10桁以上）
		regexp.MustCompile(`[\w\-.]+@[\w\-]+\.[a-zA-Z]{2,}`), // メールアドレス
		regexp.MustCompile(`[\s]{3,}`),                       // 3つ以上の連続した空白
		regexp.MustCompile(`[!?]{3,}`),                       // 過度な記号の連続
	}
	for _, pattern := range forbiddenPatterns {
		if pattern.MatchString(text) {
			return false
		}
	}
	// 4. 行数制限チェック
	if strings.Count(text, "\n") > 10 {
		return false
	}
	return true
}
