# Lizee app for Adrenaline Heroes

## Installation

pour faciliter l'installation et éviter de devoir installer golang et postgresql etc tout s'exécutera au travers de containers docker, à l'exeption du build de l'app react (je ne doute pas qu'npm soit installé sur votre machine.

### faire tout le travail avec le script de déploiement

```bash
sudo chmod +x deploy.sh
./deploy.sh
```

### build de l'app react

```bash
cd ./app && npm install && npm run build
```

### build du server et initialisation de la base postgresql

```bash
sudo docker-compose up --b
```

### première connexion à l'application:

Le serveur écoute le port 5000
sur chrome ou firefox: http://localhost:5000/ devrait charger l'application si aucune erreur ne s'est produite pendant le déploiement.

## Choix techniques

### serveur

J'ai choisi de poursuivre avec golang, c'est un langage intéressant notamment en ce qui concerne la compatibilité avec les api json.
En effet une simple structure suivi d'un tag permet de récupérer les données dans le format voulu (une structure).

#### utilisation:

Vous pouvez effectuer plusieurs requêtes:

Check if one product is available
====> GET http://localhost:5000/products/availability?product_id=6&from=2020-06-04&to=2020-06-05
Pass an order
====> POST http://localhost:5000/products/order
====> [{"product":{"availability":4,"id":1,"name":"tente trekking UL3","picture":""},"quantity":1,"from":"2020-06-03","to":"2020-06-04"}]
Get all existing categories of product
====> GET http://localhost:5000/categories

Check which products are available with category
====> GET http://localhost:5000/categories/products?categoryID=1&from=2020-06-04&to=2020-06-05

Modify quantity of corresponding product in database.
=====> POST to http://localhost:5000/availability/modifyquantity
=====> {"product_id":int, "quantity": int}

Get all available product between these dates
Notice that returned product_id is string here, to correspond to your demand.
=====> POST http://localhost:5000/availability/
=====> {"from":"2023-06-04","to":"2023-06-05"}

#### améliorations possibles :

- les requêtes envoyées pour la souscription d'une location ne sont pas lock, je n'ai pas encore eu le temps d'implémenter cette fonctionnalité, il faut ajouter un verrouillage de la base entre le moment où le serveur controle si les produits sont disponibles et celui où il insert la commande.
- Ne pas renvoyer systématiquement l'intégralité de l'objet product dans le cas de la commande erronée mais simplement l'id pour limiter la donnée en circulation. Ce sera simple à corriger.

### base de donnée

J'ai utilisé Postgresql, c'est une base de donnée assez classique pour ce genre de besoins.

#### améliorations possibles :

- Changer le mode de stockage des commandes sans affecter les performances de la base, il serait pratique de grouper touts les id des commandes dans un même champs de la table rental_order, mais je pense que les performances seraient affecter si on utilisait un tableau. à tester.
- Compléter la base de sorte que les commandes correspondent à un utilisateur.

### React.js

Parce que c'était demandé explicitement.

#### améliorations possibles :

- Utiliser redux ou un autre gestionnaire d'état pour éviter de promener les hooks à travers toute l'arborescence des composants react. Je pensais faire juste le widget au dépard, je me disais que le faire sans ne serait pas un problème.
- Ne pas renvoyer l'objet entier dans le cas de la commande mais simplement l'id, Ce sera simple à corriger.
- Créer un thème global plus élaboré.
- ajouter un prix aux items.
- changer la categorie sans avoir à effectuer une nouvelle requête.
- changer la date sans avoir à reconstituer mon panier (c'est un bug)
- meilleure gestion des erreurs
