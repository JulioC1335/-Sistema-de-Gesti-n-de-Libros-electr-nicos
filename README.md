# Biblioteca Digital (Proyecto Final-Evaluacion con el Docente)

**Realizado por :** Julio Cisneros  
**Objetivo del proyecto:**  
Desarrollar un sistema de gestión de biblioteca digital en Go, con API RESTful y frontend web para administrar libros, usuarios y préstamos de manera sencilla y funcional.

## Visión general
Este proyecto integra los conocimientos de las cuatro unidades de la materia:
1. **Unidad 1:** Análisis y diseño funcional de la estructura del sistema (DFD, paquetes, Go modules).
2. **Unidad 2:** Programación funcional en Go: structs, interfaces, manejo de errores y encapsulación.
3. **Unidad 3:** Concurrencia y pruebas: goroutines, canalización y testing de componentes.
4. **Unidad 4:** Generación de servicios web y serialización JSON: implementación de API REST y frontend con Fetch API.

## Funcionalidades principales
- **Gestión de libros:** CRUD completo (/api/books).  
- **Gestión de usuarios:** Registro y listado (/api/users).  
- **Préstamos y devoluciones:**  
  - Listar préstamos (/api/loans).  
  - Registrar un préstamo (/api/loans).  
  - Registrar una devolución (/api/returns).  
- **Interfaz web:** Página interactiva en / para administrar libros, usuarios y gestionar préstamos.  
- **Serialización JSON:**
