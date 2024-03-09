package middlewares

import "net/http"

//Check the jwt using verifyJWT
//if the jwt is invalid(no jwt found), it redirect to the login page
//if jwt valid && verified -> returns (valid,user,id)

func FilterUser(w http.ResponseWriter, r *http.Request) (string, int) {
	valid, user, id := VerifyJWT(w, r)

	if !valid {
		return "", 0
	}

	return user, id

}
