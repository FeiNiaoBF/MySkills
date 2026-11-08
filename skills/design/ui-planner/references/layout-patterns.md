# Layout Patterns

Read this in step 2 when arranging a page or naming a region's layout. The names are a shared vocabulary for structure, not a visual style or a required template. Content decides first; a name only makes the decision precise.

The dictionary has three levels. A page usually combines one page-level pattern with several section and collection patterns. State the level when a name could be mistaken for the whole page: "Bento Grid" is a section, not a page.

| Level | What it decides | Examples |
| --- | --- | --- |
| **Page** | The page shell and its main division | Sidebar Layout, Holy Grail, Master-Detail, Dashboard Grid |
| **Section** | How one band of content is arranged | Hero Section, Bento Grid, Pricing Table |
| **Collection** | How repeated items flow inside a section | Masonry, List View, Table View |

## Page layouts

| Name | What it is | Use when | Avoid when | Responsive |
| --- | --- | --- | --- | --- |
| **Single Column · 单栏阅读** | One centered reading column, usually 60–75 characters wide. | Articles, docs, changelogs, legal pages, long-form. | Comparison-heavy pages, dashboards, galleries. | Keep the measure; reduce side padding; never full-width prose. |
| **Stacked Sections · 纵向分节** | Full-width bands stacked vertically, each with one job. | Marketing pages, landing pages, product stories. | Content that needs side-by-side comparison or a persistent tool rail. | Bands stack naturally; reduce section padding; keep one idea per band. |
| **Top-nav + Content · 顶部导航** | A full-width header with content below; no persistent sidebar. | Self-contained pages, docs, repos, content sites. | Frequent context switching across many spaces. | Header collapses to a menu; content keeps a max-width column. |
| **Sidebar Layout · 侧边栏布局** | A persistent vertical navigation rail beside the main content. | Tools, dashboards, docs, multi-section apps. | Linear marketing pages, fewer than about four destinations. | Rail → icon rail → drawer/top bar; main content first in DOM. |
| **Holy Grail · 圣杯布局** | Header and footer with left navigation, main content, and a right aside. | Content-rich sites, docs with a table of contents, app pages with context. | Pages without a meaningful aside; narrow screens. | Drop the right aside first, then the left nav; stack main. |
| **Split Screen · 左右分屏** | Two halves side by side, usually visual + text, form + context, or two related views. | Auth, onboarding, comparison, product story, master-detail preview. | One half is much heavier, or the screen is narrow. | Stack; primary half first; do not squeeze both. |
| **Master-Detail · 主从布局** | A list drives a detail pane without leaving the page. | Inboxes, records, settings, browse-then-inspect flows. | Items need full-page context or the detail is rarely used. | Side-by-side on wide; list → pushed detail on mobile. |
| **Three-Pane · 三窗格** | Nav rail → item list → detail, each pane scrolling independently. | Mail, chat, triage, high-volume queues. | Low item volume or a single primary object. | Three → two panes; rail becomes drawer; detail becomes a screen. |
| **Dashboard Grid · 仪表盘网格** | A grid of independent metric cards, charts, and short lists. | Monitoring, analytics, status at a glance. | Users need to act on individual rows; long text. | 12/8/4/1 column collapse; cards stack; no internal card scrolling. |
| **Magazine / Editorial · 杂志编辑** | An asymmetric multi-column grid with large media, varied type, and a deliberate reading path. | Storytelling, features, portfolios, culture, long-form. | Dense operational content, strict comparison. | Asymmetry → single column; keep reading order; do not simply shrink. |
| **Kanban Board · 看板** | Horizontal columns by stage; cards move between columns. | Task management, pipelines, triage, production boards. | Simple lists, mobile-first viewing, records that need detail. | Columns scroll horizontally; mobile shows one stage or a list view. |
| **Feed · 动态流** | A single scrolling stream of rich items, optionally flanked by sidebars. | Social, activity, changelogs, notifications, releases. | Dense comparison or items that need columns. | Keep one column; sidebars drop below or become drawers. |
| **Floating / Centered Card · 居中卡片** | A centered content panel over a backdrop; modal-like but not an overlay. | Login, onboarding, single-task flows, waiting states. | Long forms, dense data, multi-step workflows. | Keep the card width; reduce margins; avoid a full-screen form. |
| **Full-bleed / Immersive · 全幅沉浸** | Edge-to-edge media or color with minimal chrome. | Portfolios, campaigns, visual stories, product reveals. | Dense text, tables, or persistent navigation. | Keep media edge-to-edge; put text in a readable measure; do not crop key content. |
| **Horizontal Scroll · 横向滚动** | Content advances horizontally, often with scroll-snap. | Portfolios, timelines, galleries, linear stories. | Users need comparison or random access to many items. | Provide buttons and keyboard access; fall back to a vertical stack or list. |
| **Sticky / Pinned Scroll · 钉住滚动** | A section stays pinned while its content changes or advances. | Scrollytelling, product tours, step-by-step explanations. | Short pages, mobile-first reading, reduced-motion contexts. | Disable pinning on short or narrow screens; keep a static fallback. |
| **Tabbed Workspace · 标签工作区** | Tabs switch views inside one shell. | Settings, multi-view apps, dashboards with modes. | More than about seven tabs; deep hierarchy. | Tabs scroll horizontally; move to a select/menu on narrow screens. |
| **Terminal / CLI · 终端** | Monospace, prompt-driven, dark-surface interaction. | Developer tools, logs, admin consoles, technical products. | Consumer content, visual browsing, long prose. | Keep monospace; allow wrapping; never rely on color alone. |

