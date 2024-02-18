-- main
function _init()
	-- switch upper rom to ram, render sprite 0 opaque
	poke(0x5f36,0x18)
	palt(0,false)
	-- enable keyboard
	poke(0x5f2d,1)
	keys_init()
	audio_init()
end

function _update()
	keys_update()
	audio_update()
end

function _draw()
	cls()
	keys_draw()
	audio_draw()
end
