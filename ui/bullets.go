package ui

import "github.com/fatih/color"

// Bullet is a unicode status icon for a Step
type Bullet []string

type BulletState string

var (
	RunningState  BulletState = "Running"
	CompleteState BulletState = "Complete"
	FailedState   BulletState = "Failed"

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

var ColoredAnimatedBulletSet = BulletSet{
	Failed:      Failed.WithColor(color.New(color.FgRed)),
	Future:      Future,
	Running:     AnimatedRunning6.WithColor(color.New(color.FgGreen)),
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
