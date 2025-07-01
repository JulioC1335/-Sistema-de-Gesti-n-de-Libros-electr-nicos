# Biblioteca Digital Julio Cisneros

> **Proyecto Final – Evaluación con el Docente**
>
> Sistema en Go para gestionar una biblioteca digital: registra, consulta, edita y elimina libros y usuarios, y controla préstamos y devoluciones.
> Basado en programación funcional, guarda datos en archivos JSON locales y ofrece un frontend dinámico en HTML y JavaScript.

---

## Visión general

Este proyecto integra los conocimientos de:

1. **Análisis y diseño:** Estructura modular con paquetes Go.
2. **Programación en Go:** Uso de structs, manejo de errores y módulos.
3. **Persistencia y concurrencia:** Archivos JSON para datos y patrones concurrentes.
4. **Servicios web y frontend:** API RESTful con JSON y cliente JavaScript dinámico.

---

## Funcionalidades principales

* CRUD completo de libros y usuarios
* Gestión de préstamos y devoluciones
* Búsqueda por género o autor
* Interfaz web interactiva

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
   git clone https://github.com/tu_usuario/tu_repositorio.git biblioteca
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

Este proyecto incluye **11 servicios web.**

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

---

## Tecnologías usadas

* **Backend:** Go (net/http, encoding/json)
* **Persistencia:** Archivos JSON (`data/`)
* **Frontend:** HTML5, CSS3, JavaScript (Fetch API)
* **Herramientas:** Visual Studio Code, Git/GitHub

---
