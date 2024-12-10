# Family Feud In a SHELL

**Family Feud In a SHELL** is a command-line application developed in Go, inspired by the classic game show *Family Feud*, and built using [Charms Bubble Tea](https://github.com/charmbracelet/bubbletea) for an interactive text-based experience. This app is designed to be hosted by a game master (GM), who manages the questions and answers while players attempt to guess the most popular responses.

## Features

- **Interactive Interface**: Uses Bubble Tea for a smooth and intuitive command-line interface.
- **Game Master Support**: Designed for gameplay with a GM who introduces questions, validates responses, and manages scoring.
- **Custom Question-Answer Configurations**: The GM can import a custom configuration file with questions and answers.
- **Score Tracking**: Point system to evaluate the popularity of responses, allowing teams or individual players to compete.

## Requirements

- **Go 1.20+**: Make sure Go is installed on your machine.
- **Docker** (for containerized setup)

## Installation

### Standard Installation

1. Clone the repository and install necessary dependencies:

   ```bash
   git clone <your-git-repo>
   cd family-feud-cli
   make install-dependencies
   make install-binary
   ```

2. Create your questions file

for now only `JSON` format is allowed, check the [config](config/config.json) for an example

3. Run the game:

   ```bash
   family-feud start -c <YOUR_CONFIG_PATH>
   ```

### Docker Installation

For a containerized setup:

1. Build the Docker image:

   ```bash
   docker build -t family-feud . 
   ```

2. Run the container:

   ```bash
   docker run -it -v <YOUR_CONFIG_PATH>:/app/config.json family-feud start -c /app/config.json
   ```

## Usage

To start a game session, run:

```bash
family-feud start -c <YOUR_CONFIG_PATH>
```

### Game Commands

These commands help the game master (GM) to control various aspects of the game in **Family Feud CLI**:

| Command             | Key Binding        | Description                                                        |
|---------------------|--------------------|--------------------------------------------------------------------|
| **Select Blue Team**| `B`                | Assigns points to the blue family.                                 |
| **Select Red Team** | `R`                | Assigns points to the red family.                                  |
| **Switch Team**     | `tab`              | Switch between red and blue teams                                  |
| **Win Round**       | `W`                | Declares the current round's winner, allocating points accordingly.|
| **Correct Answer**  | `1-8`              | Marks a player’s response as correct and awards points.            |
| **Incorrect Answer**| `X` or `0`         | Marks a player’s response as incorrect, no points awarded.         |
| **Next Round**      | `N`                | Moves on to the next round in the game.                            |
| **Reset Face-Off**  | `F`                | Resets the face-off for a new question.                            |
| **Quit Game**       | `Ctrl + C`         | Ends the game and exits the application.                           |

These commands provide a quick reference for GMs, making it easy to manage rounds, track scores, and control game flow with keyboard shortcuts.

## Contributing

Contributions are welcome! If you'd like to add new features or improve the interface, please feel free to open a pull request.

## TODO

- Add testing
- handle buttons/buzzer for family faceoff

