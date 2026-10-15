package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)
func ponctuation(c rune) bool {
	return strings.ContainsRune(".,!?:;", c)
}
// decouper et separer le texte en morceaux
func decouper(texte string) []string {
	caracteres := []rune(texte)
	morceaux := []string{}
	for i := 0; i < len(caracteres); {
		c := caracteres[i]
		if unicode.IsSpace(c) {
			i++
			continue
		}
		if c == '(' {
			fin := i + 1
			for fin < len(caracteres) && caracteres[fin] != ')' {
				fin++
			}
			if fin < len(caracteres) {
				morceaux = append(morceaux, string(caracteres[i:fin+1]))
				i = fin + 1
				continue
			}
		}
		// Une apostrophe à l'intérieur d'un mot reste dans le mot : don't.
		if ponctuation(c) || c == '\'' {
			morceaux = append(morceaux, string(c))
			i++
			continue
		}
		// Lire un mot jusqu'au prochain espace ou signe spécial.
		debut := i
		for i < len(caracteres) {
			lettre := caracteres[i]
			if unicode.IsSpace(lettre) ||
				ponctuation(lettre) || lettre == '(' {
				break
			}
			if lettre == '\'' {
				avant := i > debut &&
					(unicode.IsLetter(caracteres[i-1]) ||
						unicode.IsDigit(caracteres[i-1]))
				apres := i+1 < len(caracteres) &&
					(unicode.IsLetter(caracteres[i+1]) ||
						unicode.IsDigit(caracteres[i+1]))
				if !avant || !apres {
					break
				}
			}
			i++
		}
		if i == debut {
			morceaux = append(morceaux, string(c))
			i++
		} else {
			morceaux = append(morceaux, string(caracteres[debut:i]))
		}
	}
	return morceaux
}
// modifierMot applique up, low ou cap à un mot.
func modifierMot(mot, commande string) string {
	switch commande {
	case "up":
		return strings.ToUpper(mot)
	case "low":
		return strings.ToLower(mot)
	case "cap":
		lettres := []rune(strings.ToLower(mot))
		if len(lettres) > 0 {
			lettres[0] = unicode.ToUpper(lettres[0])
		}
		return string(lettres)
	}
	return mot
}
// estSigne vérifie si un morceau contient seulement de la ponctuation.
func estSigne(mot string) bool {
	if mot == "" {
		return false
	}
	for _, c := range mot {
		if !ponctuation(c) {
			return false
		}
	}
	return true
}
// estMot permet de ne pas compter la ponctuation dans (up, 2).
func estMot(mot string) bool {
	for _, c := range mot {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			return true
		}
	}
	return false
}
// corrigerTexte applique les règles à chaque ligne.
func corrigerTexte(texte string) (string, error) {
	lignes := strings.Split(texte, "\n")
	for numero, ligne := range lignes {
		resultat, err := corrigerLigne(ligne)
		if err != nil {
			return "", fmt.Errorf("ligne %d : %w", numero+1, err)
		}
		lignes[numero] = resultat
	}
	return strings.Join(lignes, "\n"), nil
}
func corrigerLigne(texte string) (string, error) {
	morceaux := decouper(texte)
	mots := []string{}
	// 1. Appliquer les commandes en lisant de gauche à droite.
	for _, morceau := range morceaux {
		if morceau == "(hex)" || morceau == "(bin)" {
			if len(mots) == 0 {
				return "", fmt.Errorf("nombre manquant avant %s", morceau)
			}
			base := 16
			if morceau == "(bin)" {
				base = 2
			}
			nombre, err := strconv.ParseInt(mots[len(mots)-1], base, 64)
			if err != nil {
				return "", fmt.Errorf(
					"nombre invalide avant %s : %w", morceau, err,
				)
			}
			mots[len(mots)-1] = strconv.FormatInt(nombre, 10)
			continue
		}
		if strings.HasPrefix(morceau, "(") &&
			strings.HasSuffix(morceau, ")") {
			interieur := strings.TrimSpace(morceau[1 : len(morceau)-1])
			parties := strings.Split(interieur, ",")
			commande := strings.TrimSpace(parties[0])
			if commande == "up" || commande == "low" || commande == "cap" {
				nombre := 1
				if len(parties) == 2 {
					n, err := strconv.Atoi(strings.TrimSpace(parties[1]))
					if err != nil || n < 0 {
						return "", fmt.Errorf("commande invalide : %s", morceau)
					}
					nombre = n
				} else if len(parties) > 2 {
					return "", fmt.Errorf("commande invalide : %s", morceau)
				}
				for i := len(mots) - 1; i >= 0 && nombre > 0; i-- {
					if estMot(mots[i]) {
						mots[i] = modifierMot(mots[i], commande)
						nombre--
					}
				}
				continue
			}
		}
		mots = append(mots, morceau)
	}
	// 2. Transformer a en an devant une voyelle ou h.
	for i := 0; i < len(mots)-1; i++ {
		if mots[i] != "a" && mots[i] != "A" {
			continue
		}
		suivant := []rune(strings.ToLower(mots[i+1]))
		if len(suivant) > 0 &&
			strings.ContainsRune("aeiouh", suivant[0]) {

			if mots[i] == "A" {
				mots[i] = "An"
			} else {
				mots[i] = "an"
			}
		}
	}
	// 3. Reconstruire le texte avec les bons espaces.
	resultat := ""
	citationOuverte := false
	dernierEtaitOuverture := false
	for _, mot := range mots {
		if mot == "'" {
			if citationOuverte {
				// Apostrophe fermante : collée au mot.
				resultat += "'"
				citationOuverte = false
			} else {
				if resultat != "" && !strings.HasSuffix(resultat, " ") {
					resultat += " "
				}
				resultat += "'"
				citationOuverte = true
			}
			dernierEtaitOuverture = citationOuverte
			continue
		}
		if estSigne(mot) {
			// Ponctuation collée à gauche.
			// Les signes successifs se regroupent.
			resultat += mot
			dernierEtaitOuverture = false
			continue
		}
		if resultat != "" && !dernierEtaitOuverture {
			resultat += " "
		}
		resultat += mot
		dernierEtaitOuverture = false
	}
	return resultat, nil
}