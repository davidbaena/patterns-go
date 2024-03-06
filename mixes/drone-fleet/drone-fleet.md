## **Problem Scenario: Autonomous Drone Fleet Management System**

### **Context:**

You've been hired by a cutting-edge drone delivery company that operates a fleet of autonomous drones. These drones are
responsible for delivering packages to customers' doorsteps. The company faces several challenges related to efficient
drone management, safety, and customer satisfaction.

### **Requirements:**

1. **Drone Creation and Cloning:**
    - The company manufactures different types of drones (e.g., quadcopters, hexacopters, fixed-wing drones).
    - When a new drone model is introduced, it should be straightforward to create variations (e.g., different payload
      capacities, flight ranges) based on an existing model.
    - Cloning a drone should copy its specifications while allowing customization of specific properties (e.g., battery
      capacity, camera type).

2. **Dynamic Drone Enhancements:**
    - Some drones require additional features or upgrades (e.g., obstacle detection, night vision, weatherproofing).
    - These enhancements should be dynamically added to the drone without altering its core functionality.
    - For instance, a basic quadcopter can be enhanced with a "Collision Avoidance" module.

3. **Mission Execution and Logging:**
    - When a delivery mission is assigned, the system should execute specific actions related to the drone (e.g.,
      takeoff, route planning, landing).
    - Detailed logs of these actions should be maintained for auditing and troubleshooting.

### **Solution Using Design Patterns:**

1. **Prototype Pattern:**
    - Each drone model will be a prototype.
    - When a new drone variation is needed (e.g., a hexacopter with extended range), we'll clone the prototype and
      customize its specifications.
    - For example, if we have a "Quadcopter" prototype, we can clone it to create a "Quadcopter Plus" with upgraded
      features.

2. **Decorator Pattern:**
    - To handle dynamic enhancements, we'll apply the decorator pattern.
    - Each enhancement (e.g., obstacle detection, night vision) will be a decorator.
    - Decorators will wrap around the base drone and add specific functionalities.
    - For instance, a "Night Vision" decorator can be applied to any drone model.

3. **Command Pattern:**
    - For mission execution, we'll use the command pattern.
    - Each mission action (e.g., takeoff, route planning) will be encapsulated as a command.
    - When a delivery mission starts, the system will execute a series of commands based on the drone's capabilities.
    - The command objects will log their actions (e.g., "Took off from Warehouse A") for auditing.
