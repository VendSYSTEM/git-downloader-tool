package main

import (
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestSecurityWorkflowPolicy(t *testing.T) {
	data, err := os.ReadFile(".github/workflows/ci.yml")
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		On struct {
			PullRequest struct {
				Types []string `yaml:"types"`
			} `yaml:"pull_request"`
		} `yaml:"on"`
		Jobs map[string]struct {
			If              string `yaml:"if"`
			Needs           string `yaml:"needs"`
			ContinueOnError string `yaml:"continue-on-error"`
			Steps           []struct {
				If              string `yaml:"if"`
				Run             string `yaml:"run"`
				ContinueOnError string `yaml:"continue-on-error"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		t.Fatal(err)
	}
	for _, event := range []string{"labeled", "unlabeled"} {
		if !containsCIEvent(workflow.On.PullRequest.Types, event) {
			t.Errorf("pull_request must run on %s", event)
		}
	}
	const policy = "${{ github.event_name == 'pull_request' && contains(github.event.pull_request.labels.*.name, 'skip:security') }}"
	for jobName, scan := range map[string]string{"trivy": "trivy fs ", "govulncheck": "govulncheck ./..."} {
		job, ok := workflow.Jobs[jobName]
		if !ok || job.If != "" || job.Needs != "" || job.ContinueOnError != "" {
			t.Errorf("%s must run independently and remain strict", jobName)
		}
		found := false
		for _, step := range job.Steps {
			if strings.Contains(step.Run, scan) && step.If == "" {
				if step.ContinueOnError != policy {
					t.Errorf("%s scan must only allow failures on labeled PRs", jobName)
				}
				found = true
			} else if step.ContinueOnError != "" {
				t.Errorf("%s non-scan steps must remain strict", jobName)
			}
		}
		if !found {
			t.Errorf("%s must run its scan without a condition", jobName)
		}
	}
	checks, ok := workflow.Jobs["checks"]
	if !ok || checks.If != "" || checks.Needs != "" || checks.ContinueOnError != "" {
		t.Error("Go checks must run independently and remain strict")
	}
	for name, job := range workflow.Jobs {
		if name == "trivy" || name == "govulncheck" {
			continue
		}
		if job.ContinueOnError != "" {
			t.Errorf("%s job must remain strict", name)
		}
		for _, step := range job.Steps {
			if step.ContinueOnError != "" {
				t.Errorf("%s steps must remain strict", name)
			}
		}
	}
}

func containsCIEvent(events []string, target string) bool {
	for _, event := range events {
		if event == target {
			return true
		}
	}
	return false
}
