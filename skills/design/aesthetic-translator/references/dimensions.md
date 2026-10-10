# 九维度词汇表

每个维度：用户的话 → 命名选项（参照物）→ 落地参数。翻译时按需查本表，给用户 2-3 个选项即可，不必穷举。

## 1. 风格（最大的决定，其余都挂在它下面）

| 用户的话 | 选项 | 参照 | 参数要点 |
|---|---|---|---|
| 极简、干净、内容为王 | Minimalist | Apple, iA Writer | 大留白、单一强调色、零装饰 |
| 严格、理性、老钱 | Swiss / International | Vercel | 严网格、超大字号、一个红色点缀 |
| 响亮、个性、拒绝忽视 | Neo-brutalism | Gumroad | 3px 黑边框、硬偏移阴影、平涂色 |
| 原始、工程师感 | Raw brutalism | Hacker News | 系统字体、下划线链接、零化妆 |
| 悬浮、轻盈、现代 | Glassmorphism | iOS, visionOS | backdrop-filter、半透明面板 |
| 安静、高级、不说出口 | Dark luxury | CRED | 近黑底、香槟金点缀、衬线标题 |
| 杂志感、叙事 | Editorial | NYT, Stripe Press | 大衬线标题、分栏、引文块 |
| 圆润、快乐 | Playful | Duolingo | 全圆角、糖果色、胖按钮 |
| 工程师宣言 | Terminal / TUI | dev tool CLIs | 等宽、深底单色高亮、ASCII/框线 |
| Google 系应用感 | Material 3 | Android | 走官方 token，不手搓 |
| 墨与纸、东方 | East-Asian ink | 见 east-asian.md | 纸墨二元、朱砂单情绪色 |

## 2. 色彩情绪（先选情绪，色值跟着来）

| 用户的话 | 选项 | 参数要点 |
|---|---|---|
| 温柔、治愈 | Pastel | 低饱和 |
| 自然、踏实 | Earthy | 陶土/沙/橄榄 |
| 富有、自信 | Jewel tones | 深宝石色 |
| 成熟、贵 | Muted / dusty | 降饱和、灰调 |
| 安全的高级 | Greyscale + 一强调色 | 灰阶做全部，一点色说话 |
| 纸和墨 | Cream & ink | off-white 底 + 墨色字 |
| 能量、年轻 | Neon on dark | 暗底电光点缀 |

色彩词汇（直接可用）：60-30-10 分配 / 单强调色 / 暖调·冷调 / 永远 off-white 不用纯白 / charcoal 不用纯黑 / 彩色阴影 / 渐变只出现在 hero。

## 3. 字体性格

| 用户的话 | 选项 | 字体 | 气质 |
|---|---|---|---|
| 中性、不抢戏 | Grotesque sans | Inter, Helvetica | 产品级 |
| 友好、创业感 | Geometric sans | Poppins, Futura | 圆润乐观 |
| 现代、技术 | Tech display | Space Grotesk | AI/web3 能量 |
| 书卷、可信 | Old-style serif | Garamond, Georgia | 几百年的阅读感 |
| 奢侈、时尚 | Fashion serif | Playfair, Didot | 厚薄对比戏剧感 |
| 厚重、踏实 | Slab serif | Roboto Slab | 砖头脚 |
| 精确、代码感 | Monospace | JetBrains Mono | 每个字母等权 |

字型词汇：超大标题 / 标题紧字距 / 正文行高宽松 / 全大写标签宽字距 / 衬线标题+无衬线正文 / 一族字体靠字重分层 / 等宽只给数字与代码。

## 4. 明暗主题（按用户的时刻选，不按心情选）

阅读重的产品 → light；媒体/金融/开发工具 → dark 可选；长会话工具 → dim soft dark；阅读产品 → warm paper（米白+墨）；可访问性优先 → high contrast。

## 5. 布局

单列限宽 720px（阅读）/ 分屏左文右图（落地页）/ 侧栏壳（工具）/ 等宽卡片阵（目录/市场）/ 杂志不等宽栏 / 全幅 hero 一句话 / bento 混合瓷砖（功能陈列）/ 浮动胶囊导航。

## 6. 细节（两个「极简」的差别全在这里）

- 圆角：0px 严肃 / 6-8px 低调现代 / 16px 友好 / 24px+ 玩具感
- 阴影：无（纯平）/ 大模糊低透明柔影 / 硬偏移无模糊（图形感）/ 多层堆叠（真实纵深）
- 卡片：描边 hairline / 柔影浮起 / 玻璃 / 渐变描边
- 按钮：胶囊 / 直角 / ghost 描边 / 3D 可按压（硬底边，点击下移）
- 收尾：可见焦点环 / hairline 分隔线替代盒子 / 单一图标集 / 暗底内描边高光

## 7. 动效（调味料：少则活，多则廉价）

默认 200ms fade + hover 微升；弹簧物理只给 playful；滚动渐入只给叙事页；拿不准就减。明确写出毫秒数与缓动，不写「流畅」。

## 8. 密度

Airy（落地页、高级品牌，费滚动）/ Balanced（多数产品）/ Compact（终端、管理面板，省留白）。与三旋钮的 DENSITY 对应。

## 9. 图文与语气

图像语言选一种并忠诚到底：3D 渲染 / 扁平矢量 / 细线图标 / 实心图标 / 真实抓拍 / 抽象渐变斑 / 颗粒噪点 / 手绘涂鸦 / 无图像纯字体。
文案语气也是设计：warm（"Oops!"）与 precise（"Request failed. Retry."）是两个产品。
