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
2. Start Motor controller server (Beware, controllers commands except Port 8080):
    ```bash
    ./motorsim -c ./cfg/motors.yaml controller serve -p 8080
    ```
### Client usage
#### Motor controller commands
- Get joints values: 
    ```bash
    ./motorsim -c ./cfg/motors.yaml controller getJoints
    ```
- Set joints values (Beware, currently set the same values for all joints):
    ```bash
    ./motorsim -c ./cfg/motors.yaml controller setJoints --j [Joint goals]
    ```
#### Motor commands
Every command should be provided with valid ID from [config file](./cfg/motors.yaml):
- Listen motor state:
    ```bash
    ./motorsim -c ./cfg/motors.yaml motor --id [ID] listen
    ```
- Set motor velocity:
    ```bash
    ./motorsim -c ./cfg/motors.yaml motor --id [ID] moveVel --vel [Velocity]
    ```

## TODO
- [ ] Prevent startup of multiple motors with same ID.