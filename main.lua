function _init()
	palt(0,false)
	poke(0x5f36,8) //map draw 0
	poke(0x5f2d,1) //enable keyboard
	key_scan_codes={20,31,26,32,8,21,34,23,35,28,36,24}
	keys={0,0,0,0,0,0,0,0,0,0,0,0}
	keys_off_clr={6,0,6,0,6,6,0,6,0,6,0,6}
	keys_on_clr= {7,5,7,5,7,7,5,7,5,7,5,7}
end

function _update()
	for k=1,12 do
		keys[k]=stat(28,key_scan_codes[k]) and 1
	end
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
