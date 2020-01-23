// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package store

import (
	"strings"
	"testing"

	"github.com/mongodb-labs/pcgc/cloudmanager"

	"github.com/10gen/mcli/internal/mocks"
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
		OpsManagerURL().
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
		OpsManagerURL().
		Return("ops_manager").
		Times(2)

	store, err := New(mockConfig)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}
	expected := "ops_manager/api/public/v1.0/"
	if store.baseURL.String() != expected {
		t.Errorf("store.baseURL = %s; want '%s'", expected, store.baseURL)
	}
}

func TestStore_apiPath(t *testing.T) {
	t.Run("ops manager", func(t *testing.T) {
		s := &Store{
			service: "ops-manager",
		}
		result := s.apiPath("localhost")
		if !strings.Contains(result, cloudmanager.APIPublicV1Path) {
			t.Errorf("apiPath() = %s; want '%s'", result, cloudmanager.APIPublicV1Path)
		}
	})
	t.Run("atlas", func(t *testing.T) {
		s := &Store{
			service: "cloud",
		}
		result := s.apiPath("localhost")
		if !strings.Contains(result, atlasAPIPath) {
			t.Errorf("apiPath() = %s; want '%s'", result, atlasAPIPath)
		}
	})
}
