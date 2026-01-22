package main

import "fmt"

type Light struct {
	isOn bool
}

func (l *Light) On() error {
	fmt.Println("Light is on.")
	l.isOn = true
	return nil
}

func (l *Light) Off() error {
	fmt.Println("Light is off.")
	l.isOn = true
	return nil
}

type Fan struct {
	isOn bool
}

func (f *Fan) On() error {
	fmt.Println("Fan is on.")
	f.isOn = true
	return nil
}
func (f *Fan) Off() error {
	fmt.Println("Fan is off.")
	f.isOn = false
	return nil
}

type Command interface {
	execute() error
}

// Concrete commands
type LightCommand struct {
	light Light
}

func (c *LightCommand) execute() error {
	if c.light.isOn {
		c.light.Off()
	} else {
		c.light.On()
	}
	return nil
}

type FanCommand struct {
	fan Fan
}

func (c *FanCommand) execute() error {
	if c.fan.isOn {
		c.fan.Off()
	} else {
		c.fan.On()
	}
	return nil
}

type AirConditioner struct {
	temperature int
}

func (c *AirConditioner) setTemperature(temperature int) error {
	fmt.Printf("Air conditioner temperature set to %d.\n", temperature)
	c.temperature = temperature
	return nil
}

type ACSetTemperatureCommand struct {
	ac          AirConditioner
	temperature int
}

func (c *ACSetTemperatureCommand) execute() error {
	return c.ac.setTemperature(c.temperature)
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) setCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) pressButton() {
	r.command.execute()
}

func main() {
	light := Light{}
	lightCommand := LightCommand{light: light}
	remoteControl := RemoteControl{}
	remoteControl.setCommand(&lightCommand)
	remoteControl.pressButton()

	fan := Fan{}
	fanCommand := FanCommand{fan: fan}
	remoteControl.setCommand(&fanCommand)
	remoteControl.pressButton()

	ac := AirConditioner{}
	acSetTemperatureCommand := ACSetTemperatureCommand{ac: ac, temperature: 25}
	remoteControl.setCommand(&acSetTemperatureCommand)
	remoteControl.pressButton()

	remoteControl.setCommand(&lightCommand)
	remoteControl.pressButton()

	remoteControl.setCommand(&fanCommand)
	remoteControl.pressButton()

}
