package postgres

const (
	INIT_LIFECYCLESTATE string = "INSERT INTO lifecyclestate (name) VALUES ('Draft'), ('New'), ('Open'), ('Analysis'), ('Design'), ('Development'), ('Integration')," +
		"('Verification'), ('Fixed'), ('Closed'), ('Retest'), ('Rejected'), ('Assigned');"
	INIT_LIFECYCLE        string = "INSERT INTO lifecycle (name,start_state_id) VALUES ('Project',1);"
	INIT_STATE_TRANSITION string = "INSERT INTO public.statetransition (lifecycle_id,from_state_id,to_state_id) VALUES (1,1,2), (1,2,4), (1,4,5), (1,5,6), (1,6,10);"
	INIT_USERS            string = "INSERT INTO public.users (first_name,last_name,email) VALUES ('Szymon','Szczawinski','szymon.szczawinski@mail.com');"
)
