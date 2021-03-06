// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package native

import (
	"reflect"
	"testing"

	"github.com/go-vela/server/database"

	"github.com/go-vela/types/library"
)

func TestNative_List(t *testing.T) {
	// setup types
	sOne := new(library.Secret)
	sOne.SetID(1)
	sOne.SetOrg("foo")
	sOne.SetRepo("bar")
	sOne.SetTeam("")
	sOne.SetName("baz")
	sOne.SetValue("foob")
	sOne.SetType("repo")
	sOne.SetImages([]string{"foo", "bar"})
	sOne.SetEvents([]string{"foo", "bar"})
	sOne.SetAllowCommand(false)

	sTwo := new(library.Secret)
	sTwo.SetID(2)
	sTwo.SetOrg("foo")
	sTwo.SetRepo("bar")
	sTwo.SetTeam("")
	sTwo.SetName("foob")
	sTwo.SetValue("baz")
	sTwo.SetType("repo")
	sTwo.SetImages([]string{"foo", "bar"})
	sTwo.SetEvents([]string{"foo", "bar"})
	sTwo.SetAllowCommand(false)

	// nolint: gosec // ignore false positive
	passphrase := "C639A572E14D5075C526FDDD43E4ECF6"

	want := []*library.Secret{sTwo, sOne}

	// setup database
	d, _ := database.NewTest()

	defer func() {
		d.Database.Exec("delete from secrets;")
		d.Database.Close()
	}()

	// run test
	s, err := New(d, passphrase)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	_ = s.Create("repo", "foo", "bar", sOne)

	// vOne, _ := decrypt([]byte(sOne.GetValue()), passphrase)
	// sOne.Value = &vOne

	_ = s.Create("repo", "foo", "bar", sTwo)

	// vTwo, _ := decrypt([]byte(sTwo.GetValue()), passphrase)
	// sTwo.Value = &vTwo

	got, err := s.List("repo", "foo", "bar", 1, 10)
	if err != nil {
		t.Errorf("List returned err: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("List is %v, want %v", got, want)
	}
}

func TestNative_List_Invalid(t *testing.T) {
	// setup database
	d, _ := database.NewTest()
	d.Database.Close()

	// nolint: gosec // ignore false positive
	passphrase := "C639A572E14D5075C526FDDD43E4ECF6"

	// run test
	s, err := New(d, passphrase)
	if err != nil {
		t.Errorf("New returned err: %v", err)
	}

	got, err := s.List("repo", "foo", "bar", 1, 10)
	if err == nil {
		t.Errorf("List should have returned err")
	}

	if got != nil {
		t.Errorf("List is %v, want nil", got)
	}
}
