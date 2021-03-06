package engine

type Map struct {
	worldMap [][]int
	midMap   [][]int
	upMap    [][]int

	sprite     []*Sprite
	numSprites int

	tex *TextureHandler
}

func NewMap(tex *TextureHandler) *Map {
	m := &Map{}
	m.tex = tex

	m.worldMap = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 6, 1, 1, 0, 0, 0, 0, 0, 1, 1, 6, 1, 1},
		{1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	m.midMap = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	m.upMap = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	return m
}

func (m *Map) LoadSprites() {
	m.sprite = []*Sprite{
		// // Doom Creator
		//NewAnimatedSprite(8, 11.5, 2, 5, m.tex.Textures[15], 5, 1, 256),

		// // line of trees for testing in front of initial view
		NewSprite(19.5, 11.5, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.5, 8, m.tex.Textures[16], 256, false, ""),
		NewSprite(17.5, 11.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(15.5, 11.5, m.tex.Textures[9], 256, false, ""),
		// // render a forest!
		NewSprite(11.5, 1.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 1.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(132.5, 1.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.5, 2, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 2, m.tex.Textures[9], 256, false, ""),
		NewSprite(13.5, 2, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.5, 2.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.25, 2.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(13.5, 2.25, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.5, 3, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 3, m.tex.Textures[9], 256, false, ""),
		NewSprite(13.25, 3, m.tex.Textures[9], 256, false, ""),
		NewSprite(10.5, 3.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.5, 3.25, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 3.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(13.25, 3.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(10.5, 4, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.5, 4, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 4, m.tex.Textures[9], 256, false, ""),
		NewSprite(13.5, 4, m.tex.Textures[14], 256, false, ""),
		NewSprite(10.5, 4.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.25, 4.5, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 4.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(13.5, 4.5, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.5, 4.25, m.tex.Textures[14], 256, false, ""),
		NewSprite(10.5, 5, m.tex.Textures[9], 256, false, ""),
		NewSprite(11.5, 5, m.tex.Textures[9], 256, false, ""),
		NewSprite(12.5, 5, m.tex.Textures[14], 256, false, ""),
		NewSprite(13.25, 5, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.5, 5, m.tex.Textures[14], 256, false, ""),
		NewSprite(11.5, 5.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(12.5, 5.25, m.tex.Textures[10], 256, false, ""),
		NewSprite(13.5, 5.25, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.5, 5.5, m.tex.Textures[10], 256, false, ""),
		NewSprite(15.5, 5.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(11.5, 6, m.tex.Textures[14], 256, false, ""),
		NewSprite(12.5, 6, m.tex.Textures[10], 256, false, ""),
		NewSprite(13.25, 6, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.25, 6, m.tex.Textures[10], 256, false, ""),
		NewSprite(15.5, 6, m.tex.Textures[14], 256, false, ""),
		NewSprite(12.5, 6.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(13.5, 6.25, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.5, 6.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(12.5, 7, m.tex.Textures[14], 256, false, ""),
		NewSprite(13.5, 7, m.tex.Textures[10], 256, false, ""),
		NewSprite(14.5, 7, m.tex.Textures[14], 256, false, ""),
		NewSprite(13.5, 7.5, m.tex.Textures[14], 256, false, ""),
		NewSprite(13.5, 8, m.tex.Textures[14], 256, false, ""),
		NewSprite(21.5, 21.5, m.tex.Textures[20], 256, false, "chainsaw"),
		NewSprite(1.5, 1.5, m.tex.Textures[21], 256, false, "shotgun"),

	}
	m.numSprites = len(m.sprite)
}

func (m *Map) GetAt(x, y int) int {
	return m.worldMap[x][y]
}

func (m *Map) GetSprites() []*Sprite {
	return m.sprite
}

func (m *Map) GetSprite(index int) *Sprite {
	return m.sprite[index]
}

func (m *Map) GetNumSprites() int {
	return m.numSprites
}

func (m *Map) getGrid() [][]int {
	return m.worldMap
}

func (m *Map) getGridUp() [][]int {
	return m.upMap
}

func (m *Map) getGridMid() [][]int {
	return m.midMap
}

func (m *Map) RemoveSprite(name string)  {
	for k, sprite := range m.sprite {
		if sprite.PickupName == name {
			m.sprite[k].ClearTexture()
		}
	}

	m.numSprites = len(m.sprite)
}

func (m *Map) getCollisionLines() []Line {
	if len(m.worldMap) == 0 || len(m.worldMap[0]) == 0 {
		return []Line{}
	}

	lines := Rect(clipDistance, clipDistance,
		float64(len(m.worldMap))-clipDistance, float64(len(m.worldMap[0]))-clipDistance)

	for x, row := range m.worldMap {
		for y, value := range row {
			if value > 0 {
				lines = append(lines, Rect(float64(x)-clipDistance, float64(y)-clipDistance,
					1.0+(2*clipDistance), 1.0+(2*clipDistance))...)
			}
		}
	}

	return lines
}
