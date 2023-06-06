#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int main(int argc, char**argv) 
{
	FILE *fp = fopen(argv[1], "r");
	if (!fp) {
	
	}
	char str[1024] = {0};
	char str1[100] = {0};
	char str2[100] = {0};
	char str3[100] = {0};
	while (fgets(str,sizeof(str), fp)) {
		sscanf(str, "%[^:]:%[^:]:%[^:]", str1, str2, str3);
		printf("%s, %s, %s\n", str1, str2, str3);
		memset(str1, 0, sizeof(str1));
		memset(str2, 0, sizeof(str2));
		memset(str3, 0, sizeof(str3));
		sleep(5);
	}
	fclose(fp);
	return 0;
}
