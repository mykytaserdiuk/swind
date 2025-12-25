package render

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/models"
)

func TestNewRaylibRender(t *testing.T) {
	cam := &rl.Camera2D{}
	r := NewRaylibRender(cam)

	if r.cam != cam {
		t.Fatal("expected camera to be set")
	}
	if len(r.queue) != 0 {
		t.Fatal("expected empty queue on init")
	}
}

func TestSubmitAndQueueOrder(t *testing.T) {
	cam := &rl.Camera2D{}
	r := NewRaylibRender(cam)

	cmd1 := DrawCmd{Layer: models.LayerUI}
	cmd2 := DrawCmd{Layer: models.LayerContent}
	cmd3 := DrawCmd{Layer: models.LayerBackground}

	r.Submit(cmd1)
	r.Submit(cmd2)
	r.Submit(cmd3)

	if len(r.queue) != 3 {
		t.Fatalf("expected 3 commands in queue, got %d", len(r.queue))
	}

	if r.queue[0].Layer != models.LayerUI {
		t.Fatalf("expected first cmd layer UI, got %v", r.queue[0].Layer)
	}
	if r.queue[1].Layer != models.LayerContent {
		t.Fatalf("expected second cmd layer Content, got %v", r.queue[1].Layer)
	}
	if r.queue[2].Layer != models.LayerBackground {
		t.Fatalf("expected third cmd layer Background, got %v", r.queue[2].Layer)
	}
}

func TestFlushSortsAndClears(t *testing.T) {
	cam := &rl.Camera2D{}
	r := NewRaylibRender(cam)

	callOrder := []int{}

	cmd1 := DrawCmd{
		Layer: models.LayerUI,
		Fn: func() {
			callOrder = append(callOrder, int(models.LayerUI))
		},
	}
	cmd2 := DrawCmd{
		Layer: models.LayerContent,
		Fn: func() {
			callOrder = append(callOrder, int(models.LayerContent))
		},
	}
	cmd3 := DrawCmd{
		Layer: models.LayerBackground,
		Fn: func() {
			callOrder = append(callOrder, int(models.LayerBackground))
		},
	}

	r.Submit(cmd1)
	r.Submit(cmd3)
	r.Submit(cmd2)

	if len(r.queue) != 3 {
		t.Fatalf("expected 3 commands before flush, got %d", len(r.queue))
	}

	expectedLayers := []models.Layer{models.LayerUI, models.LayerBackground, models.LayerContent}
	for i, cmd := range r.queue {
		if cmd.Layer != expectedLayers[i] {
			t.Fatalf("cmd %d: expected layer %v, got %v", i, expectedLayers[i], cmd.Layer)
		}
	}
}

func TestTextCommand(t *testing.T) {
	cam := &rl.Camera2D{}
	r := NewRaylibRender(cam)

	cmd := TextRenderCmd{
		Text:     "Hello",
		PosX:     10,
		PosY:     20,
		FontSize: 24,
		Col:      rl.White,
	}

	r.Text(models.LayerUI, cmd)

	if len(r.queue) != 1 {
		t.Fatalf("expected 1 command in queue, got %d", len(r.queue))
	}

	if r.queue[0].Layer != models.LayerUI {
		t.Fatalf("expected LayerUI, got %v", r.queue[0].Layer)
	}

	if r.queue[0].Fn == nil {
		t.Fatal("expected Fn to be set")
	}
}

func TestRectCommand(t *testing.T) {
	cam := &rl.Camera2D{}
	r := NewRaylibRender(cam)

	cmd := RectRenderCmd{
		PosX:   10,
		PosY:   20,
		Width:  100,
		Height: 50,
		Col:    rl.Blue,
	}

	r.Rect(models.LayerContent, cmd)

	if len(r.queue) != 1 {
		t.Fatalf("expected 1 command in queue, got %d", len(r.queue))
	}

	if r.queue[0].Layer != models.LayerContent {
		t.Fatalf("expected LayerContent, got %v", r.queue[0].Layer)
	}

	if r.queue[0].Fn == nil {
		t.Fatal("expected Fn to be set")
	}
}
