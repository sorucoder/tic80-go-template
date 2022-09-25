package tic80

import (
	"bytes"
	"reflect"
	"unsafe"
)

var (
	IO_RAM   = (*[0x18000]byte)(unsafe.Pointer(uintptr(0x00000)))
	FREE_RAM = (*[0x28000]byte)(unsafe.Pointer(uintptr(0x18000)))
)

func toCString(goString *string) unsafe.Pointer {
	var cStringBuffer bytes.Buffer
	for _, goRune := range *goString {
		if goRune > 0 {
			if goRune > 0x00 && goRune < 0x7F {
				cStringBuffer.WriteRune(goRune)
			} else {
				cStringBuffer.WriteRune('?')
			}
		}
	}
	cStringBuffer.WriteByte(0)

	cStringBytes := cStringBuffer.Bytes()
	buffer, _ := toCBuffer(&cStringBytes)
	return buffer
}

func toCBuffer(goBytes *[]byte) (buffer unsafe.Pointer, count int) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(goBytes))
	buffer = unsafe.Pointer(sliceHeader.Data)
	// For some odd reason, tinygo considers the type of reflect.SliceHeader.Len to be uintptr,
	// instead of int. Using the builtin len function instead.
	count = len(*goBytes)
	return
}

type FontOptions struct {
	transparentColors []byte
	characterWidth    int
	characterHeight   int
	fixed             bool
	scale             int
	alternateFont     bool
}

var defaultFontOptions FontOptions = FontOptions{
	transparentColors: nil,
	characterWidth:    8,
	characterHeight:   8,
	fixed:             false,
	scale:             1,
	alternateFont:     false,
}

func NewFontOptions() *FontOptions {
	options := new(FontOptions)
	*options = defaultFontOptions
	return options
}

func (options *FontOptions) AddTransparentColor(color int) *FontOptions {
	if options.transparentColors == nil {
		options.transparentColors = make([]byte, 0, 1)
	}
	options.transparentColors = append(options.transparentColors, byte(color%16))
	return options
}

func (options *FontOptions) SetCharacterSize(width, height int) *FontOptions {
	options.characterWidth = width
	options.characterHeight = height
	return options
}

func (options *FontOptions) SetScale(scale int) *FontOptions {
	options.scale = scale
	return options
}

func (options *FontOptions) ToggleFixed() *FontOptions {
	options.fixed = !options.fixed
	return options
}

func (options *FontOptions) TogglePage() *FontOptions {
	options.alternateFont = !options.alternateFont
	return options
}

type MapOptions struct {
	x                 int
	y                 int
	width             int
	height            int
	screenX           int
	screenY           int
	transparentColors []byte
	scale             int
}

var defaultMapOptions MapOptions = MapOptions{
	x:                 0,
	y:                 0,
	width:             30,
	height:            17,
	screenX:           0,
	screenY:           0,
	transparentColors: nil,
	scale:             1,
}

func NewMapOptions() *MapOptions {
	options := new(MapOptions)
	*options = defaultMapOptions
	return options
}

func (options *MapOptions) AddTransparentColor(color int) *MapOptions {
	if options.transparentColors == nil {
		options.transparentColors = make([]byte, 0, 1)
	}
	options.transparentColors = append(options.transparentColors, byte(color%16))
	return options
}

func (options *MapOptions) SetOffset(x, y int) *MapOptions {
	options.x = x
	options.y = y
	return options
}

func (options *MapOptions) SetSize(width, height int) *MapOptions {
	options.width = width
	options.height = height
	return options
}

func (options *MapOptions) SetPosition(x, y int) *MapOptions {
	options.screenX = x
	options.screenY = y
	return options
}

func (options *MapOptions) SetScale(scale int) *MapOptions {
	options.scale = scale
	return options
}

type MusicOptions struct {
	track   int
	frame   int
	row     int
	loop    bool
	sustain bool
	tempo   int
	speed   int
}

var defaultMusicOptions MusicOptions = MusicOptions{
	track:   -1,
	frame:   -1,
	row:     -1,
	loop:    true,
	sustain: false,
	tempo:   -1,
	speed:   -1,
}

func NewMusicOptions() *MusicOptions {
	options := new(MusicOptions)
	*options = defaultMusicOptions
	return options
}

func (options *MusicOptions) SetTrack(track int) *MusicOptions {
	options.track = track % 8
	return options
}

func (options *MusicOptions) SetFrame(frame int) *MusicOptions {
	options.frame = frame % 16
	return options
}

