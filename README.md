# GitDock
## app en go para controlar  las versiones
https://github.com/aratan/test.git
> /ultim ghp_TOKENxxxxxxxx aratan test main 
<p>
> /ultim ghp_TOKENxxxxxxxx aratan test dev 

La app la primera vez te detecta cambios que es la configuracion base.

el nombre de usuario y repo se tienen que llamar igual en github y en docker-hub

https://hub.docker.com/r/aratancoders/test
https://github.com/aratancoders/test

Va ha estar checkeando el la rama del github que le digas y si detecta
cambios descarga el repo y te crea una imagen en docker.

Como las aplicaciones pueden tener puertos distintos y servicios,
recomiendo que te adactes el Dockerfile y docker-componer
con tu configuración en el raiz del proyecto.


Ademos del token de github tienes que configurar la validacion con docker
docker logout
docker login

* ya puedes usar la app *
ahora con avisos por correo
