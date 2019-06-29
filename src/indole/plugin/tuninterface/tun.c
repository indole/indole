#ifdef __linux__
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <linux/if.h>
#include <linux/if_tun.h>
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