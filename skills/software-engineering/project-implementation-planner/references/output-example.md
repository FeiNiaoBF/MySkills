# Worked output example

This is a synthetic planning example, not an inspection of this repository. Scenario facts supplied for the example: the user wants a local CSV expense summary CLI; the repository is empty; Python 3.11 and columns `date,category,amount` are confirmed; only planning is authorized; real bank exports are unavailable. The missing export blocks a future bank-specific adapter, not the core CLI. All proposed paths below belong to the hypothetical project.

The plan below follows [the output contract](output-format.md). Its two slices are an example, not a required stage count or technology choice.

---

# 实施计划：本地 CSV 支出汇总工具

## 1. 目标与建议

- 目标：用户运行本地命令后，能看到 CSV 中按分类汇总的支出；输入错误时能定位到问题行。
- 范围：使用已约定的三列 CSV 和 Python 3.11；不含银行专用导入、图形界面、账户连接或自动上传。
- 建议：先做 S1，让命令能读取 CSV 并输出分类总额；再做 S2，处理输入错误。本次只交付计划，尚未授权实现。

## 2. 现状与依据

- 已确认：任务输入明确仓库为空、使用 Python 3.11，字段为 `date,category,amount`。
- 假设：金额使用十进制表示，单个文件内币种一致；若需多币种，必须先调整输入约定，不能直接相加。
- 未知项：未提供真实银行导出文件。这不阻塞 S1、S2 的约定格式实现；银行格式适配不在本次范围内。

## 3. 实施顺序

| 编号 | 可见结果 | 依赖 | 排序理由 | 验收证据 | 状态 |
| --- | --- | --- | --- | --- | --- |
| S1 | 从约定 CSV 输出按分类汇总的金额 | 无 | 先让用户能运行命令并得到汇总结果 | 两条 food 记录 12.30、7.70 汇总为 20.00，transport 3.00 单独输出 | 待实施 |
| S2 | 缺列或金额无效时，报告问题行并以非零状态退出，不输出误导性汇总 | S1 | 使用 S1 的文件读取代码，增加输入检查 | 缺少 amount 列或金额无效时返回非零状态；S1 样例仍通过 | 待实施 |

## 4. 下一切片

- 切片：S1。这里的“切片”指一块可单独开发和验证的工作。
- 行为与边界：程序接收一个文件路径，按 category 汇总 amount，并按固定顺序输出。本次只验收符合约定的输入；真实银行导出格式尚未验证。
- 前置条件：已约定运行语言和字段；没有阻塞 S1 的未决事项。
- 影响范围：拟新增 `expense_summary.py`，让用户从命令行运行；该文件尚不存在。临时 CSV 样例放系统临时目录。
- 执行步骤：
  1. 让命令接受一个 CSV 文件路径。
  2. 使用标准库读取数据和十进制金额，按分类累加并输出。
  3. 用三条固定记录检查分类、总额和输出顺序。把实际运行方式写入项目说明；若没有说明文件，则在实现范围内补充最小使用说明。
- 验证方式：拟在项目根目录执行 `python3 expense_summary.py <sample.csv>`；样例含 food 的 12.30、7.70 和 transport 的 3.00，预期退出码为 0，food 汇总为 20.00，transport 汇总为 3.00。命令尚未执行，程序尚未创建。
- 停止或重排条件：若输入包含多币种却没有币种列，先调整输入约定，再合并金额。缺少真实银行样本不影响当前约定格式的开发。

## 5. 风险与待决事项

- 尚未确认一个文件是否只含一种币种；这决定了 S1 能否直接相加。若用户需要处理多币种，先修改输入约定，再调整受影响工作的顺序。多币种支持尚未确认。
- 银行专用格式尚未确认，且不在范围内；以后需要适配时，先取得一个脱敏样本，再决定是否新增切片。

## 6. 交付记录

- 已检查：仅使用本示例提供的任务事实；未检查实际项目文件，未执行命令或测试。
- 本次产物：对话中的实施计划；未创建项目代码、计划文件或架构决策记录（ADR）。
- 变更与原因：初版计划。先实现符合约定的输入，再处理输入错误，两步分别验证。
- 下一步行动：若用户随后授权开发，从 S1 开始；不需要等待真实银行导出样本。
