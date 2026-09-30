package log

type LogInfoOption func(*LogInfo) error

//	func Type(t LogType) LogInfoOption {
//		return func(li *LogInfo) error {
//			li.Type = t
//			return nil
//		}
//	}
//

// // Media sets where the message is displayed. Format is ignored when the media is set to web
// func Media(m LogMedia) LogInfoOption {
// 	return func(li *LogInfo) error {
// 		li.Media = m
// 		return nil
// 	}
// }

// Format sets the message format. Format is ignored when the media is set to web
func Format(f FontFormat) LogInfoOption {
	return func(li *LogInfo) error {
		li.Format = f
		return nil
	}
}
