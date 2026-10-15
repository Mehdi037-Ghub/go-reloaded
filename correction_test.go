package main
import "testing"
func TestRegles(t *testing.T) {
	tests := []struct{ entree, attendu string }{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"Bonjour ! ? monde", "Bonjour!? monde"},
		{"Attends . . . maintenant", "Attends... maintenant"},
		{"don't stop,it's fine", "don't stop, it's fine"},
		{"hello (up) (low) (cap)", "Hello"},
		{"a house", "an house"},
		{"Hello!world", "Hello! world"},
		{"' hello world ' again", "'hello world' again"},
	}
	for _, test := range tests {
		obtenu, err := corrigerTexte(test.entree)
		if err != nil || obtenu != test.attendu {
			t.Errorf(
				"Entrée %q : obtenu %q, attendu %q, erreur %v",
				test.entree, obtenu, test.attendu, err,
			)
		}
	}
}