# Transformación de textos explicativos a lenguaje de marcado para mejorar la generación de discursos hablados

## Framework

- [Go version go1.18.3 linux/amd64](https://go.dev/): Lenguaje de programación multiparadigma. [Documentación](https://go.dev/doc/)

<img src="https://blog.wemake.pe/posts/2022-01-11-gopher.svg" alt="centered image"  height="100" title="image Title" /> 
<img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" alt="centered image"  height="50" title="image Title" /> 

- [Gin Web Framework](https://github.com/gin-gonic/gin): Servicio API REST. [Documentación](https://github.com/gin-gonic/gin#api-examples)

<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" alt="alt text"  height="100" title="image Title" />

## Authentication

 Se debe tener configurada la [autenticación](https://cloud.google.com/text-to-speech/docs/libraries?hl=es-419#client-libraries-usage-go) de google cloud para usar su servicio de [Text-to-Speech](https://cloud.google.com/text-to-speech/docs/create-audio-text-client-libraries?hl=es-419). Duante el desarollo del software se estan utilizando credenciales de servicio. Para una configuración global, añadir la variable de entorno en el archivo **~/.bashrc** o **~/.profile** de la siguiente forma:

        export GOOGLE_APPLICATION_CREDENTIALS="/home/user/ejemplo/credential_file.json"

## Quick Start
1. Para ejecutar la aplicación, dentro de la carpeta **.../auto_SSML_2.0/src/** ejecutar el comando:

        $ go run main.go

