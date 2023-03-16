#include <stdio.h>


void func(int *p)
{
	printf("var=%d\n", *p);
	*p = 200;
	int var = 300;
	p = &var;
}

void func2(int **p)
{
	int var = 300;
	*p = &var;
}

int main(void)
{
	int int_var = 100;
	int *ptr = &int_var;
	//int **p = &ptr;
	func(ptr);
	printf("var=%d\n", *ptr);
	func2(&ptr);
	printf("var=%d\n", *ptr);
	return 0;
}
