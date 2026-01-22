# Command Pattern

The Command pattern is a behavioral design pattern that turns a request into a stand-alone object, encapsulating all the
information about this request. This transformation allows you to parameterize methods with different requests, queue or
log requests, and even support undoable operations.

## Home Automation System: A Use Case

In the context of a home automation system, consider the following components:

1. **Remote Control (Invoker)**: You have a **remote control** with several buttons. Each button press represents a
   specific command.

2. **Commands**: Commands like turning on the light, turning off the fan, or setting the temperature on the air
   conditioner are encapsulated as objects. These are your **Commands**.

3. **Devices (Receivers)**: The devices in the home (light, fan, air conditioner) are the **Receivers**. They are the
   ones that know how to carry out the commands.

### Workflow

When you press a button on the remote control, the corresponding command is executed on the receiver.

## Conclusion

The Command pattern allows you to encapsulate a request (like "turn on the light") as an object (a LightOnCommand, for
example). This enables you to store and pass requests around within your program and even operate on them (like adding
them to a queue, logging them, or undoing them).

This pattern provides a decoupling between the object invoking the operation and the object which knows how to perform
it, offering flexibility and extensibility in dealing with these operations.
