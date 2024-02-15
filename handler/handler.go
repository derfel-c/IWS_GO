package handler

import (
	//"encoding/json"
	"fmt"
	"io"

	//"io"
	"net/http"
	//validation "github.com/insabelter/IWS_GO/validation"
)

// route to test if the server is running -> health check
func MakePingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "pong")

	}
}

// Schreibe hier den Make FeedbackHandler
func MakeRatingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body, e = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application-json")
		if e == nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(body))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, e.Error())
		}
	}
}

//Mithilfe io.ReadAll() kannst du den Body speichern
//Was soll passieren wenn dies nicht möglich ist?
//Schaut euch dazu die Rückgabewerte der Methode io.ReadAll()

// Antworte mit einer Bestätigung. Diese soll das Feedback zurück geben
// Der body ist von Typ byte[]. string() kann helfen

// Für die Schnellen mithilfe von r.FromValue("key") kannst du auf HTTP Parameter zugreifen.
//Übergebe die Gesamtbewertung als Zahl zwischen 1-10 und gebe diese aus.
