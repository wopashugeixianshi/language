#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

int main(int argc, char **argv)
{
	char *s1 = "hello";
	char *s2 = "world";
	if (strcmp(s1, s2) == 0) {
		printf("s1 == s2\n");
	} else {
		printf("s1 != s2\n");
	}
	return 0;
}
