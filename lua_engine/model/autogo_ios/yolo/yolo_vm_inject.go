package yolo

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// YoloModule 是 go-lua-vm 迁移后的模块壳。
type YoloModule struct{}

func New() *YoloModule { return &YoloModule{} }

func (m *YoloModule) Name() string { return "yolo" }

func (m *YoloModule) IsAvailable() bool { return true }

func (m *YoloModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &YoloModule{} }
