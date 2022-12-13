package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	tv := &TV{}
	controller := Controller{
		okCommand: &SwitchCommand{
			device: tv,
		},
		verticalCommand: &VolumeCommand{
			device: tv,
		},
		horizontalCommand: &ChannelCommand{
			device: tv,
		},
	}
	controller.ButtonOkHold()
	controller.ButtonLeftClick()
	controller.ButtonLeftClick()
	controller.ButtonUpClick()
	controller.ButtonRightClick()
	controller.ButtonDownClick()
	controller.ButtonDownClick()
	controller.ButtonOkClick()
	fmt.Println("----------------------")

	radio := &Radio{}
	controller = Controller{
		okCommand: &SwitchCommand{
			device: radio,
		},
		verticalCommand: &VolumeCommand{
			device: radio,
		},
		horizontalCommand: &ChannelCommand{
			device: radio,
		},
	}
	controller.ButtonOkHold()
	controller.ButtonLeftClick()
	controller.ButtonLeftClick()
	controller.ButtonUpClick()
	controller.ButtonRightClick()
	controller.ButtonDownClick()
	controller.ButtonDownClick()
	controller.ButtonOkClick()
}
