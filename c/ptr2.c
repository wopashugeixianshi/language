#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>

int main(void)
{
	char *ptr = NULL;
	char *tmp = NULL;
	ptr = tmp;
	tmp = calloc(1, 10); 
	if (!tmp) {
	
	}
	memcpy(tmp, "haha", 4);
	
	printf("%s\n", tmp);
	return 0;
}
