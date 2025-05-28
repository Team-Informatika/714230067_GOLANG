package main

import (
	"github.com/gofiber/fiber/v2"
)

type Mahasiswa struct {
	Nama    string `json:"nama"`
	NPM     string `json:"npm"`
	Umur    int    `json:"umur"`
	Jurusan string `json:"jurusan"`
}

func GetMahasiswa(c *fiber.Ctx) error {
	data := Mahasiswa{
		Nama:    "Galuh Sanjaya",
		NPM:     "714230067",
		Umur:    10,
		Jurusan: "Teknik Informatika",
	}
	return c.JSON(data)
}
func main() {
	app := fiber.New()

	app.Get("/mahasiswa", GetMahasiswa)

	app.Listen(":3000")
}
