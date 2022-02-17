package services

import (
	"errors"
	"os"
	"time"

	"github.com/fahmidyt/go-book-rental-be/src/db"
	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/fahmidyt/go-book-rental-be/src/types"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type TokenDetails struct {
	Token
	UserID      uint
	AccessUUID  string
	RefreshUUID string
	AtExpires   int64
	RtExpires   int64
}

type AuthService struct{}

func (srvc AuthService) Login(form types.LoginForm) (user models.User, token Token, err error) {
	res := db.GetDB().Model(models.User{}).Preload("Role").Preload("UserDetail").Where(models.User{Email: form.Email}).First(&user)

	// handle user not found
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res.Error = errors.New("cannot find your account. please do register instead")
		return user, token, err
	}

	// handle other error just incase
	if res.Error != nil {
		return user, token, err
	}

	// comparing password
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	// throw error
	if err != nil {
		err = errors.New("incorrect password")
		return user, token, err
	}

	if !user.Active {
		err = errors.New("please check your email to verify your account")
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := srvc.GenerateJWT(user.ID)
	if err != nil {
		return user, token, err
	}

	err = srvc.CreateRefreshToken(user.ID, tokenDetails)
	if err != nil {
		return user, token, err
	}

	// asign token
	token = tokenDetails.Token

	return user, token, err
}

func (srvc AuthService) Register(form types.RegisterForm) (user models.User, err error) {
	res := db.GetDB().Where(&models.User{Email: form.Email}).First(&user)

	// handle user not found
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err := errors.New("your email is currently registered")
		return user, err
	}

	// generate bcrypt pass
	bytePassword := []byte(form.Password)
	newPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return user, err
	}

	var role models.Role
	resRole := db.GetDB().Where(&models.Role{Name: "Customer"}).First(&role)

	if resRole.Error != nil {
		return user, resRole.Error
	}

	// PAYLOAD HERE
	userDetail := models.UserDetail{
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		PhoneNumber: form.PhoneNumber,
		DateOfBirth: form.DateOfBirth.Local(),
	}

	user = models.User{
		Email:    form.Email,
		Password: string(newPassword),
		RoleId:   role.ID,
		// TODO: send email verification
		Active:     true,
		UserDetail: userDetail,
	}

	res = db.GetDB().Create(&user)

	if res.Error != nil {
		return user, res.Error
	}

	return user, err
}

func (srvc AuthService) GenerateJWT(userID uint) (*TokenDetails, error) {
	var err error
	td := &TokenDetails{}

	td.UserID = userID

	td.AtExpires = time.Now().Add(time.Hour * 12).Unix()
	td.AccessUUID = uuid.NewString()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewString()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	// if cannot be signed
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	//if cannot be signed
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (srvc AuthService) CreateRefreshToken(userID uint, td *TokenDetails) error {
	payload := models.RefreshToken{UserID: userID, Token: td.RefreshToken}
	res := db.GetDB().Model(models.RefreshToken{}).Create(&payload)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
