package gadgetFormatter_test

import (
	. "github.com/NextThingCo/logrus-gadget-formatter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/sirupsen/logrus.v1"
)

var _ = Describe("Formatter", func() {
	var formatter *TextFormatter
	var log *logrus.Logger
	var output *LogOutput

	BeforeEach(func() {
		output = new(LogOutput)
		formatter = new(TextFormatter)
		log = logrus.New()
		log.Out = output
		log.Formatter = formatter
		log.Level = logrus.DebugLevel
	})

	Describe("Formatting output", func() {
		It("should output simple logfmt message", func() {
			formatter.DisableTimestamp = true
			log.Debug("test")
			Ω(output.GetValue()).Should(Equal("level=debug msg=test\n"))
		})

		It("should output formatted message", func() {
			formatter.DisableTimestamp = true
			formatter.ForceFormatting = true
			log.Debug("test")
			Ω(output.GetValue()).Should(Equal("DEBUG test\n"))
		})
	})

	Describe("Theming support", func() {

	})
})
