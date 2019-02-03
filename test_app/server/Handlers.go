package main
 
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/gorilla/securecookie"
)
 
var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))
 
// Handlers
 
// for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Hello")
    var body, _ = LoadFile("templates/login.html")
    fmt.Println(body)
    fmt.Fprintf(response, body)
}
 
// for POST
func LoginHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("name")
    pass := request.FormValue("password")
    redirectTarget := "/"
    if !IsEmpty(name) && !IsEmpty(pass) {
        // Database check for user data!
        _userIsValid := UserIsValid(name, pass)
 
        if _userIsValid {
            SetCookie(name, response)
            redirectTarget = "/index"
        } else {
            redirectTarget = "/register"
        }
    }
    http.Redirect(response, request, redirectTarget, 302)
}
 
// for GET
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Println("Hello Sir")
    var body, _ = LoadFile("templates/register.html")
    fmt.Fprintf(response, body)
}
 
// for POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
 
    uName := r.FormValue("username")
    email := r.FormValue("email")
    pwd := r.FormValue("password")
    confirmPwd := r.FormValue("confirmPassword")
 
    _uName, _email, _pwd, _confirmPwd := false, false, false, false
    _uName = !IsEmpty(uName)
    _email = !IsEmpty(email)
    _pwd = !IsEmpty(pwd)
    _confirmPwd = !IsEmpty(confirmPwd)
 
    if _uName && _email && _pwd && _confirmPwd {
        fmt.Fprintln(w, "Username for Register : ", uName)
        fmt.Fprintln(w, "Email for Register : ", email)
        fmt.Fprintln(w, "Password for Register : ", pwd)
        fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
    } else {
        fmt.Fprintln(w, "This fields can not be blank!")
    }
}
 
// for GET
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
    userName := GetUserName(request)
    if !IsEmpty(userName) {
        var indexBody, _ = LoadFile("templates/index.html")
        fmt.Fprintf(response, indexBody, userName)
    } else {
        http.Redirect(response, request, "/", 302)
    }
}
 
// for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
    ClearCookie(response)
    http.Redirect(response, request, "/", 302)
}
 
// Cookie
 
func SetCookie(userName string, response http.ResponseWriter) {
    value := map[string]string{
        "name": userName,
    }
    if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
        cookie := &http.Cookie{
            Name:  "cookie",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(response, cookie)
    }
}
 
func ClearCookie(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:   "cookie",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}
 
func GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("cookie"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}

func IsEmpty(data string) bool {
    if len(data) <= 0 {
        return true
    } else {
        return false
    }
}

func LoadFile(fileName string) (string, error) {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

func UserIsValid(uName, pwd string) bool {
    // DB simulation
    _uName, _pwd, _isValid := "cihanozhan", "1234!*.", false
 
    if uName == _uName && pwd == _pwd {
        _isValid = true
    } else {
        _isValid = false
    }
 
    return _isValid
}