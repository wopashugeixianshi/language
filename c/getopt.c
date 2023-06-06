#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char **argv)
{	
	int arg = 0;
	while (arg != -1) {
		arg = getopt(argc, argv, "n:et::");
		switch(arg) {
		case -1:
			break;
		case 'e':
			printf("e\n");
			break;
		case 'n':
			printf("n, %s\n", optarg);
			break;
		case 't':
			printf("t, %s\n", optarg);
			break;
		}
	}
	return 0;
}
