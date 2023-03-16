#include <stdio.h>  
#include <stdlib.h>  
#include <sys/socket.h>  
#include <unistd.h>
#include <time.h>  
#include <sys/types.h>  
#include <netinet/in.h>  
#include <netdb.h>
#include <arpa/inet.h>

void msg()  
{  
  printf("EP:scan ip startport endport\nEP:scan ip 127.0.0.1 20 2009\n");  
}

int main(int argc,char** argv)  
{  
	char *ip;  
	int startport,endport,sockfd,i;  
	struct sockaddr_in to;  
	float costtime;  
	clock_t start,end;  
	if(4!=argc)  
	{  
		msg();  
		return 0;  
	}  
	ip=argv[1];  
	startport=atoi(argv[2]);  
	endport=atoi(argv[3]);  
	if(startport<1 || endport>65535 || endport<startport)  
	{  
		printf("端口范围出错/n");  
		return 0;   
	}  
	else{
		printf("IP:%s %d-%d\n",ip,startport,endport);  
	}
	to.sin_family=AF_INET;  
	to.sin_addr.s_addr=inet_addr(ip);  
	start=clock();  
	for(i=startport;i<=endport;i++)  
	{  
		sockfd=socket(AF_INET,SOCK_STREAM,0);  
		to.sin_port=htons(i);  
		if(connect(sockfd,(struct sockaddr *)&to,sizeof(struct sockaddr)) == 0) {
			send(sockfd, "hello", 5, 0);
			printf("%s    %d\n",ip,i);
			close(sockfd);
		};
	}  
	end=clock();  
	costtime=(float)(end-start)/CLOCKS_PER_SEC;  
	printf("用时:%f秒\n",costtime);  
	return 0;  
}
