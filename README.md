# Go-reloaded

Lien GitHub : https://github.com/Amine76409/Go-reloaded.git 

Ce projet est un programme en Go. Il lit un texte dans un fichier, le corrige, et met le résultat dans un autre fichier.

Pour le lancer, on écrit : go run . sample.txt result.txt
Le fichier sample.txt est le texte de départ et le fichier result.txt est le texte corrigé.

Comment ça marche :

Le programme lit le fichier et corrige le texte ligne par ligne. Pour chaque ligne, il coupe d'abord la ligne en mots. Ensuite, il regarde s'il y a des mots entre parenthèses, comme (up) ou (hex). Ces mots disent quoi faire au mot qui est juste avant.

Avec (hex) ou (bin), on change le nombre en nombre normal. Avec (up), le mot passe en majuscules. Avec (low), il passe en minuscules. Avec (cap), il prend une majuscule au début. Si on ajoute un nombre, par exemple (up, 2), ça change les 2 mots d'avant.

Après ça, le programme remet les mots ensemble. Il colle la ponctuation au mot d'avant et met un espace après. Il colle aussi les guillemets au texte qu'ils entourent. Enfin, il change a en an quand le mot d'après commence par une voyelle ou un h.

Pour finir, il écrit le texte corrigé dans le deuxième fichier.

Un exemple :

Avant : Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
Après : Simply add 66 and 2 and you will see the result is 68.

Les fichiers du projet :

main.go lit et écrit les fichiers.
process.go corrige le texte.
tokenize.go coupe le texte en mots.
convert.go s'occupe de l'hexadécimal et du binaire.
casing.go s'occupe des majuscules et des minuscules.
punctuation.go s'occupe de la ponctuation.
quotes.go s'occupe des guillemets.
article.go change a en an.
Les fichiers qui finissent par _test.go sont les tests.

Pour lancer les tests, on écrit : go test ./...
Pour lancer le programme et voir les textes :

D’abord, on lance le programme avec : go run . sample.txt result.txt
Ensuite, pour voir le texte de départ en anglais, on écrit : Get-Content sample.txt
Puis, pour voir le texte corrigé, on écrit : Get-Content result.txt