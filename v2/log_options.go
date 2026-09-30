package log

type LogOption func(*Log) error

// Media sets where the message is displayed. Format is ignored when the media is set to web
func Media(m LogMedia) LogOption {
	return func(li *Log) error {
		li.media = m
		return nil
	}
}

// Prefix sets the log prefix
func Prefix(pfx string) LogOption {
	return func(li *Log) error {
		li.prefix = pfx
		return nil
	}
}
