# Guía para colaborar con el proyecto

Esta va a ser una pequeña guía donde se va a describir como se puede aportar en este proyecto en la traducción, corrección y actualización.

Para esto vamos a explicar como copiar lo que se ha realizado hasta el momento (realizar un fork), guardarlo en tu PC (realizar un clone), y enviarnos tus aportes (realizar un pull request).
Para que de esta forma podamos agregar tu aporte al proyecto sin que se tengan inconvenientes.

***

## Trabajando con github

El esquema de trabajo que nos propone **github** puede parecer confuso al principio, pero en realidad es bastante simple, tan sólo hay que tener presente que uno está trabajando con tres repositorios distintos y entender el rol que cada uno cumple.

**Repositorio upstream:** es el repositorio en github de la aplicación con la cual vamos a colaborar. En nuestro caso el repositorio está en https://github.com/Jackgris/build-web-application-with-golang_ES. No podremos hacer ningún cambio directamente sobre este repositorio, ya que no tendremos permisos de escritura sobre el mismo. Lo que haremos es crear nuestra propia copia (fork) del mismo.

**Repositorio origin:** es nuestra copia del repositorio que forkeamos alojada en github. A este repositorio enviaremos los cambios que efectuaremos en nuestro repositorio local.

**Repositorio local:** es el repositorio local que tendremos en nuestra propia estación de trabajo (o PC).

Comúnmente realizaremos los cambios trabajando en nuestro repositorio local. Cuando hayamos confirmado (realizando un commit) nuestros cambios localmente, los enviaremos (haciendo un push) a nuestro repositorio origin alojado en github. Luego haremos un pedido (o pull request) al repositorio upstream solicitándole que integre los cambios del repositorio origin.

Este es el esquema general de trabajo con github, que nos servirá no sólo para colaborar con esta traducción, sino con cualquier proyecto libre alojado en **github**, ya sea texto como en este caso, o aportar código.

***

## Crear una cuenta en github

Lo primero que tendremos que hacer es navegar a https://github.com/signup/free y crear una cuenta gratuita en **github**.

Luego deberás instalar git en tu equipo y configurar tu cuenta siguiendo las instrucciones para Linux, Mac OS o para windows, según el sistema operativo que estés utilizando.

***

## Hacer un fork del repositorio upstream

Navega al sitio de **github** e ingresa con tu usuario y password. Luego tendrás que navegar a https://github.com/Jackgris/build-web-application-with-golang_ES y hacer clic en **“fork”**. Con esto ya habrás creado tu repositorio origin.

***

## Clonar nuestro repositorio origin en un repositorio local

En nuestra estación de trabajo ejecutamos las siguientes instrucciones (debes ajustarlas según tu nombre de usuario, remplazando `[tu_cuenta_de_github]` por tu nombre):

    cd ~
    mkdir apps
    cd apps
    git clone git@github.com:[tu_cuenta_de_github]/build-web-application-with-golang_ES.git

De esta manera nos quedará nuestra copia en `~/apps/build-web-application-with-golang_ES`. Para traducir la documentación sigue los pasos que se encuentran abajo en la guía. 

***

## Guardar los cambios y enviar nuestras traducciones

Una vez que hemos terminado de traducir algún archivo, vamos a querer enviarlo al administrador de **wikiLibGDX_es** para que sean incluidos en producción. Para ello tendremos que confirmar los cambios en nuestro repositorio local, actualizar nuestro repositorio origin en **github**, y hacer un pedido para que estos cambios sean incluido en el repostiorio upstream.

Mediante el comando git status podemos ver los archivos que hemos modificado localmente. Con el comando **git add**, le indicamos a git cuáles son los archivos que vamos a querer confirmar (commit). Con **git commit** confirmamos los cambios en nuestro repositorio local y finalmente, con **git push origin master** subimos los cambios locales a nuestro repositorio origin. Situados en el directorio del repositorio, en nuestro ejemplo `~/apps/build-web-application-with-golang_ES`, estos serían los comandos que deberemos ejecutar:

    git status
    git add .
    git commit -m "traduje una nueva página"
    git push origin master

En caso de haber cargado un incidente para esa página, es sumamente útil incluir el texto fixes seguido del signo `#` y el número del incidente (o **issue**) en el texto que utilizamos al hacer el commit, de la siguiente manera `git commit -m "fixes #34 - traduje una nueva página"`. De esta manera **github** relacionará nuestros cambios con el incidente y lo cerrará automáticamente.

Luego navegamos a nuestro repositorio origin, en este caso `https://github.com/[nuestro-usuario]/build-web-application-with-golang_ES`, en donde ya deberías poder ver los cambios que acabamos de subir. Para solicitar al repositorio upstream que integre nuestros cambios simplemente hacemos clic en el botón pull request. Y ahora tan sólo resta esperar que el administrador del repositorio upstream (en este caso **jackgris**) revise los cambios y los impacten en la aplicación.

***

## Manteniendo nuestros repositorios actualizados

A medida que sigamos trabajando, querremos mantener actualizados nuestros repositorios (local y origin) con los cambios del repositorio upstream. Para eso configuramos el repositorio upstream como un repositorio remoto (remote) vinculado a nuestro repositorio local. Situados en el directorio del repositorio, en nuestro ejemplo `~/apps/build-web-application-with-golang_ES`, deberemos ejecutar:

    git remote add upstream git@github.com:Jackgris/build-web-application-with-golang_ES.git

Cuando querramos traer las actualizaciones de upstream, deberemos ejecutar:

    git fetch upstream
    git merge upstream/master

***

## Forma recomendada de trabajo

Estas van a ser una serie de pasos recomendados a seguir, para que no se superpongan con el trabajo las presonas que deseen contribuir con el proyecto.

