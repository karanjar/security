package main

import (
	"fmt"
	"net/http"
	"security/hanling_sensitve_data"
)

//entity
//
//user - interact with client
//client - 1. call auth server to get the aceess code . 2. call the auth server to get the access token
//auth server - return back access code to a callback url ,return back access tokento anothe r callback url
//resource server//

const (
	clientId      = "YOUR_CLIENT_ID"
	clientSecret  = "YOUR_CLIENT_SECRET"
	callbackcode  = "http://localhost:8080/code"
	callbacktoken = "http://localhost:8080/token"
)

func main() {
	str := "adcd"
	key := hanling_sensitve_data.Key

	encyrpt := hanling_sensitve_data.Encrypt(str, key)

	fmt.Println(encyrpt)

	decrytp := hanling_sensitve_data.Decrypt(encyrpt, key)

	fmt.Println(decrytp)

	/*pwd := "abcd"
	hash := hanling_sensitve_data.Hashing(pwd)

	fmt.Println(hash)*/
	//http.HandleFunc("/", handlehome)
	//
	//http.HandleFunc("/login", handleAuthServer1stCall)
	//http.HandleFunc("/code", handleCallbackcode)
	//http.HandleFunc("/token", handleCallbackaccess)
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
func handlehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<a href='http://localhost:8080/login'>Login with auth2</a>"))

}

//auth server
//this is going to generate access code

func handleAuthServer1stCall(w http.ResponseWriter, r *http.Request) {
	codeUrl := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=read&code=qwerty1234", callbackcode, clientId)

	http.Redirect(w, r, codeUrl, http.StatusFound)
}
func handleCallbackcode(w http.ResponseWriter, r *http.Request) {
	respType := r.URL.Query().Get("response_type")
	if respType != "code" {
		http.Error(w, "Unauthorised", http.StatusUnauthorized)
		return
	}
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Unauthorised", http.StatusUnauthorized)
		return
	}

	//client is going to request to aout server with
	//client id
	//client secret
	//access code
	// user info

	accessUrl := fmt.Sprintf("%s?client_id=%s&response_type=token&scope=read&token=xxx.yyy.zzz", callbacktoken, clientId)
	http.Redirect(w, r, accessUrl, http.StatusFound)

}

func handleCallbackaccess(w http.ResponseWriter, r *http.Request) {
	responseType := r.URL.Query().Get("response_type")
	if responseType != "token" {
		http.Error(w, "Unauthorised", http.StatusUnauthorized)
		return
	}
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Unauthorised", http.StatusUnauthorized)
		return
	}

	fmt.Println(token)

}
