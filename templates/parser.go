package templates

import (
	"bytes"
	"html/template"
	"log"

	"github.com/perso-proj/database"
)

func ParseEvents(events []database.Event) string {
	if len(events) < 1 {
		return ""
	}

	eventsHtml := ""
	for _, event := range events {
		tmpl, err := template.New("name").Parse(EventTempl)
		if err != nil {
			log.Fatal("Failed to parse event template")
		}
		buf := &bytes.Buffer{}
		err = tmpl.Execute(buf, event)

		eventsHtml += buf.String()
	}
	return eventsHtml

}
