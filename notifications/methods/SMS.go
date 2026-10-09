package methods

import "fmt"

type SMS struct{}

func NewSMS() SMS {
	return SMS{}
}

func (s SMS) SendMessages(arrNotifications []string) {
	fmt.Println("Уведомления успешно отправлены по SMS")
	fmt.Println()
}
