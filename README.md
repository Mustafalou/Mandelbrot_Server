
<p align=right> Mustafa Yilmaz—Imane Ourraoui</p>

<p align=right> 17292-18316 </p>

<p align=center> <strong>Mandelbrot_Server </strong>  </p>



<p align=left> <strong>Documentation de l'API du serveur</strong>  </p>

<p>Les clients et les serveurs échangent des données au moyen du protocole HTTP.</p>
<p>Les fonctions GET et POST sont utilisées pour accéder aux données du serveur.</p>
<p>GET / : Sert un formulaire qui permet aux utilisateurs de spécifier les paramètres de l'image de l'ensemble de Mandelbrot qu'ils veulent générer.</p>
<p>POST /mandelbrot : Génère et sert une image d'ensemble Mandelbrot en réponse à la demande. Les paramètres suivants doivent être inclus dans la charge utile: </p>
<p>realMin: La valeur minimale de la composante réelle du nombre complexe à tracer.<br>
realMax: La valeur maximale de la composante réelle du nombre complexe à tracer.<br>
imagMin: La valeur minimale de la composante imaginaire du nombre complexe à tracer.<br>
imagMax: La valeur maximale de la composante imaginaire du nombre complexe à tracer.<br>
iterations: Le nombre d'itérations à utiliser lors du calcul de l'ensemble de Mandelbrot.<br>
width: La largeur de l'image générée, en pixels.<br>
height: La hauteur de l'image générée, en pixels.<br>
</p>

<p align=left> <strong>Stratégie de répartition du load balancer</strong>  </p>
 
 <p>Le load balancer de notre serveur HTTP utilise l'algorithme round-robin pour distribuer les demandes entrantes à un ensemble de serveurs/workers. <br>
 Lorsqu'une demande est reçue, l'équilibreur de charge sélectionne le serveur suivant dans la liste et lui transmet la demande. Si le serveur sélectionné n'est pas disponible ou s'il n'y a plus de serveurs dans la liste, l'équilibreur de charge renvoie une erreur au client.</p>
    
<p align=left> <strong>Bibliothèques utilisés</strong>  </p>

<p>fmt : Fournit un ensemble de fonctions pour les E/S formatées avec des chaînes de caractères.</p>
<p>image : Fournit un ensemble de types et de fonctions pour travailler avec des images.</p>
<p>image/color : Fournit des types et des fonctions pour représenter et manipuler les couleurs dans les images.</p>
<p>image/png : Fournit des fonctions pour coder et décoder les images PNG.</p>
<p>math/cmplx : Fournit des fonctions permettant de travailler avec des nombres complexes.</p>
<p>net/http : Fournit un ensemble de fonctions et de types pour construire des serveurs et des clients HTTP.</p>
<p>net/http/httputil : Fournit des fonctions utilitaires pour les serveurs et clients HTTP, y compris des fonctions pour les requêtes par procuration.</p>
<p>net/url : Fournit des fonctions pour analyser et manipuler les URLs.</p>
<p>os : Fournit une interface indépendante de la plate-forme pour les fonctionnalités du système d'exploitation.</p>
<p>runtime : Fournit des fonctions pour interagir avec le runtime Go.</p>
<p>strconv : Fournit des fonctions pour convertir des chaînes de caractères vers et depuis d'autres types de données.</p>
<p>sync : Fournit des types avec lesquels vous pouvez construire une synchronisation de plus haut niveau.</p>


