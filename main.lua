function _init()
	-- switch upper rom to ram, render sprite 0 opaque
	poke(0x5f36,0x18)
	palt(0,false)
	-- enable keyboard
	poke(0x5f2d,1)
	key_scan_codes={20,31,26,32,8,21,34,23,35,28,36,24}
	keys={0,0,0,0,0,0,0,0,0,0,0,0}
	keys_off_clr={6,0,6,0,6,6,0,6,0,6,0,6}
	keys_on_clr={7,5,7,5,7,7,5,7,5,7,5,7}
	audio_init()
	audio_play(data.midc)
end

function _update()
	for k=1,12 do
		keys[k]=stat(28,key_scan_codes[k]) and 1
	end
	audio_update()
end

function _draw()
	cls()
	for k=1,12 do
		if keys[k]==1 then
			pal(k,keys_on_clr[k])
		else
			pal(k,keys_off_clr[k])
		end
	end
	map(0,0,0,0,16,8)
end
