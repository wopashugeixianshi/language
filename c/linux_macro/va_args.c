#include <stdio.h>

#define DEBUG(...) printf(__VA_ARGS__)
#define DEBUG1(fmt, ...) printf(fmt, __VA_ARGS__)

int main(void)
{
	DEBUG("hello %s\n", "world");
	DEBUG("hello\n");
	return 0;
}
