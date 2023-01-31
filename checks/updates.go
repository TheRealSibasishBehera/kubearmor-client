package checks

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/kubearmor/kubearmor-client/selfupdate"
	"strings"
)

func CheckForUpdates() error {
	client := github.NewClient(nil)

	// Fetch all releases from the GitHub API
	releases, _, err := client.Repositories.ListReleases(context.Background(),
		"kubearmor",
		"kubearmor-client",
		&github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	)
	if err != nil {
		fmt.Println("Error fetching releases:", err)
		return err
	}

	// Find the latest release with the keyword "mandatory" in its release notes
	var latestMandatoryRelease *github.RepositoryRelease
	for _, release := range releases {
		if strings.Contains(*release.Body, "mandatory") {
			latestMandatoryRelease = release
			break
		}
	}

	// Compare the current version with the latest release version
	currentVersion := selfupdate.GitSummary
	if err != nil {
		return err
	}
	if latestMandatoryRelease != nil && !strings.EqualFold(currentVersion, *latestMandatoryRelease.TagName) {
		color.HiMagenta("A mandatory update is available (current version: %s, latest release: %s).\n",
			currentVersion,
			*latestMandatoryRelease.TagName,
		)
	} else {
		fmt.Println("The client is up to date.")
	}

	return nil
}
