#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <pthread.h>

#define BUFFER_SIZE 1024
#define PORT 8888

void *packet_receiver(void *sockfd);
int main() {
  int sockfd;
  struct sockaddr_in addr;
  socklen_t len = sizeof(addr);
  char buffer[BUFFER_SIZE];

  // 创建 UDP socket
  if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) < 0) {
    perror("socket");
    exit(EXIT_FAILURE);
  }

  memset(&addr, 0, len);
  addr.sin_family = AF_INET;
  addr.sin_port = htons(PORT);
  addr.sin_addr.s_addr = htonl(INADDR_ANY);

  // 会负载均衡到所有监听的程序上
  //int optval = 1;
  //if (setsockopt(sockfd, SOL_SOCKET, SO_REUSEPORT, &optval, sizeof(optval)) == -1) {
  //  	  perror("setsockopt");
  //  	  exit(EXIT_FAILURE);
  //}

  // 绑定地址和端口
  if (bind(sockfd, (struct sockaddr *)&addr, len) < 0) {
    perror("bind");
    exit(EXIT_FAILURE);
  }
  printf("main sockfd = %d\n", sockfd);
  printf("UDP server is listening on port %d...\n", PORT);

  pthread_t tid;
  int ret = pthread_create(&tid, NULL, packet_receiver, (void *)&sockfd);
  if (ret != 0) {
		  perror("pthread_create");
		  exit(EXIT_FAILURE);
  }

  while (1) {
    // 接收消息
    int n = recvfrom(sockfd, buffer, BUFFER_SIZE, 0, (struct sockaddr *)&addr, &len);

    if (n < 0) {
      perror("recvfrom");
      exit(EXIT_FAILURE);
    }

    // 打印接收到的消息
    buffer[n] = '\0';
    printf("Received message: '%s'[main]\n", buffer);
  }

  pthread_join(tid, NULL);
  close(sockfd);
  return 0;
}

void *packet_receiver(void *sockfd) 
{
	int fd = *(int *)sockfd;
	printf("thread sockfd = %d\n", fd);
  	char buffer[BUFFER_SIZE];
	struct sockaddr_in addr;
	socklen_t len = sizeof(addr);
	while (1) {
			// 接收消息
			int n = recvfrom(fd, buffer, BUFFER_SIZE, 0, (struct sockaddr *)&addr, &len);

			if (n < 0) {
					perror("recvfrom");
					exit(EXIT_FAILURE);
			}

			// 打印接收到的消息
			buffer[n] = '\0';
			printf("Received message: '%s'[thread]\n", buffer);
	}
	
}