func (options *MusicOptions) SetRow(row int) *MusicOptions {
	options.row = row % 64
	return options
}

func (options *MusicOptions) SetTempo(tempo int) *MusicOptions {
	options.tempo = tempo%241 + 40
	return options
}

func (options *MusicOptions) SetSpeed(speed int) *MusicOptions {
	options.speed = speed%31 + 1
	return options
}

func (options *MusicOptions) ToggleLooping() *MusicOptions {
	options.loop = !options.loop
	return options
}

func (options *MusicOptions) ToggleSustain() *MusicOptions {
	options.sustain = !options.sustain
	return options
}

type PrintOptions struct {
	color         byte
	fixed         bool
	scale         int
	alternateFont bool
}

var defaultPrintOptions PrintOptions = PrintOptions{
	color:         15,
	fixed:         false,
	scale:         1,
	alternateFont: false,
}

func NewPrintOptions() *PrintOptions {
	options := new(PrintOptions)
	*options = defaultPrintOptions
	return options
}

func (options *PrintOptions) SetColor(color int) *PrintOptions {
	options.color = byte(color % 16)
	return options
}

func (options *PrintOptions) SetScale(scale int) *PrintOptions {
	options.scale = scale
	return options
}

func (options *PrintOptions) ToggleFixed() *PrintOptions {
	options.fixed = !options.fixed
	return options
}

func (options *PrintOptions) TogglePage() *PrintOptions {
	options.alternateFont = !options.alternateFont
	return options
}

type SoundEffectNote int

const NOTE_NONE SoundEffectNote = -1

const (
	NOTE_C SoundEffectNote = iota
	NOTE_C_SHARP
	NOTE_D
	NOTE_D_SHARP
	NOTE_E
	NOTE_F
	NOTE_F_SHARP
	NOTE_G
	NOTE_G_SHARP
	NOTE_A
	NOTE_A_SHARP
	NOTE_B
)

type SoundEffectOptions struct {
	id          int
	note        int
	octave      int
	duration    int
	channel     int
	leftVolume  int
	rightVolume int
	speed       int
}

var defaultSoundEffectOptions SoundEffectOptions = SoundEffectOptions{
	id:          -1,
	note:        -1,
	octave:      -1,
	duration:    -1,
	channel:     0,
	leftVolume:  15,
	rightVolume: 15,
	speed:       0,
}

func NewSoundEffectOptions() *SoundEffectOptions {
	options := new(SoundEffectOptions)
	*options = defaultSoundEffectOptions
	return options
}

func (options *SoundEffectOptions) SetId(id int) *SoundEffectOptions {
	options.id = id % 64
	return options
}

func (options *SoundEffectOptions) SetNote(note SoundEffectNote, octave int) *SoundEffectOptions {
	options.note = int(note) % 12
	options.octave = octave % 9
	return options
}

func (options *SoundEffectOptions) SetDuration(duration int) *SoundEffectOptions {
	options.duration = duration
	return options
}

func (options *SoundEffectOptions) SetChannel(channel int) *SoundEffectOptions {
	options.channel = channel % 4
	return options
}

func (options *SoundEffectOptions) SetSpeed(speed int) *SoundEffectOptions {
	if speed < -4 {
		options.speed = -4
	} else if speed > 3 {
		options.speed = 3
	} else {
		options.speed = speed
	}
	return options
}

func (options *SoundEffectOptions) SetVolume(level int) *SoundEffectOptions {
	level %= 16
	options.leftVolume = level
	options.rightVolume = level
	return options
}

func (options *SoundEffectOptions) SetStereoVolume(leftLevel, rightLevel int) *SoundEffectOptions {
	options.leftVolume = leftLevel % 16
	options.rightVolume = rightLevel % 16
	return options
}

type SpriteOptions struct {
	transparentColors []byte
	scale             int
	flip              int
	rotate            int
	width             int
	height            int
}

var defaultSpriteOptions SpriteOptions = SpriteOptions{
	transparentColors: nil,
	scale:             1,
	flip:              0,
	rotate:            0,
	width:             1,
	height:            1,
}

func NewSpriteOptions() *SpriteOptions {
	options := new(SpriteOptions)
	*options = defaultSpriteOptions
	return options
}

func (options *SpriteOptions) AddTransparentColor(color int) *SpriteOptions {
	if options.transparentColors == nil {
		options.transparentColors = make([]byte, 0, 1)
	}
	options.transparentColors = append(options.transparentColors, byte(color%16))
	return options
}

