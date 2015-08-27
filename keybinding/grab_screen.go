package keybinding

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"pkg.deepin.io/dde/daemon/keybinding/core"
	"pkg.deepin.io/lib/dbus"
)

var pressedKey string

func (m *Manager) doGrabScreen() error {
	xu, err := xgbutil.NewConn()
	if err != nil {
		return err
	}
	keybind.Initialize(xu)
	mousebind.Initialize(xu)

	err = grabKbdAndMouse(xu)
	if err != nil {
		return err
	}

	xevent.ButtonPressFun(
		func(x *xgbutil.XUtil, e xevent.ButtonPressEvent) {
			dbus.Emit(m, "KeyEvent", true, "")
			ungrabKbdAndMouse(xu)
			xevent.Quit(xu)
		}).Connect(xu, xu.RootWin())

	xevent.KeyPressFun(
		func(x *xgbutil.XUtil, e xevent.KeyPressEvent) {
			pressedKey, _ = core.FormatKeyEvent(e.State,
				int(e.Detail))
			dbus.Emit(m, "KeyEvent", true, pressedKey)
		}).Connect(xu, xu.RootWin())

	xevent.KeyReleaseFun(
		func(x *xgbutil.XUtil, e xevent.KeyReleaseEvent) {
			dbus.Emit(m, "KeyEvent", false, pressedKey)
			pressedKey = ""
			ungrabKbdAndMouse(xu)
			xevent.Quit(xu)
		}).Connect(xu, xu.RootWin())

	xevent.Main(xu)
	return nil
}

func grabKbdAndMouse(xu *xgbutil.XUtil) error {
	err := keybind.GrabKeyboard(xu, xu.RootWin())
	if err != nil {
		return err
	}

	// Ignore mouse grab error
	mousebind.GrabChecked(xu, xu.RootWin(), 0, 1, false)
	mousebind.GrabChecked(xu, xu.RootWin(), 0, 2, false)
	mousebind.GrabChecked(xu, xu.RootWin(), 0, 3, false)
	return nil
}

func ungrabKbdAndMouse(xu *xgbutil.XUtil) {
	keybind.UngrabKeyboard(xu)
	mousebind.Ungrab(xu, xu.RootWin(), 0, 1)
	mousebind.Ungrab(xu, xu.RootWin(), 0, 2)
	mousebind.Ungrab(xu, xu.RootWin(), 0, 3)
}