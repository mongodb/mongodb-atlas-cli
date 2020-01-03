package store

import (
	"testing"

	"github.com/10gen/mcli/mocks"
	"github.com/golang/mock/gomock"
)

func TestStore_New(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConfig := mocks.NewMockConfig(ctrl)

	defer ctrl.Finish()

	mockConfig.
		EXPECT().
		Service().
		Return("ops-manager").
		Times(1)

	mockConfig.
		EXPECT().
		PublicAPIKey().
		Return("1").
		Times(1)

	mockConfig.
		EXPECT().
		PrivateAPIKey().
		Return("1").
		Times(1)

	mockConfig.
		EXPECT().
		APIPath().
		Return("").
		Times(1)

	store, err := New(mockConfig)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if store.baseURL != nil {
		t.Errorf("store.baseURL = %s; want 'nil'", store.baseURL)
	}
}

func TestStore_New_WithUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConfig := mocks.NewMockConfig(ctrl)

	defer ctrl.Finish()

	mockConfig.
		EXPECT().
		Service().
		Return("ops-manager").
		Times(1)

	mockConfig.
		EXPECT().
		PublicAPIKey().
		Return("1").
		Times(1)

	mockConfig.
		EXPECT().
		PrivateAPIKey().
		Return("1").
		Times(1)

	mockConfig.
		EXPECT().
		APIPath().
		Return("ops_manager").
		Times(2)

	store, err := New(mockConfig)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if store.baseURL.String() != "ops_manager" {
		t.Errorf("store.baseURL = %s; want 'ops_manager'", store.baseURL)
	}
}
