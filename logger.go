package zerowater

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog"
)

type ZerologLoggerAdapter struct {
	logger zerolog.Logger
}

func (loggerAdapter *ZerologLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	event := loggerAdapter.logger.Err(err)

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

func (loggerAdapter *ZerologLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	event := loggerAdapter.logger.Info()

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

func (loggerAdapter *ZerologLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	event := loggerAdapter.logger.Debug()

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

func (loggerAdapter *ZerologLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	event := loggerAdapter.logger.Trace()

	if fields != nil {
		addWatermillFieldsData(event, fields)
	}

	event.Msg(msg)
}

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

func addWatermillFieldsData(event *zerolog.Event, fields watermill.LogFields) {
	for i, v := range fields {
		event.Interface(i, v)
	}
}

func NewZerologLoggerAdapter(logger zerolog.Logger) *ZerologLoggerAdapter {
	return &ZerologLoggerAdapter{
		logger: logger,
	}
}
