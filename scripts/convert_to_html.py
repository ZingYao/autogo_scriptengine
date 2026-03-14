#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Markdown 转 HTML 转换脚本
将项目中的所有 README.md 文件转换为 HTML 文档
"""

import os
import re
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

# HTML 模板
HTML_TEMPLATE = """<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{title}</title>
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
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            padding: 20px;
        }}
        
        .container {{
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 12px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.2);
            overflow: hidden;
        }}
        
        .header {{
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 40px;
            text-align: center;
        }}
        
        .header h1 {{
            font-size: 2.5em;
            margin-bottom: 10px;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
        }}
        
        .header .subtitle {{
            font-size: 1.2em;
            opacity: 0.9;
        }}
        
        .content {{
            padding: 40px;
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
        
        .toc {{
            background: #f8f9fa;
            border: 2px solid #667eea;
            border-radius: 8px;
            padding: 20px;
            margin: 20px 0;
        }}
        
        .toc h3 {{
            color: #667eea;
            margin-top: 0;
            margin-bottom: 15px;
        }}
        
        .toc ul {{
            list-style: none;
            padding-left: 0;
        }}
        
        .toc li {{
            margin-bottom: 8px;
        }}
        
        .toc a {{
            color: #667eea;
            font-weight: 500;
        }}
        
        .footer {{
            background: #2d2d2d;
            color: white;
            padding: 20px;
            text-align: center;
            margin-top: 40px;
        }}
        
        .footer a {{
            color: #667eea;
        }}
        
        @media (max-width: 768px) {{
            .content {{
                padding: 20px;
            }}
            
            .header h1 {{
                font-size: 1.8em;
            }}
            
            .header {{
                padding: 30px 20px;
            }}
        }}
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{title}</h1>
            <div class="subtitle">{subtitle}</div>
        </div>
        <div class="content">
            {content}
        </div>
        <div class="footer">
            <p>AutoGo ScriptEngine - JavaScript & Lua 脚本引擎</p>
        </div>
    </div>
</body>
</html>
"""

def convert_markdown_to_html(md_path, html_path):
    """将 Markdown 文件转换为 HTML 文件"""
    try:
        # 读取 Markdown 文件
        with open(md_path, 'r', encoding='utf-8') as f:
            md_content = f.read()
        
        # 转换 Markdown 到 HTML
        html_content = markdown.markdown(
            md_content,
            extensions=[
                'tables',
                'fenced_code',
                'codehilite',
                'toc',
                'nl2br',
                'sane_lists'
            ]
        )
        
        # 提取标题
        title_match = re.search(r'^#\s+(.+)$', md_content, re.MULTILINE)
        title = title_match.group(1) if title_match else '文档'
        
        # 生成副标题
        subtitle = f"来源: {md_path}"
        
        # 填充模板
        html = HTML_TEMPLATE.format(
            title=title,
            subtitle=subtitle,
            content=html_content
        )
        
        # 确保目标目录存在
        os.makedirs(os.path.dirname(html_path), exist_ok=True)
        
        # 写入 HTML 文件
        with open(html_path, 'w', encoding='utf-8') as f:
            f.write(html)
        
        print(f"✓ 已转换: {md_path} -> {html_path}")
        return True
    except Exception as e:
        print(f"✗ 转换失败: {md_path} - {str(e)}")
        return False

def find_all_readmes(root_dir):
    """查找所有 README.md 文件"""
    readmes = []
    for root, dirs, files in os.walk(root_dir):
        # 跳过隐藏目录和特定目录
        dirs[:] = [d for d in dirs if not d.startswith('.') and d not in ['docs', 'node_modules', '.git']]
        
        for file in files:
            if file.lower() == 'readme.md':
                readmes.append(os.path.join(root, file))
    return readmes

def main():
    """主函数"""
    print("=" * 60)
    print("Markdown 转 HTML 转换工具")
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
    success_count = 0
    fail_count = 0
    
    for md_path in readmes:
        # 计算相对路径
        rel_path = os.path.relpath(md_path, project_root)
        
        # 生成 HTML 文件路径
        html_path = os.path.join(
            project_root,
            'docs',
            rel_path.replace('.md', '.html')
        )
        
        if convert_markdown_to_html(md_path, html_path):
            success_count += 1
        else:
            fail_count += 1
    
    # 生成索引页面
    print("\n正在生成文档索引...")
    generate_index(project_root, readmes)
    
    # 输出统计信息
    print("\n" + "=" * 60)
    print("转换完成！")
    print(f"成功: {success_count} 个")
    print(f"失败: {fail_count} 个")
    print(f"总计: {len(readmes)} 个")
    print("=" * 60)
    print(f"\nHTML 文档已保存到: {os.path.join(project_root, 'docs')}")

def generate_index(project_root, readmes):
    """生成文档索引页面"""
    # 按目录分组
    grouped = {}
    for md_path in readmes:
        rel_path = os.path.relpath(md_path, project_root)
        dir_name = os.path.dirname(rel_path)
        
        if dir_name not in grouped:
            grouped[dir_name] = []
        grouped[dir_name].append(rel_path)
    
    # 生成 HTML 内容
    index_content = "<h2>文档索引</h2>\n"
    
    # 按目录排序
    for dir_name in sorted(grouped.keys()):
        index_content += f"<h3>{dir_name if dir_name != '.' else '根目录'}</h3>\n<ul>\n"
        
        for file_path in sorted(grouped[dir_name]):
            html_path = file_path.replace('.md', '.html')
            file_name = os.path.basename(file_path)
            index_content += f'  <li><a href="{html_path}">{file_name}</a></li>\n'
        
        index_content += "</ul>\n"
    
    # 生成索引页面
    index_html = HTML_TEMPLATE.format(
        title="AutoGo ScriptEngine 文档索引",
        subtitle="所有文档的导航页面",
        content=index_content
    )
    
    # 保存索引页面
    index_path = os.path.join(project_root, 'docs', 'index.html')
    os.makedirs(os.path.dirname(index_path), exist_ok=True)
    
    with open(index_path, 'w', encoding='utf-8') as f:
        f.write(index_html)
    
    print(f"✓ 已生成索引页面: {index_path}")

if __name__ == '__main__':
    main()
