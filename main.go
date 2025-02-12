package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()

	// Carrega o arquivo JSON da conta de serviço
	b, err := os.ReadFile("service_account.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de credenciais: %v", err)
	}

	// Cria a configuração JWT com o escopo do Sheets
	jwtConfig, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Erro ao criar JWTConfig: %v", err)
	}

	// Cria o cliente HTTP autenticado
	client := jwtConfig.Client(ctx)

	// Cria o serviço do Google Sheets
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Erro ao criar o serviço do Sheets: %v", err)
	}

	// ID da sua planilha
	spreadsheetId := "abc123"

	// Nome da aba, baseado na data atual (formato YYYY-MM-DD)
	today := time.Now().Format("2006-01-02")
	fmt.Println("Utilizando a aba:", today)

	// Intervalo para atualizar os dados (por exemplo, começando em A1)
	writeRange := today + "!A1"

	// Cria um ValueRange com os dados que você quer inserir/substituir
	var vr sheets.ValueRange
	vr.Values = [][]interface{}{
		{"NOME", "DATA", "LEAD", "TELEFONE", "E-MAIL"},
		{"João Silva", today, "Sim", "123456789", "joao.silva@example.com"},
		{"Maria Souza", today, "Não", "987654321", "maria.souza@example.com"},
	}

	// Se a aba já existir, substitui os dados usando Update
	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		log.Fatalf("Erro ao atualizar a aba existente: %v", err)
	}

	fmt.Println("Dados substituídos com sucesso na aba:", today)
}
