# GoMap 

A high-performance, concurrent port scanner written in [Go](https://go.dev/).

I built this project to learn Go's concurrency model. It is an simplified version of [Nmap](https://nmap.org/) that uses **Goroutines** and **Worker Pools** to scan ports its faster than a traditional single threaded scanner.

## Features

- **Concurrent Scanning:** Uses 100 background workers to scan ports in parallel.
- **No External Dependencies:** Built using only the Go standard library.
- **CLI Interface:** Easily change the target host via command line flags.

