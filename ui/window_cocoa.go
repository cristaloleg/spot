//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeWindow = *gocoa.Window

func (w *Window) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		if w.ref != nil {
			w.ref.SetTitle(w.Title)
		}
	}

	if next.Width != w.Width || next.Height != w.Height {
		w.Width = next.Width
		w.Height = next.Height
	}

	if next.Resizable != w.Resizable {
		w.Resizable = next.Resizable
		if w.ref != nil {
			w.ref.SetAllowsResizing(w.Resizable)
		}
	}

	return true
}

func (w *Window) Mount(parent spot.Control) any {
	w.ref = gocoa.NewCenteredWindow(w.Title, w.Width, w.Height)
	w.ref.SetAllowsResizing(w.Resizable)

	w.ref.MakeKeyAndOrderFront()
	w.ref.AddDefaultQuitMenu()
	return w.ref
}