func (options *SpriteOptions) SetScale(scale int) *SpriteOptions {
	options.scale = scale
	return options
}

func (options *SpriteOptions) FlipHorizontally() *SpriteOptions {
	options.flip ^= 1
	return options
}

func (options *SpriteOptions) FlipVertically() *SpriteOptions {
	options.flip ^= 2
	return options
}

func (options *SpriteOptions) Rotate90CW() *SpriteOptions {
	options.rotate = (options.rotate + 1) % 4
	return options
}

func (options *SpriteOptions) Rotate90CCW() *SpriteOptions {
	options.rotate = (options.rotate - 1) % 4
	return options
}

func (options *SpriteOptions) Rotate180() *SpriteOptions {
	options.rotate = (options.rotate + 2) % 4
	return options
}

func (options *SpriteOptions) SetSize(width, height int) *SpriteOptions {
	options.width = width
	options.height = height
	return options
}

type SyncMask int

const SYNC_ALL SyncMask = 0

const (
	SYNC_TILES SyncMask = 1 << iota
	SYNC_SPRITES
	SYNC_MAP
	SYNC_SOUND_EFFECTS
	SYNC_MUSIC
	SYNC_PALETTE
	SYNC_FLAGS
	SYNC_SCREEN
)

type TexturedTriangleOptions struct {
	useTiles             bool
	transparentColors    []byte
	useDepthCalculations bool
	z0                   int
	z1                   int
	z2                   int
}

var defaultTexturedTriangleOptions TexturedTriangleOptions = TexturedTriangleOptions{
	useTiles:             false,
	transparentColors:    nil,
	useDepthCalculations: false,
	z0:                   0,
	z1:                   0,
	z2:                   0,
}

func NewTexturedTriangleOptions() *TexturedTriangleOptions {
	options := new(TexturedTriangleOptions)
	*options = defaultTexturedTriangleOptions
	return options
}

func (options *TexturedTriangleOptions) AddTransparentColor(color int) *TexturedTriangleOptions {
	if options.transparentColors == nil {
		options.transparentColors = make([]byte, 0, 1)
	}
	options.transparentColors = append(options.transparentColors, byte(color%16))
	return options
}

func (options *TexturedTriangleOptions) SetTextureDepth(z0, z1, z2 int) *TexturedTriangleOptions {
	options.useDepthCalculations = true
	options.z0 = z0
	options.z1 = z1
	options.z2 = z2
	return options
}

func (options *TexturedTriangleOptions) ToggleTextureSource() *TexturedTriangleOptions {
	options.useTiles = !options.useTiles
	return options
}

type TraceOptions struct {
	color byte
}

var defaultTraceOptions = TraceOptions{
	color: 15,
}

func NewTraceOptions() *TraceOptions {
	options := new(TraceOptions)
	*options = defaultTraceOptions
	return options
}

func (options *TraceOptions) SetColor(color int) *TraceOptions {
	options.color = byte(color % 16)
	return options
}

//go:export btn
func rawBtn(id int32) int32

func Btn(id int) bool {
	return rawBtn(int32(id%32)) > 0
}

//go:export btnp
func rawBtnp(id, hold, period int32) bool

func Btnp(id, hold, period int) bool {
	return rawBtnp(int32(id%32), int32(hold), int32(period))
}

//go:export clip
func rawClip(x, y, width, height int32)

func Clip(x, y, width, height int) {
	rawClip(int32(x), int32(y), int32(width), int32(height))
}

//go:export cls
func rawCls(color int8)

func Cls(color int) {
	rawCls(int8(color))
}

//go:export circ
func rawCirc(x, y, radius int32, color int8)

func Circ(x, y, radius, color int) {
	rawCirc(int32(x), int32(y), int32(radius), int8(color%16))
}

//go:export circb
func rawCircb(x, y, radius int32, color int8)

func Circb(x, y, radius, color int) {
	rawCircb(int32(x), int32(y), int32(radius), int8(color%16))
}

//go:export elli
func rawElli(x, y, radiusX, radiusY int32, color int8)

func Elli(x, y, radiusX, radiusY, color int) {
	rawElli(int32(x), int32(y), int32(radiusX), int32(radiusY), int8(color%16))
}

//go:export ellib
func rawEllib(x, y, radiusX, radiusY int32, color int8)

func Ellib(x, y, radiusX, radiusY, color int) {
	rawEllib(int32(x), int32(y), int32(radiusX), int32(radiusY), int8(color%16))
}

//go:export exit
func Exit()

//go:export fget
func rawFget(sprite int32, flag int8) bool

