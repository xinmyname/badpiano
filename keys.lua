-- keys
function keys_init()
	keys={}
	keys.scan_codes={20,31,26,32,8,21,34,23,35,28,36,24}
	keys.last_state={0,0,0,0,0,0,0,0,0,0,0,0}
	keys.state={0,0,0,0,0,0,0,0,0,0,0,0}
	keys.off_color={6,0,6,0,6,6,0,6,0,6,0,6}
	keys.on_color={7,5,7,5,7,7,5,7,5,7,5,7}
end

function keys_update()
	-- ignore keypresses when ctrl or apple keys are pressed
	if (stat(28,227) or stat(28,231) or stat(28, 224) or stat(28, 228)) return

	for k=1,12 do
		keys.state[k]=stat(28,keys.scan_codes[k]) and 1 or 0
		if ((keys.last_state[k] == 0) and (keys.state[k] == 1)) then
			audio_play("midc",1,keys_audio_callback)
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

function keys_audio_callback(reason,request_id)
end
