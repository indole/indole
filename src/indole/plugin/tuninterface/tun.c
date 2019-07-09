#ifdef __linux__
#include <arpa/inet.h>
#include <fcntl.h>
#include <linux/if.h>
#include <linux/if_tun.h>
#include <net/route.h>
#include <netinet/in.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#endif

int tun_alloc(char *dev)
{
#ifdef __linux__
  struct ifreq ifr;
  int fd, e;
  if ((fd = open("/dev/net/tun", O_RDWR)) < 0)
  {
    return fd;
  }
  memset(&ifr, 0, sizeof(ifr));
  ifr.ifr_flags = IFF_TUN | IFF_NO_PI;
  strncpy(ifr.ifr_name, dev, IFNAMSIZ);
  if ((e = ioctl(fd, TUNSETIFF, (void *)&ifr)) < 0)
  {
    close(fd);
    return e;
  }
  return fd;
#else
  return -1;
#endif
}