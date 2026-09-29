# 🐉 Dragon Ball Info

Aplicação CLI desenvolvida em **Go (Golang)** que consome a [Dragon Ball API](https://dragonball-api.com/) para buscar e exibir informações detalhadas sobre os personagens do universo de Dragon Ball através de seus IDs.

---

## 🛠️ Tecnologias Utilizadas

* **[Go (Golang)](https://golang.org/)** (versão `1.27.1`): Linguagem principal utilizada para construir a lógica, requisições HTTP e manipulação de dados.
* **Pacotes Nativos do Go**:
  * `net/http`: Para realizar requisições à API REST externa.
  * `encoding/json`: Para desserializar as respostas JSON da API nas estruturas do Go.
  * `bufio` & `os`: Para leitura de entrada interativa via terminal (`stdin`).
  * `strconv` & `strings`: Para conversão de tipos (string para inteiro) e formatação de texto.

---

## ⚙️ Como Funciona

O projeto é dividido em dois pacotes principais:
1. **`main` (`main.go`)**: Gerencia o loop de interação no terminal, solicitando o ID do usuário, validando a entrada e chamando as funções do modelo.
2. **`models` (`models/`)**: Contém as estruturas de dados (`structs`) que mapeiam o JSON da API (incluindo o personagem e seu planeta de origem) e a lógica de requisição HTTP (`BuscarPersonagem`) junto com a formatação de exibição (`Exibir`).

> **⚠️ Nota sobre o idioma do Planeta:**
> 
> Você pode notar que a descrição/nome do planeta de origem às vezes aparece em **espanhol**. Decidi manter essa opção e exibi-la assim porque esse é o dado nativo retornado diretamente pela API oficial do Dragon Ball. Como os dados vêm dessa forma da fonte original, não havia alterações que eu pudesse fazer via código sem recorrer a tradutores externos, então optei por exibi-los fielmente como a API os entrega!

---

## 🚀 Como Clonar e Rodar o Projeto

### Pré-requisitos
Certifique-se de ter o **Go** instalado em sua máquina (versão compatível com `1.27.1` ou superior). Você pode verificar com:
```bash
go version
```

### Passo a passo

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/SEU-USUARIO/Dragon-Ball-info.git
   ```

2. **Entre na pasta do projeto:**
   ```bash
   cd Dragon-Ball-info
   ```

3. **Inicie o projeto (caso precise inicializar o módulo localmente):**
   ```bash
   go mod tidy
   ```

4. **Execute a aplicação:**
   ```bash
   go run main.go
   ```

5. **Interaja com o terminal:**
   * Digite o ID numérico positivo de um personagem (ex: `1` para o Goku) e pressione `Enter`.
   * Para encerrar o programa a qualquer momento, digite `sair` e pressione `Enter`.

---

## 📄 Licença

Este projeto é de uso livre para estudos e aprendizado. Sinta-se à vontade para contribuir!
