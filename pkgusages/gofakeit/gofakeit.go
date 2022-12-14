package gofakeit

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type PendingUserNotifications map[int][]*Notification
type ProcessHandler func()
type Notification struct {
	Content string
	UserId  int
}
type RecurringProcess struct {
	name      string
	interval  time.Duration
	startTime time.Time
	handler   func()
	stop      chan struct{}
}

func trigger() {
	pendingNotificationsProcess()
}

func pendingNotificationsProcess() {
	process := &RecurringProcess{}
	notifications := PendingUserNotifications{}
	handler := func() {
		collectNewUsersNotifications(notifications)
		handlePendingUsersNotifications(notifications, sendUserBatchNotificationsEmail, process)
	}
	interval := 10 * time.Second
	startTime := time.Now().Add(30 * time.Second)
	process = createRecurringProcess("Pending User Notifications", handler, interval, startTime)

	<-process.stop
}

func sendUserBatchNotificationsEmail(userId int, notifications []*Notification) {
	fmt.Printf("Sending email to user with userId %d for pending notifications %v\n", userId, notifications)
}

func handlePendingUsersNotifications(pendingNotifications PendingUserNotifications, handler func(userId int, notifications []*Notification), process *RecurringProcess) {
	userNotificationCount := 0
	for userId, notifications := range pendingNotifications {
		userNotificationCount++
		handler(userId, notifications)
		delete(pendingNotifications, userId)
	}

	if userNotificationCount == 0 {
		process.Cancel()
	}
}

func collectNewUsersNotifications(notifications PendingUserNotifications) {
	randomNotifications := getRandomNotifications()
	if len(randomNotifications) > 0 {
		notifications[randomNotifications[0].UserId] = randomNotifications
		log.Println(randomNotifications[0].Content)
	}
}

func getRandomNotifications() (notifications []*Notification) {
	rand.Seed(time.Now().UnixNano())
	userId := rand.Intn(100-10+1) + 10
	numOfNotifications := rand.Intn(5-0+1) + 0
	fmt.Printf("numOfNotifications %v\n", numOfNotifications)
	for i := 0; i < numOfNotifications; i++ {
		notifications = append(notifications, &Notification{Content: gofakeit.Paragraph(1, 2, 10, " "), UserId: userId})
	}

	return
}

func createRecurringProcess(name string, handler ProcessHandler, interval time.Duration, startTime time.Time) *RecurringProcess {
	process := &RecurringProcess{
		name:      name,
		interval:  interval,
		startTime: startTime,
		handler:   handler,
		stop:      make(chan struct{}),
	}

	go process.Start()

	return process
}

func (p *RecurringProcess) Start() {
	startTicker := &time.Timer{}
	ticker := &time.Ticker{C: nil}
	defer func() { ticker.Stop() }()

	if p.startTime.Before(time.Now()) {
		p.startTime = time.Now()
	}
	startTicker = time.NewTimer(time.Until(p.startTime))

	for {
		select {
		case <-startTicker.C:
			ticker = time.NewTicker(p.interval)
			fmt.Println("Starting recurring process")
			p.handler()
		case <-ticker.C:
			fmt.Println("Next run")
			p.handler()
		case <-p.stop:
			fmt.Println("Stoping recurring process")
			return
		}
	}
}

func (p *RecurringProcess) Cancel() {
	close(p.stop)
}
