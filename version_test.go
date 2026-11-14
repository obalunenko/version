package version

import (
	"runtime/debug"
	"testing"
)

func TestParseBuildInfo(t *testing.T) {
	const (
		revision = "0bafe32f28e1b8223f7d99f7846d4c9498b00a98"
		vcsTime  = "2026-11-14T15:22:10Z"
	)

	newInfo := func(settings ...debug.BuildSetting) *debug.BuildInfo {
		return &debug.BuildInfo{
			GoVersion: "go1.24.0",
			Path:      "example.com/app/cmd/app",
			Main:      debug.Module{Path: "example.com/app", Version: "v1.2.3"},
			Settings:  settings,
		}
	}

	tests := []struct {
		name            string
		info            *debug.BuildInfo
		wantCommit      string
		wantShortCommit string
		wantBuildDate   string
	}{
		{
			name: "clean tree",
			info: newInfo(
				debug.BuildSetting{Key: "vcs.revision", Value: revision},
				debug.BuildSetting{Key: "vcs.time", Value: vcsTime},
				debug.BuildSetting{Key: "vcs.modified", Value: "false"},
			),
			wantCommit:      revision,
			wantShortCommit: "0bafe32",
			wantBuildDate:   vcsTime,
		},
		{
			name: "modified tree",
			info: newInfo(
				debug.BuildSetting{Key: "vcs.revision", Value: revision},
				debug.BuildSetting{Key: "vcs.time", Value: vcsTime},
				debug.BuildSetting{Key: "vcs.modified", Value: "true"},
			),
			wantCommit:      revision + "+CHANGES",
			wantShortCommit: "0bafe32+CHANGES",
			wantBuildDate:   vcsTime,
		},
		{
			name:            "no vcs information",
			info:            newInfo(),
			wantCommit:      unset,
			wantShortCommit: unset,
			wantBuildDate:   unset,
		},
		{
			name: "short revision",
			info: newInfo(
				debug.BuildSetting{Key: "vcs.revision", Value: "abc"},
			),
			wantCommit:      "abc",
			wantShortCommit: "abc",
			wantBuildDate:   unset,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseBuildInfo(tt.info)

			if got.commit != tt.wantCommit {
				t.Errorf("commit = %q, want %q", got.commit, tt.wantCommit)
			}

			if got.shortcommit != tt.wantShortCommit {
				t.Errorf("shortcommit = %q, want %q", got.shortcommit, tt.wantShortCommit)
			}

			if got.builddate != tt.wantBuildDate {
				t.Errorf("builddate = %q, want %q", got.builddate, tt.wantBuildDate)
			}

			if got.version != "v1.2.3" {
				t.Errorf("version = %q, want %q", got.version, "v1.2.3")
			}

			if got.appname != "example.com/app/cmd/app" {
				t.Errorf("appname = %q, want %q", got.appname, "example.com/app/cmd/app")
			}

			if got.goversion != "go1.24.0" {
				t.Errorf("goversion = %q, want %q", got.goversion, "go1.24.0")
			}
		})
	}
}
