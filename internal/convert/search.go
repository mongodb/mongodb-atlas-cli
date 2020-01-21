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

package convert

import "github.com/mongodb-labs/pcgc/cloudmanager"

// SearchProcesses find process index by the given function
// return -1 if not found
func SearchProcesses(a []*cloudmanager.Process, f func(*cloudmanager.Process) bool) int {
	for i, p := range a {
		if f(p) {
			return i
		}
	}
	return -1
}

// SearchMembers find member index of a replica set by the given function
// return -1 if not found
func SearchMembers(a []cloudmanager.Member, f func(cloudmanager.Member) bool) int {
	for i, m := range a {
		if f(m) {
			return i
		}
	}
	return -1
}

// SearchReplicaSets find a replica set index by the given function
// return -1 if not found
func SearchReplicaSets(a []*cloudmanager.ReplicaSet, f func(*cloudmanager.ReplicaSet) bool) int {
	for i, m := range a {
		if f(m) {
			return i
		}
	}
	return -1
}
