package secure

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNewSecureStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKeyring := NewMockKeyringClient(ctrl)
	profileNames := []string{"profile1", "profile2", "profile3"}
	secureProperties := []string{"public_api_key", "private_api_key", "access_token", "refresh_token"}

	// Setup expectations for loading existing values
	mockKeyring.EXPECT().Get("atlascli_profile1", "public_api_key").Return("existing_public_1", nil)
	mockKeyring.EXPECT().Get("atlascli_profile1", "private_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile1", "access_token").Return("existing_access_1", nil)
	mockKeyring.EXPECT().Get("atlascli_profile1", "refresh_token").Return("", errors.New("not found"))

	mockKeyring.EXPECT().Get("atlascli_profile2", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile2", "private_api_key").Return("existing_private_2", nil)
	mockKeyring.EXPECT().Get("atlascli_profile2", "access_token").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile2", "refresh_token").Return("", errors.New("not found"))

	mockKeyring.EXPECT().Get("atlascli_profile3", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile3", "private_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile3", "access_token").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile3", "refresh_token").Return("", errors.New("not found"))

	store := NewSecureStoreWithClient(profileNames, secureProperties, mockKeyring)

	// Verify cache structure is initialized
	assert.NotNil(t, store.cache)
	assert.NotNil(t, store.cache["profile1"])
	assert.NotNil(t, store.cache["profile2"])
	assert.NotNil(t, store.cache["profile3"])

	// Verify existing values were loaded correctly
	assert.Equal(t, "existing_public_1", store.cache["profile1"]["public_api_key"])
	assert.Equal(t, "existing_access_1", store.cache["profile1"]["access_token"])
	assert.Equal(t, "existing_private_2", store.cache["profile2"]["private_api_key"])

	// Verify secure properties are stored
	assert.Equal(t, secureProperties, store.secureProperties)

	// Verify pending operations is empty
	assert.Empty(t, store.pendingOps)

	// Verify keyring client is set
	assert.Equal(t, mockKeyring, store.keyringClient)
}

func TestKeyringStore_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKeyring := NewMockKeyringClient(ctrl)
	profileNames := []string{"profile1"}
	secureProperties := []string{"public_api_key", "private_api_key"}

	// Setup expectations for loading (no existing values)
	mockKeyring.EXPECT().Get("atlascli_profile1", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile1", "private_api_key").Return("", errors.New("not found"))

	store := NewSecureStoreWithClient(profileNames, secureProperties, mockKeyring)

	t.Run("Set secure property", func(t *testing.T) {
		store.Set("profile1", "public_api_key", "test_public_key")

		// Verify value is in cache
		assert.Equal(t, "test_public_key", store.cache["profile1"]["public_api_key"])

		// Verify pending operation is added
		assert.Len(t, store.pendingOps, 1)
		assert.Equal(t, opSet, store.pendingOps[0].opType)
		assert.Equal(t, "profile1", store.pendingOps[0].profileName)
		assert.Equal(t, "public_api_key", store.pendingOps[0].propertyName)
		assert.Equal(t, "test_public_key", store.pendingOps[0].value)
	})

	t.Run("Set non-secure property", func(t *testing.T) {
		store.Set("profile1", "non_secure_prop", "value")

		// Verify value is NOT in cache
		assert.Empty(t, store.cache["profile1"]["non_secure_prop"])

		// Verify no new pending operation is added
		assert.Len(t, store.pendingOps, 1) // Still only the one from previous test
	})

	t.Run("Set for new profile", func(t *testing.T) {
		store.Set("new_profile", "private_api_key", "new_private_key")

		// Verify profile is created and value is set
		assert.Equal(t, "new_private_key", store.cache["new_profile"]["private_api_key"])

		// Verify pending operation is added
		assert.Len(t, store.pendingOps, 2)
	})
}

func TestKeyringStore_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKeyring := NewMockKeyringClient(ctrl)
	profileNames := []string{"profile1"}
	secureProperties := []string{"public_api_key", "private_api_key"}

	// Setup expectations for loading (no existing values)
	mockKeyring.EXPECT().Get("atlascli_profile1", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile1", "private_api_key").Return("", errors.New("not found"))

	store := NewSecureStoreWithClient(profileNames, secureProperties, mockKeyring)

	// Set a value in cache
	store.cache["profile1"]["public_api_key"] = "cached_value"

	t.Run("Get existing value", func(t *testing.T) {
		value := store.Get("profile1", "public_api_key")
		assert.Equal(t, "cached_value", value)
	})

	t.Run("Get non-existing property", func(t *testing.T) {
		value := store.Get("profile1", "non_existing")
		assert.Empty(t, value)
	})

	t.Run("Get from non-existing profile", func(t *testing.T) {
		value := store.Get("non_existing_profile", "public_api_key")
		assert.Empty(t, value)
	})
}

