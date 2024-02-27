# Smart Home Climate Control System

## Overview

Imagine a futuristic smart home equipped with an intelligent climate control system. Our system dynamically adapts to
environmental changes, optimizes energy consumption, and ensures comfort for the occupants. Let's break down how the
three design patterns come into play:

1. **State Pattern**: The state pattern allows our climate control system to transition seamlessly between different
   operational modes based on external factors (e.g., time of day, occupancy, weather conditions). Each mode represents
   a distinct state.

2. **Prototype Pattern**: The prototype pattern helps us create new instances of climate control profiles efficiently.
   We'll define prototype objects for various scenarios (e.g., "Morning Mode," "Eco Mode," "Party Mode") and clone them
   as needed.

3. **Observer Pattern**: Our smart home system involves multiple components (temperature sensors, humidity sensors, user
   interfaces). The observer pattern ensures that these components stay synchronized. When the system state changes (
   e.g., from "Morning Mode" to "Party Mode"), all relevant components receive updates.

## Details

1. **State Pattern**:
    - **States**: Our climate control system has the following states:
        - **Morning Mode**: Gentle temperature rise, natural light simulation, and soft music.
        - **Eco Mode**: Energy-efficient settings (lower temperature, minimal lighting).
        - **Party Mode**: Elevated temperature, vibrant lighting, and lively music.
    - **Context (ClimateControlSystem)**: Manages the current state and transitions between modes. When occupants wake
      up, it switches to "Morning Mode." If no one is home, it enters "Eco Mode."

2. **Prototype Pattern**:
    - **Prototype Objects**:
        - **MorningProfile**: Represents morning settings (temperature, lighting, music).
        - **EcoProfile**: Represents energy-saving settings.
        - **PartyProfile**: Represents party settings.
    - **Cloning Mechanism**: When a user customizes their preferences (e.g., adjusts the morning temperature), the
      system clones the relevant prototype and applies the modifications.

3. **Observer Pattern**:
    - **Observers**:
        - **TemperatureSensor**: Monitors room temperature.
        - **OccupancySensor**: Detects human presence.
        - **UIPanel**: Displays the current mode and allows users to switch between profiles.
    - **Subject (ClimateControlSystem)**: Notifies observers when the system state changes.
    - **Behavior**:
        - When the temperature drops unexpectedly (observed by the TemperatureSensor), the system switches to "Eco
          Mode."
        - If someone enters the room (detected by the OccupancySensor), it transitions to "Party Mode."
        - The UIPanel updates its display based on the current state.
