#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <net/if.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <string.h>
#include <sys/types.h>
#include <linux/if_tun.h>

int tun_alloc(int flags)
{
	struct ifreq ifr;
	int fd, err;
	char *clonedev = "/dev/net/tun";

	if ((fd = open(clonedev, O_RDWR)) < 0) {
		fprintf(stderr, "Failed to open \n");
		return fd;
	}

	memset(&ifr, 0, sizeof(ifr));
	ifr.ifr_flags = flags;

	if (ioctl(fd, TUNSETIFF, (void *)&ifr)) {
		perror("ioctl");
	}

	fprintf(stdout, "Open tun/tap device :%s for reading....\n", ifr.ifr_name);
	return fd;
err:
	close(fd);
	return -1;
}

int main(void)
{
	int tun_fd, nread;
	char buffer[1500];

	/*
		Flags: IFF_TUN -TUN device (no Ethernet headers)
		       IFF_TAP -TAP device
			   TFF_NO_PI - Do not provide packet information
	*/
	tun_fd = tun_alloc(IFF_TUN | IFF_NO_PI);
	if (tun_fd == -1) {
		return -1;
	}

	while (1) {
		nread = read(tun_fd, buffer, sizeof(buffer));
		if (nread < 0) {
			perror("Reading from interface");
			close(tun_fd);
			return -1;
		}
		fprintf(stdout, "Read %d bytes from tun/tap device\n", nread);
	}

	return 0;
}
