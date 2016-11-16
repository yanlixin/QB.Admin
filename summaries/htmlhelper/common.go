package htmlhelper

type SelectItem struct {
	Value    string `json:"value"`
	Text     string `json:"text"`
	Selected bool   `json:"elected"`
}

func DefultSelectItem() []SelectItem {
	var result []SelectItem
	defaultItem := SelectItem{"", "--select--", false}
	result = append(result, defaultItem)
	return result
}
