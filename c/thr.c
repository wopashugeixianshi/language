#define _GNU_SOURCE
#include <stdio.h>
#include <pthread.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/syscall.h>  
#include <sys/types.h>

#define gettid() syscall(__NR_gettid)
pthread_t tid;
void *handle_thr(void *arg) 
{
    //printf("The ID of this thread is: %ld\n", (long int)syscall(224));
	printf("thr : %ld, %ld, %ld\n", pthread_self(), tid, gettid());
	for(;;);
}

int main(void) 
{
	pthread_create(&tid, NULL, handle_thr, NULL);
	printf("main : %ld, %d\n", gettid(), getpid());
	pthread_join(tid, NULL);
	return 0;
}
