package menu

import "github.com/google/uuid"

type Manager struct {
	menuItems map[string]*Item
}

func (m *Manager) AddItem(Title string, ParentMenuID string) *Item {
	if m.menuItems == nil {
		m.menuItems = make(map[string]*Item)
	}

	menuItem := Item{
		ID:       uuid.New(),
		Title:    Title,
		Children: make(map[string]*Item),
	}

	if ParentMenuID != "" {
		m.addItemParent(&menuItem, ParentMenuID)
	}

	m.menuItems[menuItem.ID.String()] = &menuItem

	return &menuItem
}

func (m *Manager) GetItems() map[string]*Item {
	return m.menuItems
}

func (m *Manager) GetItem(MenuID string) *Item {
	return m.menuItems[MenuID]
}

func (m *Manager) addItemParent(MenuItem *Item, ParentMenuID string) {
	if ParentMenuItem, ok := m.menuItems[ParentMenuID]; ok {
		MenuItem.Parent = ParentMenuItem
		ParentMenuItem.Children[MenuItem.ID.String()] = MenuItem
	}
}

func (m Manager) Format(formatter Formatter) {
	formatter.Format(m.menuItems)
}
