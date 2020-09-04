package logging

type Formatter interface {
	Format(
		organization Organization,
		system System,
		correlationId string,
		message string) string
}