func Fget(sprite, flag int) bool {
	return rawFget(int32(sprite%512), int8(flag%8))
}

//go:export fset
func rawFset(sprite int32, flag int8, value bool)

func Fset(sprite, flag int, value bool) {
	rawFset(int32(sprite%512), int8(flag%8), value)
}

//go:export font
func rawFont(textBuffer unsafe.Pointer, x, y int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, characterWidth, characterHeight int8, fixed bool, scale int8, useAlternateFontPage bool) int32

func Font(text string, x, y int, options *FontOptions) (textWidth int) {
	if options == nil {
		options = &defaultFontOptions
	}

	transparentColorBuffer, transparentColorCount := toCBuffer(&options.transparentColors)
	textBuffer := toCString(&text)

	return int(rawFont(textBuffer, int32(x), int32(y), transparentColorBuffer, int8(transparentColorCount), int8(options.characterWidth), int8(options.characterHeight), options.fixed, int8(options.scale), options.alternateFont))
}

//go:export key
func rawKey(id int32) int32

func Key(id int) bool {
	return rawKey(int32(id)) > 0
}

//go:export keyp
func rawKeyp(id int8, hold, period int32) int32

func Keyp(id, hold, period int) bool {
	return rawKeyp(int8(id), int32(hold), int32(period)) > 0
}

//go:export line
func rawLine(x0, y0, x1, y1 float32, color int8)

func Line(x0, y0, x1, y1, color int) {
	rawLine(float32(x0), float32(y0), float32(x1), float32(y1), int8(color))
}

//go:export map
func rawMap(x, y, width, height, screenX, screenY int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, unused int32)

func Map(options *MapOptions) {
	if options == nil {
		options = &defaultMapOptions
	}

	transparentColorBuffer, transparentColorCount := toCBuffer(&options.transparentColors)

	rawMap(int32(options.x), int32(options.y), int32(options.width), int32(options.height), int32(options.screenX), int32(options.screenY), transparentColorBuffer, int8(transparentColorCount), 0)
}

//go:export memcpy
func rawMemcpy(destination, source, length int32)

func Memcpy(destination, source, length int) {
	rawMemcpy(int32(destination), int32(source), int32(length))
}

//go:export memset
func rawMemset(address, value, length int32)

func Memset(address, value, length int) {
	rawMemset(int32(address), int32(value), int32(length))
}

//go:export mget
func rawMget(x, y int32) int32

func Mget(x, y int) int {
	return int(rawMget(int32(x), int32(y)))
}

//go:export mset
func rawMset(x, y, value int32)

func Mset(x, y, value int) {
	rawMset(int32(x), int32(y), int32(value))
}

type mouseData struct {
	x       int16
	y       int16
	scrollX int8
	scrollY int8
	left    bool
	middle  bool
	right   bool
}

//go:export mouse
func rawMouse(data *mouseData)

func Mouse() (x, y int, left, middle, right bool, scrollX, scrollY int) {
	data := new(mouseData)
	rawMouse(data)

	x = int(data.x)
	y = int(data.y)
	left = data.left
	middle = data.middle
	right = data.right
	scrollX = int(data.scrollX)
	scrollY = int(data.scrollY)
	return
}

//go:export music
func rawMusic(track, frame, row int32, loop, sustain bool, tempo, speed int32)

func Music(options *MusicOptions) {
	if options == nil {
		options = &defaultMusicOptions
	}

	rawMusic(int32(options.track), int32(options.frame), int32(options.row), options.loop, options.sustain, int32(options.tempo), int32(options.speed))
}

//go:export peek
func rawPeek(address int32, bits int8) int8

func Peek(address int) byte {
	return byte(rawPeek(int32(address), 8))
}

func Peek4(address int) byte {
	return byte(rawPeek(int32(address), 4))
}

func Peek2(address int) byte {
	return byte(rawPeek(int32(address), 2))
}

func Peek1(address int) byte {
	return byte(rawPeek(int32(address), 1))
}

//go:export pix
func rawPix(x, y int32, color int8) uint8

func Pix(x, y, color int) int {
	return int(rawPix(int32(x), int32(y), int8(color%16)))
}

//go:export pmem
func Pmem(address int32, value int64) uint32

//go:export poke
func rawPoke(address int32, value, bits int8)

func Poke(address int, value byte) {
	rawPoke(int32(address), int8(value), 8)
}

func Poke4(address int, value byte) {
	rawPoke(int32(address), int8(value), 4)
}

