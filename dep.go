package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type dep struct {
	ProjectRoot string
	Latest      string
	Version     string
	Revision    string
}

func goDep() (output string, outdated int, err error) {
	cmd := exec.Command("dep", "status", "-json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		output = string(out)
		return
	}
	var deps []dep
	err = json.Unmarshal(out, &deps)
	if err != nil {
		return
	}
	var b bytes.Buffer
	for _, d := range deps {
		if d.Latest == d.Version || d.Latest == d.Revision {
			continue
		}
		latest := d.Latest
		if len(latest) >= 40 {
			latest = d.Latest[:7]
		}
		b.WriteString(fmt.Sprintf("[%s] %s (%s) => %s\n",
			d.ProjectRoot, d.Version, d.Revision[:7], latest))
		outdated++
	}
	output = strings.TrimRight(b.String(), "\n")
	return
}
