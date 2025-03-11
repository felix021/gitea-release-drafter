package src

import (
    "bytes"
    "context"
    "os"
    "testing"

    "git.andinfinity.de/gitea-release-drafter/src/config"
)

func Test_updateOrCreateDraftRelease(t *testing.T) {
    cfg := &config.DrafterConfig{
        ApiUrl:    "https://git.tianjinzhaofa.cn/",
        Token:     os.Getenv("GITEA_ACCESS_TOKEN"),
        RepoOwner: "braid2048",
        RepoName:  "wizard",
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
