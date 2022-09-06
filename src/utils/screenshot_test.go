package utils

import "testing"

func TestScreenshots(t *testing.T) {
	AutoScreenshot(`http://47.100.209.79//`, "#app")
}
