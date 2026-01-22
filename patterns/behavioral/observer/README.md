## Observer Pattern Use Case: Football Match Updates

Imagine you're developing a football match update system. In this system, you have two main components:

1. **FootballMatch**: This is the 'Subject'. Whenever there's an update in the match (like a goal scored, foul, etc.),
   the FootballMatch will notify all observers.

2. **Fans**: These are the 'Observers'. Fans subscribe to the FootballMatch to receive updates. They could be individual
   fans following the match from their devices or even news channels waiting to broadcast the latest updates.

In this scenario, whenever there's an update in the FootballMatch, all subscribed Fans get notified. This is the
Observer pattern in action. The FootballMatch doesn't need to know how many Fans are subscribed or what they do with the
updates. It just sends updates to all. Similarly, Fans don't need to constantly check FootballMatch for updates. They
just wait to be notified.

This pattern is useful in situations where a state change in one object must be broadcasted and handled by many other
objects. It provides a loosely coupled architecture with high level of abstraction.
