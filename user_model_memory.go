package main

import (
	"errors"
	"sync"
)

// 定义错误
var (
	ErrNotFound = errors.New("not found")
)

// NewUserMemoryModel 创建用户管理内存存储实例
func NewUserMemoryModel() *UserMemoryModel {
	return &UserMemoryModel{
		data:  make([]*User, 0, 16),
		index: make(map[string]int),
		lock:  new(sync.RWMutex),
	}
}

// UserMemoryModel 用户管理内存存储
type UserMemoryModel struct {
	data  []*User
	index map[string]int
	lock  *sync.RWMutex
}

// Query 查询用户数据
func (a *UserMemoryModel) Query(params UserQueryParam) ([]*User, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	if params.UserName == "" {
		return a.data, nil
	}

	var items []*User
	for _, item := range a.data {
		if item.UserName == params.UserName {
			items = append(items, item)
		}
	}

	return items, nil
}

// Get 查询指定用户数据
func (a *UserMemoryModel) Get(recordID string) (*User, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	i, ok := a.index[recordID]
	if !ok {
		return nil, ErrNotFound
	}

	return a.data[i], nil
}

// Create 创建用户数据
func (a *UserMemoryModel) Create(item User) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.data = append(a.data, &item)
	a.index[item.RecordID] = len(a.data) - 1
	return nil
}

// Update 更新用户数据
func (a *UserMemoryModel) Update(recordID string, item User) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	i, ok := a.index[recordID]
	if !ok {
		return ErrNotFound
	}
	item.RecordID = recordID
	a.data[i] = &item

	return nil
}

// Delete 删除用户数据
func (a *UserMemoryModel) Delete(recordID string) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	i, ok := a.index[recordID]
	if !ok {
		return ErrNotFound
	}

	delete(a.index, recordID)
	a.data = append(a.data[:i], a.data[i+1:]...)

	for i, item := range a.data {
		a.index[item.RecordID] = i
	}

	return nil
}
