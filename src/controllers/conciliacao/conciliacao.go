package conciliacao

import (
	"fmt"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func ConciliaXLXs(w http.ResponseWriter, r *http.Request) {
	// Abre o arquivo XLSX
	fmt.Println("Chegou")
	f, err := excelize.OpenFile("Fornecedores NBS 102024.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Obtém todos os valores da primeira planilha
	rows, err := f.GetRows("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Imprime os valores
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
