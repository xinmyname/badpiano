-- keys
function keys_init()
	keys={}
	keys.scan_codes={20,31,26,32,8,21,34,23,35,28,36,24}
	keys.last_state={0,0,0,0,0,0,0,0,0,0,0,0}
	keys.state={0,0,0,0,0,0,0,0,0,0,0,0}
	keys.off_color={6,0,6,0,6,6,0,6,0,6,0,6}
	keys.on_color={7,5,7,5,7,7,5,7,5,7,5,7}

	keys.rates={
		1.0000,
		1.0595,
		1.1225,
		1.1892,
		1.2599,
		1.3348,
		1.4142,
		1.4983,
		1.5874,
		1.6818,
		1.7818,
		1.8877
	}
end

function keys_update()
	-- ignore keypresses when ctrl or apple keys are pressed
	if (stat(28,227) or stat(28,231) or stat(28, 224) or stat(28, 228)) return

	for k=1,12 do
		keys.state[k]=stat(28,keys.scan_codes[k]) and 1 or 0
		if ((keys.last_state[k] == 0) and (keys.state[k] == 1)) then
			audio_play(patches[cur_patch],keys.rates[k])
		end
		keys.last_state[k] = keys.state[k]
	end
end

function keys_draw()
	for k=1,12 do
		if keys.state[k]==1 then
			pal(k,keys.on_color[k])
		else
			pal(k,keys.off_color[k])
		end
		
		local txt
		
	end
	map(0,0,0,0,16,8)

	for k=1,12 do
		keys.last_state[k] = keys.state[k]
	end

end
