INSERT INTO lifecyclestate (name) VALUES
	 ('Draft'),
	 ('New'),
	 ('Open'),
	 ('Analysis'),
	 ('Design'),
	 ('Development'),
	 ('Integration'),
	 ('Verification'),
	 ('Fixed'),
	 ('Closed'),
	 ('Retest'),
	 ('Rejected'),
	 ('Assigned');

INSERT INTO public.lifecycle ("name",start_state_id) VALUES
	 ('Project',1);

INSERT INTO public.statetransition (lifecycle_id,from_state_id,to_state_id) VALUES
	 (1,1,2),
	 (1,2,4),
	 (1,4,5),
	 (1,5,6),
	 (1,6,10);

INSERT INTO public.project (created,updated,"name",project_key,description,state_id,lifecycle_id,created_by_id,owner_id) VALUES
	 ('2023-11-21 19:34:03.452308+01','2023-11-21 19:34:03.452308+01','COSMOS','COSMOS','Super COSMOS Project',1,1,1,1);

INSERT INTO public.users (first_name,last_name,email) VALUES
	 ('Szymon','Szczawinski','szymon.szczawinski@mail.com');

