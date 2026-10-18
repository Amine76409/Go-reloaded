# go-reloaded

Outil en Go qui lit un fichier texte, corrige le texte, puis écrit le résultat dans un autre fichier.

## Usage

```
go run . sample.txt result.txt
```

## Règles

- `(hex)` : le mot avant est un nombre hexadécimal, il est remplacé par sa valeur décimale
- `(bin)` : le mot avant est un nombre binaire, il est remplacé par sa valeur décimale
- `(up)` : le mot avant passe en majuscules
- `(low)` : le mot avant passe en minuscules
- `(cap)` : le mot avant prend une majuscule au début
- `(up, 2)`, `(low, 2)`, `(cap, 2)` : pareil mais pour les 2 mots avant (ou un autre nombre)
- `. , ! ? : ;` sont collés au mot avant et suivis d'un espace
- `...` et `!?` restent groupés
- `' '` sont collés au mot (ou aux mots) qu'ils entourent
- `a` devient `an` devant une voyelle ou un `h`

## Fichiers

- `main.go` : lit les arguments et les fichiers
- `process.go` : applique toutes les règles ligne par ligne
- `tokenize.go` : découpe le texte en mots et marqueurs
- `convert.go` : `(hex)` et `(bin)`
- `casing.go` : `(up)`, `(low)`, `(cap)`
- `punctuation.go` : ponctuation
- `quotes.go` : apostrophes
- `article.go` : `a` en `an`
- `*_test.go` : tests unitaires
- `sample.txt` et `expected.txt` : exemple et résultat attendu

## Tests

```
go test ./...
```
