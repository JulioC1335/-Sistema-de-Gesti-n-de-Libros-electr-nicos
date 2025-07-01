# Biblioteca Digital Julio Cisneros

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22-blue?logo=go" alt="Go version" />
  <img src="https://img.shields.io/badge/Backend-Go%20net%2Fhttp-blue" alt="Backend" />
  <img src="https://img.shields.io/badge/API-RESTful-orange" alt="API RESTful" />
  <img src="https://img.shields.io/badge/Frontend-HTML_CSS_JS-lightgrey" alt="Frontend" />
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License" />
</p>

> **Proyecto Final – Evaluación con el Docente**
> **Autor:** Julio Cisneros
> **Objetivo:** Desarrollar un sistema web de gestión de biblioteca digital en Go, con API RESTful y frontend interactivo para administrar libros, usuarios y préstamos de forma clara y eficiente.

---

## Visión general

Este proyecto integra los conocimientos de:

1. **Análisis y diseño:** Estructura modular con paquetes Go.
2. **Programación en Go:** Uso de structs, manejo de errores y módulos.
3. **Persistencia y concurrencia:** Archivos JSON para datos y patrones concurrentes.
4. **Servicios web y frontend:** API RESTful con JSON y cliente JavaScript dinámico.

---

## Funcionalidades principales

* **Gestión de libros:** CRUD completo (crear, leer, actualizar, eliminar).
* **Gestión de usuarios:** Registro y listado.
* **Préstamos y devoluciones:** Registrar y finalizar préstamos.
* **Filtrado y búsqueda:** Por género o autor.
* **Interfaz web interactiva:** Panel visual para operaciones de usuario.
* **Serialización JSON:** Consumo y respuesta en formato JSON.

---

## Estructura del proyecto

* **biblioteca/**

  * `go.mod`
  * `main.go`
  * **data/**

    * `data.go`
  * **models/**

    * `book.go`
    * `user.go`
    * `loan.go`
  * **handlers/**

    * `handlers.go`
    * `loans.go`
  * **static/**

    * `index.html`
    * `style.css`
    * `script.js`

---

## Instalación

1. Instala Go (≥ 1.22):

   * **Linux/macOS:** `brew install go` o descarga desde [golang.org](https://golang.org).
   * **Windows:** Descarga el instalador desde [golang.org](https://golang.org).
2. Clona el repositorio:

   ```bash
   git clone <https://github.com/JulioC1335/-Sistema-de-Gesti-n-de-Libros-electr-nicos.git> biblioteca
   cd biblioteca
   ```
3. Inicializa el módulo Go y descarga dependencias:

   ```bash
   go mod init biblioteca
   go mod tidy
   ```

---

## Ejecución

Ejecuta el servidor desde la carpeta raíz:

```bash
go run main.go
```

Accede en tu navegador a: `http://localhost:8080`

---

## API RESTful

Este proyecto incluye **11 servicios web**, superando el mínimo de 9 requeridos.
Todas las rutas devuelven y aceptan datos en formato JSON.

1. **GET** `/api/books` – Listar todos los libros
2. **POST** `/api/books` – Crear un nuevo libro
3. **DELETE** `/api/deleteBook` – Eliminar un libro por ID
4. **GET** `/api/books/byGenre` – Filtrar libros por género
5. **GET** `/api/books/byAuthor` – Filtrar libros por autor
6. **GET** `/api/users` – Listar todos los usuarios
7. **POST** `/api/users` – Crear un nuevo usuario
8. **DELETE** `/api/deleteUser` – Eliminar un usuario por ID
9. **GET** `/api/loans` – Listar todos los préstamos
10. **POST** `/api/loans` – Registrar un nuevo préstamo
11. **POST** `/api/returnLoan` – Finalizar un préstamo

---

## Frontend

* **`index.html`**: Interfaz principal con secciones de visualización y formularios.
* **`style.css`**: Estilos modernos y responsive.
* **`script.js`**: Lógica de interacción y llamadas a la API.

<p align="center">
  <img src="https://user-images.githubusercontent.com/tu_usuario/tu_repo/screenshot.png" alt="Vista previa de la aplicación" width="600"/>
</p>

---

## Tecnologías usadas

* **Backend:** Go (net/http, encoding/json)
* **Persistencia:** Archivos JSON (`data/`)
* **Frontend:** HTML5, CSS3, JavaScript (Fetch API)
* **Herramientas:** Visual Studio Code, Git/GitHub


---
