# Worked output example

This is a synthetic planning example, not an inspection of this repository. Scenario facts supplied for the example: the user wants a local CSV expense summary CLI; the repository is empty; Python 3.11 and columns `date,category,amount` are confirmed; only planning is authorized; real bank exports are unavailable. The missing export blocks a future bank-specific adapter, not the core CLI. All proposed paths below belong to the hypothetical project.

The plan below follows [the output contract](output-format.md). Its two slices are an example, not a required stage count or technology choice.

---

# 实施计划：本地 CSV 支出汇总工具

## 1. 目标与建议

- 目标：用户运行本地命令后，能看到 CSV 中按分类汇总的支出；输入错误时能定位到问题行。
- 范围：使用已约定的三列 CSV 和 Python 3.11；不含银行专用导入、图形界面、账户连接或自动上传。
- 建议：先做 S1，用最小可运行命令贯通读取与汇总，再由 S2 补齐错误处理。本次只交付计划，尚未授权实现。

## 2. 现状与依据

- 已确认：任务输入明确仓库为空、使用 Python 3.11，字段为 `date,category,amount`。
- 假设：金额使用十进制表示，单个文件内币种一致；若需多币种，必须先调整输入约定，不能直接相加。
- 未知项：未提供真实银行导出文件。这不阻塞 S1、S2 的约定格式实现；银行格式适配不在本次范围内。

## 3. 实施顺序

| 编号 | 可见结果 | 依赖 | 排序理由 | 验收证据 | 状态 |
| --- | --- | --- | --- | --- | --- |
| S1 | 从约定 CSV 输出按分类汇总的金额 | 无 | 先验证用户可直接运行的核心路径 | 两条 food 记录 12.30、7.70 汇总为 20.00，transport 3.00 单独输出 | 待实施 |
| S2 | 对缺列或非法金额报告行号，退出失败且不输出误导性汇总 | S1 | 复用已贯通的读取路径，补齐输入失败行为 | 缺少 amount 列或输入非法金额时返回非零状态；S1 样例仍通过 | 待实施 |

## 4. 下一切片

- 切片：S1。
- 行为与边界：接收一个文件路径，按 category 汇总 amount 并稳定排序；本切片只验收符合约定的输入，不宣称已经支持真实银行导出。
- 前置条件：已约定运行语言和字段；没有阻塞 S1 的未决事项。
- 影响范围：拟新增 `expense_summary.py` 作为命令入口；这些路径尚不存在。临时 CSV 样例放系统临时目录。
- 执行步骤：
  1. 建立接收 CSV 路径的最小命令入口。
  2. 使用标准库读取数据和十进制金额，按分类累加并输出。
  3. 用三条固定记录检查分类、总额和输出顺序；把实际运行方式记入项目现有说明，若无说明则在实现范围内添加最小使用说明。
- 验证方式：拟在项目根目录执行 `python3 expense_summary.py <sample.csv>`；样例含 food 的 12.30、7.70 和 transport 的 3.00，预期退出码为 0，food 汇总为 20.00，transport 汇总为 3.00。命令尚未执行，程序尚未创建。
- 停止或重排条件：如果确认输入包含多币种且没有币种列，先解决输入约定，再实现金额合并；真实银行样本缺席不暂停当前约定格式的工作。

## 5. 风险与待决事项

- 币种约定仍是未验证假设，影响 S1 的金额语义；若用户提供跨币种需求，应修改输入约定并重排受影响工作。当前计划不声称支持多币种。
- 银行专用格式尚未确认，且不在范围内；以后需要适配时，先取得一个脱敏样本，再决定是否新增切片。

## 6. 交付记录

- 已检查：仅使用本示例提供的任务事实；未检查实际项目文件，未执行命令或测试。
- 本次产物：对话中的实施计划；未创建项目代码、计划文件或 ADR。
- 变更与原因：初版计划。把约定格式的核心路径与输入错误处理分开，使每个切片都有可观察的完成条件。
- 下一步行动：若用户随后授权开发，从 S1 开始；不需要等待真实银行导出样本。
