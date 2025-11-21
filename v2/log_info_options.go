package log

type LogInfoOption func(*LogInfo) error

// func Type(t LogType) LogInfoOption {
// 	return func(li *LogInfo) error {
// 		li.Type = t
// 		return nil
// 	}
// }

func Format(f FontFormat) LogInfoOption {
	return func(li *LogInfo) error {
		li.Format = f
		return nil
	}
}
