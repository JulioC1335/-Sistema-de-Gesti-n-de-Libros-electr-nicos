
function mostrarResultado(titulo, datos) {
  let html = `<h3>${titulo}</h3><ul>`;
  if (Array.isArray(datos)) {
    datos.forEach(item => {
      html += '<li>' + Object.entries(item).map(([k, v]) => `<b>${k}:</b> ${v}`).join(' | ') + '</li>';
    });
  } else {
    html += '<li>' + JSON.stringify(datos) + '</li>';
  }
  html += '</ul>';
  document.getElementById('resultado').innerHTML = html;
}

function cargarLibros() {
  fetch('/api/books')
    .then(res => res.json())
    .then(data => mostrarResultado("📚 Libros", data));
}
function cargarUsuarios() {
  fetch('/api/users')
    .then(res => res.json())
    .then(data => mostrarResultado("👤 Usuarios", data));
}
function cargarPrestamos() {
  fetch('/api/loans')
    .then(res => res.json())
    .then(data => mostrarResultado("📋 Préstamos", data));
}

function agregarLibro() {
  const book = {
    title: document.getElementById('titulo').value,
    author: document.getElementById('autor').value,
    stock: parseInt(document.getElementById('stock').value)
  };
  fetch('/api/books', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(book)
  })
    .then(res => res.json())
    .then(() => cargarLibros());
}
function agregarUsuario() {
  const user = {
    name: document.getElementById('nombreUsuario').value,
    password: document.getElementById('claveUsuario').value
  };
  fetch('/api/users', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(user)
  })
    .then(res => res.json())
    .then(() => cargarUsuarios());
}
function registrarPrestamo() {
  const loan = {
    user_id: parseInt(document.getElementById('usuarioID').value),
    book_id: parseInt(document.getElementById('libroID').value)
  };
  fetch('/api/loans', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(loan)
  })
    .then(res => res.json())
    .then(() => cargarPrestamos());
}
function eliminarLibro() {
  const id = document.getElementById("idLibroEliminar").value;
  fetch(`/api/deleteBook?id=${id}`, { method: 'DELETE' })
    .then(res => res.json())
    .then(() => cargarLibros());
}
function eliminarUsuario() {
  const id = document.getElementById("idUsuarioEliminar").value;
  fetch(`/api/deleteUser?id=${id}`, { method: 'DELETE' })
    .then(res => res.json())
    .then(() => cargarUsuarios());
}
function eliminarPrestamo() {
  const id = document.getElementById("idPrestamoEliminar").value;
  fetch(`/api/deleteLoan?id=${id}`, { method: 'DELETE' })
    .then(res => res.json())
    .then(() => cargarPrestamos());
}
function iniciarSesion() {
  const name = document.getElementById("loginName").value;
  const pass = document.getElementById("loginPass").value;
  fetch('/api/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, password: pass })
  })
    .then(res => {
      if (res.ok) return res.json();
      throw new Error("Credenciales incorrectas");
    })
    .then(() => {
      document.getElementById("contenido").style.display = "block";
    })
    .catch(err => alert(err.message));
}
