package main

import (
	"os"
	"log"
	"strconv"
	"net/http"
	s "strings"

	"gopkg.in/telegram-bot-api.v4"
	"github.com/PuerkitoBio/goquery"
)

var (
	timeTag, strmsg, a, m string
	sum int
	lastid = 0
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("На сегодня"),
	),
)

func check(err error) {
	if err != nil {
    	log.Fatal(err)
    }
}

func checkin(errt error) {
	if errt != nil {
        panic(errt)
    }
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
    resp.Write([]byte("Hi there! I'm Working!"))
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.ListenForWebhook("/" + bot.Token)
	http.HandleFunc("/", MainHandler)
    go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		c := make(chan string)
		c1 := make(chan string)
		c2 := make(chan string)
		c3 := make(chan string)
		c4 := make(chan string)
		c5 := make(chan string)
		c6 := make(chan string)
		c7 := make(chan string)
		idea := make(chan string)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		switch update.Message.Text {
			case "/start":
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.Send(msg)
			case "На сегодня" :
				go func () {
					msge := tgbotapi.NewMessage(update.Message.Chat.ID, "Начинаю поиск")
					sm, _ := bot.Send(msge)
		        	lastid = sm.MessageID
					for {
						edmsg := tgbotapi.NewEditMessageText(update.Message.Chat.ID, lastid, "Начинаю поиск")
						bot.Send(edmsg)
						edmsg = tgbotapi.NewEditMessageText(update.Message.Chat.ID, lastid, "Начинаю поиск..")
						bot.Send(edmsg)
		        		edmsg = tgbotapi.NewEditMessageText(update.Message.Chat.ID, lastid, "Начинаю поиск...")
						bot.Send(edmsg)
						edmsg = tgbotapi.NewEditMessageText(update.Message.Chat.ID, lastid, "Начинаю поиск....")
						bot.Send(edmsg)
						m = <- idea
						if m == "done" {
							if sum > 0 {
							    strmsg += <-c + <-c1 + <- c2 + <-c3 + <-c4 + <-c5 + <-c6 + <-c7 + "\n_Средняя по погодам_ : \t\t" + "*+" + strconv.Itoa(sum/8) + "*\n" + timeTag
							} else {
								strmsg += <-c + <-c1 + <- c2 + <-c3 + <-c4 + <-c5 + <-c6 + <-c7 + "\n_Средняя по погодам_ : \t\t" + "*" + strconv.Itoa(sum/8) + "*\n" + timeTag
							}
							edmsg = tgbotapi.NewEditMessageText(update.Message.Chat.ID, lastid, strmsg)
							edmsg.ParseMode = "markdown"
							bot.Send(edmsg)
							strmsg = ""
							sum = 0
							break
						}
			    	}
				}()
				go func () {
					go func () {
						doc, err := goquery.NewDocument("https://www.msn.com/ru-ru/weather/today/%D0%91%D0%B8%D1%88%D0%BA%D0%B5%D0%BA,%D0%91%D0%B8%D1%88%D0%BA%D0%B5%D0%BA,%D0%9A%D0%B8%D1%80%D0%B3%D0%B8%D0%B7%D0%B8%D1%8F/we-city?iso=KG&el=fuYFCItsFctEWpKyC2zWbQ%3D%3D")
					    check(err)

					    doc.Find("body .curcond").Each(func(index int, item *goquery.Selection) {
					        spanTag := item.Find(".current").Text()

					        i, errt := strconv.Atoi(spanTag)
					        checkin(errt)
					        sum += i
					        a := ""
					        if i > 0{
					        	a = "*+" + spanTag + "*"
					        } else {
					        	a = "*-" + spanTag + "*"
					        }
					        c1 <- "Bishkek www.msn.com: \t\t" + a + "\n"
					    })
					}()

					doc, err := goquery.NewDocument("https://weather.rambler.ru/v-bishkeke/today/")
				    check(err)

				    doc.Find("body .weather-detailed").Each(func(index int, item *goquery.Selection) {
				        spanTag := item.Find(".weather-now__value").Text()

				        i, errt := strconv.Atoi(spanTag)
				        checkin(errt)
				        sum += i
				        a := ""
				        if i > 0{
					       	a = "*+" + spanTag + "*"
					    } else {
					       	a = "*-" + spanTag + "*"
					    }
				        c <- "Bishkek rambler.ru: \t\t" + a + "\n"
				    })
				}()

				go func () {
					go func () {
						doc, err := goquery.NewDocument("https://rp5.ru/%D0%9F%D0%BE%D0%B3%D0%BE%D0%B4%D0%B0_%D0%B2_%D0%91%D0%B8%D1%88%D0%BA%D0%B5%D0%BA%D0%B5")
					    check(err)

					    doc.Find("body #wrapper #FheaderContent #archiveString .ArchiveTemp").Each(func(index int, item *goquery.Selection) {
					        spanTag := item.Find(".t_0").Text()

					        i, errt := strconv.Atoi(spanTag[:len(spanTag) - 4])
					        checkin(errt)
					        sum += i
					        a := ""
					        if i > 0{
					        	a = "*+" + spanTag[:len(spanTag) - 4] + "*"
					        } else {
					        	a = "*-" + spanTag[:len(spanTag) - 4] + "*"
					        }
					        c3 <- "Bishkek rp5.ru: \t\t" + a + "\n"
					    })
					}()

					doc, err := goquery.NewDocument("https://pogoda.mail.ru/prognoz/bishkek/")
				    check(err)

				    doc.Find("body").Each(func(index int, item *goquery.Selection) {
				        spanTag := item.Find(".information__content__temperature").Text()

				        spanTag = s.TrimSpace(spanTag)
				        
				        i, errt := strconv.Atoi(spanTag[:len(spanTag) - 2])
				        checkin(errt)
				        sum += i
					    a := "*" + spanTag[:len(spanTag) - 2] + "*"
				        c2 <- "Bishkek pogoda.mail.ru: \t" + a + "\n"
				    })
				}()

				go func () {
					go func () {
						doc, err := goquery.NewDocument("https://www.foreca.ru/Kyrgyzstan/Bishkek")
					    check(err)

					    doc.Find("body .left .txt-xxlarge").Each(func(index int, item *goquery.Selection) {
					        spanTag := item.Find("strong").Text()

					        i, errt := strconv.Atoi(spanTag)
					        checkin(errt)
					        sum += i
						    a := "*" + spanTag + "*"
					        c5 <- "Bishkek www.foreca.ru: \t\t" + a + "\n"
					    })
					}()
					doc, err := goquery.NewDocument("https://yandex.ru/pogoda/bishkek")
				    check(err)

				    doc.Find("body .fact .fact__temp").Each(func(index int, item *goquery.Selection) {
				        spanTag := item.Find(".temp__value").Text()

				        i, errt := strconv.Atoi(spanTag)
				        checkin(errt)
				        sum += i
				        a := ""
				        if i > 0{
					       	a = "*" + spanTag + "*"
					    } else {
					       	a = "*-" + spanTag[3:] + "*"
					    }
				        c4 <- "Bishkek yandex.ru: \t\t" + a + "\n"
				    })
				}()

				go func () {
					go func () {
						doc, err := goquery.NewDocument("http://pogoda.co.il/kyrgyzstan/bishkek")
					    check(err)

					    doc.Find("body .now_block").Each(func(index int, item *goquery.Selection) {
						spanTag := item.Find("strong").Text()

					        i, errt := strconv.Atoi(spanTag[:len(spanTag) - 2])
					        checkin(errt)
					        sum += i
					        a := "*" + spanTag[:len(spanTag) - 2] + "*"
					        c7 <- "Bishkek pogoda.co.il: \t\t" + a + "\n"
					    })
					}()

					doc, err := goquery.NewDocument("http://pogoda.desko.kg/")
				    check(err)

				    doc.Find("body").Each(func(index int, item *goquery.Selection) {
				        spanTag := item.Find(".temp_title").Text()
				        timeTag = item.Find(".fct_date").Text()

				        i, errt := strconv.Atoi(spanTag[:len(spanTag) - 3])
				        checkin(errt)
				        sum += i
				        a := "*" + spanTag[:len(spanTag) - 3] + "*"
				        c6 <- "Bishkek pogoda.desko.kg: \t" + a + "\n"
				    })
				}()

				idea <- "done"

			default :
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Моя твоя не понимать\nСорян))"))
		}
	}
}
