package zerowater_test

import (
	"bytes"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/alexdrl/zerowater"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestStdLogger_with(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	zerologger := zerolog.New(buf)
	cleanLogger := zerowater.NewZerologLoggerAdapter(zerologger)

	withLogFieldsLogger := cleanLogger.With(watermill.LogFields{"foo": "1"})

	for name, logger := range map[string]watermill.LoggerAdapter{"clean": cleanLogger, "with": withLogFieldsLogger} {
		logger.Error(name, nil, watermill.LogFields{"bar": "2"})
		logger.Info(name, watermill.LogFields{"bar": "2"})
		logger.Debug(name, watermill.LogFields{"bar": "2"})
		logger.Trace(name, watermill.LogFields{"bar": "2"})
	}

	cleanLoggerOut := buf.String()
	assert.Contains(t, cleanLoggerOut, `{"level":"info","bar":"2","message":"clean"}`)
	assert.Contains(t, cleanLoggerOut, `{"level":"info","foo":"1","bar":"2","message":"with"}`)
}

func TestStdLogger_withNil(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	zerologger := zerolog.New(buf)
	cleanLogger := zerowater.NewZerologLoggerAdapter(zerologger)

	withLogFieldsLogger := cleanLogger.With(nil)

	for name, logger := range map[string]watermill.LoggerAdapter{"clean": cleanLogger, "with": withLogFieldsLogger} {
		logger.Error(name, nil, watermill.LogFields{"bar": "2"})
		logger.Info(name, watermill.LogFields{"bar": "2"})
		logger.Debug(name, watermill.LogFields{"bar": "2"})
		logger.Trace(name, watermill.LogFields{"bar": "2"})
	}

	cleanLoggerOut := buf.String()
	assert.Contains(t, cleanLoggerOut, `{"level":"info","bar":"2","message":"clean"}`)
	assert.Contains(t, cleanLoggerOut, `{"level":"info","bar":"2","message":"with"}`)
}

func TestStdLoggerAdapter_field_with_space(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	zerologger := zerolog.New(buf)
	logger := zerowater.NewZerologLoggerAdapter(zerologger)

	logger.Info("foo", watermill.LogFields{"foo": `bar baz`})

	out := buf.String()
	assert.Contains(t, out, `"foo":"bar baz"`)
}
