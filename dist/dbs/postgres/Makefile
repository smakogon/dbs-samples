#
# dbs-samples postgres database control
#
U = serhii
H = /var/run/postgresql
D = dbs_dev

all: drop init

drop:
	psql -h $(H) -U $(U) -d $(D) < schema/_drop-all.sql

init:
	psql -h $(H) -U $(U) -d $(D) < schema/_init-all.sql

.PHONY: all drop init infra infra-cons