## Section layouts

| Name | What it is | Use when | Avoid when | Responsive |
| --- | --- | --- | --- | --- |
| **Hero Section · 首屏大横幅** | One oversized opening statement or image that states what the page is for. | Landing pages, campaigns, product introductions. | Tools and dashboards that need the task immediately; several equal priorities. | Reduce height and type; keep one primary action visible. |
| **Split Section · 图文分栏** | Text and media in two columns, often alternating left/right down the page. | Product features, stories, explanations, comparisons. | The two sides are unrelated or one is much heavier. | Stack; keep text before media or preserve reading order. |
| **Bento Grid · 便当格** | Mixed-size tiles in one grid, giving several peer items unequal emphasis. | Feature overviews, capability maps, editorial homepages. | Equal comparison, long text, more than about seven tiles. | 4 → 2 → 1 columns; keep the focal tile first; let tiles grow. |
| **Feature Grid · 功能网格** | Equal cards of icon, title, and short copy, usually three or four across. | Feature lists, benefits, service overviews. | Long copy, unequal importance, strict comparison. | 3 → 2 → 1; keep copy short; avoid an orphan row. |
| **Full-width Section · 通栏** | A band spanning the full width to isolate one statement or transition. | One key message, quote, or transition between dense sections. | Many consecutive full-width bands; content that needs a measure. | Keep edge padding and a readable measure; do not turn it into a second hero. |
| **CTA Band · 行动条** | A full-width strip that restates the offer and repeats the primary action. | End of a page, or between dense sections. | The page already has a strong hero and no new offer. | Stack copy and button; keep the action visible. |
| **Stats Band · 数据条** | A row of three to five key numbers with short labels. | Proof, impact, product scale, report summaries. | The numbers need explanation or comparison over time. | 4 → 2 → 1; keep labels with numbers; avoid tiny type. |
| **Logo Cloud · 徽标墙** | A grid or strip of customer, partner, or integration logos. | Social proof, integrations, trust. | Logos are the main content or need explanation. | Wrap or scroll; keep logos legible; avoid a uniform gray wall. |
| **Testimonial Wall · 评价墙** | A grid or masonry of customer quotes with names and roles. | Social proof, community, case studies. | Quotes are long-form or need full context. | 3 → 2 → 1; keep quotes scannable; never truncate attribution. |
| **Pricing Table · 价格表** | Two to four plans side by side with price, limits, and features. | SaaS, services, memberships. | The offer is custom or has many variables. | Stack plans; highlight one; keep the comparison readable. |
| **Comparison Table · 对比表** | Options as columns, properties as rows. | Product vs alternatives, plan features, before/after. | More than about four options or many long cells. | Horizontal scroll or stacked groups; keep row labels visible. |
| **FAQ Accordion · 常见问题** | Questions collapsed until opened. | Support, pricing, objections, onboarding. | The answer is critical to the task or must be read in order. | Keep one open at a time on mobile; do not hide essential information. |
| **Process Steps · 步骤条** | A numbered sequence of stages or steps. | Onboarding, how-it-works, workflows. | The sequence branches heavily or has more than about seven steps. | Vertical on narrow; keep numbers and order. |
| **Media Strip · 媒体条** | A horizontal row of images, videos, or cards. | Galleries, case studies, product shots. | Items need captions, comparison, or random access. | Scroll or wrap; provide controls; never autoplay with sound. |
| **Mega Footer · 大型页脚** | A multi-column footer with site map, legal, and contact. | Content-rich sites and large products. | Single-page tools with few destinations. | 4 → 2 → 1 columns; keep legal and contact reachable. |

