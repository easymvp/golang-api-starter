#!/bin/bash

# Test watcher script for Go projects
# Monitors the internal directory for changes and runs tests automatically

set -e

# Check if inotify-tools is installed
if ! command -v inotifywait &> /dev/null; then
    echo "Error: inotify-tools is required but not installed."
    echo "Please install it using your package manager:"
    echo "  For Debian/Ubuntu: sudo apt-get install inotify-tools"
    echo "  For Fedora: sudo dnf install inotify-tools"
    echo "  For macOS: brew install fswatch"
    exit 1
fi

# Directory to monitor
WATCH_DIR="./internal"
echo "Starting test watcher for directory: $WATCH_DIR"
echo "Press Ctrl+C to stop watching"

# Function to run tests
run_tests() {
    echo "File change detected. Running tests..."
    echo "----------------------------------------"
    # Ensure we're running tests from the project root directory
    # Get the directory where the script is located
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    # Navigate to the project root (parent directory of scripts)
    PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"
    # Run tests from the project root
    cd "$PROJECT_ROOT" && go test ./... -v
    echo "----------------------------------------"
    echo "Tests completed at $(date)"
    echo "Watching for changes..."
}

# Initial test run
echo "Running initial tests..."
run_tests

# Monitor for changes
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS uses fswatch instead of inotifywait
    fswatch -o "$WATCH_DIR" | while read -r f; do
        run_tests
    done
else
    # Linux uses inotifywait
    while true; do
        inotifywait -r -e modify,create,delete,move "$WATCH_DIR"
        run_tests
    done
fi
