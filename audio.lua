function audio_init()
	audio={}
	audio.base=0x4300
	audio.pos=0
	audio.buffered=0
	audio.src=nil
	audio.len=0
	audio.playing=false
end

function audio_update()
	if (not audio.playing) return

	audio.buffered = stat(108)

	if (audio.buffered<512) then
		local size = audio.pos - audio.len
		
		if (size < -256) then
			size = 256
		else
			size *= -1
			audio.playing=false
		end		
		
		for i=1,size do
			poke(audio.base+(i-1), ord(sub(audio.src,i+audio.pos,i+audio.pos)))
		end
		
		serial(0x808,audio.base,size)
		
		if (audio.playing) then
			audio.pos += size
		else
			audio.pos = 0
		end
	end
end

function audio_play(src)
	audio.src=src
	audio.len=#src
	audio.playing=true
end
