package models

import (
	"testing"
)

func TestBuildMenu(t *testing.T) {
	m := DefaultMenuList
	m.BuildMenu(0, nil)

	if len(m.menus) <= 0 {
		t.Errorf("expected 1 menu, got %v", len(m.menus))
	}
}

/*
func TestNewMenu(t *testing.T) {
	title := "learn Go"
	menu := newMenuOrFatal(t, title)
	if menu.Name != title {
		t.Errorf("expected title %q, got %q", title, menu.Name)
	}
}

func TestNewMenuEmptyMenuName(t *testing.T) {
	_, err := NewMenu("")
	if err == nil {
		t.Errorf("expected 'empty MenuName' error, got nil")
	}
}

func TestSaveMenuAndRetrieve(t *testing.T) {
	menu := newMenuOrFatal(t, "learn Go")

	m := NewMenuManager()
	m.Save(menu)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 menu, got %v", len(all))
	}
}

func TestSaveModifyAndRetrieveMenu(t *testing.T) {
	menu := newMenuOrFatal(t, "learn Go")
	m := NewMenuManager()
	m.Save(menu)

}

func TestSaveTwiceAndRetrieveMenu(t *testing.T) {
	menu := newMenuOrFatal(t, "learn Go")
	m := NewMenuManager()
	m.Save(menu)
	m.Save(menu)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 menu, got %v", len(all))
	}
}

func TestSaveAndFindMenu(t *testing.T) {
	menu := newMenuOrFatal(t, "learn Go")
	m := NewMenuManager()
	m.Save(menu)

	nt, ok := m.Find(menu.ID)
	if !ok {
		t.Errorf("didn't find MenuID")
	}
	if menu.ID != nt.ID {
		t.Errorf("")
	}

}
func TestInitMenuManager(t *testing.T) {
	m := NewMenuManager()
	m.init()
	if len(m.All()) != 3 {
		t.Errorf("expected 3 menu ,got %v", len(m.All()))
	}
}
*/
