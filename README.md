# Ultim
## app en go para controlar  las versiones
https://github.com/aratan/test.git
> /ultim ghp_TOKENxxxxxxxx aratan test main <
> /ultim ghp_TOKENxxxxxxxx aratan test dev <

La app la primera vez te detecta cambios que es la configuracion base.

Va ha estar checkeando el la rama del github que le digas y si detecta
cambios descarga el repo y te crea una imagen en docker.

Como las aplicaciones pueden tener puertos distintos y servicios,
recomiendo que te adactes el Dockerfile y docker-componer
con tu configuración en el raiz del proyecto.

Para ver el docker:
docker ps -a
docker run -it xxxxxxxx bash
