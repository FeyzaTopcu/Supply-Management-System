package handlers

import (
	"fmt"
	"net/http"

	helpers "../helpers"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uName, email, pwd, confirmPwd := "", "", "", ""
	uName = r.FormValue("username")
	email = r.FormValue("email")
	pwd = r.FormValue("password")
	confirmPwd = r.FormValue("confirmPassword")

	uNameCheck := helpers.IsEmpty(uName)
	emailCheck := helpers.IsEmpty(email)
	pwdCheck := helpers.IsEmpty(pwd)
	pwdConfirmCheck := helpers.IsEmpty(confirmPwd)

	if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
		fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
		return
	}

	if pwd == confirmPwd {
		// Save to database (username, email and password)
		fmt.Fprintln(w, "Registration successful.")
	} else {
		fmt.Fprintln(w, "Password information must be the same.")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email, pwd := "", ""
	email = r.FormValue("email")  // Data from the form
	pwd = r.FormValue("password") // Data from the form

	// Empty data checking
	emailCheck := helpers.IsEmpty(email)
	pwdCheck := helpers.IsEmpty(pwd)

	if emailCheck || pwdCheck {
		fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
		return
	}

	dbPwd := "1234!*."                 // DB simulation
	dbEmail := "feyza.topcu@gmail.com" // DB simulation

	if email == dbEmail && pwd == dbPwd {
		fmt.Fprintln(w, "Login succesful!")
	} else {
		fmt.Fprintln(w, "Login failed!")
	}

}
