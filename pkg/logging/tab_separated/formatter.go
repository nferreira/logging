package tab_separated

import (
	"fmt"
	"github.com/nferreira/logging/pkg/logging"
	"time"
)

type Formatter struct {
}

func New() logging.Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(
	organization logging.Organization,
	system logging.System,
	correlationId string,
	message string) string {
	return fmt.Sprintf(
		"%s | %s | %s | %s | %s | %s | %s | %s | - | - | %s | - | - | %s",
		organization.Id,
		organization.Name,
		organization.Unit,
		system.Environment,
		system.Id,
		system.Hostname,
		system.AppName,
		correlationId,
		f.timeNow(),
		message)
}

func (f *Formatter) timeNow() string {
	t := time.Now()
	timeNow := t.Format("20060102150405")
	return timeNow
}
