package unsafe_models

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// UnsafeModules 包含 iOS AutoGo 非安全模块。
//
// 当前 iOS AutoGo Lua 模块尚未落地，保持空集合以避免误注入 Android 模块。
var UnsafeModules = []model.Module{}
