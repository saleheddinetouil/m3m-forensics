# m3m-forensics: Basic Memory Forensics Tool in Go

This tool provides basic memory analysis functionalities, focusing on process information, network connections, and (potentially in the future) loaded modules. It leverages the `github.com/google/gops` library for process access.

## Features

* Lists running processes with details (PID, executable path, command-line arguments).
* Displays active network connections.
* [Future] Analyze loaded modules.  (This part is outlined but not fully implemented).

## Getting Started

1. **Install gops:** `go install github.com/google/gops@latest`  (Make sure gops is running)
2. **Build:** `go build`
3. **Run:** `./memforensics [process_id]` (Optional: Provide a specific PID to analyze)

## Usage

* Without arguments, the tool lists all running processes.
* Provide a PID as an argument to analyze a specific process.

## Code Structure

* `main.go`: Entry point and command-line argument handling.
* `process.go`: Functions for process analysis.
* `network.go`: Functions for network connection analysis.
* `modules.go`: (Placeholder) For future module analysis features.



## Contributing

Contributions are welcome!  (Especially for module analysis)



## License

MIT
