## Problem Statement

The football team has 5 players. Each player enters the locker room at a different time and stays for a different
duration.
The locker room can accommodate a maximum of 3 players at a time. The goal is to create a timeline of how the players
use the locker room.

The output must be a list of players entering and leaving the locker room, along with the time they spend in the locker
room:

```plaintext
Player 5 is entering the locker room.
Player 2 is entering the locker room.
Player 3 is entering the locker room.
Player 3 has left the locker room after 2 seconds.
Player 5 has left the locker room after 2 seconds.
Player 1 is entering the locker room.
Player 4 is entering the locker room.
Player 2 has left the locker room after 3 seconds.
Player 4 has left the locker room after 2 seconds.
Player 1 has left the locker room after 3 seconds.
All players have left the locker room.
```

## Player Usage Timeline

Here's an ASCII diagram representing the time each player has spent in the locker room. Each `-` represents one second.

```
Player 5: |----| 
Player 2: |-----| 
Player 3: |—|
Player 1: |-----| 
Player 4: |—|
```

## Explanation

In this diagram:

- Player 5 enters the locker room first and stays for 2 seconds.
- Player 2 enters the locker room next and stays for 3 seconds.
- Player 3 also enters the locker room at the start and stays for 2 seconds.
- After players 5 and 3 leave, player 1 enters the locker room and stays for 3 seconds.
- Player 4 also enters the locker room after players 5 and 3 leave, and stays for 2 seconds.

This diagram shows that the locker room never has more than 3 players at the same time, and all players have left the
locker room by the end.

## Conclusion

This timeline helps us understand the usage of the locker room by the players. It ensures that the locker room is never
overcrowded and is used efficiently.
