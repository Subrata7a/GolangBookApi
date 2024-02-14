package auth

import (
	"GolangBookApi/model"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var secretKey = []byte("SecretKey")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token\n")
	}

	return nil
}

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := fmt.Fprint(w, "Missing authorization header\n")
			if err != nil {
				(&model.Error{}).GetError(w, http.StatusUnauthorized, "StatusUnauthorized", "Missing authorization header")
				return
			}
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := fmt.Fprint(w, "Invalid token\n")
			if err != nil {
				(&model.Error{}).GetError(w, http.StatusUnauthorized, "StatusUnauthorized", "Invalid Token")
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var currentUser model.User

	err := json.NewDecoder(r.Body).Decode(&currentUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Details of the requested user %v\n", currentUser)

	if currentUser.Username == "adminname" && currentUser.Password == "adminpass" {
		tokenString, err := CreateToken(currentUser.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err := fmt.Errorf("no username found\n")
			if err != nil {
				fmt.Println(err)
			}
		}
		w.WriteHeader(http.StatusOK)
		_, errr := fmt.Fprint(w, tokenString)

		if errr != nil {
			fmt.Println(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprint(w, "Invalid credentials\n")

		if err != nil {
			fmt.Println(err)
		}
	}
}
