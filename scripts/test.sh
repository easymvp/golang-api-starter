#!/bin/bash

# Test script for Go projects
# Can run tests once or monitor for changes and run tests automatically

# Explicitly set to not exit on error
set +e

# Default interval in seconds
INTERVAL=5

# Parse command line arguments
WATCH_MODE=false
for arg in "$@"; do
    case $arg in
        --watch)
            WATCH_MODE=true
            shift
            ;;
        --interval=*)
            INTERVAL="${arg#*=}"
            shift
            ;;
    esac
done

# Check if watch mode is enabled and required tools are installed
if [ "$WATCH_MODE" = true ]; then
    if ! command -v inotifywait &> /dev/null && [[ "$OSTYPE" != "darwin"* ]]; then
        echo "Error: inotify-tools is required for watch mode but not installed."
        echo "Please install it using your package manager:"
        echo "  For Debian/Ubuntu: sudo apt-get install inotify-tools"
        echo "  For Fedora: sudo dnf install inotify-tools"
        echo "  For macOS: brew install fswatch"
        exit 1
    fi

    if [[ "$OSTYPE" == "darwin"* ]] && ! command -v fswatch &> /dev/null; then
        echo "Error: fswatch is required for watch mode on macOS but not installed."
        echo "Please install it using: brew install fswatch"
        exit 1
    fi
fi

# Directory to monitor (only used in watch mode)
WATCH_DIR="./internal"

# Function to run tests
run_tests() {
    local change_detected=$1

    if [ "$change_detected" = true ]; then
        echo "File change detected. Waiting $INTERVAL seconds before running tests..."
        sleep $INTERVAL
        echo "Running tests..."
    else
        echo "Running tests..."
    fi

    echo "----------------------------------------"
    # Ensure we're running tests from the project root directory
    # Get the directory where the script is located
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    # Navigate to the project root (parent directory of scripts)
    PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

    # Run tests from the project root
    # Use separate commands instead of && to continue even if tests fail
    cd "$PROJECT_ROOT"
    go test ./... -v
    # Store the test result but continue execution regardless of test success/failure
    TEST_RESULT=$?

    echo "----------------------------------------"
    echo "Tests completed at $(date)"

    # Always return success (0) regardless of test outcome
    return 0
}

# Run tests initially (with || true to ensure script continues even if this fails)
run_tests false || true

# If watch mode is enabled, monitor for changes
if [ "$WATCH_MODE" = true ]; then
    echo "Starting test watcher for directory: $WATCH_DIR"
    echo "Using interval: $INTERVAL seconds"
    echo "Press Ctrl+C to stop watching"

    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS uses fswatch instead of inotifywait
        # Use trap to ensure we don't exit the loop
        trap '' ERR
        fswatch -o "$WATCH_DIR" | while read -r f; do
            # Run tests but ensure the loop continues even if tests fail
            run_tests true || true
        done
    else
        # Linux uses inotifywait
        # Use trap to ensure we don't exit the loop
        trap '' ERR
        while true; do
            inotifywait -r -e modify,create,delete,move "$WATCH_DIR" || true
            # Run tests but ensure the loop continues even if tests fail
            run_tests true || true
        done
    fi
fi
