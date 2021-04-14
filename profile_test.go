package coinbasepro

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetProfiles(t *testing.T) {
	client := NewTestClient()
	profiles, err := client.GetProfiles()
	if err != nil {
		t.Error(err)
	}

	for _, p := range profiles {
		if p.ID == "" {
			t.Error(errors.New("profile id missing"))
		}
	}
}

func TestGetProfile(t *testing.T) {
	client := NewTestClient()
	profiles, err := client.GetProfiles()
	if err != nil {
		t.Error(err)
	}

	profile, err := client.GetProfile(profiles[0].ID)
	if err != nil {
		t.Error(err)
	}

	if profile.ID == "" {
		t.Error(errors.New("profile id missing"))
	}
}

func TestCreateProfileTransfer(t *testing.T) {
	client := NewTestClient()
	profiles, err := client.GetProfiles()
	if err != nil {
		t.Error(err)
	}

	var fromProfile, toProfile Profile
	for _, profile := range profiles {
		if profile.IsDefault {
			fromProfile = profile
			continue
		}

		if profile.Active {
			toProfile = profile
			break
		}
	}

	if toProfile.ID == "" {
		t.Skip(fmt.Sprintf("needed at least two active profiles for this test"))
	}

	// Send from first profile to second profile
	newTransfer := ProfileTransfer{
		From:     fromProfile.ID,
		To:       toProfile.ID,
		Currency: "USD",
		Amount:   "9.99",
	}

	err = client.CreateProfileTransfer(&newTransfer)
	if err != nil {
		t.Error(err)
	}
}
