package service

import (
	"fmt"
	"strconv"

	"firebase.google.com/go/messaging"
	"github.com/giancarlobastos/loteca-backend/client"
	"github.com/giancarlobastos/loteca-backend/view"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type NotificationService struct {
	firebaseClient *client.FirebaseClient
}

type NotificationType string

const (
	LOTTERY_EVENT NotificationType = "lottery_event"
	MATCH_EVENT   NotificationType = "match_event"
	POLL_EVENT    NotificationType = "poll_event"
)

func NewNotificationService(
	firebaseClient *client.FirebaseClient) *NotificationService {
	return &NotificationService{
		firebaseClient: firebaseClient,
	}
}

func (ns *NotificationService) NotifyMatchScore(lotteryId int, match *view.Match, timestamp int64) error {
	return ns.notifyMatchEvent(lotteryId, "Gol!!!", match, timestamp)
}

func (ns *NotificationService) NotifyMatchInterval(lotteryId int, match *view.Match, timestamp int64) error {
	return ns.notifyMatchEvent(lotteryId, "Intervalo", match, timestamp)
}

func (ns *NotificationService) NotifyMatchFinished(lotteryId int, match *view.Match, timestamp int64) error {
	return ns.notifyMatchEvent(lotteryId, "Fim de Jogo!", match, timestamp)
}

func (ns *NotificationService) NotifyPollEvent(lotteryId int) error {
	message := &messaging.Message{
		Data: data(lotteryId, POLL_EVENT),
		Topic: "loteca",
	}
	return ns.firebaseClient.SendMessage(message)
}

func (ns *NotificationService) NotifyNewLotteryEvent(lottery *view.Lottery) error {
	title := fmt.Sprintf("%s Disponível", *lottery.Name)

	if lottery.EstimatedPrize == nil {
		return ns.notifyLotteryEvent(*lottery.Id, title, "");
	}

	body := fmt.Sprintf("Prêmio Estimado em %s", toBRL(*lottery.EstimatedPrize))
	return ns.notifyLotteryEvent(*lottery.Id, title, body);
}

func (ns *NotificationService) NotifyLotteryResultEvent(lottery *view.Lottery) error {
	title := fmt.Sprintf("Resultado do %s", *lottery.Name)

	if *lottery.MainPrizeWinners == 0 {
		title = "ACUMULOU!! - " + title
	}

	mainPrizeResult := fmt.Sprintf("%d ganhadores com 14 acertos - Prêmio de %s", 
		*lottery.MainPrizeWinners, toBRL(*lottery.MainPrize))

	switch *lottery.MainPrizeWinners {
	case 0:
		mainPrizeResult = "Nenhum ganhador com 14 acertos!"
	case 1:
		mainPrizeResult = fmt.Sprintf("1 ganhador com 14 acertos - Prêmio de %s", toBRL(*lottery.MainPrize))
	}

	sidePrizeResult := fmt.Sprintf("%d ganhadores com 13 acertos - Prêmio de %s", 
		*lottery.SidePrizeWinners, toBRL(*lottery.SidePrize))

	switch *lottery.SidePrizeWinners {
	case 0:
		sidePrizeResult = "Nenhum ganhador com 13 acertos!"
	case 1:
		sidePrizeResult = fmt.Sprintf("1 ganhador com 13 acertos - Prêmio de %s", toBRL(*lottery.SidePrize))
	}

	body := fmt.Sprintf("%s\n%s", mainPrizeResult, sidePrizeResult)
	return ns.notifyLotteryEvent(*lottery.Id, title, body);
}

func (ns *NotificationService) NotifyLotteryUpdateEvent(lotteryId int) error {
	message := &messaging.Message{
		Data: data(lotteryId, LOTTERY_EVENT),
		Topic: "loteca",
	}
	return ns.firebaseClient.SendMessage(message)
}

func (ns *NotificationService) notifyMatchEvent(lotteryId int, title string, match *view.Match, timestamp int64) error {
	body := fmt.Sprintf("%s %d x %d %s", *match.HomeName, *match.HomeScore, *match.AwayScore, *match.AwayName)
	message := &messaging.Message{
		Data: notification(lotteryId, MATCH_EVENT, timestamp),
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: "loteca",
	}
	return ns.firebaseClient.SendMessage(message)
}

func (ns *NotificationService) notifyLotteryEvent(lotteryId int, title string, body string) error {
	message := &messaging.Message{
		Data: notification(lotteryId, LOTTERY_EVENT, 0),
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: "loteca",
	}
	return ns.firebaseClient.SendMessage(message)
}

func notification(lotteryId int, notificationType NotificationType, timestamp int64) map[string]string {
	return map[string]string{
		"type":      string(notificationType),
		"timestamp": strconv.Itoa(int(timestamp)),
		"lottery_id":  strconv.Itoa(lotteryId),
	}
}

func data(lotteryId int, notificationType NotificationType) map[string]string {
	return map[string]string{
		"type":      string(notificationType),
		"lottery_id":  strconv.Itoa(lotteryId),
	}
}

func toBRL(value float32) string {
	lang := language.MustParse("pt_BR")
	cur, _ := currency.FromTag(lang)
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(value, number.Scale(scale))
	p := message.NewPrinter(lang)
	return p.Sprintf("%v %v", currency.Symbol(cur), dec)
}
