
<p align=right> Mustafa Yilmaz—Imane Ourraoui</p>

<p align=right> 17292-18316 </p>

<p align=center> <strong>Mandelbrot_Server </strong>  </p>
A server that can compute mandelbrot using multiple cores and shows the result 


<p align=left> <strong>Documentation de l'API de votre serveur</strong>  </p>
    Les clients et les serveurs échangent des données au moyen du protocole HTTP.
    Les fonctions GET  et POST sont utilisées pour accéder aux données du serveur.
    La fonction form utilise la méthode GET.
    La fonction mandelbrot utilise la méthode POST.

<p align=left> <strong>Stratégie de répartition du load balancer</strong>  </p>
 Lorsqu'une demande est reçue, l'équilibreur de charge sélectionne le serveur suivant dans la liste et lui transmet la demande. Si le serveur sélectionné n'est pas disponible ou s'il n'y a plus de serveurs dans la liste, l'équilibreur de charge renvoie une erreur au client.
    
<p align=left> <strong>Bibliothèques utilisés</strong>  </p>

<p>net/http : Fournit un ensemble de fonctions et de types pour construire des serveurs et des clients HTTP.</p>
<p>net/http/httputil : Fournit des fonctions utilitaires pour les serveurs et les clients HTTP, y compris des fonctions pour les requêtes par procuration.</p>
<p>net/url : Fournit des fonctions pour analyser et manipuler les URLs.</p>
<p>image : Fournit un ensemble de types et de fonctions pour travailler avec des images.</p>
<p>image/color : Fournit des types et des fonctions pour représenter et manipuler les couleurs dans les images.</p>
<p>image/png : Fournit des fonctions pour coder et décoder les images PNG.</p>
<p>math/cmplx : Fournit des fonctions permettant de travailler avec des nombres complexes.</p>
<p>sync : Fournit des types avec lesquels vous pouvez construire une synchronisation de plus haut niveau.</p>
