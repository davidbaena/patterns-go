### Domains and Responsibilities

1. **Orders**
    - **Responsibility**: Handles the placement of pizza orders.
    - **Use Case**: A customer places an order for a pizza, which triggers an `OrderPlaced` event.

2. **Kitchen**
    - **Responsibility**: Prepares pizzas based on the orders received.
    - **Use Case**: Upon receiving an `OrderPlaced` event, the kitchen prepares the pizza and triggers a `PizzaPrepared` event.
    - **Business Rule**: The kitchen cannot prepare more than the maximum number of pizzas (`maxPizzas`). If the limit is reached, no more pizzas will be prepared, and a message will be logged.

3. **Delivery**
    - **Responsibility**: Manages the delivery of prepared pizzas.
    - **Use Case**: Upon receiving a `PizzaPrepared` event, the delivery service delivers the pizza to the customer.

### Use Case Example

1. **Order Placement**:
    - A customer places an order for a "Margherita" pizza.
    - The `OrderService` publishes an `OrderPlaced` event.

2. **Pizza Preparation**:
    - The `KitchenService` listens for `OrderPlaced` events.
    - Upon receiving the event, it prepares the pizza and publishes a `PizzaPrepared` event.
    - If the maximum number of pizzas (`maxPizzas`) has been prepared, no more pizzas will be prepared, and a message will be logged.

3. **Pizza Delivery**:
    - The `DeliveryService` listens for `PizzaPrepared` events.
    - Upon receiving the event, it delivers the pizza to the customer.