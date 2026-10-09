package notifications

type NotificationsMethod interface {
	SendMessages(arrNotifications []string)
}

type NotificationsModule struct {
	notificationsMethod NotificationsMethod
}

func NewNotificationsModule(
	notificationsMethod NotificationsMethod,
) NotificationsModule {
	return NotificationsModule{
		notificationsMethod: notificationsMethod,
	}
}

func (n NotificationsModule) SendMessages(arrNotifications []string) {
	n.notificationsMethod.SendMessages(arrNotifications)
}
