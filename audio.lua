-- audio
function audio_init()
	audio={}
	audio.base=0x4300
	audio.buffered=0
	audio.playing=false
	audio.num_streams=4
	audio.buf_size=256
	audio.streams={}
	audio.next_stream_idx=1
	audio.next_request_id=1

	for i=1,audio.num_streams do
		add(audio.streams, _init_audio_stream())
	end
end

function audio_update()

	audio.buffered = stat(108)

	if ((not audio.playing) or (audio.buffered>=(audio.buf_size*2))) return

	memset(audio.base, 0x80, audio.buf_size)

	local value
	local stream
	local sample
	local size

	for p=0,audio.buf_size-1 do
		value=peek(audio.base+p)

		for s=1,audio.num_streams do
			stream=audio.streams[s]
			if stream.pos>0 then
				if stream.pos<=stream.len then
					sample=data.samples[stream.request.sample_name]
					value+=(ord(sub(sample,stream.pos,stream.pos))-128)*0.8
					stream.pos+=stream.request.rate
				else
					stream.pos=0
					stream.len=0
				end
			end
		end

		if (value>255) value=255
		if (value<0) value=0

		poke(audio.base+p, value)
	end

	serial(0x808,audio.base,audio.buf_size)

	local playing=false
	for s=1,audio.num_streams do
		if (audio.streams[s].pos>0) playing=true
	end
	audio.playing=playing
end

function audio_play(sample_name,rate)
	local sample = data.samples[sample_name]
	local request={}
	request.id=audio.next_request_id
	request.sample_name=sample_name
	request.rate=rate

	local stream=audio.streams[audio.next_stream_idx]

	stream.request=request
	stream.len=#sample
	stream.pos=1

	audio.playing=true
	audio.next_request_id = (audio.next_request_id+1)%256
	audio.next_stream_idx += 1
	if (audio.next_stream_idx==audio.num_streams+1) audio.next_stream_idx=1 
end

function audio_draw()
	pal()
	print(audio.playing, 0, 72)
	print(audio.buffered, 64, 72)
	for i=1,audio.num_streams do
		local stream = audio.streams[i]
		local y=80
		print(i, (i-1)*24, y)
		y+=8
		print(flr(stream.pos), (i-1)*24, y)
		y+=8
		print(stream.len, (i-1)*24, y)
		y+=8
		if stream.pos>0 then
			local request = stream.request
			print(request.id, (i-1)*24, y)
			y+=8
			print(request.sample_name, (i-1)*24, y)
			y+=8
		end
	end
end

function _init_audio_stream()
	local stream={}
	stream.request={}
	stream.len=0
	stream.pos=0
	return stream
end
