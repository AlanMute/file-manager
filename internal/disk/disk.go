package disk

import (
	"fmt"

	"github.com/AlanMute/file-manager/pkg/util"
	"github.com/inancgumus/screen"
	"github.com/shirou/gopsutil/disk"
)

func ShowDiskInfo() {
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
	util.Pause()
}
