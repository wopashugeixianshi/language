#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/types.h>

int main() {
	int buffer_size;
	socklen_t buffer_size_len = sizeof(buffer_size);
	if (getsockopt(sock, SOL_SOCKET, SO_RCVBUF, &buffer_size, &buffer_size_len) < 0) {
  		perror("getsockopt");
    	return -1;
	}

	printf("%d\n", buffer_size);
}
