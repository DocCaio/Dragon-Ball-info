package models

import (
	"encoding/json"
	"fmt"
	"internal/strconv"
	"net/http"
	"time"
)

const baseURL = "https://dragonball-api.com/api/characters/"

type Personagens struct {
	ID      int
	Nome    string
	Raca    string
	Genero  string
	Planeta string
}
func buscarPersonagem(id int) (*Personagens, error) {
	client := &http.Client{Timeout: 10 * time.Second}
 
	resp, err := client.Get(baseURL + strconv.Itoa(id))
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



func exibir(c *Personagens) {
	fmt.Println()
	fmt.Println("===============")
	fmt.Println("ID: %\n",  c.ID)
	fmt.Printf("Nome:        %s\n", c.Nome)
	fmt.Printf("Raça:        %s\n", c.Raca)
	fmt.Printf("Gênero:      %s\n", c.Genero)

	fmt.Println("================")
	fmt.Println()
	
}