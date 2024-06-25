package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/config"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/helper"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	// mengambil data dari body request
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// mencari user berdasarkan username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username not found"}
			helper.ResponseJSON(w, http.StatusNotFound, response)
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}
	// cek validasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "invalid password"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// generate token
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-auth",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	// Mendefinisikan algoritma yang akan digunakan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    t,
		Expires:  expTime,
		HttpOnly: true,
	})

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	// mengambil data dari body request
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash password menggunakan bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashedPassword)

	// simpan data user ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusCreated, response)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Logout success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
