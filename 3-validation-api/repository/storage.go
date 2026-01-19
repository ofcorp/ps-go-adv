package repository

import (
	"encoding/json"
	"log"
	"os"
)

type Item struct {
	Email string
	Hash	string
}

type Storage struct {
	Items []Item
}

func NewStorage() *Storage {
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Println("Файл не найден, создаем новый")
		return &Storage{
			Items: make([]Item, 0),
		}
	}
	var items []Item
	err = json.Unmarshal(file, &items)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		return &Storage{
			Items: make([]Item, 0),
		}
	}
	return &Storage{
		Items: items,
	}
}

func (storage *Storage) saveData() error {
	data, err := json.Marshal(storage.Items)
	if err != nil {
		log.Println("Ошибка при сериализации данных:", err)
		return err
	}
	return storage.write(data)
}

func (storage *Storage) write(content []byte) error {
	file, err := os.Create("data.json")
	if err != nil {
		log.Println("Ошибка при создании файла:", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		log.Println("Ошибка при записи в файл:", err)
		return err
	}
	log.Println("Файл успешно записан")
	return nil
}

func (storage *Storage) AddItem(email string, hash string) error {
	storage.Items = append(storage.Items, Item{
		Email: email,
		Hash:  hash,
	})
	return storage.saveData()
}

func (storage *Storage) VerifyHash(hash string) bool {
	for i, item := range storage.Items {
		if item.Hash == hash {
			storage.Items = append(storage.Items[:i], storage.Items[i+1:]...)
			err := storage.saveData()
			if err != nil {
				log.Println("Ошибка при сохранении данных:", err)
			}
			return true
		}
	}
	return false
}
