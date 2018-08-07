PROG=dpwgen
OBJECTS=dpwgen.o dpw_file.o
CFLAGS= -g -Wall -W -pedantic
LDLIBS=
CC=gcc

all: $(PROG)

$(PROG): $(OBJECTS)

dpwgen.o: dpw_file.o dpw_macros.h dpwgen.c
dpw_file.o: dpw_macros.h dpw_file.h

.PHONY: clean
clean:
	rm -f $(PROG) $(OBJECTS)