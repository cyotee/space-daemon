// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	vault "github.com/FleekHQ/space-daemon/core/vault"
	mock "github.com/stretchr/testify/mock"
)

// Vault is an autogenerated mock type for the Vault type
type Vault struct {
	mock.Mock
}

// Retrieve provides a mock function with given fields: uuid, passphrase
func (_m *Vault) Retrieve(uuid string, passphrase string) ([]vault.VaultItem, error) {
	ret := _m.Called(uuid, passphrase)

	var r0 []vault.VaultItem
	if rf, ok := ret.Get(0).(func(string, string) []vault.VaultItem); ok {
		r0 = rf(uuid, passphrase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vault.VaultItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(uuid, passphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: uuid, passphrase, backupType, apiToken, items
func (_m *Vault) Store(uuid string, passphrase string, backupType string, apiToken string, items []vault.VaultItem) (*vault.StoredVault, error) {
	ret := _m.Called(uuid, passphrase, backupType, apiToken, items)

	var r0 *vault.StoredVault
	if rf, ok := ret.Get(0).(func(string, string, string, string, []vault.VaultItem) *vault.StoredVault); ok {
		r0 = rf(uuid, passphrase, backupType, apiToken, items)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vault.StoredVault)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, []vault.VaultItem) error); ok {
		r1 = rf(uuid, passphrase, backupType, apiToken, items)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
