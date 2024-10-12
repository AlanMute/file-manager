package xmlmenu

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

		fmt.Println("--- Работа с XML файлами ---")
		fmt.Println("1. Создать XML файл")
		fmt.Println("2. Прочитать XML файл")
		fmt.Println("3. Удалить XML файл")
		fmt.Println("4. Назад в главное меню")

		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			createXmlFile(scanner)
		case "2":
			readXmlFile(scanner)
		case "3":
			deleteXmlFile(scanner)
		case "4":
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func createXmlFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Создание XML файла ---")
	fmt.Print("Введите имя XML файла для создания (без .xml): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".xml")

	data := make(map[string]string)
	for {
		fmt.Print("Введите тег (или оставьте пустым для завершения): ")
		scanner.Scan()
		tag := scanner.Text()
		if tag == "" {
			break
		}

		fmt.Printf("Введите значение для тега <%s>: ", tag)
		scanner.Scan()
		value := scanner.Text()

		data[tag] = value
	}

	xmlContent := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root>\n"
	for tag, value := range data {
		xmlContent += fmt.Sprintf("  <%s>%s</%s>\n", tag, value, tag)
	}
	xmlContent += "</root>"

	err = os.WriteFile(fullPath, []byte(xmlContent), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи XML в файл:", err)
		util.Pause()
		return
	}

	fmt.Println("XML файл создан по пути:", fullPath)
	util.Pause()
}

func readXmlFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Чтение XML файла ---")
	fmt.Print("Введите имя XML файла для чтения (без .xml): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".xml")

	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
		util.Pause()
		return
	}

	fmt.Println("Содержимое XML файла по пути:", fullPath)
	fmt.Println(string(data))
	util.Pause()
}

func deleteXmlFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление XML файла ---")
	fmt.Print("Введите имя XML файла для удаления (без .xml): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".xml")

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
	} else {
		fmt.Println("XML файл успешно удалён по пути:", fullPath)
	}
	util.Pause()
}
