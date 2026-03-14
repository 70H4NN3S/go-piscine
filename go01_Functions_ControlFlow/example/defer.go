package main

import os

func readFile(path string) error {
	f, err := os.Open(path)
	if err != nil { return err }
	defer f.Close() // always runs, even on early return
	// ... process file ...
	return nil
}
