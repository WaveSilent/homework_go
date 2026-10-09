package methods

import "fmt"

type Email struct{}

func NewEmail() Email {
	return Email{}
}

func (e Email) SendMessages(arrNotifications []string) {
	fmt.Println("Уведомления успешно отправлены по Email")
	fmt.Println()
}