## Collection layouts

| Name | What it is | Use when | Avoid when | Responsive |
| --- | --- | --- | --- | --- |
| **Masonry · 瀑布流** | Cards flow into columns of unequal height. | Galleries, feeds, inspiration boards, mixed-height cards. | Rows must align, reading order is strict, items are equal. | 3 → 2 → 1 columns; preserve DOM order; provide a list fallback. |
| **Uniform Grid / Card Grid · 等高网格** | Equal cards arranged in rows and columns. | Catalogs, archives, libraries, product browsing. | Items have very different heights or need dense comparison. | 4 → 3 → 2 → 1; keep card heights consistent. |
| **List View · 列表** | One vertical stream of compact rows. | Inboxes, activity, settings, mobile-first records. | Visual browsing is the main job or many fields must compare. | Keep one column; rows grow; do not hide key attributes. |
| **Table View · 表格** | Rows and columns for dense structured comparison. | Admin records, inventory, analytics, CRM data. | Visual browsing matters more than comparison. | Horizontal scroll or column priority; keep headers visible. |
| **Gallery · 画廊** | A visual-first grid or flow of media. | Portfolios, photography, product shots. | Items need long copy or dense attributes. | 4 → 2 → 1; lightbox optional; keep captions and alt text. |
| **Carousel / Slider · 轮播** | One or a few items at a time with previous/next controls. | Featured content, testimonials, media highlights. | Users need to compare many items or browse at their own pace. | Show controls; pause on hover/focus; provide a static list alternative. |
| **Timeline · 时间线** | Chronological items along a line or axis. | History, activity, release notes, process. | Items are not chronological or need side-by-side comparison. | Vertical on narrow; keep dates and order; do not rely on hover. |

## Common confusions

| Pair | Distinction |
| --- | --- |
| **Bento Grid vs Card Grid vs Masonry** | Bento = mixed-size tile composition; Card Grid = equal cards; Masonry = unequal-height columns. |
| **Split Screen vs Split Section** | Split Screen divides the page or a major band; Split Section is a text/media band inside a page. |
| **Sidebar Layout vs Holy Grail vs Three-Pane** | Sidebar = one rail + main; Holy Grail = left nav + main + right aside; Three-Pane = nav + list + detail, each scrolling. |
| **Dashboard Grid vs Kanban Board** | Dashboard = glance at independent metrics; Kanban = move items through stages. |
| **Magazine vs Stacked Sections** | Magazine = asymmetric editorial composition; Stacked Sections = simple vertical bands with one job each. |
| **Feed vs List View** | Feed = rich, social or activity items; List View = compact rows for scanning and selection. |
| **Feature Grid vs Card Grid** | Feature Grid = a section of benefit or feature cards; Card Grid = a collection of independent items. |
| **Feed vs Timeline** | Feed = a scrolling stream of rich items; Timeline = chronological items along an axis. |
| **Media Strip vs Gallery** | Media Strip = a horizontal row of media; Gallery = a visual grid or flow. |
| **Full-width Section vs Full-bleed** | Full-width = a band spanning the content width; Full-bleed = edge-to-edge media or color. |
| **Sticky / Pinned Scroll vs Sticky Sidebar** | Sticky/Pinned Scroll = a page experience; Sticky Sidebar = a persistent rail while content scrolls. |

