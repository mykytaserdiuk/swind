package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func ClampToWorkArea(rect, workArea rl.Rectangle) rl.Rectangle {
	// ограничиваем размеры
	if rect.Width > workArea.Width {
		rect.Width = workArea.Width
	}
	if rect.Height > workArea.Height {
		rect.Height = workArea.Height
	}

	// clamp X
	if rect.X < workArea.X {
		rect.X = workArea.X
	}
	if rect.X+rect.Width > workArea.X+workArea.Width {
		rect.X = workArea.X + workArea.Width - rect.Width
	}

	// clamp Y
	if rect.Y < workArea.Y {
		rect.Y = workArea.Y
	}
	if rect.Y+rect.Height > workArea.Y+workArea.Height {
		rect.Y = workArea.Y + workArea.Height - rect.Height
	}

	return rect
}
