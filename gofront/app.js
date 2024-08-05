// URL de las APIs
const goApiUrl = 'http://localhost:8080/matrix';  // Cambia esto por la URL de tu API en Go
const nodeApiUrl = 'http://localhost:3000/process';  // Cambia esto por la URL de tu API en Node.js

// Función para obtener el token JWT
async function getToken() {
    const response = await fetch('http://localhost:8080/auth');  // Cambia esto por la URL de tu API que genera el token
    if (!response.ok) {
        throw new Error('No se pudo obtener el token');
    }
    const data = await response.json();
    return data.token;
}

// Función para enviar la matriz a la API en Go
async function sendMatrix(matrix) {
    try {
        const token = await getToken();
        const response = await fetch(goApiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({ matrix })
        });
        if (!response.ok) {
            throw new Error('Error al enviar la matriz');
        }
        const result = await response.json();
        return result;
    } catch (error) {
        console.error('Error:', error);
    }
}

// Función para mostrar el resultado
async function displayResult() {
    const matrixInput = document.getElementById('matrix-input').value;
    const matrix = JSON.parse(matrixInput);

    const result = await sendMatrix(matrix);
    document.getElementById('result').textContent = JSON.stringify(result, null, 2);
}

// Configurar el evento del botón
document.getElementById('send-matrix').addEventListener('click', displayResult);
