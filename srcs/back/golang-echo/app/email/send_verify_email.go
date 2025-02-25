package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

const (
	Mime                      = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	MailVerifyEndPoint        = "http://localhost:3000/api/auth/verify-email/"
	MailResetPasswordEndPoint = "http://localhost:3000/api/auth/reset-password/"
	VerifySubject             = "メールアドレスの認証"
	ResetPasswordSubject      = "PasswordReset"
)

type Config struct {
	host      string
	port      string
	username  string
	password  string
	fromEmail string
}

// シングルトンとしてのconfig
var config *Config

func init() {
	host := os.Getenv("MAILTRAP_HOST")
	port := os.Getenv("MAILTRAP_PORT")
	username := os.Getenv("MAILTRAP_USERNAME")
	password := os.Getenv("MAILTRAP_PASSWORD")
	fromEmail := os.Getenv("MAILTRAP_FROM_EMAIL")
	config = &Config{
		host:      host,
		port:      port,
		username:  username,
		password:  password,
		fromEmail: fromEmail,
	}
}

func SendResetPasswordEmail(token string, toEmail string) error {
	verifyURL := generateVerifyURL(MailResetPasswordEndPoint, token)
	htmlBody := generateResetPasswordBody(verifyURL)
	if err := sendEmail(config.username, config.password, config.host, config.port, config.fromEmail, toEmail, ResetPasswordSubject, htmlBody); err != nil {
		return err
	}
	return nil
}

func SendVerifyEmail(token string, toEmail string) error {
	verifyURL := generateVerifyURL(MailVerifyEndPoint, token)
	htmlBody := generateVerifyBody(verifyURL)
	if err := sendEmail(config.username, config.password, config.host, config.port, config.fromEmail, toEmail, VerifySubject, htmlBody); err != nil {
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

func generateVerifyBody(url string) string {
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




func generateResetPasswordBody(url string) string {
    return fmt.Sprintf(`
        <html>
        <head>
            <style>
                .success-message {
                    color: green;
                    padding: 20px;
                    border-radius: 5px;
                    background-color: #dff0d8;
                    margin-top: 20px;
                    display: none;
                }
                .error-message {
                    color: red;
                    padding: 20px;
                    border-radius: 5px;
                    background-color: #f2dede;
                    margin-top: 20px;
                    display: none;
                }
            </style>
        </head>
        <body>
            <h2>パスワードリセット</h2>
            <p>パスワードリセットのリクエストを受け付けました。以下のフォームに新しいパスワードを入力し、送信ボタンをクリックしてください：</p>
            
            <!-- 成功メッセージ -->
            <div id="successMessage" class="success-message">
                パスワードが正常に更新されました。
            </div>
            
            <!-- エラーメッセージ -->
            <div id="errorMessage" class="error-message">
                エラーが発生しました。もう一度お試しください。
            </div>

            <form id="passwordForm" onsubmit="return submitForm(event)">
                <div>
                    <label for="password">新しいパスワード：</label>
                    <input type="password" id="password" required>
                </div>
                <div>
                    <label for="confirm_password">新しいパスワード（確認）：</label>
                    <input type="password" id="confirm_password" required>
                </div>
                <button type="submit">パスワードを変更する</button>
            </form>

            <script>
            function submitForm(event) {
                event.preventDefault();
                
                // メッセージを非表示にする
                document.getElementById('successMessage').style.display = 'none';
                document.getElementById('errorMessage').style.display = 'none';

                const password = document.getElementById('password').value;
                const confirmPassword = document.getElementById('confirm_password').value;
                
                if (password !== confirmPassword) {
                    document.getElementById('errorMessage').textContent = 'パスワードが一致しません。';
                    document.getElementById('errorMessage').style.display = 'block';
                    return false;
                }

                fetch('%s', {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify({password, confirm_password})
                })
                .then(response => {
                    if (!response.ok) {
                        return response.json().then(err => Promise.reject(err));
                    }
                    return response.json();
                })
                .then(data => {
                    document.getElementById('successMessage').style.display = 'block';
                    document.getElementById('passwordForm').reset();
                })
                .catch(error => {
                    console.error('Error:', error);
                    document.getElementById('errorMessage').textContent = 
                        'エラーが発生しました：' + (error.message || '不明なエラーが発生しました');
                    document.getElementById('errorMessage').style.display = 'block';
                });

                return false;
            }
            </script>
        </body>
    </html>
    `, url)
}