- 1. Dirigirse a al repositorio upstream (https://github.com/Jackgris/git@github.com:Jackgris/build-web-application-with-golang_ES.git/) y hacer clic en la parte superior derecha donde dice **"Issues"**.
- 2. Hacer clic en **"New Issue"**, y como mínimo escribir en el título sobre que archivo se va a trabajar o crear, y si es posible alguna descripción (por ej: se va a corregir la ortografía, se va a añadir una sección, etc)  
- 3. Actualizar tu repositorio local y comenzá a trabajar.
- 4. Envia el aporte por medio de un **pull request**, en lo posible en el commit usar la recomendación descripta anteriormente sobre añadir en el texto del commit sobre el fixes y el númeto de issue.

Para realizar esta serie de paso lo más recomendable es que se trabaje unicamente sobre el archivo que se indico en el **issue** que ha abierto.
Tenga en cuenta de no subir ningun archivo temporal que puede llegar a crear su editor de texto (esto lo puede hacer sencillamente usando el comando `git status`, antes de por ejemplo realizar un `git add .`, de esta manera podrá remover  antes el archivo temporal de su repositorio local)

Si realiza esta serie de pasos, seguramente podremos trabajar de una forma mucho mas cómoda y productiva.

***

## Crear un nuevo archivo

Para crear un nuevo archivo tiene que tener en cuenta las siguiente convenciones:

- El mismo deberá tener la extensión .md
- Si el archivo va a tratar sobre una sección que no se encuentra en el indice, recuerde agregarlo en el archivo preface.md que se encuentra en la carpeta eBook del respositorio.
- Si el mismo trata sobre una de las secciones del indice, el nombre del archivo deberá seguir la siguiente convención, por ejemplo para tratar sobre la sección **1.3** el archivo deberiá llamarse `01.3.md`, para la sección **4.7** este deberá llamarse `04.7.md`. Para capítulos superiores al **9** no se tiene que tener en cuenta el **0** del inicio, por ejemplo para la sección **14.2** el archivo debería ser `14.2.md`.
Cuando se trata de la primera sección de un capítulo, esta deberá ser la número **0**, por ejemplo para el comienzo del capítulo **4** el archivo deberiá ser `04.0.md` y para el **16** el archivo debería ser `16.0.md`.
- Para escribir en el mismo tenga en cuanta la sintaxis que debe seguir para por ejemplo añadir porciones de código, que se explica en la siguiente sección.

***

## Sintaxis 

**Separar secciónes:**
Para poder realizar una separación entre secciones se deben usar solo tres * así  `***` al principio de la línea.

**Texto en negrita:**
Para hacer que un texto se vea en negrita se deben usar `**` antes y despues del texto, por ejemplo `**texto en negrita**` lo que se vería así: **texto en negrita**.

**Añadir código:**

Bueno para esto tenemos diferentes opciones:

- 1. Separar el texto del comienzo de la línea con un `TAB`.
- 2. Utilizar una comilla de este tipo **`** al comienzo del texto y otra al final en una misma línea. 
- 3. Utilizar tres comillas de este tipo **`** al comienzo y al final de una porción de texto.
- 4. En los casos anteriores se muestra el texto como código pero el mismo no se colorea ya que no le espesíficamos el tipo de código para que sepa en que forma colorearlo, por ejemplo si nosotro queremos ver algo así:

```java
private String hola = "hola";
```

Debemos hacer lo siguiente:

    ```java
    private String hola = "hola";
    ```
Como se puede ver, solo debemos añadir que tipo de código vamos a encerrar (en este caso java) despues de las primeras tres comillas.

**Titulos:**

Para añadir un titulo, va a depender el tipo de titulo que deseamos añadir a nuestro texto la cantidad de signos `#` con los que debemos iniciar la línea de texto, vamos a ver algunos ejemplos:

# Ejemplo 1

Eso es `# Ejemplo 1`

## Ejemplo 2

Eso es `## Ejemplo 2`

### Ejemplo 3

Eso es `### Ejemplo 3`

**Indices:**

Para armar algunos indices, como el indice principal que vemos en esta traducción y como voy a mostrar aquí:

- 1. Primer orden
- 2. Primer orden
    - 2.1. Segundo orden
- 3. Primer orden
    - 3.1. Segundo orden
        - 3.1.1. Tercer orden

Debería usar un código como el siguiente:

```
- 1. Primer orden
- 2. Primer orden
    - 2.1. Segundo orden
- 3. Primer orden
    - 3.1. Segundo orden
        - 3.1.1. Tercer orden
```

**Añadir imágenes y enlaces:**

Para poder agregar una imagen en el texto, primero vamos a ver un ejemplo, este es el código que se utiliza:

    ![Descripción de la imagen](PATH o ruta a la imagen)

Como se puede ver, la idea es la siguiente, comenzamos la línea con `!` y entre `[]` colocamos una descripción de la imagen, para luego colocar el `PATH` o ruta a la imagen, que puede estar dentro de el arbol de directorio de este proyecto.

Para poder añadir un enlace, debe encerrar entre `[]` el texto a mostrar y entre `()` el enlace donde desea apuntar, por ejemplo [Google](https://www.google.com.ar) esto es `[Google](https://www.google.com.ar)`

Ahora si deseamos añadir una imagen que nos dirija a un enlace, debemos hacer una especie de combinación de lo anteriormente explicado. La idea es armar un texto con enlace, solo que donde va el texto a mostrar, debemos colocar la imágen a mostrar. En otras palabras esta es la idea:

    [![Descripción de la imagen](PATH o ruta a la imagen)](enlace donde deseamos que nos dirija)

## Enlaces

- [Inicio](../README.md)
- [Indice](preface.md)