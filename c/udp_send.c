#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define PORT 8888
#define MAX_LEN 1024

int main(int argc, char *argv[]) {
    if (argc < 3) {
        printf("Usage: %s <ip_address> <message>\n", argv[0]);
        return 1;
    }

    char* ip = argv[1];
    char* message = argv[2];

    int sockfd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
    if (sockfd < 0) {
        perror("socket");
        return 1;
    }

    struct sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(PORT);
    if (inet_pton(AF_INET, ip, &addr.sin_addr) < 0) {
        perror("inet_pton");
        return 1;
    }

    int len = sendto(sockfd, message, strlen(message), 0, (struct sockaddr*)&addr, sizeof(addr));
    if (len < 0) {
        perror("sendto");
        return 1;
    }

    close(sockfd);

    printf("Sent message to %s:%d: %s\n", ip, PORT, message);

    return 0;
}

