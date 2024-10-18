package main

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type TimeData struct {
	SocialTime int64 `json:"socialTime"`
	WebTime    int64 `json:"webTime"`
}

func main() {
	router := gin.Default()

	// Endpoint per ricevere i dati e inviare il report
	router.POST("/sendReport", func(c *gin.Context) {
		var timeData TimeData
		if err := c.BindJSON(&timeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		report := generateReport(timeData.SocialTime, timeData.WebTime)
		err := sendEmail("emanuele.zuffranieri@gmail.com", report)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	router.Run(":8080")
}

func generateReport(socialTime, webTime int64) string {
	socialFormatted := formatTime(socialTime)
	webFormatted := formatTime(webTime)
	return fmt.Sprintf("Daily Report:\nTime spent on Social Media: %s\nTime spent on Web: %s\n", socialFormatted, webFormatted)
}

func formatTime(ms int64) string {
	totalSeconds := ms / 1000
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func sendEmail(to, body string) error {
	e := email.NewEmail()
	e.From = "YourApp <your-email@gmail.com>"
	e.To = []string{to}
	e.Subject = "Daily Time Report"
	e.Text = []byte(body)

	// Usare SMTP con Gmail
	return e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "nedneweuropeandream@gmail.com", "Pesciolina8183@_!", "smtp.gmail.com"))
}
