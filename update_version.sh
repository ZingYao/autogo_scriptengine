#!/bin/bash

# 设置版本号
NEW_VERSION="v0.0.13"

# 更新 examples 目录下的 go.mod 文件
echo "更新 examples 目录下的 go.mod 文件..."
find examples -name "go.mod" -type f -exec sed -i '' "s/@v[0-9]\.[0-9]*/@${NEW_VERSION}/g" {} \;

# 更新项目根目录的 go.mod 文件
echo "更新项目根目录的 go.mod 文件..."
sed -i '' "s/@v[0-9]\.[0-9]*/@${NEW_VERSION}/g" go.mod

# 更新 docs/README.md 文件
echo "更新 docs/README.md 文件..."
sed -i '' "s/@v[0-9]\.[0-9]*/@${NEW_VERSION}/g" docs/README.md

# 更新 docs/js_engine/README.md 文件
echo "更新 docs/js_engine/README.md 文件..."
sed -i '' "s/@v[0-9]\.[0-9]*/@${NEW_VERSION}/g" docs/js_engine/README.md

# 更新 docs/lua_engine/README.md 文件
echo "更新 docs/lua_engine/README.md 文件..."
sed -i '' "s/@v[0-9]\.[0-9]*/@${NEW_VERSION}/g" docs/lua_engine/README.md

# 提交更改
echo "提交更改..."
git add -A
git commit -m "更新版本为 ${NEW_VERSION}

- 更新所有 go.mod 文件
- 更新所有 README.md 文件"

# 创建并推送 tag
echo "创建并推送 tag ${NEW_VERSION}..."
git tag ${NEW_VERSION}
git push origin ${NEW_VERSION}

echo "完成！版本已更新为 ${NEW_VERSION}"
