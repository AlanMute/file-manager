package zipmenu

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/AlanMute/file-manager/pkg/util"
	"github.com/inancgumus/screen"
)

func ShowMenu(scanner *bufio.Scanner) {
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

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	zipFile, err := os.Create(archivePath)
	if err != nil {
		fmt.Println("Ошибка при создании архива:", err)
		util.Pause()
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fmt.Println("Архив создан по пути:", archivePath)
	util.Pause()
}

func addFileToZip(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Добавление файла в zip архив ---")
	fmt.Print("Введите имя архива (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	zipFile, err := os.OpenFile(archivePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии архива:", err)
		util.Pause()
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
		util.Pause()
		return
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		fmt.Println("Ошибка при получении информации о файле:", err)
		util.Pause()
		return
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println("Ошибка при создании заголовка для архива:", err)
		util.Pause()
		return
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		fmt.Println("Ошибка при создании записи в архиве:", err)
		util.Pause()
		return
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		fmt.Println("Ошибка при копировании файла в архив:", err)
		util.Pause()
		return
	}

	fmt.Println("Файл добавлен в архив:", archivePath)
	util.Pause()
}

func extractZip(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Разархивирование файла из архива ---")
	fmt.Print("Введите имя архива для разархивирования (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	zipReader, err := zip.OpenReader(archivePath)
	if err != nil {
		fmt.Println("Ошибка при открытии архива:", err)
		util.Pause()
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
				util.Pause()
				return
			}
			defer archivedFile.Close()

			content, err := io.ReadAll(archivedFile)
			if err != nil {
				fmt.Println("Ошибка при чтении файла из архива:", err)
				util.Pause()
				return
			}

			extractedFilePath := filepath.Join(documentsPath, fileName)
			err = os.WriteFile(extractedFilePath, content, 0644)
			if err != nil {
				fmt.Println("Ошибка при сохранении файла:", err)
				util.Pause()
				return
			}

			fmt.Println("Файл разархивирован и сохранен по пути:", extractedFilePath)
			fmt.Println("Содержимое файла", fileName, ":")
			fmt.Println(string(content))
			util.Pause()
			return
		}
	}

	fmt.Println("Файл", fileName, "не найден в архиве.")
	util.Pause()
}

func deleteZipAndFile(scanner *bufio.Scanner) {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("--- Удаление файла и архива ---")
	fmt.Print("Введите имя архива для удаления (без .zip): ")
	scanner.Scan()
	archiveName := scanner.Text()

	documentsPath, err := util.GetDocumentsPath()
	if err != nil {
		fmt.Println("Ошибка при получении пути до папки документов:", err)
		util.Pause()
		return
	}
	archivePath := filepath.Join(documentsPath, archiveName+".zip")

	err = os.Remove(archivePath)
	if err != nil {
		fmt.Println("Ошибка при удалении архива:", err)
		util.Pause()
		return
	}

	util.Pause()
}
