package config

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config/secure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNewStore(t *testing.T) {
	tests := []struct {
		name             string
		secureAvailable  bool
		expectProxyStore bool
	}{
		{
			name:             "secure store available - returns ProxyStore",
			secureAvailable:  true,
			expectProxyStore: true,
		},
		{
			name:             "secure store unavailable - returns insecure store",
			secureAvailable:  false,
			expectProxyStore: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockInsecure := NewMockStore(ctrl)
			mockSecure := secure.NewMockStore(ctrl)

			mockSecure.EXPECT().Available().Return(tt.secureAvailable)

			store := NewStore(mockInsecure, mockSecure)

			if tt.expectProxyStore {
				proxyStore, ok := store.(*ProxyStore)
				require.True(t, ok, "Expected ProxyStore")
				assert.Equal(t, mockInsecure, proxyStore.insecure)
				assert.Equal(t, mockSecure, proxyStore.secure)
			} else {
				assert.Equal(t, mockInsecure, store)
			}
		})
	}
}

func TestProxyStore_IsSecure(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockInsecure := NewMockStore(ctrl)
	mockSecure := secure.NewMockStore(ctrl)

	store := &ProxyStore{
		insecure: mockInsecure,
		secure:   mockSecure,
	}

	assert.True(t, store.IsSecure())
}

func TestProxyStore_PropertyRouting(t *testing.T) {
	testCases := []struct {
		propertyName string
		isSecure     bool
	}{
		{publicAPIKey, true},
		{privateAPIKey, true},
		{AccessTokenField, true},
		{RefreshTokenField, true},
		{"base_url", false},
		{"project_id", false},
		{"org_id", false},
		{"output", false},
		{"service", false},
	}

	methods := []struct {
		name     string
		testFunc func(t *testing.T, store *ProxyStore, propertyName string, isSecure bool)
	}{
		{
			name:     "GetHierarchicalValue",
			testFunc: testGetHierarchicalValue,
		},
		{
			name:     "SetProfileValue",
			testFunc: testSetProfileValue,
		},
		{
			name:     "GetProfileValue",
			testFunc: testGetProfileValue,
		},
		{
			name:     "SetGlobalValue",
			testFunc: testSetGlobalValue,
		},
		{
			name:     "GetGlobalValue",
			testFunc: testGetGlobalValue,
		},
	}

	for _, method := range methods {
		for _, tc := range testCases {
			t.Run(method.name+"_"+tc.propertyName, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				mockInsecure := NewMockStore(ctrl)
				mockSecure := secure.NewMockStore(ctrl)

				store := &ProxyStore{
					insecure: mockInsecure,
					secure:   mockSecure,
				}

				method.testFunc(t, store, tc.propertyName, tc.isSecure)
			})
		}
	}
}

func testGetHierarchicalValue(t *testing.T, store *ProxyStore, propertyName string, isSecure bool) {
	t.Helper()
	profileName := "test-profile"
	expectedValue := "test-value"

	if isSecure {
		store.secure.(*secure.MockStore).EXPECT().
			Get(profileName, propertyName).
			Return(expectedValue, nil)
	} else {
		store.insecure.(*MockStore).EXPECT().
			GetHierarchicalValue(profileName, propertyName).
			Return(expectedValue)
	}

	result := store.GetHierarchicalValue(profileName, propertyName)
	assert.Equal(t, expectedValue, result)
}

func testSetProfileValue(t *testing.T, store *ProxyStore, propertyName string, isSecure bool) {
	t.Helper()
	profileName := "test-profile"
	value := "test-value"

	if isSecure {
		store.secure.(*secure.MockStore).EXPECT().
			Set(profileName, propertyName, value).
			Return(nil)
	} else {
		store.insecure.(*MockStore).EXPECT().
			SetProfileValue(profileName, propertyName, value)
	}

	store.SetProfileValue(profileName, propertyName, value)
}

func testGetProfileValue(t *testing.T, store *ProxyStore, propertyName string, isSecure bool) {
	t.Helper()
	profileName := "test-profile"
	expectedValue := "test-value"

	if isSecure {
		store.secure.(*secure.MockStore).EXPECT().
			Get(profileName, propertyName).
			Return(expectedValue, nil)
	} else {
		store.insecure.(*MockStore).EXPECT().
			GetProfileValue(profileName, propertyName).
			Return(expectedValue)
	}

	result := store.GetProfileValue(profileName, propertyName)
	assert.Equal(t, expectedValue, result)
}

func testSetGlobalValue(t *testing.T, store *ProxyStore, propertyName string, isSecure bool) {
	t.Helper()
	value := "test-value"

	if isSecure {
		store.secure.(*secure.MockStore).EXPECT().
			Set(DefaultProfile, propertyName, value).
			Return(nil)
	} else {
		store.insecure.(*MockStore).EXPECT().
			SetGlobalValue(propertyName, value)
	}

	store.SetGlobalValue(propertyName, value)
}

func testGetGlobalValue(t *testing.T, store *ProxyStore, propertyName string, isSecure bool) {
	t.Helper()
	expectedValue := "test-value"

	if isSecure {
		store.secure.(*secure.MockStore).EXPECT().
			Get(DefaultProfile, propertyName).
			Return(expectedValue, nil)
	} else {
		store.insecure.(*MockStore).EXPECT().
			GetGlobalValue(propertyName).
			Return(expectedValue)
	}

	result := store.GetGlobalValue(propertyName)
	assert.Equal(t, expectedValue, result)
}

func TestIsSecureProperty(t *testing.T) {
	tests := []struct {
		propertyName string
		expected     bool
	}{
		{publicAPIKey, true},
		{privateAPIKey, true},
		{AccessTokenField, true},
		{RefreshTokenField, true},
		{"base_url", false},
		{"project_id", false},
		{"org_id", false},
		{"output", false},
		{"", false},
		{"random_property", false},
	}

	for _, tt := range tests {
		t.Run(tt.propertyName, func(t *testing.T) {
			result := isSecureProperty(tt.propertyName)
			assert.Equal(t, tt.expected, result)
		})
	}
}
