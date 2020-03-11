package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v29/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	versionJSONFile    = "version.json"
	defaultHTTPTimeout = 10 * time.Second
)

func main() {
	ver, err := getVersionFromFile(versionJSONFile)
	if err != nil {
		log.Fatalf("failed to get version from %s", versionJSONFile)
	}

	ghCli, err := newClient()
	if err != nil {
		log.Fatalf("failed to create github client err:%v", err)
	}

	ownerRepo := os.Getenv("GITHUB_REPOSITORY")
	if ownerRepo == "" {
		log.Fatalf("GITHUB_REPOSITORY is empty")
	}
	strs := strings.Split(ownerRepo, "/")
	if len(strs) != 2 {
		log.Fatalf("invalid GITHUB_REPOSITORY:%s", ownerRepo)
	}
	owner := strs[0]
	repo := strs[1]

	commit := os.Getenv("GITHUB_SHA")
	if commit == "" {
		log.Fatalf("GITHUB_SHA is empty")
	}

	if err := createAnnotatedTag(ghCli, owner, repo, commit, ver); err != nil {
		log.Fatalf("failed to tag current commit err:%v", err)
	}
}

func newClient() (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("GITHUB_TOKEN is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultHTTPTimeout)
	defer cancel()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc), nil
}

func getVersionFromFile(file string) (string, error) {
	var tmp struct {
		Version string `json:"version"`
	}

	if err := func() error {
		f, err := os.Open(file)
		if err != nil {
			return errors.Wrapf(err, "cannot open %s", file)
		}
		defer f.Close()

		if err := json.NewDecoder(f).Decode(&tmp); err != nil {
			return errors.Wrapf(err, "cannot decode json")
		}

		return nil
	}(); err != nil {
		return "", err
	}

	return tmp.Version, nil
}

// https://developer.github.com/v3/git/tags/#create-a-tag-object
func createAnnotatedTag(ghCli *github.Client, owner, repo, commit, tagName string) error {
	// create annotated tag
	ctx, cancel := context.WithTimeout(context.Background(), defaultHTTPTimeout)
	defer cancel()

	// create tag object
	tag, _, err := ghCli.Git.CreateTag(ctx, owner, repo, &github.Tag{
		Tag:     &tagName,
		Message: &tagName,
		Object: &github.GitObject{
			Type: github.String("commit"),
			SHA:  &commit,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "failed to create tag:%s", tagName)
	}

	// create a ref to tag
	ctx, cancel = context.WithTimeout(context.Background(), defaultHTTPTimeout)
	defer cancel()

	refName := fmt.Sprintf("refs/tags/%s", tagName)
	_, _, err = ghCli.Git.CreateRef(ctx, owner, repo, &github.Reference{
		Ref: &refName,
		Object: &github.GitObject{
			SHA: tag.SHA,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "failed to create ref to tag:%s", tagName)
	}

	return nil
}
