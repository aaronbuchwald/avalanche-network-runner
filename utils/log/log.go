// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package logutil implements various log utilities.
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalAtomicLevel zap.AtomicLevel

func init() {
	globalAtomicLevel = zap.NewAtomicLevel()
	lcfg := zap.Config{
		Level: globalAtomicLevel,

		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},

		Encoding: "json",

		// copied from "zap.NewProductionEncoderConfig" with some updates
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},

		// Use "/dev/null" to discard all
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := lcfg.Build()
	if err != nil {
		panic(err)
	}
	_ = zap.ReplaceGlobals(logger)
}

func SetGlobalLogLevel(newLevel zapcore.Level) {
	globalAtomicLevel.SetLevel(newLevel)
}
