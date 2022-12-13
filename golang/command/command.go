package command

import "fmt"

type Switchable interface {
	On()
	Off()
}

type Device interface {
	Switchable
	ChannelUp()
	ChannelDown()
	VolumeUp()
	VolumeDown()
}

type Command interface {
	Exe()
	UnExe()
}

type TV struct {
}

func (t *TV) On() {
	fmt.Println("电视机启动")
}
func (t *TV) Off() {
	fmt.Println("电视机关闭")
}

func (t *TV) ChannelUp() {
	fmt.Println("电视机频道+")
}

func (t *TV) ChannelDown() {
	fmt.Println("电视机频道-")
}

func (t *TV) VolumeUp() {
	fmt.Println("电视机音量+")
}

func (t *TV) VolumeDown() {
	fmt.Println("电视机音量-")
}

type Radio struct {
}

func (r *Radio) On() {
	fmt.Println("收音机启动")
}
func (r *Radio) Off() {
	fmt.Println("收音机关闭")
}

func (r *Radio) ChannelUp() {
	fmt.Println("收音机调频+")
}

func (r *Radio) ChannelDown() {
	fmt.Println("收音机调频-")
}

func (r *Radio) VolumeUp() {
	fmt.Println("收音机音量+")
}

func (r *Radio) VolumeDown() {
	fmt.Println("收音机音量-")
}

type SwitchCommand struct {
	device Device
}

func (s *SwitchCommand) Exe() {
	s.device.On()
}

func (s *SwitchCommand) UnExe() {
	s.device.Off()
}

type ChannelCommand struct {
	device Device
}

func (c *ChannelCommand) Exe() {
	c.device.ChannelUp()
}

func (c *ChannelCommand) UnExe() {
	c.device.ChannelDown()
}

type VolumeCommand struct {
	device Device
}

func (v *VolumeCommand) Exe() {
	v.device.VolumeUp()
}

func (v *VolumeCommand) UnExe() {
	v.device.VolumeDown()
}

type Controller struct {
	okCommand         Command
	verticalCommand   Command
	horizontalCommand Command
}

func (c *Controller) ButtonOkHold() {
	fmt.Println("长按OK键。。。")
	c.okCommand.Exe()
}

func (c *Controller) ButtonOkClick() {
	fmt.Println("单击OK键。。。")
	c.okCommand.UnExe()
}

func (c *Controller) ButtonUpClick() {
	fmt.Println("单击↑按键。。。")
	c.verticalCommand.Exe()
}

func (c *Controller) ButtonDownClick() {
	fmt.Println("单击↓按键。。。")
	c.verticalCommand.UnExe()
}

func (c *Controller) ButtonLeftClick() {
	fmt.Println("单击←按键。。。")
	c.horizontalCommand.UnExe()
}

func (c *Controller) ButtonRightClick() {
	fmt.Println("单击→按键。。。")
	c.horizontalCommand.Exe()
}
