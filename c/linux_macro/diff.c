#include <stdio.h>
#include <stdlib.h>

#define Max(max1, max2) ({ \
	typeof(max1) _num1 = (max1); \
	typeof(max2) _num2 = (max2); \
	(void) (&_num1 == &_num2); \
	_num1 > _num2 ? _num1 : _num2;})

int main(int argc, char **argv)
{
	int num1 = 100, num2 = 200;
	printf("max=%d\n", Max(atoi(argv[1]), atoi(argv[2])));
	return 0;
}
