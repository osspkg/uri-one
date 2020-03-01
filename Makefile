
HOST=root@116.202.104.72

sync:
	scp uri-one $(HOST):~/urione/
	scp -r data/ $(HOST):~/urione/
	scp -r source/ $(HOST):~/urione/