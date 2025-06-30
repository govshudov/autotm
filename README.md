# Auth AutoTM

# ByMail Service

## POST auth/mail/send
Client e mail ugratmaly

Request
```
{
    email string `json:"email"`
    device {
        UniqKey       string `json:"uniqKey" validate:"required|min_len:1"`
        Model         string `json:"model" validate:"required|min_len:1"`
        Name          string `json:"name" validate:"required|min_len:1"`
        System        string `json:"system" validate:"required|min_len:1"`
        SystemVersion string `json:"systemVersion" validate:"required|min_len:1"`        
    }
}
```

Response
```
{
    token string
}
```
Bolup biljek status code lar:

- 400 eger validation dan gecmese
- 500 service de yada database bilen baglansykly errorlar
- 429 eger client blocklanan kisilerin icinde bar bolsa 
- 403 eger client send mail limitini gecse
- 200 success bolsa 

```sql
CREATE TABLE blocked_mails(
    mail text not null default ''
);
```
- Logic
  + Hacanda client send sms yuzlenende mail_device_unique seklinde redisde ululyk arttyrmaly eger olaryn sany 5 den gecse client in sol devici 1 sagat block etmeli (yagny mail ugratmaly dal)
  + Validation edende mail i blocklanan mail lerin icinden barlamaly
  + Response de jwt token bermeli (auth ucin bir sany jwt key bolmaly) tokenin icinde device maglumatlar bolmaly
    + UniqKey
    + Model
    + Name
    + System
    + SystemVersion
---
## POST auth/sms/send
Client e sms ugratmaly

Request
```
{
    phone_number string `json:"email"`
    device {
        UniqKey       string `json:"uniqKey" validate:"required|min_len:1"`
        Model         string `json:"model" validate:"required|min_len:1"`
        Name          string `json:"name" validate:"required|min_len:1"`
        System        string `json:"system" validate:"required|min_len:1"`
        SystemVersion string `json:"systemVersion" validate:"required|min_len:1"`        
    }
}
```

Response
```
{
    token string
}
```
Bolup biljek status code lar:

- 400 eger validation dan gecmese
- 500 service de yada database bilen baglansykly errorlar
- 429 eger client blocklanan kisilerin icinde bar bolsa 
- 403 eger client send sms limitini gecse
- 200 success bolsa 

```sql
CREATE TABLE blocked_numbers(
    phone_number text not null default ''
);
```
- Logic
  + Hacanda client send sms yuzlenende mail_device_unique seklinde redisde ululyk arttyrmaly eger olaryn sany 5 den gecse client in sol devici 1 sagat block etmeli (yagny mail ugratmaly dal)
  + Validation edende mail i blocklanan mail lerin icinden barlamaly
  + Response de jwt token bermeli (auth ucin bir sany jwt key bolmaly) tokenin icinde device maglumatlar bolmaly
    + UniqKey
    + Model
    + Name
    + System
    + SystemVersion
---

## POST auth/mail/verify
Clientden code gelyar shony check etmeli

Request
```
{
    code string `json:"code"`
}
```

Response
```
{
    token string
}
```
Bolup biljek status code lar:

