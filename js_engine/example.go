package js_engine

import (
	"fmt"
	"os"
)

func main() {
	// 初始化 JavaScript 引擎
	engine := GetEngine()
	defer Close()

	// 生成 JavaScript API 文档
	docGen := NewDocumentationGenerator()

	// 生成 JavaScript 文档
	fmt.Println("生成 JavaScript API 文档...")
	err := docGen.SaveJSDocumentation("js_api.js")
	if err != nil {
		fmt.Printf("生成 JavaScript 文档失败: %v\n", err)
		return
	}
	fmt.Println("✓ JavaScript 文档已生成: js_api.js")

	// 生成 Markdown 文档
	fmt.Println("生成 Markdown API 文档...")
	err = docGen.SaveMarkdownDocumentation("js_api.md")
	if err != nil {
		fmt.Printf("生成 Markdown 文档失败: %v\n", err)
		return
	}
	fmt.Println("✓ Markdown 文档已生成: js_api.md")

	// 导出方法列表为 JSON
	fmt.Println("导出方法列表...")
	registry := engine.GetRegistry()
	jsonData, err := registry.ExportMethodsJSON()
	if err != nil {
		fmt.Printf("导出方法列表失败: %v\n", err)
		return
	}

	err = os.WriteFile("methods.json", []byte(jsonData), 0644)
	if err != nil {
		fmt.Printf("写入方法列表失败: %v\n", err)
		return
	}
	fmt.Println("✓ 方法列表已导出: methods.json")

	// 导出方法列表为 JavaScript 对象
	fmt.Println("导出方法列表为 JavaScript 对象...")
	jsObject := registry.ExportMethodsJSObject()
	err = os.WriteFile("methods.js", []byte(jsObject), 0644)
	if err != nil {
		fmt.Printf("写入 JavaScript 对象失败: %v\n", err)
		return
	}
	fmt.Println("✓ JavaScript 对象已导出: methods.js")

	fmt.Println("\n所有文档生成完成!")

	// ========== 协程模块示例 ==========
	fmt.Println("\n========== 协程模块示例 ==========")

	// 示例1: 基础协程使用
	fmt.Println("\n--- 示例1: 基础协程使用 ---")
	err = ExecuteString(`
		console.log("启动基础协程示例");
		
		// 启动一个协程
		const coroutineId = coroutine.launch(function() {
			console.log("协程开始执行");
			coroutine.sleep(1000);
			console.log("协程执行完成");
			return "协程结果";
		}, "basicCoroutine", 0);
		
		console.log("协程ID: " + coroutineId);
		
		// 等待协程完成
		coroutine.sleep(1500);
		
		// 获取协程信息
		const info = coroutine.getCoroutineInfo(coroutineId);
		console.log("协程状态: " + info.state);
		console.log("协程运行时长: " + info.duration + "ms");
	`)
	if err != nil {
		fmt.Printf("执行协程示例1失败: %v\n", err)
	}

	// 示例2: 并发执行多个任务
	fmt.Println("\n--- 示例2: 并发执行多个任务 ---")
	err = ExecuteString(`
		console.log("启动并发任务示例");
		
		// 启动多个协程
		const task1 = coroutine.launch(function() {
			console.log("任务1开始");
			coroutine.sleep(1000);
			console.log("任务1完成");
			return "任务1结果";
		}, "task1", 1);
		
		const task2 = coroutine.launch(function() {
			console.log("任务2开始");
			coroutine.sleep(1500);
			console.log("任务2完成");
			return "任务2结果";
		}, "task2", 2);
		
		const task3 = coroutine.launch(function() {
			console.log("任务3开始");
			coroutine.sleep(2000);
			console.log("任务3完成");
			return "任务3结果";
		}, "task3", 3);
		
		console.log("所有任务已启动");
		
		// 获取活跃协程数量
		const activeCount = coroutine.getActiveCoroutines();
		console.log("活跃协程数量: " + activeCount);
		
		// 等待所有任务完成
		coroutine.sleep(2500);
		
		// 获取统计信息
		const stats = coroutine.getStats();
		console.log("总任务数: " + stats.totalTasks);
		console.log("已完成: " + stats.completed);
	`)
	if err != nil {
		fmt.Printf("执行协程示例2失败: %v\n", err)
	}

	// 示例3: 使用协程池
	fmt.Println("\n--- 示例3: 使用协程池 ---")
	err = ExecuteString(`
		console.log("启动协程池示例");
		
		// 创建协程池
		const poolName = coroutine.createPool("examplePool", 3, 10);
		console.log("创建协程池: " + poolName);
		
		// 提交多个任务到协程池
		for (let i = 0; i < 5; i++) {
			const taskNum = i + 1;
			const success = coroutine.submitToPool("examplePool", function() {
				console.log("处理任务 " + taskNum);
				coroutine.sleep(500);
				console.log("任务 " + taskNum + " 完成");
				return "任务" + taskNum + "结果";
			}, taskNum);
			
			if (success) {
				console.log("任务 " + taskNum + " 已提交到协程池");
			} else {
				console.log("任务 " + taskNum + " 提交失败");
			}
		}
		
		// 等待任务完成
		coroutine.sleep(3000);
		
		// 获取协程池统计信息
		const poolStats = coroutine.getPoolStats("examplePool");
		console.log("协程池名称: " + poolStats.name);
		console.log("最大工作协程数: " + poolStats.maxWorkers);
		console.log("活跃工作协程数: " + poolStats.active);
		console.log("队列中任务数: " + poolStats.queued);
		
		// 关闭协程池
		const closeSuccess = coroutine.closePool("examplePool");
		console.log("关闭协程池: " + closeSuccess);
	`)
	if err != nil {
		fmt.Printf("执行协程示例3失败: %v\n", err)
	}

	// 示例4: 延迟执行
	fmt.Println("\n--- 示例4: 延迟执行 ---")
	err = ExecuteString(`
		console.log("启动延迟执行示例");
		
		// 设置多个延迟任务
		coroutine.delay(1000, function() {
			console.log("1秒后执行的任务");
		});
		
		coroutine.delay(2000, function() {
			console.log("2秒后执行的任务");
		});
		
		coroutine.delay(3000, function() {
			console.log("3秒后执行的任务");
		});
		
		console.log("所有延迟任务已设置");
		
		// 等待所有延迟任务完成
		coroutine.sleep(3500);
	`)
	if err != nil {
		fmt.Printf("执行协程示例4失败: %v\n", err)
	}

	// 示例5: 协程状态监控
	fmt.Println("\n--- 示例5: 协程状态监控 ---")
	err = ExecuteString(`
		console.log("启动协程状态监控示例");
		
		// 启动多个协程
		for (let i = 0; i < 3; i++) {
			const taskNum = i + 1;
			coroutine.launch(function() {
				console.log("协程 " + taskNum + " 开始执行");
				coroutine.sleep(1000 * taskNum);
				console.log("协程 " + taskNum + " 执行完成");
			}, "monitorTask" + taskNum, taskNum);
		}
		
		// 获取协程列表
		const coroutineList = coroutine.getCoroutineList();
		console.log("协程列表:");
		for (let i = 0; i < coroutineList.length; i++) {
			console.log("  - ID: " + coroutineList[i].id + 
			            ", 名称: " + coroutineList[i].name + 
			            ", 状态: " + coroutineList[i].state +
			            ", 优先级: " + coroutineList[i].priority);
		}
		
		// 等待协程完成
		coroutine.sleep(4000);
		
		// 获取最终统计信息
		const finalStats = coroutine.getStats();
		console.log("最终统计:");
		console.log("  总任务数: " + finalStats.totalTasks);
		console.log("  已完成: " + finalStats.completed);
		console.log("  失败: " + finalStats.failed);
		console.log("  已取消: " + finalStats.cancelled);
		console.log("  活跃协程: " + finalStats.active);
	`)
	if err != nil {
		fmt.Printf("执行协程示例5失败: %v\n", err)
	}

	// 示例6: 调度器使用
	fmt.Println("\n--- 示例6: 调度器使用 ---")
	err = ExecuteString(`
		console.log("启动调度器示例");
		
		// 设置调度策略
		coroutine.setScheduleStrategy("fifo");
		console.log("当前调度策略: " + coroutine.getScheduleStrategy());
		
		// 设置协程优先级
		coroutine.setPriority("highPriorityTask", 10);
		coroutine.setPriority("mediumPriorityTask", 5);
		coroutine.setPriority("lowPriorityTask", 1);
		
		// 获取优先级
		console.log("highPriorityTask 优先级: " + coroutine.getPriority("highPriorityTask"));
		console.log("mediumPriorityTask 优先级: " + coroutine.getPriority("mediumPriorityTask"));
		console.log("lowPriorityTask 优先级: " + coroutine.getPriority("lowPriorityTask"));
		
		// 启动不同优先级的协程
		coroutine.launch(function() {
			console.log("高优先级任务执行");
			coroutine.sleep(500);
		}, "highPriorityTask", 10);
		
		coroutine.launch(function() {
			console.log("中优先级任务执行");
			coroutine.sleep(500);
		}, "mediumPriorityTask", 5);
		
		coroutine.launch(function() {
			console.log("低优先级任务执行");
			coroutine.sleep(500);
		}, "lowPriorityTask", 1);
		
		// 等待完成
		coroutine.sleep(2000);
	`)
	if err != nil {
		fmt.Printf("执行协程示例6失败: %v\n", err)
	}

	// 示例7: 取消协程
	fmt.Println("\n--- 示例7: 取消协程 ---")
	err = ExecuteString(`
		console.log("启动取消协程示例");
		
		// 启动一个长时间运行的协程
		const longTaskId = coroutine.launch(function() {
			console.log("长时间任务开始");
			for (let i = 0; i < 10; i++) {
				console.log("任务进度: " + (i + 1) + "/10");
				coroutine.sleep(500);
			}
			console.log("长时间任务完成");
		}, "longTask", 0);
		
		console.log("协程ID: " + longTaskId);
		
		// 等待2秒后取消协程
		coroutine.sleep(2000);
		
		const cancelSuccess = coroutine.cancel(longTaskId);
		console.log("取消协程: " + cancelSuccess);
		
		// 等待一下
		coroutine.sleep(500);
		
		// 获取统计信息
		const stats = coroutine.getStats();
		console.log("已取消协程数: " + stats.cancelled);
	`)
	if err != nil {
		fmt.Printf("执行协程示例7失败: %v\n", err)
	}

	fmt.Println("\n========== 协程模块示例完成 ==========")
}
