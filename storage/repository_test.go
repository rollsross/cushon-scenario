package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_s_CreateAccountAndFund(t *testing.T) {
	tts := []struct {
		name    string
		prep    func(t *testing.T, r Repository) error
		wantErr string
	}{
		{
			name: "happy path - account and fund created",
			prep: func(t *testing.T, r Repository) error {
				err := r.CreateAccountAndFund(
					"00a79964-34c2-48ab-88ab-de65427cb960",
					"b19caf34-b367-40a0-a3d8-5072119d357e",
					"640aab24-4d60-4af7-bdde-5b6ab70538db",
					2500000,
				)

				return err
			},
		},
		{
			name: "unhappy path - account and fund failed to be created",
			prep: func(t *testing.T, r Repository) error {
				err := r.CreateAccountAndFund(
					"00a79964-34c2-48ab-88ab-de65427cb960",
					"",
					"640aab24-4d60-4af7-bdde-5b6ab70538db",
					2500000,
				)

				return err
			},
			wantErr: ErrFailedToCreateAccountAndFund.Error(),
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			repo, cleanup := getRepoForTest(t)
			defer cleanup()

			got := tt.prep(t, repo)

			if tt.wantErr == "" {
				assert.Nil(t, got)
			} else {
				assert.ErrorContains(t, got, tt.wantErr)
			}
		})
	}
}

func Test_s_GetAccountAndFund(t *testing.T) {
	tts := []struct {
		name    string
		prep    func(t *testing.T, r Repository) (*AccountFund, error)
		wantErr string
	}{
		{
			name: "happy path - account and fund found",
			prep: func(t *testing.T, r Repository) (*AccountFund, error) {
				err := r.CreateAccountAndFund(
					"00a79964-34c2-48ab-88ab-de65427cb960",
					"b19caf34-b367-40a0-a3d8-5072119d357e",
					"640aab24-4d60-4af7-bdde-5b6ab70538db",
					2500000,
				)
				if err != nil {
					t.Fatalf("happy path CreateAccountAndFund: %s", err.Error())
				}

				res, err := r.GetAccountAndFund("00a79964-34c2-48ab-88ab-de65427cb960")

				return res, err
			},
		},
		{
			name: "unhappy path - failed to get account and fund",
			prep: func(t *testing.T, r Repository) (*AccountFund, error) {
				err := r.CreateAccountAndFund(
					"00a79964-34c2-48ab-88ab-de65427cb960",
					"b19caf34-b367-40a0-a3d8-5072119d357e",
					"640aab24-4d60-4af7-bdde-5b6ab70538db",
					2500000,
				)
				if err != nil {
					t.Fatalf("unhappy path CreateAccountAndFund: %s", err.Error())
				}

				res, err := r.GetAccountAndFund("user")

				return res, err
			},
			wantErr: ErrNoRecordsFoundForUser.Error(),
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			repo, cleanup := getRepoForTest(t)
			defer cleanup()

			res, got := tt.prep(t, repo)

			if tt.wantErr == "" {
				assert.Equal(t, res.AccountName, "Cushon ISA")
				assert.Equal(t, res.FundName, "Cushon Equities Fund")
				assert.Equal(t, res.Balance, 2500000)
			} else {
				assert.ErrorContains(t, got, tt.wantErr)
			}
		})
	}
}