- 400 eger validation dan gecmese
- 401 token is expired
- 406 mail code is incorrect
- 423 user is blocked
- 429 try code exceeded limit
- 500 service de yada database bilen baglansykly errorlar
- 200 success (user doesn't exist go to registration)
- 202 (user already exist go to log in)


```sql
CREATE TABLE user_devices(
    user_id bigint not null,
    uniq_id text not null, 
    firebase_token text,
    name character varying(100),
    model character varying(100),
    system character varying(100),
    system_version character varying(100),
    app_version character varying(100)
);
-- users table bilen reference bolmaly
```
- Logic
  + exceeded limit check etmeli
  + gelen mail i validation etmeli
  + gelen tokenin omruni barlamaly
  + gelen kody tokendaki kod bilen denesdirmeli
  + users serviceden user blockmy dalmi check etmeli
  + users service gitmeli shular yaly user barmy barlamaly 
    + bar bolsa user_devices table e token den gelen device data leri insert etmeli on bar bolsa update etmeli (unique id boyuncha)
  + token doretmeli user on bar bolsa standart time bermeli yok bolsa registration ucin 1 sagat bermeli

## POST auth/mail/registration
User i registrasiya etmeli 

Request
```
{
    Email string `json:"email"`
    PhoneNumber `json:"phoneNumber"`
    NickName `json:"nickname"`
    FullName `json:"fullName"`
    ...
}
```

Response
```
{
    userID string
    token string
}
```
Bolup biljek status code lar:

- 400 eger validation dan gecmese
- 422 nickname is invalid
- 409 nickname is already exist
- 500 service de yada database bilen baglansykly errorlar
- 200 success (user doesn't exist go to registration)

```sql
CREATE TABLE users(
    id bigint not null,
    email character varying(100) not null,
    full_name character varying(100),
    nickname character varying(100),
    phone_number character varying(100),
);
```
- Logic
  + token middleware 
  + nickname barlamaly validation ve user service den barlamaly shular yaly nickname bar bolsa gerekli error return etmeli
  + users table insert gitmeli sonra user devices e gitmeli devices i token den almaly
  + token generate etmeli
---


# By OAuth2 Service

## Example: 
```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"net/url"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     "1014566394157-cplumj9j4v4dglhf7bn8eucb2grde3m6.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

var jwtKey = []byte("your_secret_key")

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    string `json:"id"`
	jwt.RegisteredClaims
}
func oAuthHandler(w http.ResponseWriter, r *http.Request) {
	
	url := googleOauthConfig.AuthCodeURL("random-state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	authuser := r.URL.Query().Get("authuser")
	scope := r.URL.Query().Get("scope")
	prompt := r.URL.Query().Get("prompt")

	fmt.Println("code:", code)
	fmt.Println("state:", state)
	fmt.Println("authuser:", authuser)
	fmt.Println("scope:", scope)
	fmt.Println("prompt:", prompt)

	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User info: %+v\n", userInfo)

	redirectURL := fmt.Sprintf("http://localhost:3000?success=%v&name=%s&email=%s&authuser=%s&scope=%s&prompt=%s&state=%s",
		userInfo["verified_email"],
		url.QueryEscape(fmt.Sprintf("%v", userInfo["name"])),
		url.QueryEscape(fmt.Sprintf("%v", userInfo["email"])),
		url.QueryEscape(authuser),
		url.QueryEscape(scope),
		url.QueryEscape(prompt),
		url.QueryEscape(state),
	)

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/auth/google/login", oAuthHandler).Methods("GET")
	r.HandleFunc("/auth/google/callback", googleCallbackHandler).Methods("GET")

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://10.192.5.6:3000", "http://localhost:3000", "http://localhost"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", cors(r)))
}

```
{ }
```

Response
```
{ }
```

- Logic
  + Client Get Zapros ibermeli bu get zaprosun icinde google redirect etmeli
---

## GET auth/google/callback
Google redirect api doretmeli 

Request URL Query
```
{
    code string 
    state string 
    authuser string 
    scope string 
    prompt string 
}
```

```sql
CREATE TABLE oauth2(
    full_name character varying(100),
    email text not null, 
    code text not null
);
```
- Logic
  + code boyuncha google den token almaly we token boyuncha user in maglumatlaryny almaly
  + gelen maglumatlary oauth2 table gosmaly
  + static bir frontend api redirect etmeli (url de code bolmaly) 

## POST auth/google/verify
Account Verify 

- 400 eger validation dan gecmese (code bosh bolsa yada yalnysh bolsa)
- 500 service de yada database bilen baglansykly errorlar
- 200 success 

Request 
```
{
    code string 
    device {
        UniqKey       string `json:"uniqKey" validate:"required|min_len:1"`
        Model         string `json:"model" validate:"required|min_len:1"`
        Name          string `json:"name" validate:"required|min_len:1"`
        System        string `json:"system" validate:"required|min_len:1"`
        SystemVersion string `json:"systemVersion" validate:"required|min_len:1"`        
    }
}
```
Response 
```
{
    token string 
}
```


```sql
insert into users(
  id bigint not null,
  email character varying(100) not null,
  full_name character varying(100),
  nickname character varying(100),
  phone_number character varying(100),
) values (?, ?, ?, ?, ?);
```


- Logic
  + request url de gelen cody requestde ugratmaly 
  + code boyuncha userin maglumatlaryny oauth2 den almaly yok bolsa err return
  + user in device insert etmeli bar bolsa update
  + user yok bolsa insert etmeli  
  + token generate etmeli
---

