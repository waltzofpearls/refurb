package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type dep struct {
	ProjectRoot string
	Latest      string
	Version     string
	Revision    string
}

func main() {
	var (
		deps []dep
		err  error
	)
	cmd := exec.Command("dep", "status", "-json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%s%s", out, err)
	}
	err = json.Unmarshal(out, &deps)
	if err != nil {
		log.Fatalln("error:", err)
	}
	for _, d := range deps {
		if d.Latest == d.Version || d.Latest == d.Revision {
			continue
		}
		latest := d.Latest
		if len(latest) >= 40 {
			latest = d.Latest[:7]
		}
		fmt.Printf("[%s] %s (%s) => %s\n", d.ProjectRoot, d.Version, d.Revision[:7], latest)
	}
}
