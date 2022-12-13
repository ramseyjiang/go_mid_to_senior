package snake

// SymbolSnake and SymbolEgg is used to display snake and egg position at the beginning.
// 0x2588 means a rectangle, 0x25CF means a cycle point.
const SymbolSnake = 0x2588
const SymbolEgg = 0x25CF

// FrameWidth and FrameHeight is used to change it according to your need
// When run this game, please let your terminal has enough height and width, otherwise, please adjust screenWidth and screenHeight value.
const FrameWidth = 80
const FrameHeight = 20

// FrameBorderThickness means game frame's border thickness
const FrameBorderThickness = 1

// FrameBorderVertical and others const below is used to represent game frame border symbols to make it look a bit fancy
const FrameBorderVertical = '║'
const FrameBorderHorizontal = '═'
const FrameBorderTopLeft = '╔'
const FrameBorderTopRight = '╗'
const FrameBorderBottomRight = '╝'
const FrameBorderBottomLeft = '╚'
