package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"text/template"
)

var phabricator = Phabricator{
	Host: os.Getenv("PHABRICATOR_HOST"),
	User: os.Getenv("PHABRICATOR_USER"),
	Cert: os.Getenv("PHABRICATOR_CERT"),
}

var slack = Slack{
	Token:    os.Getenv("SLACK_TOKEN"),
	Username: "Phabricator",
}

func main() {
	if err := phabricator.Connect(); err != nil {
		log.Fatalln("Failed to connect Phabricator: ", err.Error())
	}

	channel := os.Getenv("SLACK_CHANNEL")

	var t = template.Must(template.New("message").Parse(`<{{ .Uri }}|{{ .Name }}> {{ .Text }}`))

	http.HandleFunc("/story", func(w http.ResponseWriter, r *http.Request) {
		story := r.FormValue("storyID")
		text := r.FormValue("storyText")
		author := r.FormValue("storyAuthorPHID")
		phid := r.FormValue("storyData[objectPHID]")
		log.Println("New story:", story, author, phid, text)

		if phobj, err := phabricator.PhidQuery(phid); phobj != nil {
			var msg bytes.Buffer
			t.Execute(&msg, struct{ Uri, Name, Text string }{phobj["uri"], phobj["name"], text})
			slack.postMessage(channel, msg.String())
		} else {
			log.Println("Error:", err.Error())
			slack.postMessage(channel, text)
		}
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
