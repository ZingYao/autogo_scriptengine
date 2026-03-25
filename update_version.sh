#!/bin/bash

# 检查是否提供了版本号参数
if [ -z "$1" ]; then
    echo "错误：请提供版本号作为参数"
    echo "使用方法：./update_version.sh v0.0.14"
    exit 1
fi

NEW_VERSION="$1"

# 获取当前最新的 tag
CURRENT_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "")

# 检查当前 tag 是否符合预期格式（v0.0.x）
if [[ ! "$CURRENT_TAG" =~ ^v0\.0\.[0-9]+$ ]]; then
    echo "错误：当前 tag '$CURRENT_TAG' 不符合 v0.0.x 格式"
    echo "请确保当前 tag 是 v0.0.x 格式才能更新版本"
    exit 1
fi

# 提取当前版本号（去掉 v 前缀）
CURRENT_VERSION=${CURRENT_TAG#v}

# 比较版本号
if [ "$CURRENT_VERSION" = "$NEW_VERSION" ]; then
    echo "当前版本已经是 $NEW_VERSION，无需更新"
    exit 0
fi

# 查找 examples 目录下的 go.mod 文件中的包名
PACKAGE_NAME=""
for go_mod_file in examples/*/go.mod; do
    if [ -f "$go_mod_file" ]; then
        # 提取包名（如 github.com/ZingYao/autogo_scriptengine）
        PACKAGE_NAME=$(grep "^require" "$go_mod_file" | sed 's/require[[:space:]]*//g' | head -1)
        
        # 如果找到了包名，则使用它
        if [ -n "$PACKAGE_NAME" ]; then
            break
        fi
    fi
done

# 如果没有找到包名，则使用默认包名
if [ -z "$PACKAGE_NAME" ]; then
    PACKAGE_NAME="github.com/ZingYao/autogo_scriptengine"
fi

echo "检测到的包名：$PACKAGE_NAME"

# 检查当前 tag 是否包含目标包名
if [[ "$CURRENT_TAG" == *"$PACKAGE_NAME"* ]]; then
    echo "当前 tag '$CURRENT_TAG' 已包含包名 '$PACKAGE_NAME'，无需替换"
else
    # 替换 tag 中的包名
    NEW_TAG="${NEW_VERSION} (${PACKAGE_NAME})"
    echo "新的 tag 将是：$NEW_TAG"
    
    # 删除旧的 tag
    git tag -d "$CURRENT_TAG" 2>/dev/null || true
    
    # 创建新的 tag
    git tag "$NEW_TAG"
fi

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
- 更新所有 README.md 文件
- Tag: $NEW_TAG"

# 推送 tag
echo "推送 tag ${NEW_TAG}..."
git push origin "${NEW_TAG}"

echo "完成！版本已从 $CURRENT_TAG 更新为 $NEW_TAG"
