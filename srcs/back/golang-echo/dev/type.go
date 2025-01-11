package dev

type FiveThousandRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	BirthDate string `json:"birthdate"`
	Gender    string `json:"gender"`
	Sexuality string `json:"sexuality"`
	Area      string `json:"area"`
	SelfIntro string `json:"self_intro"`
	ImagePath string `json:"image_path"`
}

type FiveThousandRegisterResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type GenderType string

const (
	GMale   GenderType = "male"
	GFamale GenderType = "famale"
)

type SexualityType string

const (
	SMale       SexualityType = "male"
	SFamale     SexualityType = "famale"
	SMaleFamale SexualityType = "male/famale"
)

type MaleUsername string

const (
	Takashi MaleUsername = "takashi"
	Yutaka  MaleUsername = "yutaka"
	Koji    MaleUsername = "koji"
	Kai     MaleUsername = "kai"
	Shin    MaleUsername = "shin"
)

type FamaleUsername string

const (
	Yui    FamaleUsername = "yui"
	Mayumi FamaleUsername = "mayumi"
	Miyu   FamaleUsername = "miyu"
	Meiko  FamaleUsername = "meiko"
	Keiko  FamaleUsername = "keiko"
)


type MaleImagePath string
const (
	MaleImagePath1 MaleImagePath = "/home/appuser/uploads/images1/male1.png"
	MaleImagePath2 MaleImagePath = "/home/appuser/uploads/images1/male2.png"
	MaleImagePath3 MaleImagePath = "/home/appuser/uploads/images1/male3.png"
	MaleImagePath4 MaleImagePath = "/home/appuser/uploads/images1/male4.png"
	MaleImagePath5 MaleImagePath = "/home/appuser/uploads/images1/male5.png"
)
type FamaleImagePath string
const (
	FamaleImagePath1 FamaleImagePath = "/home/appuser/uploads/images1/famale1.png"
	FamaleImagePath2 FamaleImagePath = "/home/appuser/uploads/images1/famale2.png"
	FamaleImagePath3 FamaleImagePath = "/home/appuser/uploads/images1/famale3.png"
	FamaleImagePath4 FamaleImagePath = "/home/appuser/uploads/images1/famale4.png"
	FamaleImagePath5 FamaleImagePath = "/home/appuser/uploads/images1/famale5.png"
)