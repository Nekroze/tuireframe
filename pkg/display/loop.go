package display

import (
	"time"

	"github.com/Nekroze/tuireframe/pkg/ir"
	"github.com/gdamore/tcell"
)

type notificationChannel chan struct{}

func Screen(screen ir.Screen) (err error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if e := s.Init(); e != nil {
		return e
	}
	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()
	drawScreen(screen, s)
	s.Show()
	loopHold(s)
	return err
}

func loopHold(s tcell.Screen) {
	quitSignal := make(notificationChannel)
	go inputHandler(s, quitSignal)
	fps := time.NewTicker(time.Second / 30)
loop:
	for {
		select {
		case <-quitSignal:
			break loop
		case <-fps.C:
			s.Show()
		}
	}
	s.Fini()
}

func inputHandler(s tcell.Screen, quitSignal notificationChannel) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				if ev.Rune() == 'q' {
					close(quitSignal)
				}
				return
			case tcell.KeyEscape, tcell.KeyEnter:
				close(quitSignal)
				return
			case tcell.KeyCtrlL:
				s.Sync()
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}