## Choose with the user

- Ask about a region, not the whole page. "功能区用 Bento Grid 还是 Feature Grid？" is easier to answer than "整页换哪种布局？"
- When the page-level shell is undecided, open the bundled samples in the HTML template (Sidebar Layout, Top-nav + Content, Holy Grail, Dashboard Grid, Master-Detail) and compare them before naming regions.
- When the user has not named a pattern, offer two or three candidates and state what each changes. Do not present the whole dictionary as a menu.
- When the user names one region, change only that region. Preserve the shell, navigation, and unaffected sections.
- If a named pattern does not fit the content, say so and propose the closest fit. A name is not a reason to force content into it.
- A precise specification reads like: "页面外壳用 Top-nav + Content；首屏用 Hero Section；功能区用 Bento Grid；文章列表用 Masonry；页尾用 CTA Band + Mega Footer."

## One-to-one labels

The same name appears in four places, with no synonyms:

1. **Canvas:** a visible tag at the top-left of each named region, showing the exact English name and a short Chinese description, for example `Bento Grid · 便当格`.
2. **Layout overview:** a compact table in the notes listing each region with the same name and level. Summarize this page's choices; do not paste the whole dictionary.
3. **Region explanation:** the region's note repeats the same name and, when useful, the level.
4. **Implementation handoff:** the "Layout and interaction" section names each region with the same pattern.

The tag belongs to the region, not to a floating legend. Do not label wrappers or decorative fragments. When a region contains a collection-level pattern, label the collection area itself, not its parent section.

Keep tags visible during review so the user can point at a name and say what to change. Remove review-only tags from production code; keep the names in documentation and the handoff.

## Sources

- [Every Layout — layout primitives](https://every-layout.dev/layouts/): Stack, Box, Center, Cluster, Sidebar, Switcher, Cover, Grid, Frame, Reel, Imposter, Container. Useful for implementation-level composition, not page naming.
- [HashBuilds — Layout Patterns](https://www.hashbuilds.com/categories/layouts): Bento, Hero, Sidebar, Dashboard, Masonry, Split Screen, Card Grid, Holy Grail, Magazine, Floating, Terminal.
- [TurnConcepts — Layouts](https://turnconcepts.com/layouts): Holy Grail, Dashboard Grid, Split View, Kanban, Masonry, Hero Landing, Sidebar + Tabs, Feed/Timeline.
- [UX Patterns for Developers — Table vs List View vs Card Grid](https://uxpatterns.dev/pattern-guide/table-vs-list-vs-cards): comparison, scanning, and visual browsing tradeoffs.
- [CSS Crème — Pattern Glossary](https://csscreme.com/patterns): Hero, Bento Grid, Card Grid, Mega Footer, Pricing Table, Testimonial Wall, Feature Comparison Table, FAQ Accordion, CTA Band.
- [W3Schools — HTML & CSS Layout Patterns](https://www.w3schools.com/htmlcss/htmlcss_layout_patterns.asp): hero + media, sidebar + content, dashboard cards, marketing content blocks, footer patterns.
- [Application Layout Shells](https://github.com/jpoindexter/ux-flow-skills/blob/main/skills/flow-app-shell/SKILL.md): Sidebar + main, Top-nav + content, Three-pane, Master-detail, Dashboard grid, mobile tab-bar shell, responsive collapse order.
- [Visual Playground — Layout Patterns](https://drjulsgilliam.github.io/HTML-CSS-Visual-Playground/layout-patterns.html): reading column, split screen, masonry, sticky header/sidebar, container.
- [Editorial / Magazine Layout Reference](https://chrislemke.github.io/website_designs/designs/Editorial_Magazine_Layout.html) and [Scrollytelling](https://chrislemke.github.io/website_designs/designs/Scrollytelling.html): asymmetric grids, big type, pinned scroll, full-bleed narrative.
- [SAP Fiori List-Detail](https://www.sap.com/design-system/fiori-design-android/v25-4/layouts/list-detail/list-detail) and [Oracle Alta Page Types](https://www.oracle.com/webfolder/ux/middleware/alta/patterns/PageTypes.html): master-detail, dashboard, forms, wizards.
