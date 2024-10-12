package filemenu

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlanMute/file-manager/pkg/util"
	"github.com/inancgumus/screen"
)

func ShowMenu(scanner *bufio.Scanner) {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println("--- Работа с файлами ---")
		fmt.Println("1. Создать файл")
		fmt.Println("2. Записать в файл строку")
		fmt.Println("3. Прочитать файл")
		fmt.Println("4. Удалить файл")
		fmt.Println("5. Назад в главное меню")

		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			createFile(scanner)
		case "2":
			writeToFile(scanner)
		case "3":
			readFile(scanner)
		case "4":
			deleteFile(scanner)
		case "5":
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func createFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Создание файла ---")
	fmt.Print("Введите имя файла для создания: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		util.Pause()
		return
	}
	defer file.Close()

	fmt.Println("Файл создан по пути:", fullPath)
	util.Pause()
}

func writeToFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Запись строки в файл ---")
	fmt.Print("Введите имя файла для записи: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		util.Pause()
		return
	}
	defer file.Close()

	fmt.Print("Введите строку для записи: ")
	scanner.Scan()
	text := scanner.Text()

	if _, err := file.WriteString(text + "\n"); err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	} else {
		fmt.Println("Строка записана в файл по пути:", fullPath)
	}
	util.Pause()
}

func readFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Чтение файла ---")
	fmt.Print("Введите имя файла для чтения: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
		util.Pause()
		return
	}

	fmt.Println("Содержимое файла по пути:", fullPath)
	fmt.Println(string(data))
	util.Pause()
}

func deleteFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление файла ---")
	fmt.Print("Введите имя файла для удаления: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
	} else {
		fmt.Println("Файл успешно удалён по пути:", fullPath)
	}
	util.Pause()
}
