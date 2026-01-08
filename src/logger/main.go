package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/dusted-go/logging/prettylog"
)

// nvim causes terminal to switch to rawline mode, which messes up line endings on Windows.
// This writer forces CRLF line endings to avoid that issue.
type ForceCrlfWriter struct {
	io.Writer
	writer io.Writer
}

func (w *ForceCrlfWriter) Write(p []byte) (n int, err error) {
	pCrlf := make([]byte, 0, len(p)+len(p)/10)
	for i := range p {
		if p[i] == '\n' {
			pCrlf = append(pCrlf, '\r')
		}
		pCrlf = append(pCrlf, p[i])
	}
	return w.writer.Write(pCrlf)
}

func PrepareLogger(isDebug bool) {
	logLevel := slog.LevelInfo

	if isDebug {
		logLevel = slog.LevelDebug
	}

	log := slog.New(prettylog.New(
		&slog.HandlerOptions{
			Level:     logLevel,
			AddSource: isDebug,
		},
		prettylog.WithDestinationWriter(&ForceCrlfWriter{
			writer: os.Stderr,
		}),
		prettylog.WithColor(),
	))

	slog.SetDefault(log)
}
