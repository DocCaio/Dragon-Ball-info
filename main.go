package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Dragon-Ball-info/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Digite o ID do personagem (ou 'sair'):")
		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		if strings.EqualFold(entrada, "sair") {
			fmt.Println("Até mais!")
			return
		}

		id, err := strconv.Atoi(entrada)
		if err != nil || id <= 0 {
			fmt.Println("ID inválido. Digite um número inteiro positivo.")
			continue
		}


		personagem, err := models.BuscarPersonagem(id)
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		
		models.Exibir(personagem)
	}
}