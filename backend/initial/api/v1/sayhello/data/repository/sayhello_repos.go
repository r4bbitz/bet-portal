package repository

import (
	"github.com/r4bbitz/SayHello/api/v1/sayhello/model"
	"time"
)

// PromotionRepository is the interface
type SayHelloRepository interface {
	Get(timeInput string) (model.Message, error)
}
type sayHelloRepository struct {
	model.Message
}
func NewSayHelloRepository() SayHelloRepository {
	return &sayHelloRepository{
	}
}
func (repo *sayHelloRepository)  Get(timeInput string) (model.Message, error) {
	location := time.FixedZone("UTC +07:00", +7*60*60)
	timeLayOut := "15:04"
	now := time.Now().In(location)
	yourTime := ""

	if timeInput == "" {
		yourTime = now.Format(timeLayOut)

	} else {
		yourTime = timeInput
	}
	yourTimeParsed, _ := time.ParseInLocation(timeLayOut, yourTime, location)
	wordsToSayConfig := []struct {
		startTime string
		endTime   string
		words     string
	}{
		{"04:59", "11:59", "Good Morning"},
		{"11:59", "17:59", "Good Afternoon"},
		{"17:59", "20:59", "Good Evening"},
		{"20:59", "23:59", "Good Night"},
		{"23:59", "04:59", "Good Night"},
	}
	say := ""
	for _, wordsConfig := range wordsToSayConfig {

		startTime, _ := time.ParseInLocation(timeLayOut, wordsConfig.startTime, location)
		endTime, _ := time.ParseInLocation(timeLayOut, wordsConfig.endTime, location)
		if inTimeSpan(startTime, endTime, yourTimeParsed) {
			say = wordsConfig.words
			break
		}
	}
	data := model.Message{}
	data.Message=say
	data.YourTime=yourTime
	return data,nil
}


func inTimeSpan(start, end, check time.Time) bool {
	if check.After(end) {
		return false
	}
	if end.After(start) {
		return check.After(start)
	}
	return check.Before(start)
}
