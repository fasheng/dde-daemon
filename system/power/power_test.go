/**
 * Copyright (C) 2016 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package power

import (
	. "github.com/smartystreets/goconvey/convey"
	"pkg.deepin.io/lib/dbus"
	"testing"
)

func Test_getValidName(t *testing.T) {
	Convey("getValidName", t, func() {
		names := []string{"BAT0", "test.t", "test:t", "test-t", "test.1:2-3.4:5-6"}
		for _, name := range names {
			path := dbus.ObjectPath("/battery_" + getValidName(name))
			t.Log(path)
			So(path.IsValid(), ShouldBeTrue)
		}
	})
}
