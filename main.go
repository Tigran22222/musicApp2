package main

import "fmt"





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
   }



   
func (storage*AlbumStorage) DeleteAlbum(){
	if len(storage.Albums) == 0{
		fmt.Println("Нет альбомов для удаления")
		return
	}
	fmt.Println("Выберите альбом для удаления")
	for i, album := range storage.Albums{
		fmt.Printf("%d. %s\n", i+1, album.Name)
	}

var choice int
fmt.Println("Введите номер альбома")
_, err:= fmt.Scan(&choice)
if err != nil || choice<1 || choice > len(storage.Albums){
	fmt.Println("Ошибка. Введите корректный номер альбома")
	return
}
storage.Albums = append(storage.Albums[:choice-1], storage.Albums[choice:]...)
fmt.Println("Альбом успешно удален")

}

func(storage*AlbumStorage) DeleteMusicInAlbum(){
if len(storage.Albums)==0 {
	fmt.Println("Нет песен для удаления")
	return
}

fmt.Println("Выберите альбом, из которого хотите удалить песни")
for i, album := range storage.Albums{
	fmt.Printf("%d. %s\n", i+1, album.Name)
}
 
var albumChoice int
fmt.Print("Введите номер алььбома")
_, err := fmt.Scan(&albumChoice)
if err!= nil || albumChoice <1 || albumChoice >len(storage.Albums){
	fmt.Println("Ошибка. Введите корректный номер альбома")
	return
}
chosenAlbum := &storage.Albums[albumChoice-1]

if len(chosenAlbum.Musics) == 0{
	fmt.Printf("В этом альбоме нет песен для удаления")
	return
}

fmt.Println("Выберите песню для удаления")
for i, music:= range chosenAlbum.Musics{
	fmt.Printf("%d. %s\n", i+1, music.Name)
}

var musicChoice int
fmt.Print("Введите номер песни")
_, err = fmt.Scan(&musicChoice)
if err != nil || musicChoice < 1 || musicChoice > len(chosenAlbum.Musics){
	fmt.Println("Ошибка. Введите корректный номер песни")
	return
}
chosenAlbum.Musics = append (chosenAlbum.Musics[:musicChoice-1], chosenAlbum.Musics[musicChoice:]...)
fmt.Println("Песня успешно удалена")

}




   func (storage *AlbumStorage) AddAlbum(Name string, createdAt int) error {
	for _, album := range storage.Albums {
	 if album.Name == Name {
	  return fmt.Errorf("Ошибка: альбом уже существует")
	 }
	}
	newAlbum := Album{
	 Name:      Name,
	 CreatedAt: createdAt,
	 Musics:    []Music{},
	}
	storage.Albums = append(storage.Albums, newAlbum)
	return nil
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
	fmt.Println("Введите номер альбома:")
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 {
	 fmt.Println("Ошибка")
	 return
	}
   

	var musicName string

	for{
	fmt.Println("Введите название песни:")
	var musicName string
	fmt.Scan(&musicName)
	if len(musicName) < 3{
		fmt.Println("Ошибка. Название песни слишком короткое, введите хотя бы 3 символа")
	}else if len(musicName) <=100 {
		fmt.Println("Ошибка. Название песни слишкоим длинное, введите название меньше 100 символов")
	} else{
		break
	}

	}

	fmt.Println("Введите описание песни")
	var caption string
	fmt.Scan(&caption)
   
	selectedAlbum := &storage.Albums[choice-1]
	newMusic := Music{Name: musicName, Caption: caption}
	selectedAlbum.Musics = append(selectedAlbum.Musics, newMusic)
   
	fmt.Println("Песня успешно добавлена в альбом!")
   
   }
   

   func (storage *AlbumStorage) ShowAlbums() {
	if len (storage.Albums)==0{
		fmt.Println("В библиотеке пока нет альбомов")
		return
	}
	fmt.Println("Список альбомов")
	for i, album := range storage.Albums {
	 fmt.Printf("%d. %s (Год: %d)\n", i+1, album.Name, album.CreatedAt)
	}
   }
  

   func (storage *AlbumStorage) ShowMusicInChosenAlbum() {
	fmt.Println("Выберите альбом для просмотра песен")
	for i, album := range storage.Albums {
	 fmt.Printf("%d. %s\n", i+1, album.Name)
	}
   
	var choice int
	fmt.Printf("Введите номер альбома:")
	_, err:=fmt.Scan(&choice)
	if err!= nil || choice<1 || choice>len(storage.Albums){
		fmt.Println("Ошибка. Введите корректный номер альбома")
		return
	}
	//косяк
   
	chosenAlbum := storage.Albums[choice-1]
   if len(chosenAlbum.Musics)==0{
	fmt.Println("В этом альбоме пока нет песен")
   }


	fmt.Printf("Песни в альбоме: %s\n", chosenAlbum.Name)
	for i, music := range chosenAlbum.Musics {
	 fmt.Printf("%d. %s\n", i+1, music.Name,)
	}
   }
   

   func main() {
	storage := AlbumStorage{}
	storage.Albums = getTestData()
	play(storage)
   }
   
   func play(storage AlbumStorage) {
	for {
	 fmt.Println("\nМеню")
	 fmt.Println("1-Добавить альбом")
	 fmt.Println("2-Добавить песню в альбом")
	 fmt.Println("3-Показать все альбомы")
	 fmt.Println("4-Показать все песни в выбранном альбоме")
	 fmt.Println("5-Удалить альбом")
	 fmt.Println("6-Удалить песню из альбома")
	 fmt.Println("0-Выход")
   
	 var choice int
	 fmt.Println("Введите любое число из пункта меню и введите ")
	 fmt.Scan(&choice)
	 switch choice {
	 
	case 1:
	  var Name string
	  var CreatedAt int
	  fmt.Println("Введите название альбома")
	  _, err := fmt.Scan(&Name)
	  //обработка ввода пример
	  if err != nil {
	   fmt.Println("введено неверное значение попробуйте снова")
	   continue
	  }
	  
	  fmt.Println("Введите год выпуска альбома")
	  fmt.Scan(&CreatedAt)
	  if err != nil{
		fmt.Println("Ошибка: введите корректное сисло для года")
	  }
   
	  err = storage.AddAlbum(Name, CreatedAt)
	  if err != nil {
	   fmt.Println(err.Error())
	  } else {
	   fmt.Println("Новый албом добавлен")
	  }
   
	 case 2:
	  storage.AddMusic()
   
	 case 3:
	  storage.ShowAlbums()
   
	 case 4: storage.ShowMusicInChosenAlbum()

	 case 5: storage.DeleteAlbum()

	 case 6: storage.DeleteMusicInAlbum()

   
	 case 0:
	  fmt.Println("Выход из программы")
	  return
	 default:
	  fmt.Println("Некорректный ввод: Выберите пункт меню от 0 до 6")
   
	 }
	}
   }


   func getTestData() []Album {
	album := Album{
	 Name:      "yaderka",
	 CreatedAt: 123,
	 Musics: []Music{
	  {
	   Name:    "test1",
	   Caption: "hksdbjksbnfc",
	  },
	  {
	   Name:    "test2",
	   Caption: "fnjdsnvkjsd",
	  },
	 },
	}
	album2 := Album{
	 Name:      "bionikl",
	 CreatedAt: 43534,
	 Musics: []Music{
	  {
	   Name:    "rakha",
	   Caption: "hksdbjcvdcvdfksbnfc",
	  },
	  {
	   Name:    "ikarus",
	   Caption: "d df d",
	  },
	 },
	}
	return []Album{album, album2}
   }