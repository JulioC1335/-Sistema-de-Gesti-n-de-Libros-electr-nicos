// Este archivo contiene las funciones JavaScript que llaman a la API
// y actualizan el contenido de la página para mostrar resultados.

// Función genérica para mostrar resultados en el contenedor "resultado"
function mostrarResultado(titulo, datos) {
  // Construyo el HTML inicial con un título y una lista
  let html = `<h3>${titulo}</h3><ul>`;
  
  // Si los datos son un array, itero y muestro cada objeto
  if (Array.isArray(datos)) {
    datos.forEach(item => {
      // Convierto cada par clave-valor a un string legible
      html += '<li>' +
        Object.entries(item)
          .map(([k, v]) => `<b>${k}:</b> ${v}`)
          .join(' | ')
        + '</li>';
    });
  } else {
    // Si no es un array (por ejemplo, una respuesta simple), lo muestro tal cual
    html += '<li>' + JSON.stringify(datos) + '</li>';
  }
  
  html += '</ul>';
  
  // Inserto el HTML generado dentro del elemento con id 'resultado'
  document.getElementById('resultado').innerHTML = html;
}

// Carga la lista de libros desde el endpoint y la muestra
function cargarLibros() {
  fetch('/api/books')                  // Petición GET a /api/books
    .then(res => res.json())           // Convierto la respuesta a JSON
    .then(data => mostrarResultado("📚 Libros", data)); // Muestro los libros
}

// Carga la lista de usuarios desde el endpoint y la muestra
function cargarUsuarios() {
  fetch('/api/users')                  // Petición GET a /api/users
    .then(res => res.json())           // Convierto la respuesta a JSON
    .then(data => mostrarResultado("👤 Usuarios", data)); // Muestro los usuarios
}

// Carga la lista de préstamos desde el endpoint y la muestra
function cargarPrestamos() {
  fetch('/api/loans')                  // Petición GET a /api/loans
    .then(res => res.json())           // Convierto la respuesta a JSON
    .then(data => mostrarResultado("📋 Préstamos", data)); // Muestro los préstamos
}

// Envía un nuevo libro al servidor y luego recarga la lista
function agregarLibro() {
  // Creo el objeto libro con los valores ingresados
  const book = {
    title: document.getElementById('titulo').value,
    author: document.getElementById('autor').value,
    stock: parseInt(document.getElementById('stock').value)
  };
  
  fetch('/api/books', {                // Petición POST a /api/books
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(book)         // Envío el libro como JSON
  })
    .then(() => cargarLibros());       // Recargo la lista de libros
}

// Envía un nuevo usuario al servidor y luego recarga la lista
function agregarUsuario() {
  const user = { name: document.getElementById('nombreUsuario').value };
  
  fetch('/api/users', {                // Petición POST a /api/users
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(user)         // Envío el usuario como JSON
  })
    .then(() => cargarUsuarios());     // Recargo la lista de usuarios
}

// Envía un nuevo préstamo al servidor y luego recarga la lista
function registrarPrestamo() {
  const loan = {
    user_id: parseInt(document.getElementById('usuarioID').value),
    book_id: parseInt(document.getElementById('libroID').value)
  };
  
  fetch('/api/loans', {                // Petición POST a /api/loans
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(loan)         // Envío el préstamo como JSON
  })
    .then(() => cargarPrestamos());    // Recargo la lista de préstamos
}

// Marca un préstamo existente como devuelto y recarga la lista
function finalizarPrestamo() {
  const id = document.getElementById("loanID").value;
  
  fetch(`/api/returnLoan?id=${id}`, {  // Petición POST a /api/returnLoan
    method: 'POST'
  })
    .then(res => {
      if (!res.ok) throw new Error("No encontrado"); // Manejo de error
      return res.json();
    })
    .then(() => cargarPrestamos())      // Recargo la lista de préstamos
    .catch(e => alert("Error: " + e.message)); // Muestro alerta en caso de fallo
}

// Elimina un libro por su ID y luego recarga la lista
function eliminarLibro() {
  const id = document.getElementById("idLibroEliminar").value;
  
  fetch(`/api/deleteBook?id=${id}`, {  // Petición DELETE a /api/deleteBook
    method: 'DELETE'
  })
    .then(() => cargarLibros());        // Recargo la lista de libros
}

// Elimina un usuario por su ID y luego recarga la lista
function eliminarUsuario() {
  const id = document.getElementById("idUsuarioEliminar").value;
  
  fetch(`/api/deleteUser?id=${id}`, {  // Petición DELETE a /api/deleteUser
    method: 'DELETE'
  })
    .then(() => cargarUsuarios());      // Recargo la lista de usuarios
}
