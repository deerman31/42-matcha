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
	GFemale GenderType = "female"
)

type SexualityType string

const (
	SMale       SexualityType = "male"
	SFemale     SexualityType = "female"
	SMaleFemale SexualityType = "male/female"
)

type MaleUsername string

const (
	Takashi MaleUsername = "takashi"
	Yutaka  MaleUsername = "yutaka"
	Koji    MaleUsername = "koji"
	Kai     MaleUsername = "kai"
	Shin    MaleUsername = "shin"
)

type FemaleUsername string

const (
	Yui    FemaleUsername = "yui"
	Mayumi FemaleUsername = "mayumi"
	Miyu   FemaleUsername = "miyu"
	Meiko  FemaleUsername = "meiko"
	Keiko  FemaleUsername = "keiko"
)


type MaleImagePath string
const (
	MaleImagePath1 MaleImagePath = "/home/appuser/uploads/images1/male1.png"
	MaleImagePath2 MaleImagePath = "/home/appuser/uploads/images1/male2.png"
	MaleImagePath3 MaleImagePath = "/home/appuser/uploads/images1/male3.png"
	MaleImagePath4 MaleImagePath = "/home/appuser/uploads/images1/male4.png"
	MaleImagePath5 MaleImagePath = "/home/appuser/uploads/images1/male5.png"
)
type FemaleImagePath string
const (
	FamaleImagePath1 FemaleImagePath = "/home/appuser/uploads/images1/female1.png"
	FamaleImagePath2 FemaleImagePath = "/home/appuser/uploads/images1/female2.png"
	FamaleImagePath3 FemaleImagePath = "/home/appuser/uploads/images1/female3.png"
	FamaleImagePath4 FemaleImagePath = "/home/appuser/uploads/images1/female4.png"
	FamaleImagePath5 FemaleImagePath = "/home/appuser/uploads/images1/female5.png"
)