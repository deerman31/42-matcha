package auth

// User はデータベースのユーザー情報を表す構造体
type User struct {
	ID            int
	Username      string
	PasswordHash  string
	isOnline      bool
	isRegistered  bool
	isPreparation bool
}
