package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/inancgumus/screen"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	mainMenu()
}

func mainMenu() {
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
			showDiskInfo()
		case "2":
			fileMenu(scanner)
		case "3":
			jsonMenu(scanner)
		case "4":
			xmlMenu(scanner)
		case "5":
			zipMenu(scanner)
		case "6":
			fmt.Println("Выход из программы.")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func showDiskInfo() {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Информация о дисках ---")

	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("Ошибка при получении информации о дисках: %v\n", err)
		return
	}

	for _, partition := range partitions {
		fmt.Printf("Имя диска: %s\n", partition.Device)
		fmt.Printf("Тип файловой системы: %s\n", partition.Fstype)

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Ошибка при получении информации о диске %s: %v\n", partition.Device, err)
			continue
		}
		fmt.Printf("Общий размер: %v байт\n", usage.Total)
		fmt.Printf("Свободно: %v байт\n", usage.Free)
		fmt.Printf("Использовано: %.2f%%\n", usage.UsedPercent)
		fmt.Println("----------------------")
	}
	pause()
}

func fileMenu(scanner *bufio.Scanner) {
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

func getDocumentsPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	documentsPath := filepath.Join(homeDir, "Documents")
	return documentsPath, nil
}

func createFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Создание файла ---")
	fmt.Print("Введите имя файла для создания: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		pause()
		return
	}
	defer file.Close()

	fmt.Println("Файл создан по пути:", fullPath)
	pause()
}

func writeToFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Запись строки в файл ---")
	fmt.Print("Введите имя файла для записи: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		pause()
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
	pause()
}

func readFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Чтение файла ---")
	fmt.Print("Введите имя файла для чтения: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
		pause()
		return
	}

	fmt.Println("Содержимое файла по пути:", fullPath)
	fmt.Println(string(data))
	pause()
}

func deleteFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление файла ---")
	fmt.Print("Введите имя файла для удаления: ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}

	fullPath := filepath.Join(documentsPath, filename)

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
	} else {
		fmt.Println("Файл успешно удалён по пути:", fullPath)
	}
	pause()
}

/// JSON

func jsonMenu(scanner *bufio.Scanner) {
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

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
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
		pause()
		return
	}

	err = os.WriteFile(fullPath, fileData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи JSON в файл:", err)
		pause()
		return
	}

	fmt.Println("JSON файл создан по пути:", fullPath)
	pause()
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

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
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
		pause()
		return
	}

	err = os.WriteFile(fullPath, fileData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи JSON в файл:", err)
		pause()
		return
	}

	fmt.Println("Объект сериализован в JSON и записан по пути:", fullPath)
	pause()
}

func readJsonFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Чтение JSON файла ---")
	fmt.Print("Введите имя JSON файла для чтения (без .json): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".json")

	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
		pause()
		return
	}

	fmt.Println("Содержимое JSON файла по пути:", fullPath)
	fmt.Println(string(data))
	pause()
}

func deleteJsonFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление JSON файла ---")
	fmt.Print("Введите имя JSON файла для удаления (без .json): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".json")

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
	} else {
		fmt.Println("JSON файл успешно удалён по пути:", fullPath)
	}
	pause()
}

// XML

func xmlMenu(scanner *bufio.Scanner) {
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

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
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
		pause()
		return
	}

	fmt.Println("XML файл создан по пути:", fullPath)
	pause()
}

func readXmlFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Чтение XML файла ---")
	fmt.Print("Введите имя XML файла для чтения (без .xml): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".xml")

	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
		pause()
		return
	}

	fmt.Println("Содержимое XML файла по пути:", fullPath)
	fmt.Println(string(data))
	pause()
}

func deleteXmlFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление XML файла ---")
	fmt.Print("Введите имя XML файла для удаления (без .xml): ")
	scanner.Scan()
	filename := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	fullPath := filepath.Join(documentsPath, filename+".xml")

	err = os.Remove(fullPath)
	if err != nil {
		fmt.Println("Данного файла не существует")
	} else {
		fmt.Println("XML файл успешно удалён по пути:", fullPath)
	}
	pause()
}

