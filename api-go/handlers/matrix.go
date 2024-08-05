package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gonum.org/v1/gonum/mat"
)

func MatrixHandler(c *fiber.Ctx) error {
	var request models.MatrixRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Imprimir la matriz recibida para depuración
	fmt.Println("Received matrix:", request.Matrix)

	// Convertir la matriz a formato gonum
	rows := len(request.Matrix)
	cols := len(request.Matrix[0])
	data := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i*cols+j] = request.Matrix[i][j]
		}
	}
	matrix := mat.NewDense(rows, cols, data)

	// Realizar la factorización QR
	var qr mat.QR
	qr.Factorize(matrix)

	// Obtener la matriz Q
	q := mat.NewDense(rows, rows, nil) // Matriz cuadrada para almacenar Q
	qr.QTo(q)

	// Obtener la matriz R
	r := mat.NewDense(rows, cols, nil) // Matriz rectangular para almacenar R
	qr.RTo(r)

	// Convertir las matrices resultantes a arrays de arrays de float64
	qResult := make([][]float64, q.RawMatrix().Rows)
	for i := 0; i < q.RawMatrix().Rows; i++ {
		qResult[i] = q.RawRowView(i)
	}

	rResult := make([][]float64, r.RawMatrix().Rows)
	for i := 0; i < r.RawMatrix().Rows; i++ {
		rResult[i] = r.RawRowView(i)
	}

	// Crear una respuesta con las matrices Q y R
	response := map[string]interface{}{
		"Q": qResult,
		"R": rResult,
	}

	// Enviar la matriz resultante a la API en Node.js
	resultJSON, err := json.Marshal(response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot encode JSON"})
	}

	// Generar un token JWT
	token, err := utils.GenerateJWT()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot generate JWT"})
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/process", bytes.NewBuffer(resultJSON))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Solicitud para la API de Node.js Erronea"})
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "No se envía request a Node.js API"})
	}
	defer resp.Body.Close()

	var stats map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Decodificar la respuesta de la API de Node.js Erroneo"})
	}

	return c.JSON(stats)
}
