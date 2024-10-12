package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AlanMute/file-manager/internal/disk"
	"github.com/AlanMute/file-manager/internal/filemenu"
	"github.com/AlanMute/file-manager/internal/jsonmenu"
	"github.com/AlanMute/file-manager/internal/xmlmenu"
	"github.com/AlanMute/file-manager/internal/zipmenu"
	"github.com/inancgumus/screen"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println("\n--- Главное меню ---")
		fmt.Println("1. Информация о логических дисках")
		fmt.Println("2. Работа с файлами")
		fmt.Println("3. Работа с JSON файлами")
		fmt.Println("4. Работа с XML файлами")
		fmt.Println("5. Работа с zip архивами")
		fmt.Println("6. Выход")

		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			disk.ShowDiskInfo()
		case "2":
			filemenu.ShowMenu(scanner)
		case "3":
			jsonmenu.ShowMenu(scanner)
		case "4":
			xmlmenu.ShowMenu(scanner)
		case "5":
			zipmenu.ShowMenu(scanner)
		case "6":
			fmt.Println("Выход из программы.")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