func TestKeyringStore_DeleteKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKeyring := NewMockKeyringClient(ctrl)
	profileNames := []string{"profile1"}
	secureProperties := []string{"public_api_key", "private_api_key"}

	// Setup expectations for loading (no existing values)
	mockKeyring.EXPECT().Get("atlascli_profile1", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile1", "private_api_key").Return("", errors.New("not found"))

	store := NewSecureStoreWithClient(profileNames, secureProperties, mockKeyring)

	// Set some values in cache
	store.cache["profile1"]["public_api_key"] = "profile1_value1"
	store.cache["profile1"]["private_api_key"] = "profile1_value2"

	t.Run("Delete existing key", func(t *testing.T) {
		store.DeleteKey("profile1", "public_api_key")

		// Verify value is removed from cache
		_, exists := store.cache["profile1"]["public_api_key"]
		assert.False(t, exists)

		// Verify other values remain
		assert.Equal(t, "profile1_value2", store.cache["profile1"]["private_api_key"])

		// Verify pending operation is added
		assert.Len(t, store.pendingOps, 1)
		assert.Equal(t, opDelete, store.pendingOps[0].opType)
		assert.Equal(t, "profile1", store.pendingOps[0].profileName)
		assert.Equal(t, "public_api_key", store.pendingOps[0].propertyName)
	})

	t.Run("Delete from non-existing profile", func(t *testing.T) {
		store.DeleteKey("non_existing", "public_api_key")

		// Verify pending operation is still added
		assert.Len(t, store.pendingOps, 2)
	})
}

func TestKeyringStore_DeleteProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKeyring := NewMockKeyringClient(ctrl)
	profileNames := []string{"profile1", "profile2"}
	secureProperties := []string{"public_api_key", "private_api_key"}

	// Setup expectations for loading (no existing values)
	mockKeyring.EXPECT().Get("atlascli_profile1", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile1", "private_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile2", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile2", "private_api_key").Return("", errors.New("not found"))

	store := NewSecureStoreWithClient(profileNames, secureProperties, mockKeyring)

	// Set some values in cache
	store.cache["profile1"]["public_api_key"] = "value1"
	store.cache["profile2"]["private_api_key"] = "value2"

	t.Run("Delete existing profile", func(t *testing.T) {
		store.DeleteProfile("profile1")

		// Verify profile is removed from cache
		_, exists := store.cache["profile1"]
		assert.False(t, exists)

		// Verify other profiles remain
		assert.Equal(t, "value2", store.cache["profile2"]["private_api_key"])

		// Verify pending operation is added
		assert.Len(t, store.pendingOps, 1)
		assert.Equal(t, opDeleteProfile, store.pendingOps[0].opType)
		assert.Equal(t, "profile1", store.pendingOps[0].profileName)
	})
}

func TestKeyringStore_Save(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockKeyring := NewMockKeyringClient(ctrl)
	profileNames := []string{"profile1"}
	secureProperties := []string{"public_api_key", "private_api_key"}

	// Setup expectations for loading (no existing values)
	mockKeyring.EXPECT().Get("atlascli_profile1", "public_api_key").Return("", errors.New("not found"))
	mockKeyring.EXPECT().Get("atlascli_profile1", "private_api_key").Return("", errors.New("not found"))

	store := NewSecureStoreWithClient(profileNames, secureProperties, mockKeyring)

	t.Run("Save successful operations", func(t *testing.T) {
		// Add some pending operations
		store.Set("profile1", "public_api_key", "new_public_key")
		store.Set("profile1", "private_api_key", "new_private_key")
		store.DeleteKey("profile1", "old_key")

		assert.Len(t, store.pendingOps, 3)

		// Setup expectations for Save operation
		mockKeyring.EXPECT().Set("atlascli_profile1", "public_api_key", "new_public_key").Return(nil)
		mockKeyring.EXPECT().Set("atlascli_profile1", "private_api_key", "new_private_key").Return(nil)
		mockKeyring.EXPECT().Delete("atlascli_profile1", "old_key").Return(nil)

		err := store.Save()
		require.NoError(t, err)

		// Verify pending operations are cleared
		assert.Empty(t, store.pendingOps)
	})

	t.Run("Save with error", func(t *testing.T) {
		// Add a pending operation
		store.Set("profile1", "public_api_key", "failing_key")

		// Mock an error
		mockKeyring.EXPECT().Set("atlascli_profile1", "public_api_key", "failing_key").Return(errors.New("keyring error"))

		err := store.Save()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "keyring error")

		// Verify pending operations are NOT cleared on error
		assert.Len(t, store.pendingOps, 1)
	})

	t.Run("Save delete profile operation", func(t *testing.T) {
		// Clear any previous operations
		store.pendingOps = []pendingOperation{}

		store.DeleteProfile("test_profile")

		// Mock successful delete all
		mockKeyring.EXPECT().DeleteAll("atlascli_test_profile").Return(nil)

		err := store.Save()
		require.NoError(t, err)

		// Verify pending operations are cleared
		assert.Empty(t, store.pendingOps)
	})
}

func TestKeyringStore_Available(t *testing.T) {
	t.Parallel()

	t.Run("Available when keyring works", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockKeyring := NewMockKeyringClient(ctrl)
		// Mock successful get call
		mockKeyring.EXPECT().Get("atlascli_default", "demo_secret").Return("my_demo_secret", nil)

		store := NewSecureStoreWithClient([]string{"default"}, []string{"demo_secret"}, mockKeyring)
		assert.True(t, store.Available())
	})

	t.Run("Available when keyring works, but no profiles", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockKeyring := NewMockKeyringClient(ctrl)
		// Mock successful get call
		mockKeyring.EXPECT().Get("atlascli_default", "test").Return("test", nil)

		store := NewSecureStoreWithClient([]string{}, []string{"demo_secret"}, mockKeyring)
		assert.True(t, store.Available())
	})

	t.Run("Not available when keyring fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockKeyring := NewMockKeyringClient(ctrl)
		// Mock failed get call
		mockKeyring.EXPECT().Get("atlascli_default", "test").Return("", errors.New("keyring error"))

		store := NewSecureStoreWithClient([]string{}, []string{}, mockKeyring)
		assert.False(t, store.Available())
	})
}
