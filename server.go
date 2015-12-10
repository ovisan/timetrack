package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
	"github.com/unrolled/secure"
	"html/template"
	"log"
	"net/http"
)

var sessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

//Parse template files
var templates = template.Must(template.ParseFiles("templates/main.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/login.html"))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "SessionName")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	username, found := session.Values["username"]
	if !found || username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	templates.ExecuteTemplate(w, "main", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("inputEmail")
	pass := r.FormValue("inputPassword")
	if email == "visan.ovidiu@gmail.com" && pass == "Test" {
		session, err := sessionStore.Get(r, "SessionName")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// Get the previously flashes, if any.
		// if flashes := session.Flashes(); len(flashes) > 0 {
		// 	// Use the flash values.
		// } else {
		// 	// Set a new flash.
		// 	session.AddFlash("Hello, flash messages world!")
		// }
		session.Values["username"] = email
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/", 302)
		templates.ExecuteTemplate(w, "main", nil)
	} else {
		templates.ExecuteTemplate(w, "login", nil)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "SessionName")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	session.Values["username"] = ""
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func main() {

	sessionStore.Options = &sessions.Options{
		// Path:     "/",
		MaxAge: 600,
		// HttpOnly: true,
	}
	secureMiddleware := secure.New(secure.Options{
		// AllowedHosts:          []string{"localhost", "localhost"},
		SSLRedirect: true,
		SSLHost:     "localhost:8043",
		// SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		// STSSeconds:            315360000,
		// STSIncludeSubdomains:  true,
		// STSPreload:            true,
		// FrameDeny:             true,
		// ContentTypeNosniff:    true,
		// BrowserXssFilter:      true,
		// ContentSecurityPolicy: "default-src 'self'",
		// PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
	})

	router := mux.NewRouter()

	// serving the static files, css
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/login", loginHandler)

	http.Handle("/", router)
	log.Println("Listening...")

	chain := alice.New(secureMiddleware.Handler).Then(router)
	// HTTPS
	// To generate a development cert and key, run the following from your *nix terminal:
	// go run $GOROOT/src/pkg/crypto/tls/generate_cert.go --host="localhost"
	http.ListenAndServeTLS(":8043", "cert.pem", "key.pem", chain)

}