func Poke2(address int, value byte) {
	rawPoke(int32(address), int8(value), 2)
}

func Poke1(address int, value byte) {
	rawPoke(int32(address), int8(value), 1)
}

//go:export print
func rawPrint(textBuffer unsafe.Pointer, x, y int32, color, fixed, scale, useAlternateFontPage int8) int32

func Print(text string, x, y int, options *PrintOptions) int {
	if options == nil {
		options = &defaultPrintOptions
	}

	textBuffer := toCString(&text)

	var optionFixed int8
	if options.fixed {
		optionFixed = 1
	}

	var optionAlternateFont int8
	if options.alternateFont {
		optionAlternateFont = 1
	}

	return int(rawPrint(textBuffer, int32(x), int32(y), int8(options.color), optionFixed, int8(options.scale), optionAlternateFont))
}

//go:export rect
func rawRect(x, y, width, height int32, color int8)

func Rect(x, y, width, height, color int) {
	rawRect(int32(x), int32(y), int32(width), int32(height), int8(color%16))
}

//go:export rectb
func rawRectb(x, y, width, height int32, color int8)

func Rectb(x, y, width, height, color int) {
	rawRectb(int32(x), int32(y), int32(width), int32(height), int8(color%16))
}

//go:export sfx
func rawSfx(id, note, octave, duration, channel, volumeLeft, volumeRight, speed int32)

func Sfx(options *SoundEffectOptions) {
	if options == nil {
		options = &defaultSoundEffectOptions
	}

	rawSfx(int32(options.id), int32(options.note), int32(options.octave), int32(options.duration), int32(options.channel), int32(options.leftVolume), int32(options.rightVolume), int32(options.speed))
}

//go:export spr
func rawSpr(id, x, y int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, scale, flip, rotate, width, height int32)

func Spr(id, x, y int, options *SpriteOptions) {
	if options == nil {
		options = &defaultSpriteOptions
	}

	transparentColorBuffer, transparentColorCount := toCBuffer(&options.transparentColors)

	rawSpr(int32(id), int32(x), int32(y), transparentColorBuffer, int8(transparentColorCount), int32(options.scale), int32(options.flip), int32(options.rotate), int32(options.width), int32(options.height))
}

//go:export sync
func rawSync(mask int32, bank, toCart int8)

func Sync(mask SyncMask, bank int, toCart bool) {
	var toCartValue int8
	if toCart {
		toCartValue = 1
	}

	rawSync(int32(mask), int8(bank), toCartValue)
}

//go:export ttri
func rawTtri(x0, y0, x1, y1, x2, y2, u0, v0, u1, v1, u2, v2 float32, useTiles int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, z0, z1, z2 float32, depth bool)

func Ttri(x0, y0, x1, y1, x2, y2, u0, v0, u1, v1, u2, v2 int, options *TexturedTriangleOptions) {
	if options == nil {
		options = &defaultTexturedTriangleOptions
	}

	transparentColorBuffer, transparentColorCount := toCBuffer(&options.transparentColors)

	var useTilesValue int32
	if options.useTiles {
		useTilesValue = 1
	}

	rawTtri(float32(x0), float32(y0), float32(x1), float32(y1), float32(x2), float32(y2), float32(u0), float32(v0), float32(u1), float32(v1), float32(u2), float32(v2), useTilesValue, transparentColorBuffer, int8(transparentColorCount), float32(options.z0), float32(options.z1), float32(options.z2), options.useDepthCalculations)
}

//go:export time
func Time() float32

//go:export trace
func rawTrace(messageBuffer unsafe.Pointer, color int8)

func Trace(message string, options *TraceOptions) {
	if options == nil {
		options = &defaultTraceOptions
	}

	messageBuffer := toCString(&message)

	rawTrace(messageBuffer, int8(options.color))
}

//go:export tri
func rawTri(x0, y0, x1, y1, x2, y2 float32, color int8)

func Tri(x0, y0, x1, y1, x2, y2, color int) {
	rawTri(float32(x0), float32(y0), float32(x1), float32(y1), float32(x2), float32(y2), int8(color))
}

//go:export trib
func rawTrib(x0, y0, x1, y1, x2, y2 float32, color int8)

func Trib(x0, y0, x1, y1, x2, y2, color int) {
	rawTrib(float32(x0), float32(y0), float32(x1), float32(y1), float32(x2), float32(y2), int8(color))
}

//go:export tstamp
func Tstamp() uint32

//go:linkname Start _start
func Start()

//go:export main.main
func main() {}
