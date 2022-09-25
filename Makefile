# Build dependencies
GO = tinygo
WASM_OPT = wasm-opt
TIC80 = tic80

# Whether to build for debugging instead of release
DEBUG = 1

# Compilation flags
GOFLAGS = -target ./target.json -panic print
ifeq ($(DEBUG), 1)
	GOFLAGS += -opt 1
else
	GOFLAGS += -opt z -no-debug
endif

# wasm-opt flags
WASM_OPT_FLAGS = -Oz --zero-filled-memory --strip-producers

all: build cartridge run

build:
	$(GO) build $(GOFLAGS) -o out.wasm .
ifneq ($(DEBUG), 1)
ifeq ($(shell command -v $(WASM_OPT)),'')
	@echo Tip: $(WASM_OPT) was not found. Install it from binaryen for smaller builds!
else
	$(WASM_OPT) $(WASM_OPT_FLAGS) out.wasm -o out.wasm
endif
endif

cartridge:
	$(TIC80) --cli --fs . --cmd 'load cart.wasmp & import binary out.wasm & save game.tic & exit'

run:
	$(TIC80) --fs . --cmd 'load game.tic & run & exit'

clean:
	rm out.wasm
	rm game.tic