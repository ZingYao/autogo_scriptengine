#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Markdown 转 HTML 转换脚本（带侧边栏版本）
将项目中的所有 README.md 文件转换为带左侧菜单的 HTML 文档
"""

import os
import re
import json
from pathlib import Path
import subprocess
import sys

# 检查是否安装了 markdown 库
try:
    import markdown
except ImportError:
    print("错误: 未安装 markdown 库")
    print("请运行: pip install markdown")
    sys.exit(1)

# HTML 模板（带侧边栏）
HTML_TEMPLATE = """<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AutoGo ScriptEngine 文档</title>
    <style>
        * {{
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }}
        
        body {{
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
            line-height: 1.6;
            color: #333;
            background: #f5f7fa;
            height: 100vh;
            overflow: hidden;
        }}
        
        .app-container {{
            display: flex;
            height: 100vh;
        }}
        
        /* 左侧边栏 */
        .sidebar {{
            width: 300px;
            background: linear-gradient(180deg, #667eea 0%, #764ba2 100%);
            color: white;
            overflow-y: auto;
            flex-shrink: 0;
            box-shadow: 2px 0 10px rgba(0,0,0,0.1);
        }}
        
        .sidebar-header {{
            padding: 30px 20px;
            text-align: center;
            border-bottom: 1px solid rgba(255,255,255,0.2);
        }}
        
        .sidebar-header h1 {{
            font-size: 1.8em;
            margin-bottom: 10px;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
        }}
        
        .sidebar-header .subtitle {{
            font-size: 0.9em;
            opacity: 0.9;
        }}
        
        .search-box {{
            padding: 15px 20px;
            border-bottom: 1px solid rgba(255,255,255,0.1);
        }}
        
        .search-box input {{
            width: 100%;
            padding: 10px 15px;
            border: none;
            border-radius: 20px;
            background: rgba(255,255,255,0.2);
            color: white;
            font-size: 0.9em;
            outline: none;
            transition: all 0.3s ease;
        }}
        
        .search-box input::placeholder {{
            color: rgba(255,255,255,0.6);
        }}
        
        .search-box input:focus {{
            background: rgba(255,255,255,0.3);
            box-shadow: 0 0 0 2px rgba(255,255,255,0.2);
        }}
        
        .search-results {{
            display: none;
            padding: 10px 20px;
            border-bottom: 1px solid rgba(255,255,255,0.1);
        }}
        
        .search-results.show {{
            display: block;
        }}
        
        .search-result-item {{
            padding: 8px 12px;
            cursor: pointer;
            border-radius: 6px;
            margin-bottom: 5px;
            transition: all 0.3s ease;
        }}
        
        .search-result-item:hover {{
            background: rgba(255,255,255,0.15);
        }}
        
        .search-result-item .title {{
            font-size: 0.9em;
            margin-bottom: 3px;
        }}
        
        .search-result-item .path {{
            font-size: 0.75em;
            opacity: 0.7;
        }}
        
        .sidebar-content {{
            padding: 20px 0;
        }}
        
        .sidebar-section {{
            margin-bottom: 20px;
        }}
        
        .sidebar-section-title {{
            padding: 10px 20px;
            font-size: 0.85em;
            text-transform: uppercase;
            letter-spacing: 1px;
            opacity: 0.7;
            font-weight: 600;
        }}
        
        .sidebar-menu {{
            list-style: none;
        }}
        
        .sidebar-menu-item {{
            padding: 12px 20px;
            cursor: pointer;
            transition: all 0.3s ease;
            border-left: 3px solid transparent;
            font-size: 0.95em;
        }}
        
        .sidebar-menu-item:hover {{
            background: rgba(255,255,255,0.1);
            border-left-color: rgba(255,255,255,0.5);
        }}
        
        .sidebar-menu-item.active {{
            background: rgba(255,255,255,0.2);
            border-left-color: white;
            font-weight: 600;
        }}
        
        .sidebar-menu-item .icon {{
            margin-right: 8px;
            opacity: 0.7;
        }}
        
        /* 右侧内容区 */
        .main-content {{
            flex: 1;
            overflow-y: auto;
            padding: 40px;
            position: relative;
        }}
        
        .content-wrapper {{
            max-width: 900px;
            margin: 0 auto;
            background: white;
            border-radius: 12px;
            box-shadow: 0 4px 20px rgba(0,0,0,0.1);
            padding: 50px;
        }}
        
        /* 悬浮目录 */
        .toc-float {{
            position: fixed;
            right: 20px;
            top: 100px;
            width: 250px;
            max-height: 70vh;
            overflow-y: auto;
            background: white;
            border-radius: 8px;
            box-shadow: 0 4px 20px rgba(0,0,0,0.15);
            padding: 15px;
            z-index: 1000;
            display: none;
        }}
        
        .toc-float.show {{
            display: block;
        }}
        
        .toc-float-header {{
            font-weight: 600;
            color: #667eea;
            margin-bottom: 10px;
            padding-bottom: 8px;
            border-bottom: 2px solid #667eea;
            font-size: 0.95em;
        }}
        
        .toc-float-item {{
            padding: 6px 10px;
            cursor: pointer;
            transition: all 0.3s ease;
            border-left: 2px solid transparent;
            font-size: 0.85em;
            line-height: 1.4;
        }}
        
        .toc-float-item:hover {{
            background: #f8f9fa;
            border-left-color: #667eea;
        }}
        
        .toc-float-item.level-1 {{
            font-weight: 600;
            color: #667eea;
        }}
        
        .toc-float-item.level-2 {{
            padding-left: 20px;
            font-weight: 500;
            color: #555;
        }}
        
        .toc-float-item.level-3 {{
            padding-left: 35px;
            font-weight: 400;
            color: #666;
            font-size: 0.9em;
        }}
        
        .toc-float-item.level-4 {{
            padding-left: 50px;
            font-weight: 400;
            color: #777;
            font-size: 0.85em;
        }}
        
        .toc-float::-webkit-scrollbar {{
            width: 6px;
        }}
        
        .toc-float::-webkit-scrollbar-track {{
            background: #f1f1f1;
        }}
        
        .toc-float::-webkit-scrollbar-thumb {{
            background: #667eea;
            border-radius: 3px;
        }}
        
        .toc-float::-webkit-scrollbar-thumb:hover {{
            background: #764ba2;
        }}
        
        .content-header {{
            margin-bottom: 30px;
            padding-bottom: 20px;
            border-bottom: 2px solid #667eea;
        }}
        
        .content-header h1 {{
            color: #667eea;
            font-size: 2.2em;
            margin-bottom: 10px;
        }}
        
        .content-header .breadcrumb {{
            color: #999;
            font-size: 0.9em;
        }}
        
        h1 {{
            color: #667eea;
            border-bottom: 3px solid #667eea;
            padding-bottom: 10px;
            margin-top: 40px;
            margin-bottom: 20px;
        }}
        
        h2 {{
            color: #764ba2;
            border-bottom: 2px solid #764ba2;
            padding-bottom: 8px;
            margin-top: 30px;
            margin-bottom: 15px;
        }}
        
        h3 {{
            color: #555;
            margin-top: 25px;
            margin-bottom: 12px;
        }}
        
        h4 {{
            color: #666;
            margin-top: 20px;
            margin-bottom: 10px;
        }}
        
        p {{
            margin-bottom: 15px;
            line-height: 1.8;
        }}
        
        a {{
            color: #667eea;
            text-decoration: none;
            transition: all 0.3s ease;
        }}
        
        a:hover {{
            color: #764ba2;
            text-decoration: underline;
        }}
        
        code {{
            background: #f4f4f4;
            padding: 2px 6px;
            border-radius: 4px;
            font-family: "Monaco", "Menlo", "Ubuntu Mono", monospace;
            font-size: 0.9em;
            color: #e83e8c;
        }}
        
        pre {{
            background: #2d2d2d;
            color: #f8f8f2;
            padding: 20px;
            border-radius: 8px;
            overflow-x: auto;
            margin: 20px 0;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        }}
        
        pre code {{
            background: none;
            color: inherit;
            padding: 0;
            font-size: 0.95em;
        }}
        
        blockquote {{
            border-left: 4px solid #667eea;
            padding-left: 20px;
            margin: 20px 0;
            color: #666;
            background: #f8f9fa;
            padding: 15px 20px;
            border-radius: 0 8px 8px 0;
        }}
        
        ul, ol {{
            margin-bottom: 15px;
            padding-left: 30px;
        }}
        
        li {{
            margin-bottom: 8px;
        }}
        
        table {{
            width: 100%;
            border-collapse: collapse;
            margin: 20px 0;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }}
        
        th {{
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 12px;
            text-align: left;
            font-weight: 600;
        }}
        
        td {{
            padding: 12px;
            border: 1px solid #ddd;
        }}
        
        tr:nth-child(even) {{
            background: #f8f9fa;
        }}
        
        tr:hover {{
            background: #e9ecef;
        }}
        
        hr {{
            border: none;
            height: 2px;
            background: linear-gradient(90deg, #667eea, #764ba2);
            margin: 40px 0;
        }}
        
        img {{
            max-width: 100%;
            height: auto;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }}
        
        /* 滚动条样式 */
        .sidebar::-webkit-scrollbar,
        .main-content::-webkit-scrollbar {{
            width: 8px;
        }}
        
        .sidebar::-webkit-scrollbar-track,
        .main-content::-webkit-scrollbar-track {{
            background: rgba(0,0,0,0.1);
        }}
        
        .sidebar::-webkit-scrollbar-thumb,
        .main-content::-webkit-scrollbar-thumb {{
            background: rgba(255,255,255,0.3);
            border-radius: 4px;
        }}
        
        .sidebar::-webkit-scrollbar-thumb:hover,
        .main-content::-webkit-scrollbar-thumb:hover {{
            background: rgba(255,255,255,0.5);
        }}
        
        /* 移动端适配 */
        @media (max-width: 768px) {{
            .app-container {{
                flex-direction: column;
            }}
            
            .sidebar {{
                width: 100%;
                height: auto;
                max-height: 40vh;
            }}
            
            .content-wrapper {{
                padding: 20px;
            }}
            
            .content-header h1 {{
                font-size: 1.5em;
            }}
            
            .toc-float {{
                display: none !important;
            }}
        }}
    </style>
</head>
<body>
    <div class="app-container">
        <!-- 左侧边栏 -->
        <div class="sidebar">
            <div class="sidebar-header">
                <h1>📚 AutoGo ScriptEngine</h1>
                <div class="subtitle">JavaScript & Lua 脚本引擎</div>
            </div>
            <div class="search-box">
                <input type="text" id="search-input" placeholder="🔍 搜索文档..." oninput="handleSearch(this.value)">
            </div>
            <div class="search-results" id="search-results"></div>
            <div class="sidebar-content">
                {sidebar_content}
            </div>
        </div>
        
        <!-- 右侧内容区 -->
        <div class="main-content">
            <div class="content-wrapper">
                <div id="content-area">
                    <!-- 内容将通过 JavaScript 动态加载 -->
                </div>
            </div>
            <!-- 悬浮目录 -->
            <div class="toc-float" id="toc-float">
                <div class="toc-float-header">📋 目录</div>
                <div id="toc-float-content">
                    <!-- 目录内容将通过 JavaScript 动态生成 -->
                </div>
            </div>
        </div>
    </div>
    
    <script>
        // 文档数据
        const documents = {documents_data};
        
        // 当前选中的文档
        let currentDoc = null;
        
        // 初始化
        document.addEventListener('DOMContentLoaded', function() {{
            // 默认加载第一个文档
            const firstDoc = Object.keys(documents)[0];
            if (firstDoc) {{
                loadDocument(firstDoc);
            }}
        }});
        
        // 加载文档
        function loadDocument(docId, headingId = null) {{
            const doc = documents[docId];
            if (!doc) return;
            
            // 更新侧边栏选中状态
            document.querySelectorAll('.sidebar-menu-item').forEach(item => {{
                item.classList.remove('active');
            }});
            
            const activeItem = document.querySelector(`[data-doc-id="${{docId}}"]`);
            if (activeItem) {{
                activeItem.classList.add('active');
            }}
            
            // 如果指定了标题 ID，滚动到该位置
            let content = doc.content;
            if (headingId) {{
                // 等待内容加载后滚动
                setTimeout(() => {{
                    const headingElement = document.getElementById(headingId);
                    if (headingElement) {{
                        headingElement.scrollIntoView({{ behavior: 'smooth', block: 'start' }});
                    }}
                }}, 100);
            }}
            
            // 更新内容区
            const contentArea = document.getElementById('content-area');
            contentArea.innerHTML = `
                <div class="content-header">
                    <h1>${{doc.title}}</h1>
                    <div class="breadcrumb">${{doc.breadcrumb}}</div>
                </div>
                ${{content}}
            `;
            
            // 为所有标题添加 ID，方便跳转
            addHeadingIds();
            
            // 生成右侧悬浮目录
            generateFloatingTOC(doc.headings);
            
            // 滚动到顶部
            if (!headingId) {{
                document.querySelector('.main-content').scrollTop = 0;
            }}
            
            currentDoc = docId;
        }}
        
        // 为标题添加 ID
        function addHeadingIds() {{
            const headings = document.querySelectorAll('#content-area h1, #content-area h2, #content-area h3, #content-area h4');
            headings.forEach((heading, index) => {{
                const text = heading.textContent.trim();
                const id = 'heading-' + index;
                heading.id = id;
            }});
        }}
        
        // 生成右侧悬浮目录
        function generateFloatingTOC(headings) {{
            const tocFloat = document.getElementById('toc-float');
            const tocContent = document.getElementById('toc-float-content');
            
            if (!headings || headings.length === 0) {{
                tocFloat.classList.remove('show');
                return;
            }}
            
            // 生成目录 HTML
            let tocHTML = '';
            headings.forEach(heading => {{
                const level = heading.level;
                const text = heading.text;
                const id = heading.id;
                
                tocHTML += `<div class="toc-float-item level-${{level}}" onclick="scrollToHeading('${{id}}')">${{text}}</div>`;
            }});
            
            tocContent.innerHTML = tocHTML;
            tocFloat.classList.add('show');
        }}
        
        // 滚动到指定标题
        function scrollToHeading(headingId) {{
            const headingElement = document.getElementById(headingId);
            if (headingElement) {{
                headingElement.scrollIntoView({{ behavior: 'smooth', block: 'start' }});
            }}
        }}
        
        // 搜索功能
        function handleSearch(query) {{
            const searchResults = document.getElementById('search-results');
            
            if (!query.trim()) {{
                searchResults.classList.remove('show');
                return;
            }}
            
            query = query.toLowerCase();
            const results = [];
            
            // 搜索文档标题和内容
            for (const [docId, doc] of Object.entries(documents)) {{
                // 搜索标题
                if (doc.title.toLowerCase().includes(query)) {{
                    results.push({{
                        docId: docId,
                        title: doc.title,
                        path: doc.breadcrumb,
                        matchType: 'title'
                    }});
                }}
                // 搜索内容
                else if (doc.content.toLowerCase().includes(query)) {{
                    results.push({{
                        docId: docId,
                        title: doc.title,
                        path: doc.breadcrumb,
                        matchType: 'content'
                    }});
                }}
            }}
            
            // 显示搜索结果
            if (results.length > 0) {{
                searchResults.innerHTML = results.map(result => `
                    <div class="search-result-item" onclick="loadDocument('${{result.docId}}')">
                        <div class="title">${{result.title}}</div>
                        <div class="path">${{result.path}}</div>
                    </div>
                `).join('');
                searchResults.classList.add('show');
            }} else {{
                searchResults.innerHTML = '<div style="padding: 10px 20px; opacity: 0.7;">未找到相关文档</div>';
                searchResults.classList.add('show');
            }}
        }}
    </script>
</body>
</html>
"""

def convert_markdown_to_html(md_path):
    """将 Markdown 文件转换为 HTML 内容"""
    try:
        # 读取 Markdown 文件
        with open(md_path, 'r', encoding='utf-8') as f:
            md_content = f.read()
        
        # 提取标题
        title_match = re.search(r'^#\s+(.+)$', md_content, re.MULTILINE)
        title = title_match.group(1) if title_match else '文档'
        
        # 提取所有标题（一级、二级、三级等）
        headings = []
        heading_pattern = r'^(#{1,4})\s+(.+)$'
        for match in re.finditer(heading_pattern, md_content, re.MULTILINE):
            level = len(match.group(1))
            text = match.group(2).strip()
            headings.append({
                'level': level,
                'text': text,
                'id': f"heading-{len(headings)}"
            })
        
        # 转换 Markdown 到 HTML
        html_content = markdown.markdown(
            md_content,
            extensions=[
                'tables',
                'fenced_code',
                'codehilite',
                'nl2br',
                'sane_lists'
            ]
        )
        
        return {
            'title': title,
            'content': html_content,
            'headings': headings
        }
    except Exception as e:
        print(f"✗ 转换失败: {md_path} - {str(e)}")
        return None

def find_all_readmes(root_dir):
    """查找所有 README.md 文件"""
    readmes = []
    for root, dirs, files in os.walk(root_dir):
        # 跳过隐藏目录和特定目录
        dirs[:] = [d for d in dirs if not d.startswith('.') and d not in ['docs', 'node_modules', '.git', 'scripts']]
        
        for file in files:
            if file.lower() == 'readme.md':
                readmes.append(os.path.join(root, file))
    return readmes

def generate_sidebar(project_root, readmes, documents_data):
    """生成侧边栏内容"""
    # 按目录分组
    grouped = {}
    for md_path in readmes:
        rel_path = os.path.relpath(md_path, project_root)
        dir_name = os.path.dirname(rel_path)
        
        if dir_name not in grouped:
            grouped[dir_name] = []
        grouped[dir_name].append(rel_path)
    
    # 生成侧边栏 HTML
    sidebar_html = ''
    
    # 定义图标
    icons = {
        '.': '🏠',
        'js_engine': '⚡',
        'lua_engine': '🌙',
        'js_engine/model': '📦',
        'lua_engine/model': '📦'
    }
    
    # 按目录排序
    sorted_dirs = sorted(grouped.keys())
    
    # 根目录优先
    if '.' in sorted_dirs:
        sorted_dirs.remove('.')
        sorted_dirs.insert(0, '.')
    
    for dir_name in sorted_dirs:
        icon = icons.get(dir_name, '📄')
        display_name = dir_name if dir_name != '.' else '根目录'
        
        sidebar_html += f'<div class="sidebar-section">\n'
        sidebar_html += f'<div class="sidebar-section-title">{icon} {display_name}</div>\n'
        sidebar_html += f'<ul class="sidebar-menu">\n'
        
        for file_path in sorted(grouped[dir_name]):
            # 生成文档 ID
            doc_id = file_path.replace('/', '_').replace('.md', '')
            
            # 获取文档数据
            doc_data = documents_data.get(doc_id, {})
            doc_title = doc_data.get('title', '文档')
            
            # 生成菜单项
            sidebar_html += f'  <li class="sidebar-menu-item" data-doc-id="{doc_id}" onclick="loadDocument(\'{doc_id}\')">\n'
            sidebar_html += f'    <span class="icon">📄</span>{doc_title}\n'
            sidebar_html += f'  </li>\n'
        
        sidebar_html += f'</ul>\n'
        sidebar_html += f'</div>\n'
    
    return sidebar_html

def generate_documents_data(project_root, readmes):
    """生成文档数据"""
    documents = {}
    
    for md_path in readmes:
        rel_path = os.path.relpath(md_path, project_root)
        doc_id = rel_path.replace('/', '_').replace('.md', '')
        
        # 转换 Markdown
        doc_data = convert_markdown_to_html(md_path)
        if doc_data:
            documents[doc_id] = {
                'title': doc_data['title'],
                'breadcrumb': rel_path,
                'content': doc_data['content'],
                'headings': doc_data['headings']
            }
    
    return documents

def main():
    """主函数"""
    print("=" * 60)
    print("Markdown 转 HTML 转换工具（带侧边栏版本）")
    print("=" * 60)
    
    # 获取项目根目录
    script_dir = os.path.dirname(os.path.abspath(__file__))
    project_root = os.path.dirname(script_dir)
    
    print(f"\n项目根目录: {project_root}")
    
    # 查找所有 README.md 文件
    print("\n正在查找所有 README.md 文件...")
    readmes = find_all_readmes(project_root)
    print(f"找到 {len(readmes)} 个 README.md 文件\n")
    
    # 转换所有文件
    print("正在转换文档...")
    documents_data = generate_documents_data(project_root, readmes)
    print(f"成功转换 {len(documents_data)} 个文档\n")
    
    # 生成侧边栏
    print("正在生成侧边栏...")
    sidebar_content = generate_sidebar(project_root, readmes, documents_data)
    
    # 生成 HTML
    print("正在生成 HTML 文件...")
    html = HTML_TEMPLATE.format(
        sidebar_content=sidebar_content,
        documents_data=json.dumps(documents_data, ensure_ascii=False)
    )
    
    # 保存 HTML 文件
    docs_dir = os.path.join(project_root, 'docs')
    os.makedirs(docs_dir, exist_ok=True)
    
    output_path = os.path.join(docs_dir, 'index.html')
    with open(output_path, 'w', encoding='utf-8') as f:
        f.write(html)
    
    print(f"✓ 已生成: {output_path}")
    
    # 输出统计信息
    print("\n" + "=" * 60)
    print("转换完成！")
    print(f"文档数量: {len(documents_data)} 个")
    print("=" * 60)
    print(f"\nHTML 文档已保存到: {output_path}")
    print("\n使用浏览器打开该文件即可查看文档！")

if __name__ == '__main__':
    main()
