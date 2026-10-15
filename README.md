
## 1. Présentation du projet
L'objectif est de créer un programme qui permet de lire un fichier texte, de modifier son contenu selon plusieurs règles, puis d'enregistrer le résultat dans un nouveau fichier.
Le programme fonctionne depuis le terminal en utilisant deux arguments :
- Le nom du fichier contenant le texte à corriger.
- Le nom du fichier dans lequel enregistrer le texte corrigé.
## 2. Fonctionnalités
Le programme applique les règles suivantes.
### Conversion hexadécimale `(hex)`
Convertit le nombre hexadécimal précédent en nombre décimal.
Exemple :
`1E (hex) files were added`
Résultat :
`30 files were added`
### Conversion binaire `(bin)`
Convertit le nombre binaire précédent en nombre décimal.
Exemple :
`It has been 10 (bin) years`
Résultat :
`It has been 2 years`
### Transformation en majuscules `(up)`
Transforme le mot précédent en majuscules.
Exemple :
`Ready, set, go (up) !`
Résultat :
`Ready, set, GO!`
### Transformation en minuscules `(low)`
Transforme le mot précédent en minuscules.
Exemple :
`I should stop SHOUTING (low)`
Résultat :
`I should stop shouting`
### Majuscule initiale `(cap)`
Met une majuscule à la première lettre du mot précédent et les autres lettres en minuscules.
Exemple :
`Welcome to the Brooklyn bridge (cap)`
Résultat :
`Welcome to the Brooklyn Bridge`
### Transformation de plusieurs mots
Les commandes `(up, n)`, `(low, n)` et `(cap, n)` permettent de modifier plusieurs mots précédents.
Exemple :
`This is so exciting (up, 2)`
Résultat :
`This is SO EXCITING`
### Correction de la ponctuation
Les signes `. , ! ? : ;` doivent être collés au mot précédent et séparés du mot suivant par un espace.
Exemple :
`I was sitting over there ,and then BAMM !!`
Résultat :
`I was sitting over there, and then BAMM!!`
### Groupes de ponctuation
Les groupes comme `...` et `!?` restent ensemble.
Exemple :
`I was thinking ... You were right`
Résultat :
`I was thinking... You were right`
### Correction des apostrophes
Les apostrophes utilisées pour encadrer une citation doivent être collées aux mots situés à l'intérieur.
Exemple :
`I am exactly how they describe me: ' awesome '`
Résultat :
`I am exactly how they describe me: 'awesome'`
Cette règle fonctionne également lorsque plusieurs mots sont entre les apostrophes.
### Transformation de `a` en `an`
Le mot `a` devient `an` lorsque le mot suivant commence par une voyelle (`a`, `e`, `i`, `o`, `u`) ou par `h`.
Exemple :
`There it was. A amazing rock!`
Résultat :
`There it was. An amazing rock!`
## 3. Organisation du projet
Le projet contient les fichiers suivants :
- **main.go** : point d'entrée du programme. Il récupère les arguments, lit le fichier d'entrée et écrit le résultat.
- **correction.go** : contient les fonctions permettant de modifier et de corriger le texte.
- **correction_test.go** : contient les tests pour vérifier le fonctionnement du programme.
- **go.mod** : définit le module Go du projet.
- **sample.txt** : contient un exemple de texte à corriger.
- **README.md** : explique le fonctionnement et l'utilisation du projet.
Le fichier `result.txt` est créé lors de l'exécution du programme.
## 4. Technologies utilisées
- Langage : Go
- Éditeur de code : Visual Studio Code
- Bibliothèque standard de Go : `os`, `fmt`, `strings`, `strconv` et `unicode`
## 5. Lancer le programme
Ouvrir le dossier du projet dans VS Code, puis ouvrir le terminal.
Exécuter :
```bash
go run . sample.txt result.txt
```
Cette commande :
1. Lance le programme Go.
2. Lit le contenu du fichier `sample.txt`.
3. Applique les règles de correction.
4. Enregistre le résultat dans `result.txt`.
Pour consulter le résultat, ouvrir directement le fichier `result.txt` dans VS Code.
## 6. Exemple d'utilisation
Contenu de `sample.txt` :
```text
1E (hex) files were added.
This is so exciting (up, 2) !
There it was. A amazing rock!
```
Après exécution, le fichier `result.txt` contient :
```text
30 files were added.
This is SO EXCITING!
There it was. An amazing rock!
```
## 7. Tests du programme
Le projet contient un fichier `correction_test.go` permettant de vérifier plusieurs fonctionnalités.
Pour lancer les tests :
```bash
go test -v
```
Si les tests réussissent, Go affiche `PASS`.
Pour effectuer une vérification supplémentaire du code :
```bash
go vet .
Les tests portent notamment sur les conversions, les transformations de mots, la ponctuation, les apostrophes et les articles.
## 8. Fonctionnement du code
Le programme fonctionne en plusieurs étapes :
1. `main.go` récupère les deux noms de fichiers donnés dans le terminal.
2. Le contenu du fichier d'entrée est lu avec `os.ReadFile`.
3. La fonction `corrigerTexte` traite le texte ligne par ligne.
4. La fonction `decouper` sépare le texte en mots, commandes et ponctuation.
5. Les fonctions de correction appliquent les transformations demandées.
6. Le texte est reconstruit avec les espaces et la ponctuation corrigés.
7. `os.WriteFile` enregistre le résultat dans le fichier de sortie.
