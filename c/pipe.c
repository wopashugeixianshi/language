#include <stdio.h>  
#include <unistd.h>  
#include <string.h>  
#include <errno.h>  
#include <sys/wait.h>
int main()  
{  
    int fd[2]; 
    int ret = pipe(fd);  
    if (ret == -1)  
    {  
        perror("pipe error\n");  
        return 1;  
    } 
	while(1);
    pid_t id = fork();  
    if (id == 0)  
    {//child  
        int i = 0;  
        close(fd[0]);  
        char *child = "I am  child!";  
        while (i<10)  
        {  
            ret = write(fd[1], child, strlen(child) + 1);  
			printf("ret = %d\n", ret);
            sleep(2);  
            i++;  
        }  
    }  
    else if (id>0)  
    {//father  
        close(fd[1]);  
        char msg[100];  
        int status = 0;  
        int j = 0;  
        while (j<2)  
        {  
            memset(msg, '\0', sizeof(msg));  
            ssize_t s = read(fd[0], msg, sizeof(msg));  
            if (s>0)  
            {  
                msg[s - 1] = '\0';  
            }  
            printf("%s  %d\n", msg, j);  
            j++;  
        }  
        //写方还在继续,而读方已经关闭它的读端  
        close(fd[0]);  
        //pid_t ret = waitpid(id, &status, 0);  
        //printf("exitsingle(%d),exit(%d)\n", status & 0xff, (status >> 8) & 0xff);  
        //低八位存放该子进程退出时是否收到信号  
        //此低八位子进程正常退出时,退出码是多少  
    }  
    else  
    {//error  
        perror("fork error\n");  
        return 2;  
    }  

	sleep(20);
    return  0;  
}  
