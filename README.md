# Multi motor controller
Simulation of a multi-motor controller for [Aisprid technical challenge](doc/Aisprid_ingenieur_robotique_challenge_technique.pdf).

## Compile
In root directory, call `make` to generate binary file. To clean up, call `make clean`.
## Usage
### Configure simulation
- Edit [config file](./cfg/motors.yaml) with simulation parameters. This file should be provided to every commands.
### Servers startup
1. Start Motor servers (with valid IDs from [config file](./cfg/motors.yaml)):
    ```bash
    ./motorsim -c ./cfg/motors.yaml motor --id [ID] serve
    ```
2. Start Motor controller server (Beware, if modifying port controllers commands except Port 8080):
    ```bash
    ./motorsim -c ./cfg/motors.yaml controller serve
    ```
### Client usage
#### Motor controller commands
- Get joints values: 
    ```bash
    ./motorsim -c ./cfg/motors.yaml controller getJoints
    ```
- Set joints values:
    ```bash
    ./motorsim -c ./cfg/motors.yaml controller setJoints [Joint goals]
    ```
#### Motor commands
Every command should be provided with valid ID from [config file](./cfg/motors.yaml):
- Listen motor state:
    ```bash
    ./motorsim -c ./cfg/motors.yaml motor --id [ID] listen
    ```
- Set motor velocity:
    ```bash
    ./motorsim -c ./cfg/motors.yaml motor --id [ID] moveVel -- [Velocity]
    ```

## TODO
- [ ] Closed loop control.
- [ ] Share Motor controller port with controller commands.
- [ ] Use coroutines for gRPC services callbacks.
- [ ] Fix get Motor State service (Behavior not robust).
- [ ] Fix Controller SetJoint command CLI omitting multiples input values.
- [ ] Send error responses to clients of gRPC services.