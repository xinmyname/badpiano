-- main
function _init()
	-- switch upper rom to ram, render sprite 0 opaque
	poke(0x5f36,0x18)
	palt(0,false)
	-- enable keyboard
	poke(0x5f2d,1)
	keys_init()
	audio_init()
	patches={
		"epiano",
		"harpsi",
		"jazzy",
		"midc",
		"organ",
		"piano"
	}
	cur_patch=6
end

function _update()
	keys_update()
	audio_update()
	if btnp(0) then
		cur_patch=cur_patch-1
		if (cur_patch==0) cur_patch=6
	end
	if btnp(1) then
		cur_patch=cur_patch+1
		if (cur_patch==7) cur_patch=1
	end
end

function _draw()
	cls()
	keys_draw()
	pal()
	local l=#patches[cur_patch]
	local x=64-((l*8)/2)
	print("\^w\^t"..patches[cur_patch], x, 72, 12)
end
