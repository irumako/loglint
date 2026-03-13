package loglint

type argPos int

const (
	firstArg argPos = iota
	secondArg
	thirdArg
)

// Поддерживаемые логгеры:
// log/slog;
// go.uber.org/zap
var messageArgPosByFunction = map[string]argPos{
	// log/slog
	"log/slog.Debug":                  firstArg,
	"log/slog.DebugContext":           secondArg,
	"log/slog.Info":                   firstArg,
	"log/slog.InfoContext":            secondArg,
	"log/slog.Warn":                   firstArg,
	"log/slog.WarnContext":            secondArg,
	"log/slog.Error":                  firstArg,
	"log/slog.ErrorContext":           secondArg,
	"log/slog.LogAttrs":               thirdArg,
	"log/slog.Log":                    thirdArg,
	"(*log/slog.Logger).Debug":        firstArg,
	"(*log/slog.Logger).DebugContext": secondArg,
	"(*log/slog.Logger).Info":         firstArg,
	"(*log/slog.Logger).InfoContext":  secondArg,
	"(*log/slog.Logger).Warn":         firstArg,
	"(*log/slog.Logger).WarnContext":  secondArg,
	"(*log/slog.Logger).Error":        firstArg,
	"(*log/slog.Logger).ErrorContext": secondArg,
	"(*log/slog.Logger).Log":          thirdArg,
	"(*log/slog.Logger).LogAttrs":     thirdArg,

	// zap
	"(*go.uber.org/zap.Logger).Debug":  firstArg,
	"(*go.uber.org/zap.Logger).Info":   firstArg,
	"(*go.uber.org/zap.Logger).Warn":   firstArg,
	"(*go.uber.org/zap.Logger).Error":  firstArg,
	"(*go.uber.org/zap.Logger).DPanic": firstArg,
	"(*go.uber.org/zap.Logger).Panic":  firstArg,
	"(*go.uber.org/zap.Logger).Fatal":  firstArg,
	"(*go.uber.org/zap.Logger).Log":    secondArg,

	// zap sugared
	"(*go.uber.org/zap.SugaredLogger).Debug":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Debugf":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).Debugw":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).Info":    firstArg,
	"(*go.uber.org/zap.SugaredLogger).Infof":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Infow":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Warn":    firstArg,
	"(*go.uber.org/zap.SugaredLogger).Warnf":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Warnw":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Error":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Errorf":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).Errorw":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).DPanic":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).DPanicf": firstArg,
	"(*go.uber.org/zap.SugaredLogger).DPanicw": firstArg,
	"(*go.uber.org/zap.SugaredLogger).Panic":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Panicf":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).Panicw":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).Fatal":   firstArg,
	"(*go.uber.org/zap.SugaredLogger).Fatalf":  firstArg,
	"(*go.uber.org/zap.SugaredLogger).Fatalw":  firstArg,
}
