package src

import (
	"bytes"
	"context"
	"git.andinfinity.de/gitea-release-drafter/src/config"
	"testing"
)

func Test_updateOrCreateDraftRelease(t *testing.T) {
	cfg := &config.DrafterConfig{
		ApiUrl:    "https://git.tianjinzhaofa.cn/",
		Token:     "68b91e8988b6179f081ab1a488d606c8f2d8778c",
		RepoOwner: "felix",
		RepoName:  "butnoice",
	}
	ctx := context.WithValue(context.Background(), "key", "value")
	a, err := NewAction(&ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	configReader := bytes.NewReader([]byte{})
	repoCfg, err := config.ReadRepoConfig(configReader, "main")
	if err != nil {
		panic(err)
	}
	updateOrCreateDraftRelease(a, repoCfg)
}
