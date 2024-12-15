package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

var mots = []string{"golang", "developpeur", "pendu", "template", "serveur"}
var tmpl = template.Must(template.ParseFiles(
	filepath.Join("templates", "index.html"),
	filepath.Join("templates", "start.html"),
	filepath.Join("templates", "win.html"),
	filepath.Join("templates", "lose.html"),
	filepath.Join("templates", "scoreboard.html"),
))

type Jeu struct {
	Mot            string
	Tentatives     string
	Essais         int
	MotAffiche     string
	NomUtilisateur string
	Score          int
}

var jeu Jeu

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", debutHandler)
	http.HandleFunc("/hangman", penduHandler)
	http.HandleFunc("/win", victoireHandler)
	http.HandleFunc("/lose", defaiteHandler)
	http.HandleFunc("/scoreboard", tableauHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func debutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		jeu = NouveauJeu()
		jeu.NomUtilisateur = r.FormValue("username")
		http.Redirect(w, r, "/hangman", http.StatusSeeOther)
	} else {
		tmpl.ExecuteTemplate(w, "start.html", nil)
	}
}

func penduHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		lettre := r.FormValue("letter")
		if !strings.Contains(jeu.Tentatives, lettre) {
			jeu.Tentatives += lettre
			if strings.Contains(jeu.Mot, lettre) {
				jeu.MotAffiche = mettreAJourMotAffiche(jeu.Mot, jeu.Tentatives)
				if jeu.MotAffiche == jeu.Mot {
					http.Redirect(w, r, "/win", http.StatusSeeOther)
					return
				}
			} else {
				jeu.Essais--
				if jeu.Essais == 0 {
					http.Redirect(w, r, "/lose", http.StatusSeeOther)
					return
				}
			}
		}
	}
	tmpl.ExecuteTemplate(w, "index.html", jeu)
}

func victoireHandler(w http.ResponseWriter, r *http.Request) {
	jeu.Score += 10
	tmpl.ExecuteTemplate(w, "win.html", jeu)
}

func defaiteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "lose.html", jeu)
}

func tableauHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "scoreboard.html", jeu)
}

func NouveauJeu() Jeu {
	mot := mots[rand.Intn(len(mots))]
	return Jeu{
		Mot:        mot,
		Tentatives: "",
		Essais:     10,
		MotAffiche: strings.Repeat("_", len(mot)),
		Score:      0,
	}
}

func mettreAJourMotAffiche(mot, tentatives string) string {
	affiche := ""
	for _, char := range mot {
		if strings.ContainsRune(tentatives, char) {
			affiche += string(char)
		} else {
			affiche += "_"
		}
	}
	return affiche
}
