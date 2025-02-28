package src

import (
	"code.gitea.io/sdk/gitea"
	"github.com/sethvargo/go-githubactions"
	"regexp"
	"strings"
)

type Changelog map[string][]*gitea.PullRequest

// GenerateChangelog fetches all the pull requests merged into the default branch since the last release and groups them by label. note that duplicates might occur if a pull request has multiple labels.
func GenerateChangelog(c *gitea.Client, owner string, repo string, lastRelease *gitea.Release) (*Changelog, error) {
	changelogByLabels := make(Changelog)

	// FIXME: use pagination
	prs, _, err := c.ListRepoPullRequests(owner, repo, gitea.ListPullRequestsOptions{
		State: gitea.StateClosed,
	})
	if err != nil {
		return nil, err
	}

	for _, pr := range prs {
		// only consider merged pull requests. note that we can't filter by that in the API
		if pr.HasMerged {
			// if there was a release, only take into account pull requests that have been merged after that
			if lastRelease == nil || lastRelease != nil && pr.Merged.After(lastRelease.CreatedAt) {
				for _, l := range pr.Labels {
					_, ok := changelogByLabels[l.Name]

					if ok {
						changelogByLabels[l.Name] = append(changelogByLabels[l.Name], pr)
					} else {
						changelogByLabels[l.Name] = []*gitea.PullRequest{pr}
					}
				}

				if len(pr.Labels) == 0 {
					githubactions.Warningf("PR #%d doesn't have any labels, guess one", pr.ID)
					label := readLabelFromTitle(pr.Title)
					changelogByLabels[label] = append(changelogByLabels[label], pr)
				}
			}
		}
	}

	return &changelogByLabels, nil
}

// if title is formatted like "label: xxx" or "[label] xxx" we should return "label" otherwise "(uncategorized)"
func readLabelFromTitle(title string) string {
	title = strings.TrimSpace(title)
	// Define the regex pattern to match "label: xxx" or "[label] xxx"
	pattern := `^(([^]]+):|\[([^]]+)\])`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Find the matches
	matches := re.FindStringSubmatch(title)

	// If matches are found, return the label
	if len(matches) > 2 {
		if matches[2] != "" {
			return matches[2]
		}
		return matches[3]
	}

	// If no matches
	return "(uncategorized)"
}