// ZIP

func zipMenu(scanner *bufio.Scanner) {
	for {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println("--- Работа с zip архивами ---")
		fmt.Println("1. Создать zip архив")
		fmt.Println("2. Добавить файл в zip архив")
		fmt.Println("3. Разархивировать файл")
		fmt.Println("4. Удалить архив")
		fmt.Println("5. Назад в главное меню")

		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			createZipArchive(scanner)
		case "2":
			addFileToZip(scanner)
		case "3":
			extractZip(scanner)
		case "4":
			deleteZipAndFile(scanner)
		case "5":
			return // Возврат в главное меню
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func createZipArchive(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Создание zip архива ---")
	fmt.Print("Введите имя архива для создания (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	zipFile, err := os.Create(archivePath)
	if err != nil {
		fmt.Println("Ошибка при создании архива:", err)
		pause()
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fmt.Println("Архив создан по пути:", archivePath)
	pause()
}

func addFileToZip(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Добавление файла в zip архив ---")
	fmt.Print("Введите имя архива (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	zipFile, err := os.OpenFile(archivePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии архива:", err)
		pause()
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fmt.Print("Введите имя файла для добавления в архив (полный путь или в папке 'Документы'): ")
	scanner.Scan()
	fileName := scanner.Text()
	filePath := filepath.Join(documentsPath, fileName)

	fileToZip, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла для архивации:", err)
		pause()
		return
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		fmt.Println("Ошибка при получении информации о файле:", err)
		pause()
		return
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println("Ошибка при создании заголовка для архива:", err)
		pause()
		return
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		fmt.Println("Ошибка при создании записи в архиве:", err)
		pause()
		return
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		fmt.Println("Ошибка при копировании файла в архив:", err)
		pause()
		return
	}

	fmt.Println("Файл добавлен в архив:", archivePath)
	pause()
}

func extractZip(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Разархивирование файла из архива ---")
	fmt.Print("Введите имя архива для разархивирования (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	zipReader, err := zip.OpenReader(archivePath)
	if err != nil {
		fmt.Println("Ошибка при открытии архива:", err)
		pause()
		return
	}
	defer zipReader.Close()

	fmt.Println("Файлы в архиве:")
	for _, file := range zipReader.File {
		fmt.Println(file.Name)
	}

	fmt.Print("\nВведите имя файла для разархивирования: ")
	scanner.Scan()
	fileName := scanner.Text()

	for _, file := range zipReader.File {
		if file.Name == fileName {
			archivedFile, err := file.Open()
			if err != nil {
				fmt.Println("Ошибка при открытии файла из архива:", err)
				pause()
				return
			}
			defer archivedFile.Close()

			content, err := io.ReadAll(archivedFile)
			if err != nil {
				fmt.Println("Ошибка при чтении файла из архива:", err)
				pause()
				return
			}

			extractedFilePath := filepath.Join(documentsPath, fileName)
			err = os.WriteFile(extractedFilePath, content, 0644)
			if err != nil {
				fmt.Println("Ошибка при сохранении файла:", err)
				pause()
				return
			}

			fmt.Println("Файл разархивирован и сохранен по пути:", extractedFilePath)
			fmt.Println("Содержимое файла", fileName, ":")
			fmt.Println(string(content))
			pause()
			return
		}
	}

	fmt.Println("Файл", fileName, "не найден в архиве.")
	pause()
}

func deleteZipAndFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление файла и архива ---")
	fmt.Print("Введите имя архива для удаления (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := getDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	err = os.Remove(archivePath)
	if err != nil {
		fmt.Println("Ошибка при удалении архива:", err)
		pause()
		return
	}

	pause()
}

func pause() {
	fmt.Println("\nНажмите Enter для продолжения...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
