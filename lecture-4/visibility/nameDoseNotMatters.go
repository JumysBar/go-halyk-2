package visibility

import (
	"fmt"
	"time"
)

// secretsFields структура с приватными полями
type secretsFields struct {
	isBan            bool
	recommendations  []string
	createStructTime time.Time
}

// VisibilityExample структура нашего объекта взаимодействия
type VisibilityExample struct {
	Name  string
	Age   int // Если вы большой оптимист - используйте uint
	Email string

	secretsFields
}

// NewVisibilityExample создаёт экземпляр объекта с начальными данными
func NewVisibilityExample() *VisibilityExample {
	return &VisibilityExample{
		secretsFields: secretsFields{
			isBan:            false,
			recommendations:  nil,
			createStructTime: time.Now(),
		},
	}
}

// privateMethodForSetRecommendations сервис для получения данных о рекомендациях
func (v *VisibilityExample) privateMethodForSetRecommendations() ([]string, error) {
	// это может быть обращения к сервису, который обратится к базе данных
	return []string{"Golang", "K8s", "Strawberry"}, nil
}

// PublicMethodForIsBan публичный метода, к которому мы можем обратиться за пределами данного пакета
func (v *VisibilityExample) PublicMethodForIsBan() {
	v.isBan = true
}

// PublicMethodFroRecommendationsMethode сервис для присваивания значения для рекомендаций
func (v *VisibilityExample) PublicMethodFroRecommendationsMethode() error {
	r, err := v.privateMethodForSetRecommendations()
	if err != nil {
		return fmt.Errorf("PublicMethodFroRecommendationsMethode: %w", err)
	}

	v.recommendations = r[0:2]
	return nil
}

// GetRecommendations метод для получения данных из приватного поля recommendations
func (v *VisibilityExample) GetRecommendations() []string {
	return v.recommendations
}
