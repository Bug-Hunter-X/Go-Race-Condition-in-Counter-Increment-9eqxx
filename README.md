# Go Race Condition Example

This repository demonstrates a race condition in Go.  Multiple goroutines concurrently increment a shared counter variable without proper synchronization, leading to an incorrect final count.

## The Problem

The `bug.go` file contains a simple program that launches multiple goroutines to increment a counter.  Because the increment operation (`counter++`) is not atomic and is not protected by a mutex or other synchronization mechanism, race conditions can occur, resulting in the final counter value being lower than expected.

## The Solution

The `bugSolution.go` file demonstrates how to fix the race condition using a mutex.  The mutex ensures that only one goroutine can access and modify the counter at a time, preventing race conditions and producing the correct result.

## How to Run

1. Clone this repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` to see the race condition in action.  You'll likely get a counter value less than 1000.
4. Run `go run bugSolution.go` to see the correct implementation.  This will consistently output a counter value of 1000.
