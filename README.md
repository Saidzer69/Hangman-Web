Hangman-Web-Gamify
Objectif
Ce projet consiste à créer une version web du jeu du pendu avec une interface utilisateur améliorée et des fonctionnalités supplémentaires pour rendre le jeu plus interactif et amusant.

Structure des fichiers
hangman-web/
│
├── main.go
├── templates/
│   ├── index.html
│   ├── start.html
│   ├── win.html
│   ├── lose.html
│   └── scoreboard.html
└── static/
    └── style.css
Contenu des fichiers
main.go
Le fichier main.go contient le code du serveur en Go. Il gère les routes, les requêtes HTTP et la logique du jeu.

templates/
Le dossier templates/ contient les fichiers HTML pour les différentes pages de l'application :

index.html : Page principale du jeu.
start.html : Page de démarrage où l'utilisateur entre son nom.
win.html : Page affichée lorsque l'utilisateur gagne.
lose.html : Page affichée lorsque l'utilisateur perd.
scoreboard.html : Page affichant le tableau des scores.
static/
Le dossier static/ contient les fichiers CSS pour le style de l'interface utilisateur :

style.css : Fichier CSS principal.
Fonctionnalités
Page de démarrage : Permet à l'utilisateur d'entrer son nom et de commencer une nouvelle partie.
Jeu du pendu : L'utilisateur devine les lettres du mot caché. Le nombre d'essais restants est affiché.
Page de victoire : Affichée lorsque l'utilisateur devine correctement le mot.
Page de défaite : Affichée lorsque l'utilisateur utilise tous ses essais sans deviner le mot.
Tableau des scores : Affiche les scores des utilisateurs.
Instructions pour exécuter le projet
Cloner le dépôt :

git clone <URL_DU_DEPOT>
cd hangman-web
Lancer le serveur Go :

go run main.go
Ouvrir le navigateur et accéder à l'URL suivante : http://localhost:8080

Technologies utilisées
Go : Langage de programmation utilisé pour le serveur.
HTML : Utilisé pour la structure des pages web.
CSS : Utilisé pour le style des pages web.
