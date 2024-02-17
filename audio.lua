-- audio
function audio_init()
	audio={}
	audio.base=0x4300
	audio.pos=0
	audio.buffered=0
	audio.src=nil
	audio.len=0
	audio.playing=false
	audio.num_streams=4
	audio.streams={}
	audio.next_stream_idx=1
	audio.next_request_id=1

	for i=1,audio.num_streams do
		add(audio.streams, _init_audio_stream())
	end
end

function audio_legacy_update()
	if (not audio.playing) return

	audio.buffered = stat(108)

	if audio.buffered < 512 then
		local size = audio.pos - audio.len
		
		if size < -256 then
			size = 256
		else
			size *= -1
			audio.playing=false
		end		
		
		for i=1,size do
			poke(audio.base+(i-1), ord(sub(audio.src,i+audio.pos,i+audio.pos)))
		end
		
		serial(0x808,audio.base,size)
		
		if audio.playing then
			audio.pos += size
		else
			audio.pos = 0
		end
	end
end

function audio_legacy_play(src)
	audio.src=src
	audio.len=#src
	audio.playing=true
end

function audio_update()
	if (not audio.playing) return
end

function audio_play(sample_name,rate,callback)
	local request={}
	request.id=audio.next_request_id
	request.sample_name=sample_name
	request.rate=rate
	request.callback=callback

	local stream=audio.streams[audio.next_stream_idx]

	if stream.playing then
		callback("aborted", request.id)
	end

	stream.request=request
	stream.playing=true

	callback("started", request.id)

	audio.next_request_id += 1
	audio.next_stream_idx += 1
	if (audio.next_stream_idx==audio.num_streams+1) audio.next_stream_idx=1 
end

function audio_draw()
	pal()
	for i=1,audio.num_streams do
		local stream = audio.streams[i]
		print(i, (i-1)*24, 72)
		print(stream.playing, (i-1)*24, 80)
		if stream.playing then
			local request = stream.request
			print(request.id, (i-1)*24, 88)
			print(request.sample_name, (i-1)*24, 96)
		end
	end
end

function _init_audio_stream()
	local stream={}
	stream.request={}
	stream.playing=false
	return stream
end
