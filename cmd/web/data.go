// Filename: cmd/web/data.go
package main

import (
	"github.com/Aazan-Iqbal/hello/internal/models"
)

type templateData struct {
	Question *models.Question
	Flash    string
}
