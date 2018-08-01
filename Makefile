PROG=dpwgen
OBJECTS=
CFLAGS= -g -Wall -O3
LDLIBS=
CC=gcc

$(PROG): $(OBJECTS)

clean:
	rm -f $(PROG) $(PROG).o