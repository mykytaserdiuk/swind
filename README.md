# SWIND

A simple 2D typing game built with **Go** and **Raylib**. Players match text input against a target word and receive a percentage score.

## About

SWIND is a desktop application with a menu scene and a game scene. The game displays a target word, and the player types to match it. The system compares the input character-by-character and calculates the match percentage (0-100%).

## Technology Stack

- **Language**: Go 1.25.0
- **Graphics**: Raylib 0.55.1
- **Architecture**: Scene-based with event bus pattern
- **UI Framework**: Raygui (Raylib GUI components)

## Key Implementation Details

- **Event Bus**: Decoupled communication between scenes and UI components using pub/sub pattern
- **Scene Management**: Two scenes (Menu and Game) with lifecycle management
- **Text Matching**: Unicode-aware character comparison for accurate scoring
- **Layered Rendering**: Command-based rendering with layer sorting for correct draw order
- **Input Handling**: Real-time keyboard input with uppercase normalization

## Getting Started

## Gameplay

1. Click the "Exit" button on the menu to start the game
2. A target words appears
3. Type the words in the text input field
4. Score updates in real-time as you type

Unit tests cover text matching logic and rendering command queuing.

