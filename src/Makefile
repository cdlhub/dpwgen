PROG=dpwgen
OBJECTS=dpwgen.o dpw_file.o
CFLAGS= -g -Wall -W -pedantic --std=c99
LDLIBS=
CC=gcc

all: $(PROG)

$(PROG): $(OBJECTS)

%.o: %.c
	$(CC) -c $(CFLAGS) $*.c

dpwgen.o: dpw_file.o dpw_macros.h
dpw_file.o: dpw_macros.h dpw_file.h

.PHONY: clean
clean:
	rm -f $(PROG) $(OBJECTS)