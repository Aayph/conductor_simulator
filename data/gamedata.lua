screenwidth = 0
screenheight = 0

traveltime = 8
note_jiggl_speed = 0.4
max_bad_v = 0.1

note_start_y = -0.1
note_end_y = 0.65

note_count = 0
notes_done = 0

level_age = 0
level_time = 0

level_play = 0
main_menu_state = 0

good_value = 10
bad_value = 50
good_count = 0
bad_count = 0

bad_min_time = 2
bad_max_time = 4

max_bad_notes = 0

song_time = 0

notes = {}
batontip = nil


lane_timer = {}
lane_time = 1.5

function get_relative_score()
	f = (bad_count * bad_value - good_count * good_value) / (max_bad_notes * bad_value)
	f = (f + 1) * 0.5
	return f
end

function get_score()
	f = bad_count * bad_value - good_count * good_value
	if f < 0 then
		f = 0
	end
	return f
end