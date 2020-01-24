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

package search

import (
	"github.com/mongodb-labs/pcgc/cloudmanager"
)

func StringInSlice(a []string, x string) bool {
	for _, b := range a {
		if b == x {
			return true
		}
	}
	return false
}

// Processes return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, Processes returns n and false
func Processes(a []*cloudmanager.Process, f func(*cloudmanager.Process) bool) (int, bool) {
	for i, p := range a {
		if f(p) {
			return i, true
		}
	}
	return len(a), false
}

// Members return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, Members returns n and false
func Members(a []cloudmanager.Member, f func(cloudmanager.Member) bool) (int, bool) {
	for i, m := range a {
		if f(m) {
			return i, true
		}
	}
	return len(a), false
}

// Members return the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// returns the first true index. If there is no such index, Members returns n and false
func ReplicaSets(a []*cloudmanager.ReplicaSet, f func(*cloudmanager.ReplicaSet) bool) (int, bool) {
	for i, m := range a {
		if f(m) {
			return i, true
		}
	}
	return len(a), false
}

// ClusterExists return true if a cluster exists for the given name
func ClusterExists(c *cloudmanager.AutomationConfig, name string) bool {
	_, found := ReplicaSets(c.ReplicaSets, func(r *cloudmanager.ReplicaSet) bool {
		return r.ID == name
	})

	return found
}
