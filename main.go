package main

import (
	"fmt"
)

type Album struct {
	Name      string
	CreatedAt int
	Musics    []Music
}

type Music struct {
	Name    string
	Caption string
}

type AlbumStorage struct {
	Albums []Album
	Musics []Music
}

func (storage AlbumStorage) ShowMusicInChoosenAlbum() {
	panic("unimplemented")
}

func (storage *AlbumStorage) AddAlbum(Name string, createdAt int) {
	for _, album := range storage.Albums {
		if album.Name == Name {
			fmt.Println("Ошибка: альбом уже существует")
			return
		}
	}
	newAlbum := Album{
		Name:      Name,
		CreatedAt: createdAt,
		Musics:    []Music{},
	}
	storage.Albums = append(storage.Albums, newAlbum)
	fmt.Println("Новый албом добавлен")
}

/*
func (storage*AlbumStorage) AddMusic(authorName string, musicName string, caption string){

	for i := range storage.Albums{ //Используем range, чтобы получить индексы альбомов.
		if storage.Albums[i].AuthorName == authorName{ // Проверяем, совпадает ли имя исполнителя с переданным authorName.
			newMusic := Music{Name: musicName, Caption: caption} //Создаем новую структуру Music.
			storage.Albums[i].Musics = append(storage.Albums[i].Musics, newMusic) //Добавляем песню в массив Musics альбома.
			fmt.Println("Песня добавлена в альбом")
			return
		}
	}
	fmt.Println("Альбом не найден.") //Если альбом не найден, выводим сообщение:
	}

*/

func (storage *AlbumStorage) AddMusic() {
	if len(storage.Albums) == 0 {
		fmt.Println("Ошибка, добавьте хотя бы один альбом")
	}

	fmt.Println("Выберите альбом для добавления песни")
	for i, album := range storage.Albums {
		fmt.Printf("%d. %s(Год: %d)\n", i+1, album.Name, album.CreatedAt)
	}

	var choice int
	fmt.Print("Введите номер альбома:")
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 {
		fmt.Println("Ошиька")
		return
	}

	fmt.Println("Введите название песни:")
	var musicName string
	fmt.Scan(&musicName)

	fmt.Println("Введите описание песни")
	var caption string
	fmt.Scan(&caption)

	selectedAlbum := &storage.Albums[choice-1]
	newMusic := Music{Name: musicName, Caption: caption}
	selectedAlbum.Musics = append(selectedAlbum.Musics, newMusic)

	fmt.Println("Песня успешно добавлена в альбом!")

}

func (storage *AlbumStorage) ShowAlbums() {
	for i, album := range storage.Albums {
		fmt.Printf("%d. %s (Год: %d)\n", i+1, album.Name, album.CreatedAt)
	}
}

func (storage *AlbumStorage) ShowMusicInAlbum() {
	for i, album := range storage.Albums {
		fmt.Printf("%d. %s", i+1, album.Musics)
	}
}

func (storage *AlbumStorage) ShowMusicInChosenAlbum() {

	fmt.Println("Выберите альбом для просмотра песен")
	for i, album := range storage.Albums {
		fmt.Printf("%d. %s", i+1, album.Name)
	}

	var choice int
	fmt.Printf("Введите номер альбома:")

	chosenAlbum := storage.Albums[choice-1]

	fmt.Printf("Песни в альбоме: %s\n", chosenAlbum.Name)
	for i, music := range chosenAlbum.Musics {
		fmt.Printf("%d. %s\n", i+1, music.Name)
	}
}

func main() {
	storage := AlbumStorage{}
	for {
		fmt.Println("\nМеню")
		fmt.Println("1-Добавить альбом")
		fmt.Println("2-Добавить песню в альбом")
		fmt.Println("3-Показать все альбомы")
		fmt.Println("4-Показать все пенси в альбоме")
		fmt.Println("5-Показать все песни в выбранном альбоме")
		fmt.Println("0-Выход")

		var choice int
		fmt.Println("Введите любое число из пункта меню и введите ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var Name string
			var CreatedAt int
			fmt.Print("Введите имя исполнителя")
			fmt.Scan(&Name)
			fmt.Print("Введите год выпуска альбома")
			fmt.Scan(&CreatedAt)
			storage.AddAlbum(Name, CreatedAt)

		case 2:
			storage.AddMusic()

		case 3:
			storage.ShowAlbums()

		case 4:
			storage.ShowMusicInAlbum()
		case 5:
			storage.ShowMusicInChoosenAlbum()

		case 0:
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Некорректный ввод, попробуйте снова")

		}
	}
}