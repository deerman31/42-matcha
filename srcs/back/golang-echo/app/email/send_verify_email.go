package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

const (
	Mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	Subject =  "メールアドレスの認証"
)

type Config struct {
	host      string
	port      string
	username  string
	password  string
	fromEmail string
	endpoint  string
}

// シングルトンとしてのconfig
var config *Config

func init() {
	host := os.Getenv("MAILTRAP_HOST")
	port := os.Getenv("MAILTRAP_PORT")
	username := os.Getenv("MAILTRAP_USERNAME")
	password := os.Getenv("MAILTRAP_PASSWORD")
	fromEmail := os.Getenv("MAILTRAP_FROM_EMAIL")
	endpoint := os.Getenv("MAIL_VERIFY_ENDPOINT_URL")
	config = &Config{
		host:      host,
		port:      port,
		username:  username,
		password:  password,
		fromEmail: fromEmail,
		endpoint:  endpoint,
	}
}

func SendVerifyEmail(token string, toEmail string) error {
	// HTMLメールの本文を作成

	verifyURL := generateVerifyURL(config.endpoint, token)

	htmlBody := generateHtmlBody(verifyURL)

	if err := sendEmail(config.username, config.password, config.host, config.port, config.fromEmail, toEmail, Subject, htmlBody); err != nil {
		return err
	}

	return nil
}

func sendEmail(username, password, host, port, fromEmail, toEmail, subject, htmlBody string) error {
	toEmails := []string{toEmail}

	// メールヘッダーの作成
	headers := make(map[string]string)
	headers["From"] = fromEmail
	headers["To"] = strings.Join(toEmails, ",")
	headers["Subject"] = subject

	// メッセージの組み立て
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += Mime + htmlBody

	// SMTP認証
	auth := smtp.PlainAuth(
		"",
		username,
		password,
		host,
	)

	// メール送信
	err := smtp.SendMail(
		host+":"+port,
		auth,
		fromEmail,
		toEmails,
		[]byte(message),
	)
	if err != nil {
		return err
	}
	return nil
}

func generateVerifyURL(endpoint, token string) string {
	return fmt.Sprintf("%s%s", endpoint, token)
}

func generateHtmlBody(url string) string {
	return fmt.Sprintf(`
        <html>
        <body>
            <h2>メール認証</h2>
            <p>以下のリンクをクリックしてメールアドレスを認証してください：</p>
            <p><a href="%s">メールアドレスを認証する</a></p>
            <p>このリンクの有効期限は1時間です。</p>
            <hr>
            <p>リンクがクリックできない場合は、以下のURLをブラウザにコピー&ペーストしてください：</p>
            <p>%s</p>
        </body>
        </html>
    `, url, url)
}
