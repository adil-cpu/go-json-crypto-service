package generator

import "testing"

func TestGenerateJSON(t *testing.T) {
	data := GenerateJSON()

	if len(data) == 0 {
		t.Error("Ожидались данные, но карта пуста")
	}

	// проверяем, есть ли в JSON хотя бы 1 ключ-значение
	found := false
	for _, v := range data {
		if v != nil {
			found = true
			break
		}
	}
	if !found {
		t.Error("JSON сгенерирован, но не содержит значений")
	}
}
