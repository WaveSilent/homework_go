package methods

import "fmt"

type Push struct{}

func NewPush() Push {
	return Push{}
}

func (p Push) SendMessages(arrNotifications []string) {
	fmt.Println("Уведомления успешно отправлены по push-уведомлению")
	fmt.Println()
}
