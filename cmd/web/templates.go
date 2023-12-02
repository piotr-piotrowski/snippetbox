package main

import "snippetbox.pp.com/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
