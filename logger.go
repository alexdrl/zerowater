package zerowater

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog"
)

type ZerologLoggerAdapter struct {
	logger zerolog.Logger
}

// Logs an error message.
func (loggerAdapter *ZerologLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	event := loggerAdapter.logger.Err(err)

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Logs an info message.
func (loggerAdapter *ZerologLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	event := loggerAdapter.logger.Info()

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Logs a debug message.
func (loggerAdapter *ZerologLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	event := loggerAdapter.logger.Debug()

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Logs a trace.
func (loggerAdapter *ZerologLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	event := loggerAdapter.logger.Trace()

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

// Creates new adapter wiht the input fields as context.
func (loggerAdapter *ZerologLoggerAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	if fields == nil {
		return loggerAdapter
	}

	subLog := loggerAdapter.logger.With()

	for i, v := range fields {
		subLog = subLog.Interface(i, v)
	}

	return &ZerologLoggerAdapter{
		logger: subLog.Logger(),
	}
}

// Gets a new zerolog adapter for use in the watermill context.
func NewZerologLoggerAdapter(logger zerolog.Logger) *ZerologLoggerAdapter {
	return &ZerologLoggerAdapter{
		logger: logger,
	}
}

func addWatermillFieldsData(event *zerolog.Event, fields watermill.LogFields) {
	for i, v := range fields {
		event.Interface(i, v)
	}
}
