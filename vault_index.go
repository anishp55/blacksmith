package main

import (
	"encoding/json"
)

type VaultIndex struct {
	parent *Vault
	path   string
	Data   map[string]interface{}
}

func (v *Vault) GetIndex(path string) (*VaultIndex, error) {
	Data, exists, err := v.Get(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		Data = make(map[string]interface{})
	}
	return &VaultIndex{
		parent: v,
		path:   path,
		Data:   Data,
	}, nil
}

func (vi *VaultIndex) Save() error {
	return vi.parent.Put(vi.path, vi.Data)
}

func (vi *VaultIndex) JSON() string {
	b, err := json.Marshal(vi.Data)
	if err != nil {
		return `{"error":"json failed"}`
	}

	return string(b)
}