package abstractFactory

// 継承ではなくインターフェースによるポリモフィズムを用いる
type item interface {
	toString() string
}

// インターフェースの埋込も可能
type link interface {
	item
}

// インターフェースのポリモフィズムでは構造体間の共通実装やプロパティを
// 利用することができないため、インターフェースを実装した構造体を埋め込む
// ことで実現する
type tray interface {
	item
	AddToTray(item item)
}

type baseTray struct {
	tray []item
}

func (bt *baseTray) AddToTray(item item) {
	bt.tray = append(bt.tray, item)
}

type page interface {
	AddToContent(item item)
	Output() string
}

type basePage struct {
	content []item
}

func (bp *basePage) AddToContent(item item) {
	bp.content = append(bp.content, item)
}

type factory interface {
	CreateLink(caption, url string) link
	CreateTray(caption string) tray
	CreatePage(title, author string) page
}

type mdLink struct {
	caption, url string
}

func (ml *mdLink) toString() string {
	return "[" + ml.caption + "](" + ml.url + ")"
}

type mdTray struct {
	// 共通処理とプロパティを扱える構造体、かつインターフェースによる
	// ポリモフィズムも可能
	baseTray
	caption string
}

func (mt *mdTray) toString() string {
	tray := "- " + mt.caption + "\n"
	for _, item := range mt.tray {
		tray += item.toString() + "\n"
	}
	return tray
}

type mdPage struct {
	basePage
	title, author string
}

func (mp *mdPage) Output() string {
	content := "title: " + mp.title + "\n"
	content += "author: " + mp.author + "\n"
	for _, item := range mp.content {
		content += item.toString() + "\n"
	}
	return content
}

type MdFactory struct {
}

func (mf *MdFactory) CreateLink(caption, url string) link {
	return &mdLink{caption, url}
}
func (mf *MdFactory) CreateTray(caption string) tray {
	return &mdTray{caption: caption}
}
func (mf *MdFactory) CreatePage(title, author string) page {
	return &mdPage{title: title, author: author}
}
