package systems

import (
	"game2d/components"
)

type RenderSystem struct{}

func (r *RenderSystem) Update(Game *Game) {
	// Render entities on the screen
	for _, entity := range Game.Entities {
		if positionComponent := entity.GetComponent(&components.PositionComponent{}); positionComponent != nil {
			continue
			//position := positionComponent.(*components.PositionComponent)
			// Render the entity at its position
			// Replace this with your rendering code
			//fmt.Printf("Rendering entity %d at position (%.2f, %.2f)\n", entity.ID, position.X, position.Y)
		}
	}
}
