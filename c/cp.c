#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct {
	int id;
	char *name;
}people_t;

int main(void)
{
	people_t people;
	people_t tmp;
	people.id = 100;
	people.name = calloc(1, 10);
	memcpy(people.name, "haha", 4);
	memcpy(&tmp, &people, sizeof(people));

	printf("%d, %s, addr: %p\n", people.id, people.name, people.name);
	free(people.name);
	printf("%d, %s, addr: %p\n", tmp.id, tmp.name, tmp.name);
	return 0;
}
