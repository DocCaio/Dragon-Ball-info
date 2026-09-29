package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const BaseURL = "https://dragonball-api.com/api/characters/"


type Planet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	IsDestroyed bool   `json:"isDestroyed"`
	DeletedAt   string `json:"deletedAt"`
}

type Personagens struct {
	ID      int    `json:"id"`
	Nome    string `json:"name"`
	Raca    string `json:"race"`
	Genero  string `json:"gender"`
	Planeta Planet `json:"originPlanet"` 
}

func BuscarPersonagem(id int) (*Personagens, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(BaseURL + strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("erro na requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("personagem com id %d não encontrado", id)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("a API respondeu com status %d", resp.StatusCode)
	}

	var c Personagens
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, fmt.Errorf("erro ao ler o JSON: %w", err)
	}
	return &c, nil
}

func Exibir(c *Personagens) {
	fmt.Println()
	fmt.Println("===============")
	fmt.Printf("ID:      %d\n", c.ID)
	fmt.Printf("Nome:    %s\n", c.Nome)
	fmt.Printf("Raça:    %s\n", c.Raca)
	fmt.Printf("Gênero:  %s\n", c.Genero)
	fmt.Printf("Planeta: %s\n", c.Planeta.Name) 
	fmt.Println("===============")
	fmt.Println()
}