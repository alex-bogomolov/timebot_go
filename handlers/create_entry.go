package handlers

import (
	"database/sql"
	"github.com/alex-bogomolov/timebot_go/models"
	"github.com/alex-bogomolov/timebot_go/sender"
	"github.com/nlopes/slack"
	"regexp"
	"time"
	"strings"
	"fmt"
)

const batchNewEntryStringRegexp = "^(?:\\s*.* \\d?\\d:\\d\\d .+\\s*)*$"
const newEntryStringRegexp = "^ *(.*) (\\d?\\d:[0-5]\\d) ([^\\s](?:.|\\s)*[^\\s])$"

func handleNewEntry(message *slack.Msg) {
	messageText := regexp.MustCompile(batchNewEntryStringRegexp).FindString(message.Text)

	entryStrings := strings.Split(messageText, "\n")

	entriesCreated := 0

	for _, entryString := range entryStrings {
		entryString := strings.TrimSpace(entryString)

		if matched, err := regexp.MatchString(newEntryStringRegexp, entryString); err == nil && !matched {
			continue
		}

		err := createEntry(entryString, message.User)

		if err != nil {
			handleError(message.User, err)
			continue
		}

		entriesCreated++
	}

	if entriesCreated == 0 {
		sender.SendMessage(message.User, "No time entries were created.")
	} else if entriesCreated == 1 {
		sender.SendMessage(message.User, "One time entry was created.")
	} else {
		sender.SendMessage(message.User, fmt.Sprintf("%d time entries were created.", entriesCreated))
	}
}

func createEntry(entryString, userId string) error {
	newEntryRegexp := regexp.MustCompile(newEntryStringRegexp)
	matches := newEntryRegexp.FindStringSubmatch(entryString)

	projectName := matches[1]
	entryTime := matches[2]
	minutes, err := parseTime(entryTime)

	if err != nil {
		return err
	}

	details := matches[3]
	user, err := models.FindUser(userId)

	if err != nil {
		return err
	}

	project, err := models.FindProjectByNameOrAlias(projectName)

	if err != nil {
		return err
	}

	timeEntry := models.TimeEntry{
		UserId:    user.ID,
		Date:      models.NewDate(time.Now()),
		ProjectId: project.ID,
		Details:   sql.NullString{details, true},
		Minutes:   minutes,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Time:      entryTime,
	}

	err = timeEntry.Create()

	if err != nil {
		return err
	}

	return nil
}
