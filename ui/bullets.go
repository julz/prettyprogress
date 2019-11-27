package ui

// Bullet is a unicode status icon for a Step
type Bullet []string

var (
	Failed      Bullet = []string{"✗"}
	Future      Bullet = []string{" "}
	Running     Bullet = []string{"►"}
	Downloading Bullet = []string{"↡"}
	Uploading   Bullet = []string{"↟"}
	Complete    Bullet = []string{"✓"}

	AnimatedRunning  Bullet = []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}
	AnimatedRunning2 Bullet = []string{"◴", "◷", "◶", "◵"}
	AnimatedRunning3 Bullet = []string{"◐", "◓", "◑", "◒"}
	AnimatedRunning4 Bullet = []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
	AnimatedRunning5 Bullet = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	AnimatedRunning6 Bullet = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
)

type BulletSet struct {
	Failed      Bullet
	Future      Bullet
	Running     Bullet
	Downloading Bullet
	Uploading   Bullet
	Complete    Bullet
}

var DefaultBulletSet = BulletSet{
	Failed:      Failed,
	Future:      Future,
	Running:     Running,
	Downloading: Downloading,
	Uploading:   Uploading,
	Complete:    Complete,
}

var AnimatedBulletSet = BulletSet{
	Failed:      Failed,
	Future:      Future,
	Running:     AnimatedRunning6,
	Downloading: Downloading,
	Uploading:   Uploading,
	Complete:    Complete,
}

type Color interface {
	Sprint(a ...interface{}) string
}

func (b Bullet) WithColor(color Color) Bullet {
	colored := make(Bullet, len(b))
	for i, c := range b {
		colored[i] = color.Sprint(c)
	}

	return colored
}
