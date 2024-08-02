package gui

import (
	"testing"
)

func TestMacButton_ButtonRender(t *testing.T) {
	var factory GUIFactory

	//创建 Windows 风格控件
	factory = WindowFactory{}
	button := factory.CreateButton()
	button.ButtonRender()
	textBox := factory.CreateTextBox()
	textBox.TextBoxRender()

	//创建 Mac 风格控件
	factory = MacFactory{}
	button = factory.CreateButton()
	button.ButtonRender()
	textBox = factory.CreateTextBox()
	textBox.TextBoxRender()
}
