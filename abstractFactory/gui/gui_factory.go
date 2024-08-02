package gui

import "github.com/marmotedu/iam/pkg/log"

type Button interface {
	ButtonRender()
}

type TextBox interface {
	TextBoxRender()
}

type WindowButton struct {
}

func (w WindowButton) ButtonRender() {
	log.Infof("window button")
}

type MacButton struct {
}

func (m MacButton) ButtonRender() {
	log.Infof("mac button")
}

type WindowTextBox struct {
}

func (w WindowTextBox) TextBoxRender() {
	log.Infof("window TextBox")
}

type MacTextBox struct{}

func (m MacTextBox) TextBoxRender() {
	log.Infof("mac TextBox")
}

// GUIFactory 抽象工厂
type GUIFactory interface {
	CreateButton() Button
	CreateTextBox() TextBox
}

// WindowFactory 具体工厂
type WindowFactory struct {
}

func (f WindowFactory) CreateButton() Button {
	return WindowButton{}
}

func (f WindowFactory) CreateTextBox() TextBox {
	return WindowTextBox{}
}

// MacFactory 具体工厂
type MacFactory struct {
}

func (f MacFactory) CreateButton() Button {
	return MacButton{}
}
func (f MacFactory) CreateTextBox() TextBox {
	return MacTextBox{}
}
