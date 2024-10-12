package jsonmenu

import (
	"bufio"
	"encoding/json"
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

		fmt.Println("--- Работа с JSON файлами ---")
		fmt.Println("1. Создать JSON файл")
		fmt.Println("2. Создать объект и сериализовать в JSON")
		fmt.Println("3. Прочитать JSON файл")
		fmt.Println("4. Удалить JSON файл")
		fmt.Println("5. Назад в главное меню")

		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			createJsonFile(scanner)
		case "2":
			serializeToJson(scanner)
		case "3":
			readJsonFile(scanner)
		case "4":
			deleteJsonFile(scanner)
		case "5":
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func createJsonFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Создание JSON файла ---")
	fmt.Print("Введите имя JSON файла для создания (без .json): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".json")

	data := make(map[string]string)
	for {
		fmt.Print("Введите ключ (или оставьте пустым для завершения): ")
		scanner.Scan()
		key := scanner.Text()
		if key == "" {
			break
		}

		fmt.Print("Введите значение для ключа ", key, ": ")
		scanner.Scan()
		value := scanner.Text()

		data[key] = value
	}

	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сериализации данных в JSON:", err)
		util.Pause()
		return
	}

	err = os.WriteFile(fullPath, fileData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи JSON в файл:", err)
		util.Pause()
		return
	}

	fmt.Println("JSON файл создан по пути:", fullPath)
	util.Pause()
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func serializeToJson(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Сериализация объекта в JSON ---")
	fmt.Print("Введите имя JSON файла для создания (без .json): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".json")

	var person Person
	fmt.Print("Введите имя: ")
	scanner.Scan()
	person.Name = scanner.Text()

	fmt.Print("Введите возраст: ")
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &person.Age)

	fmt.Print("Введите email: ")
	scanner.Scan()
	person.Email = scanner.Text()

	fileData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сериализации объекта в JSON:", err)
		util.Pause()
		return
	}

	err = os.WriteFile(fullPath, fileData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи JSON в файл:", err)
		util.Pause()
		return
	}

	fmt.Println("Объект сериализован в JSON и записан по пути:", fullPath)
	util.Pause()
}

func readJsonFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Чтение JSON файла ---")
	fmt.Print("Введите имя JSON файла для чтения (без .json): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".json")

	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
		util.Pause()
		return
	}

	fmt.Println("Содержимое JSON файла по пути:", fullPath)
	fmt.Println(string(data))
	util.Pause()
}

func deleteJsonFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление JSON файла ---")
	fmt.Print("Введите имя JSON файла для удаления (без .json): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".json")

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
	} else {
		fmt.Println("JSON файл успешно удалён по пути:", fullPath)
	}
	util.Pause()
}
