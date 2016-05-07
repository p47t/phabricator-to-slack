package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/kardianos/service"
	"github.com/yinghau76/phabricator-to-slack"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	var phabricator = ph2slack.Phabricator{
		Host: os.Getenv("PHABRICATOR_HOST"),
		User: os.Getenv("PHABRICATOR_USER"),
		Cert: os.Getenv("PHABRICATOR_CERT"),
	}

	var slack = ph2slack.Slack{
		Token:    os.Getenv("SLACK_TOKEN"),
		Username: "Phabricator",
	}

	if err := phabricator.Connect(); err != nil {
		logger.Error("Failed to connect Phabricator: ", err.Error())
		return
	}

	channel := os.Getenv("SLACK_CHANNEL")

	var t = template.Must(template.New("message").Parse(`<{{ .Uri }}|{{ .Name }}> {{ .Text }}`))

	http.HandleFunc("/story", func(w http.ResponseWriter, r *http.Request) {
		// Reconnect everytime to refresh token
		phabricator.Connect()

		story := r.FormValue("storyID")
		text := r.FormValue("storyText")
		author := r.FormValue("storyAuthorPHID")
		phid := r.FormValue("storyData[objectPHID]")
		logger.Info("New story:", story, author, phid, text)

		if phobj, err := phabricator.PhidQuery(phid); phobj != nil {
			var msg bytes.Buffer
			t.Execute(&msg, struct{ Uri, Name, Text string }{phobj["uri"], phobj["name"], text})
			slack.PostMessage(channel, msg.String())
		} else {
			logger.Error("Error:", err.Error())
			slack.PostMessage(channel, text)
		}
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "phabricator-to-slack",
		DisplayName: "phabricator-to-slack",
		Description: "Passing Phabricator notifications to Slack",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		verb := args[0]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				log.Fatalln("Failed to install:", err)
			}
			log.Printf("Service \"%s\" installed.\n", svcConfig.DisplayName)
		case "uninstall":
			err = s.Uninstall()
			if err != nil {
				log.Fatalln("Failed to uninstall:", err)
			}
			log.Printf("Service \"%s\" uninstalled.\n", svcConfig.DisplayName)
		}
		return
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
