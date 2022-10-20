package models

type ProtonArg map[string]any

func New(self map[string]any, data map[string]any) ProtonArg {
	m := make(map[string]any)
	m["self"] = self
	m["data"] = data
	return m
}

func (protonArg ProtonArg) GetSelf() map[string]any {
	return protonArg["self"].(map[string]any)
}

func (protonArg ProtonArg) GetData() map[string]any {
	return protonArg["data"].(map[string]any)
}
