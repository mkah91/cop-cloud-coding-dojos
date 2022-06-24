# CoP Cloud Coding Dojos

### VS Code Remote - Dev Containers

Follow these steps to open the development container using the VS Code Remote - Containers extension:

1. If this is your first time using a development container, please ensure your system meets the pre-reqs (i.e. have Docker installed) in the [getting started steps](https://aka.ms/vscode-remote/containers/getting-started).
1. Make sure you have the [VS Code Remote - Containers extension](https://aka.ms/vscode-remote/containers) installed.
1. To use this repository you can open a locally cloned copy of the code with VS Code:

   - Clone this repository to your local filesystem.
   - Press <kbd>F1</kbd> and select the **Remote-Containers: Open Folder in Container...** command.
   - Select the cloned copy of this folder, wait for the container to start, and try things out!
   
## Things to try

Once the dev container is running, you'll be able to work with it like you would locally.

> **Note:** This container runs as a non-root user with sudo access by default. Comment out `"remoteUser": "vscode"` in `.devcontainer/devcontainer.json` if you'd prefer to run as root.

Before you start please make sure you are able to push and pull changes from GitHub e.g. with the following command:
```sh
gh auth login
```
Also make sure the enviroment variables like in the env.sample file are adapted and set.

Some things to try:
1. Run Hello World
    ```sh
    # inside root folder
    go run ./hello
    # or inside hello folder
    go run .
    ```
2. Run Hello World Test
    ```sh
    # inside root folder
    go test ./hello/...
    # or inside hello folder
    go test ./...
    ```
3. Debug your code
    - Set a break point in your code
    - Hit `F5` to start the debugger (make sure you are using the right laucher)
    - See the debugger stop the execution at your breakpoint
4. Use (mob)[https://mob.sh/]
    - Make sure you don't have any uncommited changes
    - Make sure your enviroment variables are set correctly
    - Make sure you are able to push and pull changes from GitHub
    - Run `mob start 5` to start the git handover tool for remote pair/mob programming.
    - See a timer got started at https://timer.mob.sh/itk-cloud-coding-dojo. Make sure your name appears in the History. Else check your enviroment variables again.
    - Run `mob stop` to stop the git handover tool (the timer will continue)

## Go Cheat Sheat
https://devhints.io/go
