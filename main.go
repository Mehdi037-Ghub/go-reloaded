package main
import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Utilisation : go run . entree.txt sortie.txt")
		return
	}
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}
	resultat, err := corrigerTexte(string(contenu))
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Erreur d'ecriture :", err)
	}
}