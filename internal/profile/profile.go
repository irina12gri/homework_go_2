package profile

import (
	"strconv"
	"strings"
)

// Блок profile — объединяющая задача.
// Здесь нужно закрепить:
// string normalization, int, bool, if,
// форматирование строки и простую бизнес-логику.

// BuildUserCard собирает карточку пользователя.
//
// TODO: верните строку "name=<name> age=<age> group=<group> status=<status>". Имя очистите по краям, приведите к нижнему регистру и сделайте первый символ заглавным; пустое имя замените на "Unknown". Возраст от 18 — adult, иначе minor; active=true — active, иначе inactive.
func BuildUserCard(name string, age int, active bool) string {
	group := "minor"
	if age >= 18 {
		group = "adult"
	}
	status := "inactive"
	if active == true {
		status = "active"
	}
	nameCorrected1 := strings.TrimSpace(name)
	nameCorrected2 := strings.Title(strings.ToLower(nameCorrected1))
	if nameCorrected2 == "" {
		nameCorrected2 = "Unknown"
	}
	return "name=" + nameCorrected2 + " age=" + strconv.Itoa(age) + " group=" + group + " status=" + status
}
