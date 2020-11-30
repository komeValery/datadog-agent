// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probe

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/DataDog/datadog-agent/pkg/security/secl/eval"
)

func TestMkdirJSON(t *testing.T) {
	tr, err := NewTimeResolver()
	if err != nil {
		t.Fatal(err)
	}
	e := NewEvent(&Resolvers{TimeResolver: tr})
	e.Process = ProcessContext{
		Pid: 123,
		Tid: 456,
		UID: 8,
		GID: 9,
	}
	e.Mkdir = MkdirEvent{
		FileEvent: FileEvent{
			Inode:       33,
			PathnameStr: "/etc/passwd",
		},
		Mode: 0777,
	}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}

	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()

	var i interface{}
	err = d.Decode(&i)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAbsolutePath(t *testing.T) {
	model := &Model{}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "/var/log/*"}); err != nil {
		t.Fatalf("shouldn't return an error: %s", err)
	}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "~/apache/httpd.conf"}); err == nil {
		t.Fatal("should return an error")
	}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "../../../etc/apache/httpd.conf"}); err == nil {
		t.Fatal("should return an error")
	}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "/etc/apache/./httpd.conf"}); err == nil {
		t.Fatal("should return an error")
	}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "*/"}); err == nil {
		t.Fatal("should return an error")
	}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "/1/2/3/4/5/6/7/8/9/10/11/12/13/14/15/16"}); err == nil {
		t.Fatal("should return an error")
	}
	if err := model.ValidateField("open.filename", eval.FieldValue{Value: "f59226f52267c120c1accfe3d158aa2f201ff02f45692a1c574da29c07fb985ef59226f52267c120c1accfe3d158aa2f201ff02f45692a1c574da29c07fb985e"}); err == nil {
		t.Fatal("should return an error")
	}
}
