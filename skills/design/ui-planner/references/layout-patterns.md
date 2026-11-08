# Layout Patterns

Read this in step 2 when arranging a page or naming a region's layout. The names are a shared vocabulary for structure, not a visual style or a required template. Content decides first; a name only makes the decision precise.

## Levels

| Level | Patterns | What the name decides |
| --- | --- | --- |
| Page | Sidebar Layout, Split Screen | The page shell and its main division |
| Section | Hero Section, Bento Grid, Full-width Section | How one band of content is arranged |
| List | Masonry | How a repeated collection flows inside a section |

A page usually combines one page-level pattern with several section and list patterns. State the level when a name could be mistaken for the whole page: "Bento Grid" is a section, not a page.

## Patterns

### Hero Section · 首屏大横幅

One oversized opening statement or image that tells the reader what the page is for before any detail.

- **Use when:** a landing page, campaign, product introduction, or first screen needs one message and one primary action.
- **Avoid when:** the user needs the task immediately (tools, dashboards), several items have equal priority, or the first screen must show data or controls.
- **Responsive:** keep one message and one primary action; reduce height and type scale instead of hiding the action below the fold.
- **Label:** `Hero Section · 首屏大横幅 · 区块级`.

### Bento Grid · 便当格

Mixed-size tiles in one grid, so several peer items get unequal emphasis without separate sections.

- **Use when:** a feature overview, capability map, editorial homepage, or dashboard needs a focal item plus supporting items.
- **Avoid when:** items need equal comparison, text is long, there are more than about seven tiles, or strict reading order matters.
- **Responsive:** collapse to two columns and then one; keep the focal tile first; let tiles grow instead of clipping content.
- **Label:** `Bento Grid · 便当格 · 区块级`.

### Masonry · 瀑布流

Cards flow into columns of unequal height, so a collection grows downward without a rigid row grid.

- **Use when:** an image gallery, article feed, inspiration board, or mixed-height cards benefit from browsing.
- **Avoid when:** rows must align, reading order is strict, or items have equal importance and short text.
- **Responsive:** cap columns (three, then two, then one); preserve DOM order; provide a list alternative for keyboard and screen-reader users; do not lazy-load the first screen.
- **Label:** `Masonry · 瀑布流 · 列表级`.

### Split Screen · 左右分屏

The page or a section divides into two halves, usually image and text or form and context, read against each other.

- **Use when:** signing in, onboarding, comparing two things, telling a story with one strong visual, or pairing a form with live context.
- **Avoid when:** one half is much heavier than the other, the screen is narrow, or the two halves are unrelated.
- **Responsive:** stack the halves; put the primary half first; keep the secondary half after the action; do not shrink both to unreadable widths.
- **Label:** `Split Screen · 左右分屏 · 页面级或区块级`.

### Sidebar Layout · 侧边栏布局

A persistent vertical navigation rail beside the main content, available while the content changes.

- **Use when:** a tool, documentation site, dashboard, settings area, or multi-section app needs stable navigation.
- **Avoid when:** the page is a linear marketing page, has fewer than about four destinations, or the rail would crowd the content on narrow screens.
- **Responsive:** turn the rail into a drawer or top bar; keep main content first in DOM order; the rail must not trap focus or cover content.
- **Label:** `Sidebar Layout · 侧边栏布局 · 页面级`.

### Full-width Section · 通栏

A band that spans the full width to isolate one statement, call to action, or transition.

- **Use when:** one key message, call to action, quote, or transition deserves a break from the surrounding dense sections.
- **Avoid when:** several full-width bands would compete, content needs a reading measure, or the band would duplicate a hero.
- **Responsive:** keep edge padding and a readable text measure; do not let the band become a second hero.
- **Label:** `Full-width Section · 通栏 · 区块级`.

## Choose with the user

- Ask about a region, not the whole page. "功能区用 Bento Grid 还是三条并列？" is easier to answer than "整页换哪种布局？"
- When the user has not named a pattern, offer two or three candidates and state what each changes. Do not present all six as a menu.
- When the user names one region, change only that region. Preserve the shell, navigation, and unaffected sections.
- If a named pattern does not fit the content, say so and propose the closest fit. A name is not a reason to force content into it.
- A precise specification reads like: "页面外壳用 Sidebar Layout；主内容区从上到下是 Hero Section、Bento Grid、Split Screen；文章列表用 Masonry；页尾用 Full-width Section."

## One-to-one labels

The same name appears in three places, with no synonyms:

1. **Canvas:** a visible tag at the top-left of each named region, showing the exact English name and a short Chinese description, for example `Bento Grid · 便当格`.
2. **Region explanation:** the region's note repeats the same name and, when useful, the level.
3. **Implementation handoff:** the "Layout and interaction" section names each region with the same pattern.

The tag belongs to the region, not to a floating legend. Do not label a wrapper or a decorative fragment. When a region contains a list-level pattern, label the list area itself, not its parent section.

Keep tags visible during review so the user can point at a name and say what to change. Remove review-only tags from production code; keep the names in documentation and the handoff.
